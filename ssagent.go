package main

import (
	"bytes"
	"flag"
	"fmt"
	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func parse(htmlpage []byte) [][]string {
	path := xmlpath.MustCompile("//*[@class='portfolio-item']")
	hostPath := xmlpath.MustCompile("div/div/h4[1]/span[1]")
	portPath := xmlpath.MustCompile("div/div/h4[2]")
	pwPath := xmlpath.MustCompile("div/div/h4[3]/span[1]")
	mdPath := xmlpath.MustCompile("div/div/h4[4]")

	root, err := xmlpath.ParseHTML(bytes.NewReader(htmlpage))
	if err != nil {
		log.Fatal(err)
	}
	it := path.Iter(root)
	serverConfig := [][]string{}
	for it.Next() {
		node := it.Node()
		//fmt.Println("Found:", node)
		var host, port, passwd, method string
		var ok bool
		if host, ok = hostPath.String(node); ok {
			fmt.Println("HOST:", host)
		}
		if port, ok = portPath.String(node); ok {
			port = port[7:]
			fmt.Println("PORT:", port)
		}
		if passwd, ok = pwPath.String(node); ok {
			fmt.Println("PASSWORD:", passwd)
		}
		if method, ok = mdPath.String(node); ok {
			method = method[7:]
			fmt.Println("METHOD:", method)
		}
		item := []string{host + ":" + port, passwd, method}
		serverConfig = append(serverConfig, item)

	}
	return serverConfig
}

func getSSConfig(ssconfig *ss.Config) {
	//os.Setenv("HTTP_PROXY", "http://proxy.houston.hpecorp.net:8080")
	res, err := http.Get("http://fast.ishadow.online/")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s", robots)
	config := parse(robots)
	log.Printf("%v", config)
	ssconfig.ServerPassword = config
}

func main() {
	log.SetOutput(os.Stdout)

	var cmdLocal string
	var cmdConfig ss.Config
	//var printVer bool

	flag.StringVar(&cmdLocal, "b", "", "local address, listen only to this address if specified")
	flag.IntVar(&cmdConfig.Timeout, "t", 300, "timeout in seconds")
	flag.IntVar(&cmdConfig.LocalPort, "l", 1080, "local socks5 proxy port")
	flag.BoolVar((*bool)(&debug), "d", false, "print debug message")
	flag.BoolVar(&cmdConfig.Auth, "A", false, "one time auth")

	flag.Parse()

	getSSConfig(&cmdConfig)

	parseServerConfig(&cmdConfig)

	run(cmdLocal + ":" + strconv.Itoa(cmdConfig.LocalPort))
}
