const { ethers } = require('ethers');
const prompt = require('prompt-sync')(); 

const privateKey = prompt("Enter the wallet private key:");

const provider = new ethers.providers.JsonRpcProvider("https://eth-sepolia.g.alchemy.com/v2/YOUR-API");

const wallet = new ethers.Wallet(privateKey, provider);

const recipientAddress = prompt(" Enter recipient address: ");
const amountInEther = prompt("Enter the amount of ether:");

async function sendEther() {
    try {
        console.log("Sending transaction...");
        const tx = await wallet.sendTransaction({
            to: recipientAddress,
            value: ethers.utils.parseEther(amountInEther)  
        });

        console.log("Transaction hash:", tx.hash);
        console.log("We are waiting for the confirmation of the transaction...");
        await tx.wait();  
        console.log("The transaction was successfully confirmed!");

    } catch (error) {
        console.error("Error sending transaction:", error);
    }
}

sendEther().catch(console.error);
