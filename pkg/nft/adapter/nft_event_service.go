package adapter

import (
	"artemb/nft/pkg/config"
	repo "artemb/nft/pkg/db/repo"
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type NFTEventService struct {
	repo *repo.CollectionsRepository
	cfg  *config.Config
}

func NewNFTEventService(r *repo.CollectionsRepository, cfg *config.Config) NFTEventService {
	return NFTEventService{repo: r, cfg: cfg}
}

func (s *NFTEventService) RunService() error {
	client, err := ethclient.Dial(s.cfg.NFT.Provider)
	if err != nil {
		return err
	}

	contractAddress := common.HexToAddress(s.cfg.NFT.ContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return err
	}
	abiToken := `[ { "inputs": [ { "internalType": "address", "name": "initialOwner", "type": "address" } ], "stateMutability": "nonpayable", "type": "constructor" }, { "inputs": [ { "internalType": "address", "name": "owner", "type": "address" } ], "name": "OwnableInvalidOwner", "type": "error" }, { "inputs": [ { "internalType": "address", "name": "account", "type": "address" } ], "name": "OwnableUnauthorizedAccount", "type": "error" }, { "anonymous": false, "inputs": [ { "indexed": false, "internalType": "address", "name": "collection", "type": "address" }, { "indexed": false, "internalType": "string", "name": "name", "type": "string" }, { "indexed": false, "internalType": "string", "name": "symbol", "type": "string" } ], "name": "CollectionCreated", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": true, "internalType": "address", "name": "previousOwner", "type": "address" }, { "indexed": true, "internalType": "address", "name": "newOwner", "type": "address" } ], "name": "OwnershipTransferred", "type": "event" }, { "anonymous": false, "inputs": [ { "indexed": false, "internalType": "address", "name": "collection", "type": "address" }, { "indexed": false, "internalType": "address", "name": "recipient", "type": "address" }, { "indexed": false, "internalType": "uint256", "name": "tokenId", "type": "uint256" }, { "indexed": false, "internalType": "string", "name": "tokenUri", "type": "string" } ], "name": "TokenMinted", "type": "event" }, { "inputs": [ { "internalType": "address", "name": "nftAddress", "type": "address" }, { "internalType": "string", "name": "name", "type": "string" }, { "internalType": "string", "name": "symbol", "type": "string" } ], "name": "createCollection", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [ { "internalType": "address", "name": "collectionAddress", "type": "address" }, { "internalType": "address", "name": "recipient", "type": "address" }, { "internalType": "uint256", "name": "tokenId", "type": "uint256" }, { "internalType": "string", "name": "tokenUri", "type": "string" } ], "name": "mint", "outputs": [], "stateMutability": "payable", "type": "function" }, { "inputs": [], "name": "owner", "outputs": [ { "internalType": "address", "name": "", "type": "address" } ], "stateMutability": "view", "type": "function" }, { "inputs": [], "name": "renounceOwnership", "outputs": [], "stateMutability": "nonpayable", "type": "function" }, { "inputs": [ { "internalType": "address", "name": "newOwner", "type": "address" } ], "name": "transferOwnership", "outputs": [], "stateMutability": "nonpayable", "type": "function" } ]`
	tokenAbi, err := abi.JSON(strings.NewReader(abiToken))

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err) // will fail fully
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log

			unpacked, err := tokenAbi.Unpack("CollectionCreated", vLog.Data)
			if err == nil {
				errDB := s.repo.AddCollection(context.Background(),
					fmt.Sprintf("%v", unpacked[0]),
					fmt.Sprintf("%v", unpacked[1]),
					fmt.Sprintf("%v", unpacked[2]))
				if errDB != nil {
					fmt.Println("error saving collection DB: " + errDB.Error())
				}
				continue
			}

			unpacked, err = tokenAbi.Unpack("TokenMinted", vLog.Data)
			if err == nil {
				errDB := s.repo.AddMinted(context.Background(),
					fmt.Sprintf("%v", unpacked[0]),
					fmt.Sprintf("%v", unpacked[1]),
					fmt.Sprintf("%v", unpacked[2]),
					fmt.Sprintf("%v", unpacked[3]))
				if errDB != nil {
					fmt.Println("error saving collection DB: " + errDB.Error())
				}
				continue
			}
		}
	}
}
