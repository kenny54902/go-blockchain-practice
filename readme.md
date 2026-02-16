透過 Project-Based Learning 方式，使用 Golang 嘗試實作區塊鏈，學習 Golang 語言與區塊鏈相關知識。 

Learning Golang and blockchain concepts through project-based learning.

## 參考資源
- https://github.com/liuchengxu/blockchain-tutorial
- https://jeiwan.net/

## part 1 基本原型
- 區塊鏈是一個公開的分佈式數據庫
- 區塊包含區塊的hash和前一個區塊的hash,以及儲存的資料
- 比特幣中區塊由區塊頭和交易清單組成
- 區塊鏈透過前一區塊的hash將區塊串接成鏈狀結構
- 區塊鏈中必須至少有一個區塊,其中第一個區塊稱為創世塊

## part 2 工作證明
- 要向區塊鏈中加入區塊必須完成工作證明
  - 在比特幣中工作證明為找到一個區塊的hash, 該hash必須滿足一些必要條件
  - 完成工作證明的人會獲得獎勵 (挖礦獲得幣)
- Hashcash 
  - 取一些公開資料加上一個counter(nonce)後取hash,直到滿足特定條件

## part 3 持久化和CLI
- LMDB (Lightning Memory Mapped Database)
  - key-value 嵌入式数据库
  - 內存映射,讀取快速
  - 完整ACID事務支持
- LevelDB
 - key-value 嵌入式数据库
 - 寫入效能高
 - 支援原子更新
- RDBMS
 - 區塊鏈不需要複雜查詢
 - key-valuey足以應付且效能較好
 
## part 4 交易1
- UTXO模型 (Unspent Transaction Output)
 - UTXO => 未花費的交易輸出
 - transaction 由 input 和 output構成
 - input 指向之前某筆交易產生的輸出(可以花的錢)
 - output 包含金額和鎖定腳本(如收款人的公鑰HASH), 
 - output 透過鎖定腳本判斷誰有權引用其作為input (花錢) 
 - 未被消耗的output => UTXO
 - 總餘額 = 擁有者有權解鎖且尚未被花費的 UTXO 總和。
 - UTXO不可分割, 發生找零時產出新的output

## part 5 地址
- 一組公私鑰的公鑰經過運算轉換為字符串即為地址
- 持有私鑰即擁有對應公鑰的控制權(幣的所有權)
- 數位簽章
  - 驗證身份
  - 確保資料未被篡改
  - 不可否認性
  - 私鑰簽章, 公鑰驗證
  - 簽署交易的摘要(涵蓋發送者,接收者,金額)
  - 簽名放在交易的input上
    - 簽署者確認交易
    - 作為解鎖引用的輸的key, (有權使用引用的輸出)
