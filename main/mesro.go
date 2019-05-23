package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"mesro/sysinfo"
	"net"
	"net/http"
	"os"
)

var (
	Version   = "<version>"
	BuildTime = "<buildTime>"
)

func main() {
	ip := ""
	netaddr, _ := net.InterfaceAddrs()
	networkIp, _ := netaddr[1].(*net.IPNet)
	if !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
		ip = networkIp.IP.String()
	}
	port := flag.String("p", "9999", "HTTP listen port\nCustom Port '-p port'\nExample: ./mesro -p 9999")
	version := flag.Bool("v", false, "Show version and BuildTime")
	info := flag.Bool("a", false, "Show system info\ninclude system info and cpu/memory/net/io ")
	flag.Bool("Mesro Portal", false, "http://"+ip+":9999")
	flag.Bool("Start Collect", false, "parameter 'n' is report name\nparameter 't' is collect duration(minute)\nExample: http://"+ip+":9999/start?n=test&t=30")
	flag.Bool("Stop Collect", false, "Example: http://"+ip+":9999/stop")
	flag.Bool("View Report", false, "Open browser: http://"+ip+":9999/report")
	flag.Bool("Close mesro", false, "Close self\nExample: http://"+ip+":9999/close")
	flag.Parse()

	if *info {
		sysinfo.Info()
		os.Exit(0)
	}

	if *version {
		fmt.Println("Version: " + Version)
		fmt.Println("BuildTime: " + BuildTime)
		os.Exit(0)
	}

	fmt.Println("Mesro Portal : http://" + ip + ":" + *port)
	fmt.Println("Start Collect: http://" + ip + ":" + *port + "/start?n=testname&t=30")
	fmt.Println("Stop Collect: http://" + ip + ":" + *port + "/stop")
	fmt.Println("View Report :  http://" + ip + ":" + *port + "/report")
	fmt.Println("Close mesro:  http://" + ip + ":" + *port + "/close")

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//设定请求url不存在的返回值
	r.NoRoute(NoResponse)

	//index
	r.LoadHTMLGlob("web/index.html")
	r.Static("/js", "web/js")
	r.Static("/fonts", "web/fonts")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

//NoResponse 请求的url不存在，返回404
func NoResponse(c *gin.Context) {
	//返回404状态码
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not exists!",
	})
}
