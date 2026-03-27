# Test four type of gRPC calling method

* EchoMeow -> Unary RPC
* ManyMeow -> Server streaming RPC
* ImpatientMeow -> Client streaming RPC
* Conversation -> Bidirectional streaming RPC

```
2026/03/27 16:40:49 EchoMeow Calling...
2026/03/27 16:40:50 Hello! ^›⩊‹^ ੭
2026/03/27 16:40:50 ManyMeow Calling...
2026/03/27 16:40:50 你 ^•𖥦•^.ᐟ
2026/03/27 16:40:50 好 ฅ(´꒳ `ฅ)ꪆ
2026/03/27 16:40:50 嗎 ^•𖥦•^.ᐟ
2026/03/27 16:40:50 ? ฅ(´꒳ `ฅ)ꪆ
2026/03/27 16:40:50 ImpatientMeow Calling...
2026/03/27 16:40:50 I LOVE CAT ฅ(´꒳ `ฅ)ꪆ
2026/03/27 16:40:50 Conversation Calling...
2026/03/27 16:40:50 I ^›⩊‹^ ੭
2026/03/27 16:40:50 LOVE (＾• ω •＾)
2026/03/27 16:40:50 CAT ^›⩊‹^ ੭
```