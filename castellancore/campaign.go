package castellancore

import (
	"log"
	"fmt"
	"castellan/data"
)

func registerCampaign(GuildID string, name string)  string {
	_, err := data.RegisterCampaign(GuildID, name)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Campaign %s successfully registered for Guild %s with weather type %s", name, GuildID)
} 
