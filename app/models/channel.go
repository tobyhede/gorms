package models

//import "github.com/nu7hatch/gouuid"
import (
	j "gorp/app/jobs"
	"github.com/garyburd/redigo/redis"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/jobs/app/jobs"
    "crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"time"
)


func unique() string {
	bytes := []byte("dg239asdkh6yasckjg")
    hasher := sha256.New()
    hasher.Write(bytes)
    return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

type Channel struct {
	Key			string
	Messages	[]Message
	Msg 		map[string]string
	conn		redis.Conn
}

type Message struct {
	Timestamp 	string
	Message 	string
}

// read 4
// write 2
// execute 1
// read-write 6

func NewChannel(conn redis.Conn, key string) *Channel {
	log.Print("NewChannel")

	log.Print("=====================")
	log.Print(unique())

	// local user = redis.call('HGET', 'key', 'auth');
	// local wild = redis.call('HGET', 'key', '');
	// return user or wild;

	authKey := fmt.Sprintf("auth:%s", key)
	res, err := conn.Do("HSET", authKey, "user", "6")

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}

	res, err = redis.Int(conn.Do("HGET", authKey, "user"))

	log.Print(err == redis.ErrNil)

	if err == redis.ErrNil {

	}

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}

	log.Print(res)

	return &Channel{conn: conn, Key: key}
}

func (c *Channel) Get() {

	//messages, err := redis.Strings(c.conn.Do("LRANGE", c.Key, 0, 10))
	messages, err := redis.Strings(c.conn.Do("ZREVRANGE", c.Key, 0, 10, "WITHSCORES"))

	//m := make(map[string]string)
	log.Print(len(messages)/2)
	m := make([]Message, 0, len(messages)/2)
	log.Print("*****************")
	log.Print("*****************")

	log.Print(m)


	for index,element := range messages {
		if index % 2 == 1 {
			//m[element] = messages[index-1]
			log.Print(element)
			msg := &Message{Timestamp: element, Message: messages[index-1]}
			log.Print(msg)
			m = append(m, *msg)
		}
	}
	
	//messages, err := redis.Values(c.conn.Do("ZREVRANGE", c.Key, 0, 10, "WITHSCORES"))
	log.Print(m)
	c.Messages = m
	//c.Msg = m
	//var m []struct {
        //Message  string
		//Timestamp string
	//}

	//if err := redis.ScanSlice(messages, &m); err != nil {
		//panic(err)
	//}
	//log.Print(m)

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}
}

func (c *Channel) Pop() {

	//str := `local result = redis.call('LRANGE','%s',0,10); redis.call('LTRIM','%s',1,-10); return result;`

	//script := redis.NewScript(0, fmt.Sprintf(str, c.Key, c.Key))

	////messages, err := redis.Strings(script.Do(c.conn))

	////c.Messages = messages

	//if err != nil {
		//log.Fatal(err)
		//revel.ERROR.Fatal(err)
	//}
}

func (c Channel) Append(message string) string {
	log.Print("AppendChannel")
	log.Print(c.Key)
	log.Print(message)

	score := time.Now().UnixNano() / 1e6
	_, err := c.conn.Do("ZADD", c.Key, score, message)
	log.Print("((score))")
	log.Print(score)
	//time := time.Unix(0, score)

	//i, _ := strconv.ParseFloat("1.3835413118343698e+18", 64)
	//log.Print(int64(i))
	//time := time.Unix(0, i)
	//log.Print(time)

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}
	
	jobs.Now(j.CreateMeta{})

	return "OK"
}


func (c Channel) GetMessages() []Message {
	return c.Messages
}

