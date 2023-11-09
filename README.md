# Test APP

THA for Solidity exp engineer

Develop a smart contract(-s) on Solidity for deploying a NFT collection (ERC721) with some arguments (name, symbol). The smart contract should emit the following events:

CollectionCreated(address collection, name, symbol)

TokenMinted(address collection, address recipient, tokenId, tokenUri)

Develop a simple backend server with in-memory storage to handle emitted events and serve it via HTTP.

Develop a front end demo application that interacts with the smart contract and has the following functionality:

Create a new NFT collection with specified name and symbol (from user input);

Mints a new NFT with specified collection address (only created on 3.a), tokenId, tokenUri.

## Project Structure

Video: https://youtu.be/hKWqgVXSd-8

`/contract` - include smart-contract codes with hardhat deployment

`/pkg` - golang code for backend and listening emited events

`/web` - html/js simple UI

Deployed contract on sepolia is https://sepolia.etherscan.io/address/0xcf23b9C1a4B7351D965195c32B84c4f9412a3248, includes minting and create collections methods.

For `hardhat` - prepare `.env` file with PRIVATE_KEY, TENDERLY_URL and TENDERLY_ACCESS_KEY

Project includes 2 smart contracts: 

`NFTCollection.sol` - is actual NFT

`NFTMarketplace.sol` - some kind of proxy contract, which emit events for create_collection and mint methods.

## Run project

`make up` - will run UI (backend removed, hence wss required signed contract with valid certs)

`go mod vendor`

`go build -o api -v ./cmd/api`

`chmod +x api`

`./api server -c config.yaml` - will run API server locally (reason why it's not in Docker above)

You should see message with `Starting server on port:8080`

UI available on port: `http://localhost:8000/` and API: `localhost:8080`

API Endpoints:
`localhost:8080/collections` will return all added collections, i.e.

```json
{
    "collections": [
        {
        "ID": 1,
        "Address": "0x010C6F93B7625788C90c4aE50ad5A5E4727a907C",
        "Symbol": "APES1",
        "Name": "MY_COLLECTION4",
        "Updated": 1699539219276,
        "CreatedAt": 1699539219
        }
    ]
}
```

`localhost:8080/minted`

```json
{
    "minted": [
        {
            "ID": 1,
            "Collection": "0x3Abd554B1b32fC219965F739Cd227f782FF443b0",
            "Recipient": "0x3Abd554B1b32fC219965F739Cd227f782FF443b0",
            "TokenID": "2",
            "TokenURL": "https://google.com",
            "Updated": 1699540515810,
            "CreatedAt": 1699540515
        }
    ]
}
```

## TODO

There is no unit-test/integration tests for API layer.

Web UI is too simple, just for testing purpose.