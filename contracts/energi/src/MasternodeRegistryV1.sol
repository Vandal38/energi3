// Copyright 2019 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// Energi Governance system is the fundamental part of Energi Core.

// NOTE: It's not allowed to change the compiler due to byte-to-byte
//       match requirement.
pragma solidity 0.5.9;
//pragma experimental SMTChecker;

import { GlobalConstants } from "./constants.sol";
import { IGovernedContract, GovernedContract } from "./GovernedContract.sol";
import { IBlockReward } from "./IBlockReward.sol";
import { IMasternodeRegistry } from "./IMasternodeRegistry.sol";

/**
 * Genesis hardcoded version of MasternodeRegistry
 *
 * NOTE: it MUST NOT change after blockchain launch!
 */
contract MasternodeRegistryV1 is
    GlobalConstants,
    GovernedContract,
    IBlockReward,
    IMasternodeRegistry
{
    // IMasternodeRegistry
    //---------------------------------
    function announce(address masternode, uint32 ipv4address) external {}
    function denounce() external {}
    function heartbeat(uint sw_features) external {}
    function invalidate(address masternode) external {}
    function validate(address masternode) external {}
    function isValid(address masternode) external view returns(bool) {}
    function count() external view returns(uint active, uint total, uint max_of_all_times) {}

    // IGovernedContract
    //---------------------------------
    constructor(address _proxy) public GovernedContract(_proxy) {}
    function migrate(IGovernedContract) external requireProxy {}

    // IBlockReward
    //---------------------------------
    function reward(uint) external payable {
    }

    function getReward(uint _blockNumber)
        external view
        returns(uint amount)
    {
        if (_blockNumber > 0) {
            amount = REWARD_MASTERNODE_V1;
        }
    }

    // Safety
    //---------------------------------
    function () external payable {
        revert("Not supported");
    }
}
