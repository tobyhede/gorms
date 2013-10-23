package models

//import "github.com/nu7hatch/gouuid"
import (
	"github.com/garyburd/redigo/redis"
	"github.com/robfig/revel"
	"github.com/robfig/revel/modules/jobs/app/jobs"
	j "gorp/app/jobs"
	"log"
	"fmt"
)

type Channel struct {
	Key			string
	Messages	[]string
	conn		redis.Conn
}

func NewChannel(conn redis.Conn, key string) *Channel {
	log.Print("NewChannel")
	return &Channel{conn: conn, Key: key}
}

func (c *Channel) Get() {
	log.Print("Get")
	log.Print(c.conn)
	log.Print(c.Key)

	messages, err := redis.Strings(c.conn.Do("LRANGE", c.Key, 0, 10))
	c.Messages = messages

	log.Print(messages)

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}
}

func (c *Channel) Pop() {

	str := `local result = redis.call('LRANGE','%s',0,10); redis.call('LTRIM','%s',1,-10); return result;`

	log.Print(fmt.Sprintf(str, c.Key, c.Key))
	script := redis.NewScript(0, fmt.Sprintf(str, c.Key, c.Key))

	messages, err := redis.Strings(script.Do(c.conn))

	c.Messages = messages

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}
}

func (c Channel) Append(message string) string {
	log.Print("AppendChannel")
	log.Print(c.Key)
	log.Print(message)

	res, err := c.conn.Do("RPUSH", c.Key, message)

	log.Print(res)
	log.Print(err)

	if err != nil {
		log.Fatal(err)
		revel.ERROR.Fatal(err)
	}
	
	jobs.Now(j.CreateMeta{})

	return "OK"
}


func (c Channel) GetMessages() []string {
	return c.Messages
}

func (c Channel) String() string {
	return "Hello"
}
