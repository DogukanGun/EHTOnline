// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "@openzeppelin/contracts@4.9.3/token/ERC20/ERC20.sol";

contract TestCrossChainToken is ERC20 {
    constructor() ERC20("TestCrossChainToken", "TCCT") {
        _mint(msg.sender, 1000000 * 10**decimals()); // Mint 1,000,000 tokens to deployer
    }

}
