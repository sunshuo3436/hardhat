//A simple token function,
//including token issuance,
//transfers and balance inquiries
pragma solidity ^0.8.24;

import "hardhat/console.sol";

contract Token {

    string public name = "My Hardhat Token";
    string public symbol = "MBT";//Names and symbols of tokens

    // fixed issue, stored in an unsigned integer
    uint256 public totalSupply = 1000000;

    // An address type variable is used to store ethereum accounts.
    address public owner;

    // A mapping is a key/value map. Store each account balance.
    mapping(address => uint256) balances;

    // Event to track token transfers
    event Transfer(address indexed from, address indexed to, uint256 value);

    //Only once when the contract is created.
    constructor() {
    //TotalSupply is assign to sender(which account deploy the smart contract
        balances[msg.sender] = totalSupply;
        owner = msg.sender;
    }
    //External:only can call from outside
    function transfer(address to, uint256 amount) external {
        // Check if the transaction sender has enough tokens.
        require(balances[msg.sender] >= amount, "Not enough tokens");
        console.log(
            "Transferring from %s to %s %s tokens",
            msg.sender,
            to,
            amount
        );

        // Transfer the amount.
        balances[msg.sender] -= amount;
        balances[to] += amount;

        // Emit Transfer event
        emit Transfer(msg.sender, to, amount);
    }

    function balanceOf(address account) external view returns (uint256) {
        return balances[account];
    }
}
