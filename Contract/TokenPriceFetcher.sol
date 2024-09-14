// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";

contract TokenPriceFetcher is ChainlinkClient {
    using Chainlink for Chainlink.Request;

    address private oracle;
    bytes32 private jobId;
    uint256 private fee;

    uint256 public price; // قیمت دریافتی

    address public tokenSwapContract; // آدرس قرارداد TokenSwap

    constructor(address _tokenSwapContract) {
        setChainlinkToken(0x779877A7B0D9E8603169DdbD7836e478b4624789);
        oracle = 0x6090149792dAAeE9D1D568c9f9a6F6B46AA29eFD; // آدرس اوراکل در شبکه Sepolia
        jobId = "ca98366cc7314957b8c012c72f05aeeb"; // Job ID برای دریافت uint256
        fee = 0.1 * 10 ** 18; // هزینه به لینک توکن
        tokenSwapContract = _tokenSwapContract; // تنظیم آدرس قرارداد TokenSwap
    }

    function requestPrice(string memory apiUrl, string memory jsonPath, int times) public returns (bytes32 requestId) {
        Chainlink.Request memory request = buildChainlinkRequest(jobId, address(this), this.fulfill.selector);
        request.add("get", apiUrl);
        request.add("path", jsonPath);
        request.addInt("times", times); // ضرب در ضریب برای تبدیل به Wei
        return sendChainlinkRequestTo(oracle, request, fee);
    }

    function fulfill(bytes32 _requestId, uint256 _price) public recordChainlinkFulfillment(_requestId) {
        price = _price;
        // ارسال قیمت به قرارداد TokenSwap
        TokenSwap(tokenSwapContract).updatePrice(_price);
    }
}

interface TokenSwap {
    function updatePrice(uint256 _price) external;
}
