const { ethers } = require('ethers');
const prompt = require('prompt-sync')();  

const privateKey = prompt("Enter the wallet private key: ");

const provider = new ethers.providers.JsonRpcProvider("https://data-seed-prebsc-1-s1.binance.org:8545/");

const wallet = new ethers.Wallet(privateKey, provider);


const recipientAddress = prompt(" Enter recipient address: ");
const amountInBNB = prompt("Enter the amount of BNB: ");

async function sendBNB() {
    try {
        console.log("Sending transaction...");  
        
        const tx = await wallet.sendTransaction({
            to: recipientAddress,  
            value: ethers.utils.parseEther(amountInBNB)  
        });

        console.log("Transaction hash:", tx.hash);  
        console.log("The transaction was successfully confirmed!");  
        
        await tx.wait(); 
        console.log("The transaction was successfully confirmed!"); 

    } catch (error) {
        console.error("Error sending transaction:", error);
    }
}

sendBNB().catch(console.error);
