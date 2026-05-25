// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

interface IERC721 {
    function transferFrom(address from, address to, uint tokenId) external;
}

contract EnglishAuction {
    // Events
    event Start(uint startTime, uint endTime);
    event Bid(address indexed bidder, uint value);
    event End(address highestBidder, uint value);
    event Withdraw(address indexed bidder, uint value);

    // Auction state
    bool public started;
    bool public ended;
    uint public endTime;
    uint public highestBid;
    address public highestBidder;
    mapping(address => uint) public allBids;

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call the function");
        _;
    }

    // Constructor variables
    address payable public immutable owner;
    IERC721 public immutable nft;
    uint public immutable nftId;

    constructor(address _nft, uint _nftId) {
        owner = payable(msg.sender);
        nft = IERC721(_nft);
        nftId = _nftId;
    }

    function start(uint _openingBid, uint _duration) external onlyOwner {
        require(!started, "Auction has already started");
        highestBid = _openingBid;
        endTime = block.timestamp + _duration;
        nft.transferFrom(owner, address(this), nftId);
        started = true;
        emit Start(block.timestamp, endTime);
    }

    function bid() external payable {
        require(started, "Auction has not started");
        require(msg.value > highestBid, "Bid price is lower than the current highest bid");
        require(block.timestamp < endTime, "Auction has ended");
        allBids[highestBidder] += highestBid;
        highestBid = msg.value;
        highestBidder = msg.sender;
        emit Bid(msg.sender, msg.value);
    }

    function end() external onlyOwner {
        require(started, "Auction has not started");
        require(block.timestamp >= endTime, "Auction has not ended");
        require(!ended, "Auction has already ended");
        ended = true;
        nft.transferFrom(address(this), highestBidder, nftId);
        owner.transfer(highestBid);
        emit End(highestBidder, highestBid);
    }

    function withdraw() external {
        uint value = allBids[msg.sender];
        allBids[msg.sender] = 0;
        if (value > 0) {
            payable(msg.sender).transfer(value);
        }
        emit Withdraw(msg.sender, value);
    }
}
