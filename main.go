package main

import (
	"encoding/json"
	"fmt"
	//	"net/http"
	_ "net/http/pprof"
	"os"
	"syscall"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"github.com/ninefive/coral/middleware"
	"github.com/ninefive/coral/models"
	"github.com/ninefive/coral/pkg/setting"
	"github.com/ninefive/coral/pkg/version"
	"github.com/ninefive/coral/routers"

	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("config", "c", "", "user center interface config file path.")
	v   = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()
	if *v {
		info := version.GetInfo()
		marshalled, err := json.MarshalIndent(&info, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	//init config
	if err := setting.Init(*cfg); err != nil {
		panic(err)
	}

	//init db
	models.DB.Init()
	defer models.DB.Close()

	//set run mode
	gin.SetMode(viper.GetString("runmode"))

	//create the gin engine
	g := gin.New()

	//routers
	routers.Init(
		g,
		middleware.Logging(),
		middleware.RequestId(),
	)

	endPoint := viper.GetString("addr")
	readTimeout := time.Duration(viper.GetInt("read_timeout"))
	writeTimeout := time.Duration(viper.GetInt("write_timeout"))
	maxHeaderBytes := viper.GetInt("max_header_bytes")

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, g)
	server.BeforeBegin = func(add string) {
		log.Infof("Actual pid is %d", syscall.Getpid())
	}

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(server.ListenAndServe().Error())
}
