// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// TODO ERC721Enumerable might be more fit
contract NFTCollection is ERC721, ERC721URIStorage, Ownable {
    constructor(
        address initialOwner,
        string memory name,
        string memory symbol
    ) Ownable(initialOwner) ERC721(name, symbol) {}

    function tokenURI(
        uint256 tokenId
    )
        public
        view
        virtual
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function mint(
        address recipient,
        string memory tokenUri,
        uint256 tokenId
    ) public returns (uint256) {
        require(recipient != address(0), "ERC721: mint to the zero address");

        _mint(recipient, tokenId);
        _setTokenURI(tokenId, tokenUri);

        return tokenId;
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721, ERC721URIStorage) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
