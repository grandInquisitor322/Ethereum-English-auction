# Ethereum English Auction

A fully functional English auction smart contract built with Solidity and Foundry, with a Go REST API for interacting with the blockchain.

## Tech Stack
- **Solidity** — Smart contract development
- **Foundry** — Contract compilation, testing, and deployment
- **OpenZeppelin** — ERC721 NFT standard
- **Go** — REST API backend
- **Gin** — Go web framework
- **go-ethereum (Geth)** — Go Ethereum client and bindings

## How It Works
The auction allows an owner to put an NFT up for auction. Bidders can place bids in ETH, and when the auction ends, the highest bidder receives the NFT and the owner receives the ETH. Losing bidders can withdraw their funds.

## Smart Contract Features
- Start auction with opening bid and duration
- Place bids in ETH
- End auction — transfers NFT to winner, ETH to owner
- Withdraw funds for non-winning bidders
- onlyOwner access control
- ERC721 NFT integration via IERC721 interface

## API Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /auction/start | Start the auction |
| POST | /auction/end | End the auction |
| POST | /auction/withdraw | Withdraw losing bids |
| POST | /bids | Place a bid |

## Running Locally
1. Start local testnet: `anvil`
2. Deploy NFT: `forge create src/MyNFT.sol:MyNFT --rpc-url http://localhost:8545 --unlocked --from <account> --broadcast --constructor-args <account>`
3. Deploy auction: `forge create src/EnglishAuction.sol:EnglishAuction --rpc-url http://localhost:8545 --unlocked --from <account> --broadcast --constructor-args <nft_address> 1`
4. Start API: `cd api && go run .`
