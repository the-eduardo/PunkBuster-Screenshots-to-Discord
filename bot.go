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
	ChannelID = "YOUR_Screenshots_ChannelID" // You should create a New Channel! The bot will spam a lot of images!

)
func init() { flag.Parse() }
		// TODO: Create a new function that delete the local files (we don't want to resend the same files every time)
		DisgordMain() {

	sc, _ := discordgo.New("Bot " + *BotToken)
	sc.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot job is done, closing the session!")
		sc.Close()
	})

	for x := 0; x < len(fileVerify()); x++ {
		//fmt.Println("Files found")
		filePath := "./downloads/" + fileVerify()[x]
		log.Println("Sending file: ", fileVerify()[x])
		myFile, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Cannot open the file: %v", err)
		}
		_, err = sc.ChannelFileSend(ChannelID, filePath, myFile)
		if err != nil {
			log.Fatalf("Cannot send the file to the Channel: %v", err)
		}
	}

	err := sc.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer sc.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, os.Kill)
	<-stop
	log.Println("Graceful shutdown")
	stop <- os.Interrupt

}

func fileVerify() []string { // Check if the file exists and return as io.Reader
	// List all files on local path downloads/
	files, err := os.ReadDir("./Downloads/")
	if err != nil {
		log.Fatalf("[BOT] Error on verifying Local Files: %v", err)
	}

	if len(files) == 0 {
		log.Panic("No files found")
	}

	var LocalFiles []string
	for _, f := range files {
		//fmt.Println(f.Name())
		LocalFiles = append(LocalFiles, f.Name())
	}

	return LocalFiles
}
