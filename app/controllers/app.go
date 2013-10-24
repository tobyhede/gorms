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

func init() {
	//log.Print("Init")
	//initPool()
	revel.OnAppStart(Init)
	revel.InterceptMethod((*App).Connect, revel.BEFORE)
	revel.InterceptMethod((*App).Close, revel.FINALLY)
}
