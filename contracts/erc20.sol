// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MyToken {
    event Transfer(address indexed from, address indexed to, uint256 value);

    function transfer(address to, uint256 value) public returns (bool) {
        emit Transfer(msg.sender, to, value);
        return true;
    }
}
