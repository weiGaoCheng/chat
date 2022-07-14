package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-gin-chat/models"
	"go-gin-chat/routes"
	"go-gin-chat/views"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func init() {
	var workDir string

	if len(os.Args) > 1 {
		workDir, _ = filepath.Abs(filepath.Dir(os.Args[1]))
	} else {
		workDir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}
	fmt.Println(workDir)
	viper.AddConfigPath(workDir + "/conf/.")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("cfg file error: %s\n", err)
		os.Exit(1)
	}
	//viper.SetConfigType("json") // 设置配置文件的类型
	//
	//if err := viper.ReadConfig(bytes.NewBuffer(conf.AppJsonConfig)); err != nil {
	//	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	//		// Config file not found; ignore error if desired
	//		log.Println("no such config file")
	//	} else {
	//		// Config file was found but another error was produced
	//		log.Println("read config error")
	//	}
	//	log.Fatal(err) // 读取配置文件失败致命错误
	//}

	models.InitDB()
}

func main() {
	// 关闭debug模式
	gin.SetMode(gin.ReleaseMode)

	port := viper.GetString(`app.port`)
	router := routes.InitRoute()

	//加载模板文件
	router.SetHTMLTemplate(views.GoTpl)

	log.Println("监听端口", "http://127.0.0.1:"+port)

	http.ListenAndServe(":"+port, router)
}
