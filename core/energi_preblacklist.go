// Copyright 2019 The Energi Core Authors
// This file is part of the Energi Core library.
//
// The Energi Core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Energi Core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Energi Core library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"errors"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"

	energi_abi "energi.world/core/gen3/energi/abi"
	energi_params "energi.world/core/gen3/energi/params"
)

const (
	preblFeeGasLimit uint64 = 500000
)

var (
	energiBLProposeID types.MethodID
)

func init() {
	bl_abi, err := abi.JSON(strings.NewReader(energi_abi.IBlacklistRegistryABI))
	if err != nil {
		panic(err)
	}
	copy(energiBLProposeID[:], bl_abi.Methods["propose"].Id())
}

var (
	pbCleanupTimeout = time.Minute
	pbPeriod         = time.Hour

	ErrPreBlacklist = errors.New("preliminary blacklisted")
)

type preBlacklist struct {
	proposed    map[common.Address]time.Time
	nextCleanup time.Time
	timeNow     func() time.Time
}

func newPreBlacklist() *preBlacklist {
	return &preBlacklist{
		proposed:    make(map[common.Address]time.Time),
		nextCleanup: time.Now().Add(pbCleanupTimeout),
		timeNow:     time.Now,
	}
}

func (pb *preBlacklist) cleanupRoutine(now time.Time) {
	if pb.nextCleanup.After(now) {
		return
	}

	pb.nextCleanup = now.Add(pbCleanupTimeout)
	//---
	for k, v := range pb.proposed {
		if now.Sub(v) > pbPeriod {
			delete(pb.proposed, k)
		}
	}
}

func (pb *preBlacklist) processTx(pool *TxPool, tx *types.Transaction) error {
	now := pb.timeNow()

	pb.cleanupRoutine(now)

	// Check if know preliminary blacklist
	//---
	sender, err := types.Sender(pool.signer, tx)
	if err != nil {
		log.Debug("Pre-blacklist sender error", "err", err)
		return err
	}

	if pb.isActive(sender, now) {
		log.Debug("Pre-blacklisted sender", "sender", sender)
		return ErrPreBlacklist
	}

	// Check if the call is valid
	//---
	pb.processProposal(pool, sender, now, tx)

	return nil
}

func (pb *preBlacklist) isActive(sender common.Address, now time.Time) bool {
	if t, ok := pb.proposed[sender]; ok && now.Sub(t) <= pbPeriod {
		return true
	}

	return false
}

func (pb *preBlacklist) processProposal(
	pool *TxPool,
	sender common.Address,
	now time.Time,
	tx *types.Transaction,
) {
	// Check if a new blacklist proposal
	//---
	if to := tx.To(); to == nil || *to != energi_params.Energi_BlacklistRegistry {
		return
	}
	if method := tx.MethodID(); method != energiBLProposeID {
		return
	}
	// DBL-10 - only enable for EBI proposals which have zero cost by design
	if tx.GasPrice().Cmp(common.Big0) != 0 {
		return
	}

	//---
	var target common.Address
	callData := tx.Data()
	copy(target[:], callData[16:36])

	// Do not reset timeout, if already known!
	if _, ok := pb.proposed[target]; ok {
		return
	}

	statedb := pool.currentState.Copy()

	if IsWhitelisted(statedb, target) {
		log.Warn("Skipping preliminary blacklist for whitelisted target",
			"target", target.Hex(), "sender", sender.Hex())
		return
	}

	msg := types.NewMessage(
		sender,
		tx.To(),
		tx.Nonce(),
		tx.Value(),
		tx.Gas(),
		tx.GasPrice(),
		callData,
		false,
	)

	bc := pool.chain.(*BlockChain)
	if bc == nil {
		log.Debug("PreBlacklist on missing blockchain")
		return
	}
	vmc := bc.GetVMConfig()
	ctx := NewEVMContext(msg, bc.CurrentHeader(), bc, &sender)
	ctx.GasLimit = preblFeeGasLimit
	evm := vm.NewEVM(ctx, statedb, bc.Config(), *vmc)

	gp := new(GasPool).AddGas(tx.Gas())
	output, _, failed, err := ApplyMessage(evm, msg, gp)
	if failed || err != nil {
		log.Debug("PreBlacklist failure at execution",
			"sender", sender, "target", target.Hex(), "err", err, "output", output)
		return
	}

	if len(output) != len(common.Hash{}) {
		log.Debug("PreBlacklist at unpack",
			"sender", sender, "target", target.Hex(), "output", output)
		return
	}

	// New pre-blacklist item
	//---
	log.Warn("New preliminary blacklist", "target", target.Hex(), "sender", sender.Hex())
	pb.proposed[target] = now
	pool.removeBySenderLocked(target)
}

func (pb *preBlacklist) filterBlocks(blocks types.Blocks) types.Blocks {
	pb.cleanupRoutine(pb.timeNow())

	for i, b := range blocks {
		if _, ok := pb.proposed[b.Coinbase()]; ok {
			return blocks[:i]
		}
	}

	return blocks
}

//=============================================================================

func (pool *TxPool) PreBlacklistHook(blocks types.Blocks) types.Blocks {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	return pool.preBlacklist.filterBlocks(blocks)
}

func (pool *TxPool) RemoveBySender(sender common.Address) bool {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	return pool.removeBySenderLocked(sender)
}

func (pool *TxPool) removeBySenderLocked(sender common.Address) bool {
	res := false

	if txs, ok := pool.pending[sender]; ok {
		for _, tx := range txs.Flatten() {
			txhash := tx.Hash()
			log.Trace("Removing by sender", "txhash", txhash, "sender", sender)
			pool.removeTx(txhash, true)
		}

		res = true
	}

	if txs, ok := pool.queue[sender]; ok {
		for _, tx := range txs.Flatten() {
			txhash := tx.Hash()
			log.Trace("Removing by sender", "txhash", txhash, "sender", sender)
			pool.removeTx(txhash, true)
		}

		res = true
	}

	return res
}