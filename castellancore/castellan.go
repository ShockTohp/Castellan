package castellancore

import (
	"flag"
	"os"
	"os/signal"
	"log"
	"time"
	//"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	
)

var s *discordgo.Session
var sc *gocron.Scheduler
var (
	GuildID     	= flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globall")
	BotToken		= flag.String("token", "1136002891132452875", "Bot access token")
	RemoveCommands	= flag.Bool("rmcmd", true, "Remove all commands after shutting down or not") 
)

func init() {
	flag.Parse()
}

func init() {
	sc = gocron.NewScheduler(time.UTC)
	sc.StartAsync()
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

var (
	integerOptionMinValue     = 1.0
	dmPermission              = false
	defaultMemberPermissions  = discordgo.PermissionManageServer

	commands = []*discordgo.ApplicationCommand{
		{
			Name: "help",
			Description: "Outputs the help message.",
		},
		{
			Name: "weather",
			Description: "Rolls for the next unrolled day and gives a short report.",
		},
		{
			Name:        "weather-report",
			Description: "Delivers the weather report for a given day and location.",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "date",
					Description: "The date you wish to retrieve.",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "location",
					Description: "Give a specific locations weather, if configured.",
					Required:    false,
				},
			},
		},
		/*{
			Name: "schedule-report",
			Description: "Schedule a weather report",
			Options: []*discordgo.ApplicationCommandOption{
				{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "time",
				Description: "Time you want to recieve the report",
				Required: true,
				},
				{
				Type: discordgo.ApplicationCommandOptionInteger,
				Name: "timezone",
				Description: "Your timezone",
				Required: true,
				},
			},
		},*/
		{
			Name:        "register",
			Description: "Register a new campaign in this server.",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "campaign-name",
					Description: "The name of your campaign",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "weather-system",
					Description: "The weather system for your campaign. THIS CAN NOT BE CHANGED. Defaults to Yoon Suin.",
					Required:    false,
				},
			},
		},
	}
	


	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: helpMessage,
				},
			})

		},
		"weather": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: weather(i.Interaction.GuildID),
				},
			})
		},
		"weather-report": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// Get the value from the option map.
			// When the option exists, ok = true
			if option, nok := optionMap["date"]; nok {
				if loc, lok := optionMap["location"]; lok {
					msg := weatherReport(i.Interaction.GuildID, option.StringValue(), loc.StringValue())
					s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg ,
				},
			})
				} else {
					msg := weatherReport(i.Interaction.GuildID, option.StringValue())
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg ,
				},
			})
		}

			}

					},
		/*"schedule-report": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			weatherMessage := func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			channel := i.Interaction.ChannelID
			_, err := s.ChannelMessageSend(channel, weather(i.Interaction.GuildID))
			if err != nil {
				log.Fatal(err)
			}
			log.Println("tried to send message")
						}


			for _, opt := range options {
				optionMap[opt.Name] = opt
			}
			if opt, ok := optionMap["seconds"]; ok {
				var timeout int = int(opt.IntValue())
				//timeout = time.Duration(opt.IntValue()) 
				sc.Every(timeout).Second().Do(weatherMessage, s, i)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Message Scheduled", 
					},
				})
			}
		},*/
			"register": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// Get the value from the option map.
			// When the option exists, ok = true
			if opt, ok := optionMap["campaign-name"]; ok {
				// Option values must be type asserted from interface{}.
				// Discordgo provides utility functions to make this simple.
				var msg string;
				if wth, wok := optionMap["weather-system"]; wok {
					msg = registerCampaign(i.Interaction.GuildID, opt.StringValue(), wth.StringValue())
				} else {
				msg = registerCampaign(i.Interaction.GuildID, opt.StringValue(), "Yoon Suin")
				}
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
				},
			})
		}
			
		},
		
	}
)




func init() {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}



func Run() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range registeredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")

}



