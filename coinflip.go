package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	token string
)

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
	for i := 0; i < 100; i++ {

	}
	if strings.ToLower(mess.Content) == "!!flip" {
		go func() {
			ds.ChannelMessageSend(mess.ChannelID, mess.Author.Mention()+
				" "+flip())
		}()
	}
}

func flip() string {
	o := big.NewInt(time.Now().Unix())
	i, err := rand.Int(rand.Reader, o)
	if err != nil {
		return "this cryptographically secure random value is really bad"
	}
	println(i.Uint64() % 2)
	if i.Uint64()%2 == 0 {
		return "HEADS!"
	}
	return "TAILS!"
}
