const { ethers } = require("hardhat");
require("dotenv").config();

const tokenAddress = "0x8BA4Da9c3dfA6d895e6A1E0d72f6F64906769F54"
const contractName = "NFTMarketplace"
const collection_addr = "0x5418E4EBe49fA4F39171f4b18687472FF8726FEa"
const receiver = "0x5418E4EBe49fA4F39171f4b18687472FF8726FEa"
const token_id = 3
const token_url = "google.com"

async function main() {
    provider = new ethers.providers.JsonRpcProvider(process.env.TENDERLY_URL);

    const token = await ethers.getContractAt(contractName, tokenAddress);
    await token.mint(collection_addr, receiver, token_id, token_url);
    console.log("Minted");
}

// Do the thing!
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });