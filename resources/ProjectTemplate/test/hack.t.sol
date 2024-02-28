// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8.10;

import "forge-std/Test.sol";
import "/Users/z/Documents/x/interface/interface.sol";

contract Hack is Test {
    function setUp() public {
        vm.createSelectFork("theNet");
    }

    function testMe() public {
        uint256 height = block.number;
        console.log("height:", height);
        assertTrue(height > 0);
    }
}
