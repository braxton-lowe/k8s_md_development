package main

import (
	"fmt" //to print messages to stdout
	"log" //logging :)
	//our web server that will host the mock
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"os"
	"path/filepath"
)

var configuration []byte
var secret []byte

func Response(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello")
}

func Status(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "ok")
}

func ReadConfig() {
	fmt.Println("reading config...")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("current config cwd", path)

	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files) // contains a list of all files in the current directory

	config, e := ioutil.ReadFile("/configz/config.json")
	if e != nil {
		fmt.Printf("Error reading config file: %v\n", e)
		os.Exit(1)
	}
	configuration = config
	fmt.Println("config loaded!")

}

func ReadSecret() {
	fmt.Println("reading secret...")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("current secret cwd", path)

	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files) // contains a list of all files in the current directory

	s, e := ioutil.ReadFile("/secretz/secret.json")
	if e != nil {
		fmt.Printf("Error reading secret file: %v\n", e)
		os.Exit(1)
	}
	secret = s
	fmt.Println("secret loaded!")

}

func main() {

	fmt.Println("starting...")
	ReadConfig()
	ReadSecret()
	router := fasthttprouter.New()
	router.GET("/", Response)
	router.GET("/status", Status)

	log.Fatal(fasthttp.ListenAndServe(":5000", router.Handler))
}
