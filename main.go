package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Shoji-Nakasu/discord-to-slack/config"
	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + config.Config.Token) //"Bot"という接頭辞がないと401 unauthorizedエラーが起きる
	if err != nil {
		fmt.Println("error:start\n", err)
		return
	} //on message
	dg.AddHandler(messageCreate)
	err = dg.Open()
	if err != nil {
		fmt.Println("error:wss\n", err)
		return
	}
	fmt.Println("BOT Running...") //シグナル受け取り可にしてチャネル受け取りを待つ（受け取ったら終了）
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM) 
	<- sc
	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	nick := m.Author.Username
	member, err := s.State.Member(m.GuildID, m.Author.ID)
	if err == nil && member.Nick != "" {
		nick = member.Nick
	}
	fmt.Println("< " + m.Content + " by " + nick)
	if m.Content == "test1" {
		s.ChannelMessageSend(m.ChannelID, "test1")
		fmt.Println(">test1")
	}
	if strings.Contains(m.Content, "test2") {
		s.ChannelMessageSend(m.ChannelID, "test2")
		fmt.Println("> test2")
	}
}
