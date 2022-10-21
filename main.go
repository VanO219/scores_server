package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/takama/daemon"
)

const (
	// name of the service
	name        = "scoresd_server"
	description = "scoresd_server service"
)

var (
	stdlog, errlog *log.Logger
	doneContext    context.Context
	doneFunc       context.CancelFunc
)

func init() {
	stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (usage string, err error) {
	usage = fmt.Sprintf("Usage: %s install | remove | start | stop | status", name)

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	doneContext, doneFunc = context.WithCancel(context.Background())
	//errs = manage()
	//if errs != nil {
	//	errs = errors.Wrap(errs, "ошибка формирования данных для работы сервера")
	//	errlog.Fatalln(errs)
	//}

	//authd := app.NewAuthD(doneContext)
	//go func() {
	//	err = authd.Start()
	//	if err != nil {
	//		errlog.Fatalln(err)
	//	}
	//}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	allDone := func() {
		doneFunc()
		//err = authd.Stop()
		//if err != nil {
		//	errlog.Println(err)
		//}
	}

	for {
		select {
		case killSignal := <-interrupt:
			stdlog.Println("Got signal:", killSignal)
			allDone()
			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
}

func main() {
	//fmt.Println("GOMAXPROCS \t", os.Getenv("GOMAXPROCS"))
	//fmt.Println("runtime.GOMAXPROCS \t", runtime.GOMAXPROCS(6))
	//fmt.Println("runtime.NumCPU \t", runtime.NumCPU())
	srv, err := daemon.New(name, description, daemon.SystemDaemon)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}

	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
}
