// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity 0.8.13;

import "forge-std/Test.sol";
import "/Users/z/Documents/x/interface/interface.sol";
import "src/target_handler.sol";
import "./hackbase.sol";

//  this file is for invariant test

contract HackInvariant is HackBase {
    bytes4[] targetHandlerSelectors;

    function setUp() public {
        super.setUpTargetFromCreationCode("./src/target_creation_code.txt");

        // init hander
        handler = new TargetHandler(target);
        targetContract(address(handler));
        delete targetHandlerSelectors;

        targetHandlerSelectors.push(handler.toFallback.selector);
        // add more selectors...

        // set targetSelector
        targetSelector(FuzzSelector({addr: address(handler), selectors: targetHandlerSelectors}));
    }

    // function invariant_TokenSupplyIsAlwaysZero() public {
    //     assertEq(0, target.totalSupply());
    // }
}
