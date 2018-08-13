package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

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

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/health/ping")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)

	}
	return errors.New("Cannot connect to the router.")
}
