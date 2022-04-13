package main

import (
	"ThinkGOChat/gateservice"
	"ThinkGOChat/roomservice"
	"ThinkGOChat/thinkutils/logger"
	"ThinkGOChat/worldservice"
	"fmt"
	"github.com/lonng/nano"
	"github.com/lonng/nano/component"
	"github.com/lonng/nano/serialize/json"
	"github.com/lonng/nano/session"
	"github.com/urfave/cli"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"runtime"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
	mainServices = &component.Components{}
)

//type MainService struct {
//	component.Base
//}

func runMain(args *cli.Context) error {
	log.Info("Run Register Center")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		return nil
	}

	szListen := fmt.Sprintf("127.0.0.1:%d", cfg.Section("register-center").Key("port").MustInt())
	log.Info("Listen for %s", szListen)

	nano.Listen(szListen,
		nano.WithMaster(),
		nano.WithComponents(mainServices),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithDebugMode(),
	)

	return nil
}

func runGate(args *cli.Context) error {
	log.Info("Run Gate Server")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		return nil
	}

	szRegisterCenter := fmt.Sprintf("%s:%d", cfg.Section("register-center").Key("host").String(), cfg.Section("register-center").Key("port").MustInt())
	log.Info("Register Center addr: %s", szRegisterCenter)

	szListen := fmt.Sprintf("127.0.0.1:%d", cfg.Section("gate-server").Key("port").MustInt())
	log.Info("Listen for %s", szListen)

	szGate := fmt.Sprintf("127.0.0.1:%d", cfg.Section("gate-server").Key("gate-port").MustInt())
	log.Info("Websocket addr %s", szGate)

	session.Lifetime.OnClosed(gateservice.OnSessionClosed)

	nano.Listen(szListen,
		nano.WithAdvertiseAddr(szRegisterCenter),
		nano.WithClientAddr(szGate),
		nano.WithComponents(gateservice.Services),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithIsWebsocket(true),
		nano.WithWSPath("/nano"),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
		nano.WithDebugMode(),
	)

	return nil
}

func runWorld(args *cli.Context) error {
	log.Info("Run World Server")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		return nil
	}

	szRegisterCenter := fmt.Sprintf("%s:%d", cfg.Section("register-center").Key("host").String(), cfg.Section("register-center").Key("port").MustInt())
	log.Info("Register Center addr: %s", szRegisterCenter)

	szListen := fmt.Sprintf("127.0.0.1:%d", cfg.Section("world-server").Key("port").MustInt())
	log.Info("Listen for %s", szListen)

	session.Lifetime.OnClosed(worldservice.OnSessionClosed)

	// Startup Nano server with the specified listen address
	nano.Listen(szListen,
		nano.WithAdvertiseAddr(szRegisterCenter),
		nano.WithComponents(worldservice.Services),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithDebugMode(),
	)

	return nil
}

func runRoom(args *cli.Context) error {
	log.Info("Run Room Server")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		return nil
	}

	szRegisterCenter := fmt.Sprintf("%s:%d", cfg.Section("register-center").Key("host").String(), cfg.Section("register-center").Key("port").MustInt())
	log.Info("Register Center addr: %s", szRegisterCenter)

	szListen := fmt.Sprintf("127.0.0.1:%d", cfg.Section("room-server").Key("port").MustInt())
	log.Info("Listen for %s", szListen)

	session.Lifetime.OnClosed(roomservice.OnSessionClosed)

	// Startup Nano server with the specified listen address
	nano.Listen(szListen,
		nano.WithAdvertiseAddr(szRegisterCenter),
		nano.WithComponents(roomservice.Services),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithDebugMode(),
	)

	return nil
}



func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.Info("Hello World")

	app := cli.NewApp()
	app.Name = "ThinkGOChat"
	app.Author = "Thinkman"
	app.Email = "Thinkman Wang"
	app.Description = "China Best Chat Server"
	app.Commands = []cli.Command{
		{
			Name: "main",
			Action: runMain,
		},
		{
			Name: "gate-server",
			Action: runGate,
		},
		{
			Name: "world-server",
			Action: runWorld,
		},
		{
			Name: "room-server",
			Action: runRoom,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Error("Startup server error %+v", err)
	}

	//cfg, err := ini.Load("app.ini")
	//if err != nil {
	//	return
	//}
	
}
