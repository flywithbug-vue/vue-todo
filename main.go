package main

import (
	"flag"
	log "github.com/flywithbug/log4go"
	"todo-go/config"
	"todo-go/core/jwt"
	"todo-go/core/mongo"
	"todo-go/server"
)

//log 启动配置
func SetLog() {
	w := log.NewFileWriter()
	w.SetPathPattern("./log/log-%Y%M%D.log")
	c := log.NewConsoleWriter()
	c.SetColor(true)
	log.Register(w)
	log.Register(c)
	log.SetLevel(1)
	log.SetLayout("2006-01-02 15:04:05")
}

func main() {
	//配置文件
	configPath := flag.String("config", "config.json", "Configuration file to use")
	flag.Parse()
	err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatal("读取配置文件错误:", err.Error())
	}
	conf := config.Conf()

	//signingKey read
	jwt.ReadSigningKey(conf.PrivateKeyPath, conf.PublicKeyPath)

	SetLog()
	defer log.Close()

	//mongodb启动连接
	mongo.DialMgo(conf.DBConfig.Url)
	go func() {
		//静态文件服务
		server.StartWeb(conf.WebPort, conf.StaticPath)
	}()
	//启动ApiServer服务
	server.StartApi(conf.ApiPort, conf.RouterPrefix, conf.AuthPrefix)
}
