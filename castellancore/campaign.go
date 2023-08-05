package castellancore

import (
	"log"
	"fmt"
	"castellan/data"
)

func registerCampaign(GuildID, name, wt string)  string {
	if !data.IsValidWeatherSystem(wt) {
		return fmt.Sprintf("%s is not a valid weather system. Please check valid weather system list in the documentation.", wt);
	}

	log.Println(fmt.Sprintf("tried to register a campagin with weather type %s", wt))
	_, err := data.RegisterCampaign(GuildID, name, wt)
	log.Println(fmt.Sprintf("Registered a campagin with weather type %s", wt))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Campaign %s successfully registered for Guild %s with weather type %s", name, GuildID, wt)
} 
