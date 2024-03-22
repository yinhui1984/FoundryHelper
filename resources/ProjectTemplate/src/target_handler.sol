// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity 0.8.13;

import "forge-std/Base.sol";
import "forge-std/StdCheats.sol";
import "forge-std/StdUtils.sol";
import "/Users/z/Documents/x/interface/interface.sol";

// Change IECR20 to the interface you want to test
interface ITargetContract is IERC20 {}

contract TargetHandler is CommonBase, StdCheats, StdUtils {
    ITargetContract public target;

    constructor(ITargetContract _target) {
        target = _target;
    }

    receive() external payable {}

        function toFallback(uint256 amount) external payable {
        amount=bound(amount, 0, address(this).balance);
        (bool ok,)=address(target).call{value: amount}("");
        require(ok, "send value failed");
    }
}
