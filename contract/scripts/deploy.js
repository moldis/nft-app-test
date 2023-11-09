const { ethers } = require("hardhat");
require("dotenv").config();

const contractName = "NFTMarketplace"

async function main() {
    const [deployer] = await ethers.getSigners();
    provider = new ethers.providers.JsonRpcProvider(process.env.TENDERLY_URL);

    console.log("Deploying with " + deployer.address)
    const factory = await ethers.getContractFactory(contractName, deployer);
    console.log("Deploying token...");

    const atomToken = await factory.deploy(deployer.address);
    await atomToken.deployed();
    console.log("Smart Contract deployed to:", atomToken.address);
}

// Do the thing!
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });