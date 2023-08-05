package castellancore

import (
	"fmt"
	"time"
	"errors"
	
	"castellan/data"
)

type weatherHexReporter struct {
	date string
	weatherDetailMap map[string]*data.WeatherDetail
}


func NewWeatherHexReporter(d string, wdp map[string]*data.WeatherDetail) *weatherHexReporter {
	return &weatherHexReporter{
		date: d,
		weatherDetailMap: wdp,
	}
}


func GetReporterForCampaignDate(cId int, date string) (*weatherHexReporter, error) {
	details := data.GetWeatherDetailsForCampaignDate(cId, date)
	if len(details) > 0 {
		return NewWeatherHexReporter(date, details), nil
	}
	testDate, _ := time.Parse(data.DateLayout, date)
	lastDate := data.GetLatestCampaignDate(cId)
	if testDate.After(lastDate) {
		return nil, errors.New("That date is further in the future than has been rolled. Currently, future dates are not supported. Please use the '/weather' command to generate weather up to this date.")
	}
	firstDate := data.GetEarliestCampaignDate(cId)
	if testDate.Before(firstDate) {
			return nil, errors.New("That date is further in the past than the campaign's start. Currently, generating weather for the past is not supported. Please try not to time travel.")
	}
	return nil, errors.New("That date does not seem to be in the campaign records. Contact the developer, it's probably his fault.")

}


func (whr weatherHexReporter) detailReport(loc ...string) string{
	if len(loc) > 0 {
		if det, detok := whr.weatherDetailMap[loc[0]]; detok {
			return fmt.Sprintf("**Weather report for %s**\n*%s*\n %s", whr.date, det.Name(), det.Details())
		}
		return fmt.Sprintf("%s Is not a recognized location, please check the known locations list", loc);
	}
	return whr.String()
}


func (whr weatherHexReporter) String() string {
	var locationStr = ""
	var name = ""
	for city, det := range whr.weatherDetailMap {
		name = det.Name()
		locationStr = locationStr + fmt.Sprintf("\t%s: %s\n", city, det.Details()) 
	}
	return fmt.Sprintf("**Weather report for %s**\n*%s*\n %s", whr.date, name, locationStr)
}