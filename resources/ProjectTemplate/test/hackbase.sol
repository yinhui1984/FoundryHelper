// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity 0.8.13;

import "forge-std/Test.sol";
import "/Users/z/Documents/x/interface/interface.sol";
import "src/target_handler.sol";

contract HackBase is Test {
    ITargetContract target;
    TargetHandler handler;
    address deployedTargetAddr;

    function CreateFromCreationCode(string memory filePath) internal returns (address) {
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
        require(addr != address(0), "Failed to create contract from creation code");
        return addr;
    }

    function setUpTargetFromCreationCode(string memory filePath) internal {
        deployedTargetAddr = CreateFromCreationCode(filePath);
        console.log("target deployed:", deployedTargetAddr);
        target = ITargetContract(deployedTargetAddr);
    }

    receive() external payable {}
}
