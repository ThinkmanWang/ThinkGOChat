package main

import (
	"ThinkGOChat/thinkutils/logger"
	"fmt"
	"github.com/lonng/nano"
	"github.com/lonng/nano/examples/cluster/chat"
	"github.com/lonng/nano/examples/cluster/gate"
	"github.com/lonng/nano/examples/cluster/master"
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
)

func runRegisterCenter(args *cli.Context) error {
	log.Info("Run Register Center")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		return nil
	}

	szListen := fmt.Sprintf("127.0.0.1:%d", cfg.Section("register-center").Key("port").MustInt())
	log.Info("Listen for %s", szListen)

	nano.Listen(szListen,
		nano.WithMaster(),
		nano.WithComponents(master.Services),
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

	nano.Listen(szListen,
		nano.WithAdvertiseAddr(szRegisterCenter),
		nano.WithClientAddr(szGate),
		nano.WithComponents(gate.Services),
		nano.WithSerializer(json.NewSerializer()),
		nano.WithIsWebsocket(true),
		nano.WithWSPath("/nano"),
		nano.WithCheckOriginFunc(func(_ *http.Request) bool { return true }),
		nano.WithDebugMode(),
	)

	return nil
}

func runWorld(args *cli.Context) error {

	return nil
}

func runChat(args *cli.Context) error {
	log.Info("Run Chat Server")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		return nil
	}

	szRegisterCenter := fmt.Sprintf("%s:%d", cfg.Section("register-center").Key("host").String(), cfg.Section("register-center").Key("port").MustInt())
	log.Info("Register Center addr: %s", szRegisterCenter)

	szListen := fmt.Sprintf("127.0.0.1:%d", cfg.Section("room-server").Key("port").MustInt())
	log.Info("Listen for %s", szListen)

	session.Lifetime.OnClosed(chat.OnSessionClosed)

	// Startup Nano server with the specified listen address
	nano.Listen(szListen,
		nano.WithAdvertiseAddr(szRegisterCenter),
		nano.WithComponents(chat.Services),
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
			Name: "register-center",
			Action: runRegisterCenter,
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
			Action: runChat,
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
