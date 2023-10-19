package castellancore

import(
	"fmt"
	"castellan/data"
	"log"
	"time"
)

const (
	hexFlower ="Hex Flower"
	rDice = "Dice"
)

func weather(GuildID string) string {
	c, err := data.GetCampaignByGuild(GuildID)
	if err != nil {
		return "Unable to retrieve campaign information, have you intialized a campaign?"
	}
	var r = c.WeatherSystem().ResolutionType.Name
	switch {
	case r == hexFlower:
		if wr, ok := openCampaignResolvers[c.Id()]; ok {
			return wr.resolve()
		}
		wr := NewWeatherHexResolver(c, data.GetWeatherHexesForSystem(c.WeatherSystem().Id), data.GetLastWeatherMarkerForCampaign(c.Id()) );
		registerResolver(wr)
		return wr.resolve();
	case r == rDice:
		return "Dice tables not yet implemented"
	default:
		return fmt.Sprintf("This campaign has an unknown weather resolution sytem: %s. Please contact the developer.", r)
	}
}


func dailyWeatherReport(cid int, date string) string {
	c, err := data.GetCampaignById(cid)
	if err != nil {
		return "Error retrieving campaign information"
	}
	if c.CampaignHasReportForDate(date) {
		return weatherReport(string(c.GuildId()), date)
	}
	lastDate := data.GetLatestCampaignDate(c.Id())
	testDate, derr := time.Parse(data.DateLayout, date)
	if derr != nil {
		return "Error parsing today's date"
	}
	time := testDate.Sub(lastDate)
	daysApart := (time.Hours() / 24)
	if daysApart < 0 {
		return "A wizard has wrecked your calendar. Today appears to be in the past but has no weather information."
	} 
	if daysApart > 1 {
		return fmt.Sprintf("Today is %d days further in the future than the last weather day. Please manually use the weather command to catch up", daysApart)
	}
	wr := NewWeatherHexResolver(c, data.GetWeatherHexesForSystem(c.WeatherSystem().Id), data.GetLastWeatherMarkerForCampaign(c.Id()))
	wr.resolve()
	return weatherReport(string(c.GuildId()), date)
}



func weatherReport(GuildID, date string, loc ...string) string {
	c, err := data.GetCampaignByGuild(GuildID)
	if err != nil {
		log.Println(err)
		return "Unable to retrieve campaign information, have you intialized a campaign?"
	}
	var r = c.WeatherSystem().ResolutionType.Name
	switch {
	case r == hexFlower:
		wr, err := GetReporterForCampaignDate(c.Id(), date)
		if err == nil {
			if (len(loc) > 0) {
			return wr.detailReport(loc[0]);
			}
			return wr.detailReport()
		}
		return fmt.Sprintf("%s", err)
	case r == rDice:
		return "Dice tables not yet implemented"
	default:
		return fmt.Sprintf("This campaign has an unknown weather resolution sytem: %s. Please contact the developer.", r)
	}
}


var openCampaignResolvers = map[int]*weatherHexResolver{}

func registerResolver(wr *weatherHexResolver) {
	log.Println(fmt.Sprintf("Registered new resolver for campagin %s", wr.Campaign.Id()))
	openCampaignResolvers[(wr.Campaign.Id())] = wr;
}

