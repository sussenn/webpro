package tool

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

//Session初始化
func InitSession(engine *gin.Engine) {
	config := GetConfig().RedisConfig
	//?
	store, err := redis.NewStore(10, "tcp", config.Addr+":"+config.Port, "", []byte("secret"))
	if err != nil {
		log.Println("InitSession() err: ", err)
	}
	engine.Use(sessions.Sessions("mySession", store))
}

//Set
func SetSession(ctx *gin.Context, key interface{}, value interface{}) error {
	session := sessions.Default(ctx)
	if session == nil {
		return nil
	}
	session.Set(key, value)
	return session.Save()
}

//Get
func GetSession(ctx *gin.Context, key interface{}) interface{} {
	session := sessions.Default(ctx)
	return session.Get(key)
}
