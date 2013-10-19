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
}

func (c App) Index() revel.Result {
    return c.Render()
}

func init_pool() {
    maxIdle := 5
    idleTimeout :=  240 * time.Second

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
    }
}

func init() {
    log.Print("Init")
    init_pool()
}

//pool = &redis.Pool{
 //MaxIdle: 3,
 //IdleTimeout: 240 * time.Second,
 //Dial: func () (redis.Conn, error) {
     //c, err := redis.Dial("tcp", server)
     //if err != nil {
         //return nil, err
     //}
     //if _, err := c.Do("AUTH", password); err != nil {
         //c.Close()
         //return nil, err
     //}
     //return c, err
 //},
    //TestOnBorrow: func(c redis.Conn, t time.Time) error {
        //_, err := c.Do("PING")
     //return err
    //},
//}


