#  Game Room Socket

- Support both `ws` and `wss`
  - `ws` on port 8080
  - `wss` on port 8081

## Connect

`host/game/room/{roomId}`

## Request and Response

All requests and responses are json

```json
{
  "topic": "",
  "attributes": {},
  "payload": null
}
```

## [Topics](socket/topics)