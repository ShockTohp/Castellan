package castellancore
import (
	"castellan/data"
	"math/rand"
	"fmt"
	"errors"
	"time"
	"log"
	//"strconv"
)


type weatherResolver struct {
	campaignId int
	WeatherSystem data.WeatherSystem
	resolutionType string
	weatherTypes map[int]data.WeatherType
	weatherHexes map[int]data.WeatherHex
	lastMarker *data.WeatherMarker
	weatherTypeMap map[int]string
	lastDiceRoll int
}

type validResolutionTypes string
const (
	hexFlower = "Hex Flower"
	dice = "Dice"
)

func (wr weatherResolver) resolve() (string, error) {
	switch {
	case wr.resolutionType == hexFlower:
		return wr.resolveHexFlower(), nil
	case wr.resolutionType == dice:
		return "dice resolution", nil
	default:
		return "", errors.New(fmt.Sprintf("%s if not a valid resolution type", wr.resolutionType)) 
	}
} 

func (wr weatherResolver) resolveHexFlower() string {
	if wr.lastMarker == nil {
		return wr.resolveFirstHexFlower()
	} else {
		rand.Seed(time.Now().UnixNano())
		diceRoll := rand.Intn(6 - 1 + 1) + 1
		nextHex, err := wr.getNextHex(diceRoll)
		if err != nil {
			return "Sorry, we encountered an unspecified error. Please report this to the admin."
		}
		newDate := wr.lastMarker.Date.Add(24 * time.Hour)
		wm := data.WeatherMarker{0, wr.campaignId, wr.WeatherSystem.Id, nextHex, newDate, diceRoll}
		wm.WriteMarker()
		wr.lastMarker = data.GetLastWeatherMarkerForCampaign(wr.campaignId)
		wr.registerResolver()
		return wr.weatherReport()
	}
}

func (wr weatherResolver) getNextHex(dr int) (int, error)  {
	switch {
		case dr == 1:
			return wr.weatherHexes[wr.lastMarker.CurrentHex].One, nil
		case dr == 2:
			return wr.weatherHexes[wr.lastMarker.CurrentHex].Two, nil
		case dr == 3:
			return wr.weatherHexes[wr.lastMarker.CurrentHex].Three, nil
		case dr == 4:
			return wr.weatherHexes[wr.lastMarker.CurrentHex].Four, nil
		case dr == 5:
			return wr.weatherHexes[wr.lastMarker.CurrentHex].Five, nil
		case dr == 6:
			return wr.weatherHexes[wr.lastMarker.CurrentHex].Six, nil
		default:
			return -1, errors.New(fmt.Sprintf("Dice roll out of bounds"))
	}
}

func (wr weatherResolver) resolveFirstHexFlower() string {
	diceRoll := 0 //this means that this is the first marker in this chain
	firstHex := wr.WeatherSystem.StartingHex;
	wm := data.WeatherMarker{0, wr.campaignId, wr.WeatherSystem.Id, firstHex, time.Now(), diceRoll}
	wm.WriteMarker()
	wr.lastMarker = data.GetLastWeatherMarkerForCampaign(wr.campaignId)
	wr.registerResolver()
	return wr.weatherReport()
}

func (wr weatherResolver) weatherReport() string {
	return fmt.Sprintf ("Weather report for %s. The Dice roll today was a  %d. The hex is currently: %d. That means the current weather is %s", wr.lastMarker.DateString(), 
	wr.lastMarker.LastDiceRoll,
	wr.lastMarker.CurrentHex, 
	wr.getWeatherTypeForLastMarker())
}

func dateString(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
}

var openCampaignResolvers = map[string]weatherResolver{}

func (wr weatherResolver) registerResolver() {
	openCampaignResolvers[string(wr.campaignId)] = wr;
}

func buildWeatherResolverForCampaign(cId int, weatherName string) {
	log.Println("adding new resolver")
	wr := weatherResolver{}
	wr.campaignId = cId;
	wr.WeatherSystem = data.GetWeatherSystemByName(weatherName);
	wr.resolutionType = wr.WeatherSystem.ResolutionType
	wr.weatherTypes = data.GetWeatherTypesForSystem(wr.WeatherSystem.Id)
	wr.weatherHexes = data.GetWeatherHexesForSystem(wr.WeatherSystem.Id)
	wr.lastMarker = data.GetLastWeatherMarkerForCampaign(cId)
	wr.weatherTypeMap = map[int]string{}
	for _, wt := range wr.weatherTypes {
		wr.weatherTypeMap[wt.Id] = wt.Name
	}
	wr.registerResolver();
}


func (wr weatherResolver) getWeatherTypeForLastMarker() string {
	var hexId = wr.lastMarker.CurrentHex
	var weatherType = wr.weatherHexes[hexId].WeatherTypeId
	return wr.weatherTypeMap[weatherType]
}



func resolveWeather(cId int, wn string) (string, error) {
	if wr, ok := openCampaignResolvers[string(cId)]; ok {
		return wr.resolve()
	} else {
		buildWeatherResolverForCampaign(cId, wn)
		return resolveWeather(cId, wn)
	}
}

func weather(GuildID, wn string) string {
	cId, err := data.GetCampaignIdByGuild(GuildID)
	if err != nil {
		return "error getting weather, have you intiliazed a campaign?"
	}
	msg, _ := resolveWeather(cId, wn)
	return msg
}

func getTodaysWeatherReport(campaignId int) string {
	return "Todays Weather Report is not functional"
}

func dailyWeather(id, wType string) (string) {
	return "dailyWeather not yet functional"
	
}