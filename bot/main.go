package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	glogger "cloud.google.com/go/logging"
	"github.com/bwmarrin/discordgo"
	scm "github.com/ethanent/discordgo-scm"
	"github.com/joho/godotenv"
	api "hushclan.com/api/database"
	"hushclan.com/api/logging"
	"hushclan.com/app"
)

func main() {
	app := &app.App{}
	var err error

	// godotenv

	if err := godotenv.Load(); err != nil {
		log.Fatal("could not load env file: ", err)
	}

	// loading environment variables

	if app.Envs.Token = os.Getenv("BOT_TOKEN"); app.Envs.Token == "" {
		log.Fatal("bot token is not set")
	}
	if app.Envs.Guild = os.Getenv("GUILD_ID"); app.Envs.Guild == "" {
		log.Fatal("guild is not set")
	}
	if app.Envs.ProjectID = os.Getenv("PROJECT_ID"); app.Envs.ProjectID == "" {
		log.Fatal("project id is not set")
	}
	if app.Envs.LogName = os.Getenv("LOG_NAME"); app.Envs.LogName == "" {
		log.Fatal("project id is not set")
	}

	// creating logger

	client, err := glogger.NewClient(context.Background(), app.Envs.ProjectID)
	if err != nil {
		log.Fatal("could not connect to logging: ", err)
	}
	defer client.Close()

	app.Log = &logging.Log{
		Logger: client.Logger(app.Envs.LogName),
	}

	// connecting to database

	app.Database = &api.Database{}
	if err := app.Database.Connect(); err != nil {
		app.Log.Critical("could not connect to database", err)
		log.Fatal("could not connect to database: ", err)
	}

	// creating Discord session

	app.Session, err = discordgo.New("Bot " + app.Envs.Token)
	if err != nil {
		app.Log.Critical("could not start discord session", err)
		log.Fatal("could not start discord session: ", err)
	}

	// creating slash commands manager

	app.Manager = scm.NewSCM()

	// creating and populating features

	app.PopulateSCM()

	// handling on join and registering commands

	app.Session.AddHandler(app.HandleGuildJoin)
	app.Session.AddHandler(app.HandleGuildLeave)

	// start the Discord session

	err = app.Session.Open()
	if err != nil {
		app.Log.Critical("could not open connection", err)
		log.Fatal("could not open connection: ", err)
	}

	// await sysexit

	app.Log.Info("bot has started")
	log.Print("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// cleanly close down the Discord session

	app.Session.Close()
}
