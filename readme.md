透過 Project-Based Learning 方式，使用 Golang 嘗試實作區塊鏈，學習 Golang 語言與區塊鏈相關知識。 

Learning Golang and blockchain concepts through project-based learning.

## 參考資源
- https://github.com/liuchengxu/blockchain-tutorial
- https://jeiwan.net/

## part 1 基本原型
- 建立Block
- 建立BlockChain

## part 2 工作證明
- 建立ProofOfWork
- 計算 hash,尋找符合目標條件的nonce和hash

## part 3 持久化和CLI
- 儲存區塊到database
  - 將block序列化,解序列化
  - key => value
  - 32 bytes block hash  => block
  - l => hash of last block 
- 建立CLI
  - addblock
  - printchain