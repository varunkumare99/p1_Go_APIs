Level 1: Simple Handlers

Create /welcome → returns { "msg": "Welcome Varun" }

Create /status → returns { "status": "ok" }

Level 2: Dynamic Responses

Create /echo?msg=hello → should return { "echo": "hello" }

Create /square?num=4 → should return { "result": 16 }

Level 3: JSON Input (POST)

Create /add endpoint (POST)

Accept JSON: { "a": 5, "b": 10 }

Respond: { "sum": 15 }

Level 4: Error handling

Level 5: logging

