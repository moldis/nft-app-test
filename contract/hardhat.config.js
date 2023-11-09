require("@nomiclabs/hardhat-ethers");
require("dotenv").config();

var tdly = require("@tenderly/hardhat-tenderly");
tdly.setup({ automaticVerifications: true });

const privateKey = process.env.PRIVATE_KEY;
const tenderlyUrl = process.env.TENDERLY_URL;

module.exports = {
  solidity: "0.8.20",
  defaultNetwork: "sepolia",
  networks: {
    sepolia: {
      url: tenderlyUrl,
      accounts: [`0x${privateKey}`],
      gasPrice: 20000000000,
      gas: 6000000,
      chainId: 11155111,
    },
  },
  tenderly: {
    project: "crypto",
    username: "moldis",
    privateVerification: true,
  },
};