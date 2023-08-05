package castellancore

import (
	"fmt"
	
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


func GetReporterForCampaignDate(cId int, date string) *weatherHexReporter {
	details := data.GetWeatherDetailsForCampaignDate(cId, date)
	return NewWeatherHexReporter(date, details)
}


func (whr weatherHexReporter) detailReport(loc []string) string{
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