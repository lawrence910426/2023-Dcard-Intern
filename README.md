# 2023-Dcard-Intern
實作了一個共用 Key-Value 的列表系統，功能如下
1. 使用 `gRPC` 新增列表到 `Redis` 儲存庫
2. Restful API 部分採用 `go` 實作
3. 以 `Docker-Compose` 啟動 `Redis` 與 `Go`
4. 為了避免編譯 `proto` 時的環境問題，`Dockerfile` 會自動編譯 `proto`
5. 在 `/api/test` 下執行 `go test` 即可執行 Integration Test

# Get 設計
關於 `GetHead`/`GetPage` 的部分，會從 `url` 的 `query_string` 傳入 `ListKey`/`PageKey`，並從資料庫中找出對應的 `NextPageKey` 與 `Article`。

![](https://i.imgur.com/d83dhYt.png)

![](https://i.imgur.com/Xz8Ad9A.png)

# Set 設計
採用 `gRPC` 來呼叫 api server，Client (Algo team) 端僅需傳入一連串的 article，即可建立列表，並在最後回傳 `GetHead` 用的 Token。

![](https://i.imgur.com/BMpaSvT.png)

# 儲存系統的選擇
考慮到不需要永久儲存資料，使用 `MySQL` 或 `PostgreSQL` 可能有些過於笨重。因此，輕量化的 `Redis` 是一個有效處理高吞吐量任務的選擇。如果命令 `Redis` 不使用 WAL (Write Ahead Logging)，儘管在斷電後無法立即恢復資料，但可以提高處理速度，而且失去的資料只需讓算法團隊重新生成並寫入列表系統即可。

綜上所述，我認為在這個架構下，`Redis` 是一個相對合理的選擇。儘管失去了永久儲存能力，但它對於提高效能有非常明顯的幫助。

# 自動清理
`Redis` 提供 `EXPIRE` 這個指令，用於自動清理不需要的 Key-Value pair，既然每組資料最多只需要儲存一天，那麼讓 `Redis` 在一天後自動清理即可。

![](https://i.imgur.com/TvzEAxN.png)

# Test
利用 `docker-compose.yml` 啟動 `go` 與 `Redis` 後，再執行 `go test`，即可執行 Integration test。

![](https://i.imgur.com/NKIEKF3.png)

![](https://i.imgur.com/g2E0vYp.png)
