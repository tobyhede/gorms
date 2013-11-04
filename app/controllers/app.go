package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	Pool *redis.Pool
)

type App struct {
	*revel.Controller
	Conn redis.Conn
}

func (c App) Index() revel.Result {
	return c.Render()
}


func getConfig(config string) int {
	value, found := revel.Config.Int(config)
	
	if !found {
		revel.ERROR.Fatal(fmt.Sprintf("Configuration [%s] not found", config))
	}

	return value
}

func Init() {
	log.Print("==== INIT ====")

	maxIdle			:= getConfig("redis.pool.maxidle")
	idleTimeout 	:= time.Duration(getConfig("redis.pool.timeout")) * time.Second
	maxActive 		:= getConfig("redis.pool.maxactive")
	port 			:= fmt.Sprintf(":%d", getConfig("redis.port"))
	
	Pool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,

		Dial: func() (redis.Conn, error) {
			log.Print("Dial")
			conn, err := redis.Dial("tcp", port)
			if err != nil {
				log.Fatal(err)
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (c *App) Connect() revel.Result {
	c.Conn = Pool.Get()
	return nil
}

func (c *App) Close() revel.Result {
	c.Conn.Close()
	return nil
}

func (c *App) ValidateId() revel.Result {
	//if id := c.Params.Get("id"); id != "" {
		//c.Validation.Required(!strings.HasPrefix(id, "auth:"))

		//if c.Validation.HasErrors() {
			//log.Fatal("Invalid Key")
			//revel.ERROR.Fatal("Invalid Key")
		//}
	//}
	return nil
}

func (c *App) Validate() revel.Result {

	//auth := c.Request.Header["Authorization"]
	
	//log.Print(auth)

	//str := `local result = redis.call('LRANGE','%s',0,10); redis.call('LTRIM','%s',1,-10); return result;`
	//script := redis.NewScript(0, fmt.Sprintf(str, c.Key, c.Key))

	//messages, err := redis.Strings(script.Do(c.conn))

	//c.Messages = messages

	//if err != nil {
		//log.Fatal(err)
		//revel.ERROR.Fatal(err)
	//}


	return nil
}


func init() {
	//log.Print("Init")
	//initPool()
	revel.OnAppStart(Init)
	revel.InterceptMethod((*App).ValidateId, revel.BEFORE)
	revel.InterceptMethod((*App).Connect, revel.BEFORE)
	revel.InterceptMethod((*App).Validate, revel.BEFORE)
	revel.InterceptMethod((*App).Close, revel.FINALLY)
}
