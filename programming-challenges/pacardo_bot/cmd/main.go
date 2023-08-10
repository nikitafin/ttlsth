package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

type CommandsCallback func(session *discordgo.Session,
	interaction *discordgo.InteractionCreate)

type CommandHandlers = map[string]CommandsCallback

var commandsHandlers CommandHandlers

var BotToken string
var GuildID string
var ApplicationID string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't load .env file")
	}

	ApplicationID = os.Getenv("APPLICATION_ID")
	if ApplicationID == "" {
		log.Fatal("Cannot read APPLICATION_ID")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	if BotToken == "" {
		log.Fatal("Can't read BOT_TOKEN")
	}
	GuildID = os.Getenv("GUILD_ID")
	if GuildID == "" {
		log.Fatal("Cannot read GUILD_ID")
	}

	botSession, createBotErr := discordgo.New("Bot " + BotToken)
	if createBotErr != nil {
		log.Fatal("Can't create bot session")
	}
	botSession.Identify.Intents |= discordgo.IntentGuildMessages
	botSession.Identify.Intents |= discordgo.IntentMessageContent

	botSession.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		}
	})

	_, err = botSession.ApplicationCommandCreate(ApplicationID, GuildID, &discordgo.ApplicationCommand{
		Name:        "buttons",
		Description: "descriptions",
	})
	if err != nil {
		log.Fatalf("Cannot create slash command: %v", err)
	}

	botSessionOpenErr := botSession.Open()
	if botSessionOpenErr != nil {
		log.Fatalf("Cannot open the session: %v", botSessionOpenErr)
	}
	defer botSession.Close()

	fmt.Println("Bot is running. Press ctrl+c to close")
	sc := make(chan os.Signal, 1)
	signal.Notify(
		sc,
		syscall.SIGINT,
		syscall.SIGTERM,
		os.Interrupt,
		syscall.SIGTERM,
	)
	<-sc
}

func init() {
	commandsHandlers = make(CommandHandlers)
	commandsHandlers["buttons"] = ButtonsCallback
}
func ButtonsCallback(session *discordgo.Session, create *discordgo.InteractionCreate) {}

//func ButtonsCallback(session *discordgo.Session, create *discordgo.InteractionCreate) {
//	components := []discordgo.MessageComponent{
//		discordgo.ActionsRow{
//			Components: []discordgo.MessageComponent{
//				discordgo.Button{
//					// Label is what the user will see on the button.
//					Label: "Yes",
//					// Style provides coloring of the button. There are not so many styles tho.
//					Style: discordgo.SuccessButton,
//					// Disabled allows bot to disable some buttons for users.
//					Disabled: false,
//					// CustomID is a thing telling Discord which data to send when this button will be pressed.
//					CustomID: "fd_yes",
//				},
//				discordgo.Button{
//					Label:    "No",
//					Style:    discordgo.DangerButton,
//					Disabled: false,
//					CustomID: "fd_no",
//				},
//				discordgo.Button{
//					Label:    "I don't know",
//					Style:    discordgo.LinkButton,
//					Disabled: false,
//					// Link buttons don't require CustomID and do not trigger the gateway/HTTP event
//					URL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
//					Emoji: discordgo.ComponentEmoji{
//						Name: "🤷",
//					},
//				},
//			},
//		},
//		discordgo.ActionsRow{
//			Components: []discordgo.MessageComponent{
//				discordgo.SelectMenu{
//					CustomID:    "select",
//					Placeholder: "Bet",
//					MinValues:   nil,
//					MaxValues:   1,
//					Options: []discordgo.SelectMenuOption{
//						{
//							Label:       "WIN",
//							Value:       "win",
//							Description: "TEMICK EZ WIN",
//						},
//						{
//							Label:       "LOOOOSE",
//							Value:       "loose",
//							Description: "TEMICK EZ LOOSE",
//							Default:     true,
//						},
//					},
//				},
//			},
//		},
//	}
//
//	data := discordgo.InteractionResponseData{
//		Content:    "Message TEXT",
//		Components: components,
//		Flags:      discordgo.MessageFlagsSuppressEmbeds,
//	}
//
//	response := discordgo.InteractionResponse{
//		Type: discordgo.InteractionResponseChannelMessageWithSource,
//		Data: &data,
//	}
//
//	err := session.InteractionRespond(create.Interaction, &response)
//	if err != nil {
//		panic(err)
//	}
//}
