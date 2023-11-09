// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

interface INFTCollection {
    function  mint(
        address recipient,
        string memory tokenUri,
        uint256 tokenId
    ) external returns (uint256);
}

contract NFTMarketplace is Ownable {
    event CollectionCreated(
        address collection,
        string name,
        string symbol
    );

    event TokenMinted(
        address collection,
        address recipient,
        uint256 tokenId,
        string tokenUri
    );

    struct Collection {
        string name;
        string symbol;
    }

    mapping(address => Collection) private collections;

    constructor(
        address initialOwner
    ) Ownable(initialOwner){}

    function createCollection(
        address nftAddress,
        string memory name,
        string memory symbol
    ) public {
        require(IERC721(nftAddress).supportsInterface(type(IERC721).interfaceId) == true);
        collections[nftAddress] = Collection(name, symbol);

        emit CollectionCreated(nftAddress, name, symbol);
    }

    function mint(
        address collectionAddress,
        address recipient,
        uint256 tokenId,
        string memory tokenUri
    ) public payable {
        INFTCollection nft = INFTCollection(collectionAddress);
        nft.mint(recipient, tokenUri, tokenId);

        emit TokenMinted(collectionAddress, recipient, tokenId, tokenUri);
    }
}
