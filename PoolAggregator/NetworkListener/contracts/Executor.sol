// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Executor {

    address owner;

    constructor(){
        owner = msg.sender;
    }

    modifier onlyOwner(){
        require(msg.sender == owner,"You are not owner");
        _;
    }

    event CallAggregator();

    function startProcess() external onlyOwner  {
        emit CallAggregator();
    }

}
