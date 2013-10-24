package main

import (
    "github.com/garyburd/redigo/redis"
    "regexp"    
    "strings"
    "time"
)

var pool = &redis.Pool{  //http://godoc.org/github.com/garyburd/redigo/redis#Pool
             MaxIdle: 3,
             IdleTimeout: 240 * time.Second,
             Dial: func () (redis.Conn, error) {
                 c, err := redis.Dial("tcp", ":6379")
                 if err != nil {
                     return nil, err
                 }
                 return c, err
             },
    				TestOnBorrow: func(c redis.Conn, t time.Time) error {
    				    _, err := c.Do("PING")
                     return err
    			    },
         }

type DataAccess struct {
  conn redis.Conn
}

func NewDataAccess() *DataAccess {
  dao := new(DataAccess)
  dao.conn = pool.Get()//redis.Dial("tcp", ":6379")
  return dao
}

func (self *DataAccess) escape(name string) (string){
  wordEx, _ := regexp.Compile(`\W`)
  dashEx, _ := regexp.Compile(`-+`)

  safe := wordEx.ReplaceAllString(name, "-")
  return dashEx.ReplaceAllString(strings.ToLower(safe), "-")
}

func (self *DataAccess) CardsForSet(set string) (string, error){
  name := self.escape(set)
  var err error
  t, err := self.conn.Do("SMEMBERS", "set-cards-" + name)
  if err != nil{
    return "[]", err
  }
    
  guids, _ := redis.Strings(t, nil)
  
  keys := make([]interface{}, len(guids))
  for i := 0; i < len(guids); i++ {
    keys[i] = "card-" + guids[i]
  }
  
  t2, err := self.conn.Do("MGET", keys...)  
  if err != nil{
    return "[]", err
  }  
  items, _ := redis.Strings(t2, nil)

  return "[" + strings.Join(items, ",") + "]", nil
}

func (self *DataAccess) Close(){
  self.conn.Close()
}