package castellancore
import (
	"castellan/data"
	"castellan/util"
	"fmt"
	"errors"
	"time"
	"log"
	//"strconv"
)


type weatherHexResolver struct {
	Campaign *data.Campaign
	WeatherSystem *data.WeatherSystem
	weatherHexes map[int]*data.WeatherHex
	marker *data.WeatherMarker
}

func NewWeatherHexResolver(c *data.Campaign, wh map[int]*data.WeatherHex, m *data.WeatherMarker) (*weatherHexResolver) {
	return &weatherHexResolver{
	Campaign:  c,
	WeatherSystem: c.WeatherSystem(),
	weatherHexes: wh,
	marker: m,
	}
}

func (wr *weatherHexResolver) resolve() string {
	if wr.marker == nil {
		wr.resolveFirstHexFlower()
		return wr.weatherReport()
	} else {
		diceRoll := util.RollDSix()
		nextHex, err := wr.getNextHex(diceRoll)
		if err != nil {
			return "Sorry, we encountered an dice rolling error. Please report this to the admin. "
		}
		wr.advance(nextHex, diceRoll)
		defer wr.marker.WriteMarker(wr.Campaign.Id(), wr.WeatherSystem.Id)
		return wr.weatherReport()
	}
}

func (wr *weatherHexResolver) advance(hex, ldr int) {
	wr.marker = wr.marker.Advance(hex, ldr)
}

func (wr weatherHexResolver) getNextHex(dr int) (int, error)  {
	var hex = wr.getCurrentHex()
	switch {
		case dr == 1:
			return hex.One, nil
		case dr == 2:
			return hex.Two, nil
		case dr == 3:
			return hex.Three, nil
		case dr == 4:
			return hex.Four, nil
		case dr == 5:
			return hex.Five, nil
		case dr == 6:
			return hex.Six, nil
		default:
			return -1, errors.New(fmt.Sprintf("Dice roll out of bounds"))
	}
}

func (wr *weatherHexResolver) resolveFirstHexFlower() *weatherHexResolver {
	log.Println("Resolving first hex flower")
	diceRoll := 0 //this means that this is the first marker in this chain
	firstHex := wr.WeatherSystem.StartingHex;
    nm := data.NewWeatherMarker(0, firstHex, diceRoll, time.Now())
	defer nm.WriteMarker(wr.Campaign.Id(), wr.WeatherSystem.Id)
	wr.marker = nm
	return wr
}

func (wr weatherHexResolver) weatherReport() string {
	return fmt.Sprintf("Weather report for %s. The Dice roll today was a  %d. The hex is currently: %d. That means the current weather is %s", wr.marker.DateString(), 
	wr.marker.GetLastDiceRoll(),
	wr.marker.GetMarkedHex(), 
	wr.getCurrentMarkedWeatherName())
}

func dateString(date time.Time) string {
	return fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
}

func (wr weatherHexResolver) getCurrentMarkedWeatherName() string {
	return wr.getCurrentHex().GetWeatherName()
}



func dailyWeather(id, wType string) (string) {
	return "Daily Weather not yet functional"
	
}

func (wr weatherHexResolver) getCurrentHex() *data.WeatherHex {
	var hexId = wr.marker.GetMarkedHex()
	var hex = wr.weatherHexes[hexId]
	if hex == nil {
		log.Fatal("No weather hexes for systems")
	}
	return hex
}