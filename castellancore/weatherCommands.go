package castellancore

import(
	"fmt"
	"castellan/data"
	"log"
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
	var r = c.WeatherSystem.ResolutionType.Name
	switch {
	case r == hexFlower:
		if wr, ok := openCampaignResolvers[c.Id]; ok {
			return wr.resolve()
		}
		wr := NewWeatherHexResolver(c, data.GetWeatherHexesForSystem(c.WeatherSystem.Id), data.GetLastWeatherMarkerForCampaign(c.Id) );
		registerResolver(wr)
		return wr.resolve();
	case r == rDice:
		return "Dice tables not yet implemented"
	default:
		return fmt.Sprintf("This campaign has an unknown weather resolution sytem: %s. Please contact the developer.", r)
	}
}




func weatherReport(GuildID, date string, loc ...string) string {
	c, err := data.GetCampaignByGuild(GuildID)
	if err != nil {
		log.Println(err)
		return "Unable to retrieve campaign information, have you intialized a campaign?"
	}
	var r = c.WeatherSystem.ResolutionType.Name
	switch {
	case r == hexFlower:
		wr, err := GetReporterForCampaignDate(c.Id, date)
		if err == nil {
		return wr.detailReport(loc[0]);
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
	log.Println(fmt.Sprintf("Registered new resolver for campagin %s", wr.Campaign.Id))
	openCampaignResolvers[(wr.Campaign.Id)] = wr;
}

