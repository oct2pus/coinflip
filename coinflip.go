package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "bot token")
}

func main() {
	flag.Parse()
	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		return
	}
	bot.AddHandler(messageCreate)
	// Open Bot
	err = bot.Open()
	if err != nil {
		fmt.Printf("Error openning connection: %v\n", err)
	}

	bot.UpdateStatus(0, "!!flip")

	// wait for ctrl+c to close.
	signalClose := make(chan os.Signal, 1)

	signal.Notify(signalClose,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
		os.Kill)
	<-signalClose

	bot.Close()
}

func messageCreate(ds *discordgo.Session, mess *discordgo.MessageCreate) {
	if mess.Author.Bot {
		return
	}
	if strings.ToLower(mess.Content) == "!!flip" {
		go func() {
			ds.ChannelMessageSend(mess.ChannelID, mess.Author.Mention()+
				" "+flip())
		}()
	}
}

func flip() string {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(2)
	if i == 0 {
		return "HEADS"
	}
	return "TAILS"
}
