G.O.R.M.S
=========

Go Redis Messaging System.

High performance messaging using Revel and Redis.


How-to
---------
G.O.R.M.S. allows messages to be published to a **Channel** identified by an **ID**.

Messages can be any valid JSON-friendly string. 

### GET
#### GET channels/:id

Returns a JSON array of the last 10 messages for the **ID** Channel.

```
  /channels/vtha/
  [
    "blah",
    "vtha"
  ]
```

### Append
#### POST channels/:id
#### GET channels/:id/append

Appends a message to the **ID** Channel.

Expects a paramaeter called *message*.

Returns OK on success.

GET/append isn't by-the-book REST, but it is certainly easier to test.


```
    /channels/vtha/append?message={"time":"%202013-10-19T21:18:45",%20"message":"blah"}
    "OK"
  
    curl -X POST --data "message=hello" http://gorms.io/channels/vtha/
    "OK"
```

### Pop
#### GET channels/:id/pop

Returns a JSON array of the last 10 messages for the **ID** Channel.
The messages are removed permanently from the channel.

The use of GET here is wacky.

```
    /channels/vtha/pop
    [
      "blah",
      "vtha"
    ]
```


