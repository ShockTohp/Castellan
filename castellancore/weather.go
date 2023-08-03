package castellancore
import (
	"castellan/data"
	"log"
	"fmt"
)



func weather() (string) {
 weather := data.GetMarkedHex()	
 return fmt.Sprintf("%s", weather);
}

func getTodaysWeatherReport(campaignId int) string {
	report, err := data.GetTodaysWeatherReport(campaignId)
	if err != nil {
		log.Fatal(err)
	}
	return report
}