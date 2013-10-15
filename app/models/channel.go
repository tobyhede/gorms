package models

//import "github.com/nu7hatch/gouuid"
import (
    "github.com/garyburd/redigo/redis"
    "log"
)

type Channel struct {
    Key         string
    Messages    []string
}

func NewChannel() *Channel {
    return &Channel{}
}

func GetChannel(conn redis.Conn, key string) *Channel {
    c := NewChannel()
    log.Print("GetChannel")
    log.Print(key)

    messages, err := redis.Strings(conn.Do("LRANGE", key, 0, 10))
    c.Messages = messages

    if err != nil {
        log.Fatal(err)
    }

    return c
}

func AppendChannel(conn redis.Conn, key, data string) string {
    log.Print("AppendChannel")
    log.Print(key)
    log.Print(data)

    res, err := conn.Do("RPUSH", key, data)

    log.Print(res)
    log.Print(err)

    if err != nil {
        log.Fatal(err)
    }
    return "OK"
}


func (c Channel) GetMessages() []string {
    return c.Messages
}

func (c Channel) String() string {
    return "Hello"
}
