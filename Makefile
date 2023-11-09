up:
	docker-compose up -d --force-recreate

sc_compile:
	cd contract && npx hardhat compile

sc_deploy:
	cd contract && npx hardhat run scripts/deploy.js --network sepolia

sc_deploy_collection:
	cd contract && npx hardhat run scripts/deploy_collection.js --network sepolia

sc_estimate:
	cd contract && npx hardhat run scripts/estimate.js --network main

sc_mint:
	cd contract && npx hardhat run scripts/mint.js --network sepolia

sc_mint_2:
	cd contract && npx hardhat run scripts/mint_collection.js --network sepolia