package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)


var (
	BotToken = flag.String("token", "SECRET.TOKEN_HERE", "Bot token")

)

func init() { flag.Parse() }

func main() {
	contentTests := "Pruu"

	sc, _ := discordgo.New("Bot " + *BotToken)
	sc.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})
	msg, msgerr := sc.ChannelMessageSend("YOUR_CHANNEL_ID", contentTests)
	if msgerr != nil {
		log.Panic(msgerr)
	}
	fmt.Println(msg)

	err := sc.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer sc.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, os.Kill)
	<-stop
	log.Println("Graceful shutdown")

}

func fileVerify() []string {
	// List all files on local path downloads/
	files, err := os.ReadDir("./Downloads/")
	if err != nil {
		log.Fatal("[BOT] Error on verifying Local Files", err)
	}

	if len(files) == 0 {
		panic("No files found")
	}

	var LocalFiles []string
	for _, f := range files {
		fmt.Println(f.Name())
		LocalFiles = append(LocalFiles, f.Name())
	}

	return LocalFiles
}
