require("@nomicfoundation/hardhat-toolbox");

const SEPOLIA_PRIVATE_KEY = "661260045b66da8aa4b5e6ac127f07021bee203136ee769fcaebb97c2eeb3708";

module.exports = {
  solidity: "0.8.24",
  networks: {
    sepolia: {
      url: "https://sepolia.infura.io/v3/c0fd902e8c32475f85909278c7352314",
      accounts: [SEPOLIA_PRIVATE_KEY]
    }
  }
};
