# arweave
## arweave is created to combine all in-game assets and store data about NFTs for [Mintbase](https://www.mintbase.io/)
### [Arweave](https://www.arweave.org/) is a new type of storage that backs data with sustainable and perpetual endowments, allowing users and developers to truly store data forever – for the very first time.
## It is deployed to Azure Function App. So it is serverless and sclable.
### There is a Makefile rule to deploy everything to Azure
## An [example](https://arweave.net/0WlrNa__1FatCyZuaNB7rNXAQxAKiLzXZgj3P9rv7vI) of the data stored in arweave is:
```json
{
   "media":"https://arweave.net/We5kpKyMpIRYXYOMvnlSfOAirzBWVXyuS1kBkoKUxkI",
   "media_hash":"We5kpKyMpIRYXYOMvnlSfOAirzBWVXyuS1kBkoKUxkI",
   "tags":[
      
   ],
   "title":"Raccoon🦝 #1318",
   "description":"This is a Coaty World Raccoon. It brings a lot of tokens 🪙",
   "extra":[
      
   ],
   "store":"coatyworld.mintspace2.testnet",
   "type":"NEP171",
   "category":null
}
```
## Transaction id is passed to the [Mintbase](https://www.mintbase.io/) smart contract. [Here is an example](https://explorer.testnet.near.org/transactions/GFHq8FBLJD97UBpPP2GnxjxXHmAAuYoTTGnNwwJosDdW)

