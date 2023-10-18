// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Executor {
    address owner;

    constructor(){
        owner = msg.sender;
    }

    event StartExecution();

    modifier onlyOwner {
        require(msg.sender == owner,"You are not owner");
        _;
    }

    function getOwner() public returns(address){
        return owner;
    }

    function setOwner(address _owner) public onlyOwner{
        owner = _owner;
    }

    function startExecution() public onlyOwner{
        emit StartExecution();
    }
}
