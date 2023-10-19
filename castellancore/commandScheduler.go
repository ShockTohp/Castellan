package castellancore

import (
	"time"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
)
const twentyFourHrLayout = "16:04"
const amPmLayout = "04:00 PM"
var weatherMessage = func(s *discordgo.Session, channelId string, guildId int) {
			_, err := s.ChannelMessageSend(channelId, dailyWeatherReport(guildId, time.Now().String()))
		
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Sent daily weather message.")
}


type Command struct {
	campaignId string
	guildId string
	channelId string
	commmandName string
	time time.Time
	timezone time.Location
	params map[string]interface{}
}


func NewCommand(c, cn, g, ch string, t time.Time, tz time.Location, p map[string]interface{}) *Command {
	return &Command {
		campaignId: c,
		guildId: g,
		channelId: ch,
		commmandName: cn,
		time: t,
		timezone: tz,
		params: p,
	}
}

func writeScheduledCommandToDatabase() {
	
}

func loadScheduledCommands() {
	
}

func ScheduleWeatherReport(s *discordgo.Session, channelId, guildId, time string, sc *gocron.Scheduler) {
	sc.Every(1).Day().At(time).Do(weatherMessage, s, channelId, guildId)
}
