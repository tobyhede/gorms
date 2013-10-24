package controllers

import (
	"github.com/robfig/revel"
	"github.com/garyburd/redigo/redis"
	"time"
	"log"
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

func initPool() {
	maxIdle := 5
	idleTimeout := 240 * time.Second

	Pool = &redis.Pool{
		MaxIdle: maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			log.Print("Dial")
			conn, err := redis.Dial("tcp", ":6379")
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
	initPool()
	revel.InterceptMethod((*App).Connect, revel.BEFORE)
	revel.InterceptMethod((*App).Close, revel.FINALLY)
}
