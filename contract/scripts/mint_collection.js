const { ethers } = require("hardhat");
require("dotenv").config();

const tokenAddress = "0xA623dc0B40A58B8827ca9F9Fb3E56A50CC3097aF"
const contractName = "NFTCollection"
const receiver = "0xA623dc0B40A58B8827ca9F9Fb3E56A50CC3097aF"
const token_id = 3
const token_url = "google.com"

async function main() {
    provider = new ethers.providers.JsonRpcProvider(process.env.TENDERLY_URL);

    const token = await ethers.getContractAt(contractName, tokenAddress);
    await token.mint(receiver, token_url, token_id);
    console.log("Minted");
}

// Do the thing!
main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });