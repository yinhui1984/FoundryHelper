// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity 0.8.13;

import "forge-std/Test.sol";
import "/Users/z/Documents/x/interface/interface.sol";
import "src/target_handler.sol";
import "./hackbase.sol";

contract Hack is HackBase {
    uint256 public initBalance = 10 ether;

    function setUp() public {
        // 如果离线使用,不要fork
        // vm.createSelectFork("theNet");
        super.setUpTargetFromCreationCode("./src/target_creation_code.txt");

        // deal(address(this), initBalance);
        // deal(address(target), address(this), initBalance);
    }

    function testContractName() public {
        console.log("token name:", target.name());
        assertTrue(keccak256(abi.encodePacked(target.name())) != keccak256(abi.encodePacked("")));
    }
}
