// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity 0.8.13;

import "forge-std/Test.sol";
import "/Users/z/Documents/x/interface/interface.sol";
import "src/target_handler.sol";



contract Hack is Test {
    ITargetContract target;
    TargetHandler handler;
    address deployedTargetAddr;

    uint256 public initBalance = 10 ether;

    receive() external payable {}

    function createFromCreationCode(string memory filePath) public returns (address) {
        // 可以使用环境变量传入
        // string memory creationCodeEnv = "THE_CREATION_CODE";
        // bytes memory creationCode = bytes(vm.envBytes(creationCodeEnv));

        // 也可以文件读取,需要在foundry.toml中加入权限 比如粗暴的fs_permissions = [{ access = "read-write", path = "./"}]
        // 这里的路径是相对于项目根的路径, 不是相对于代码文件
        string memory content = vm.readFile(filePath);
        bytes memory creationCode = vm.parseBytes(content);


        require(creationCode.length > 0, "No creation code");
        address addr;
        assembly {
            addr := create(0, add(creationCode, 0x20), mload(creationCode))
        }
        require(addr != address(0), "Failed to create contract");
        return addr;
    }

    function setUp() public {
        // 如果离线使用,不要fork
        // 如果有invariant test, 不建议fork, 否则非常慢
        // vm.createSelectFork("theNet");

        deployedTargetAddr = createFromCreationCode("./src/target_creation_code.txt");
        console.log("target deployed:", deployedTargetAddr);
        target = ITargetContract(deployedTargetAddr);

        handler = new TargetHandler(target);
        deal(address(handler), initBalance);

        // 设置目标合约的地址,非常重要, 不然foundry在做不变性测试会把我们所有的合约都当成目标合约
        targetContract(address(handler));

        // 设置目标合约的selector,非常重要, 不然foundry会对测试合约的所有可调用函数进行测试
        // 而我们只需要测试我们关心的函数, 并且有些函数会破坏我们要测试的“不可变性”
        bytes4[] memory selectors = new bytes4[](1);
        selectors[0] = handler.toFallback.selector;
        // selectors[1] = handler.toDeposit.selector;
        // add more...
        targetSelector(FuzzSelector({addr: address(handler), selectors: selectors}));
    }

    function testContractName() public {
        console.log("token name:", target.name());
        assertTrue(keccak256(abi.encodePacked(target.name())) != keccak256(abi.encodePacked("")));
    }

    function invariant_test_1() public {
        uint256 amount = 1 ether;
        uint256 balanceBefore = address(this).balance;
        uint256 targetBalanceBefore = address(target).balance;
        handler.toFallback{value: amount}(amount);
        uint256 balanceAfter = address(this).balance;
        uint256 targetBalanceAfter = address(target).balance;
        assertTrue(balanceBefore - amount == balanceAfter);
        assertTrue(targetBalanceBefore + amount == targetBalanceAfter);
    }
}
