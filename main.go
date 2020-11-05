package main

import (
	"flag"
	"fmt"
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/db"
	"go-gin-boilerplate/middlewares"
	"go-gin-boilerplate/server"
	"os"
	"os/signal"
	"syscall"
)

func postProcessFuncRegister() {
	// do something before http server down, deal with signals such as Ctrl-C, SIGKILL
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGHUP:
				fmt.Println("hungup signal received")
			case syscall.SIGINT:
				fmt.Println("Ctrl-C signal received")
				recycleResources()
				exit_chan <- 0
			case syscall.SIGTERM:
				fmt.Println("sigterm signal received")
				recycleResources()
				exit_chan <- 0
			}
		}
	}()

	go func() {
		code := <-exit_chan
		os.Exit(code)
	}()
}

func preProcessFuncRegister() {
	db.Init()

	middlewares.SchedulerRegister()

	middlewares.RedisInit()
}

func recycleResources() {
	db.CloseDb()
}

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	preProcessFuncRegister()
	postProcessFuncRegister()
	server.Init()
}
