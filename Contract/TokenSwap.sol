// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract TokenSwap is ChainlinkClient {
    using Chainlink for Chainlink.Request;

    address public owner;
    uint256 public swapFee;
    uint256 public price;

    address private oracle;
    bytes32 private jobId;
    uint256 private fee;

    event SwapExecuted(address indexed fromToken, address indexed toToken, uint256 amountIn, uint256 amountOut);
    event PriceUpdated(uint256 price);

    constructor(address _linkToken) {
        owner = msg.sender;
        swapFee = 10;
        
        setChainlinkToken(0x779877A7B0D9E8603169DdbD7836e478b4624789);
        oracle = 0x6090149792dAAeE9D1D568c9f9a6F6B46AA29eFD;
        jobId = "ca98366cc7314957b8c012c72f05aeeb"; 
        fee = (1 * LINK_DIVISIBILITY) / 10; 
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can execute this");
        _;
    }

    function requestPrice(string memory apiUrl, string memory jsonPath, int times) public onlyOwner returns (bytes32 requestId) {
        Chainlink.Request memory request = buildChainlinkRequest(jobId, address(this), this.fulfill.selector);
        request.add("get", apiUrl);
        request.add("path", jsonPath);
        request.addInt("times", times);
        return sendChainlinkRequestTo(oracle, request, fee);
    }

    function fulfill(bytes32 _requestId, uint256 _price) public recordChainlinkFulfillment(_requestId) {
        price = _price;
        emit PriceUpdated(price);
    }

    function swapTokens(address tokenA, address tokenB, uint256 amountA) public {
        require(IERC20(tokenA).transferFrom(msg.sender, address(this), amountA), "Transfer failed");

        uint256 amountB = getSwapAmount(amountA);
        require(IERC20(tokenB).balanceOf(address(this)) >= amountB, "Insufficient liquidity");

        require(IERC20(tokenB).transfer(msg.sender, amountB), "Transfer failed");
        emit SwapExecuted(tokenA, tokenB, amountA, amountB);
    }

    function getSwapAmount(uint256 amountA) internal view returns (uint256) {
        return (amountA * price) / (1 ether) - (amountA * swapFee / 1000);
    }

    function addLiquidity(address token, uint256 amount) public onlyOwner {
        require(IERC20(token).transferFrom(msg.sender, address(this), amount), "Transfer failed");
    }

    function setSwapFee(uint256 _swapFee) public onlyOwner {
        require(_swapFee <= 100, "Swap fee too high");
        swapFee = _swapFee;
    }
}
