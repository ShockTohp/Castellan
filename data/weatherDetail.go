package data

import (
	"fmt"
	"time"
)

const weatherDetailTable = "detailedWeatherHexReports"
type WeatherDetail struct {
	name string
	location string
	details string
}


func NewWeatherDetail(n, l, d string) *WeatherDetail{
	return &WeatherDetail{
		name: n,
		location: l,
		details: d,
	}
}

func (wd WeatherDetail) Name() string {
	return wd.name
}

func (wd WeatherDetail) Location() string {
	return wd.location
}

func (wd WeatherDetail) Details() string {
	return wd.details
}

func GetWeatherDetailsForCampaignDate(cId int, date string) map[string]*WeatherDetail{
	tableq :=  fmt.Sprintf("SELECT weatherName, location, details FROM %s WHERE campaignId = %d AND weatherDate LIKE \"%s\";", weatherDetailTable, cId, date);
	rows, err := runQuery(tableq)
	defer rows.Close()
	checkerr(err)

	details := map[string]*WeatherDetail{}
		for rows.Next() {
		currentType := WeatherDetail{}
		err = rows.Scan(&currentType.name, &currentType.location, &currentType.details) 
		checkerr(err)
		details[currentType.location] = &currentType
	}

	 return details;	

}

func GetLatestCampaignDate(cId int) time.Time {
	tableq :=  fmt.Sprintf("SELECT weatherDate FROM %s WHERE campaignId = %d ORDER BY weatherDate DESC LIMIT 1;", weatherDetailTable, cId);
	rows, err := runQuery(tableq)
	defer rows.Close()
	checkerr(err)

	dates := make([]time.Time, 0)
		for rows.Next() {
		var curDateString string
		err = rows.Scan(&curDateString) 
		checkerr(err)
		curDate, derr := time.Parse(dateLayout, curDateString)
		checkerr(derr)
		dates = append(dates, curDate)
	}

	 return dates[0];	

}

func GetEarliestCampaignDate(cId int) time.Time {
	tableq :=  fmt.Sprintf("SELECT weatherDate FROM %s WHERE campaignId = %d ORDER BY weatherDate ASC LIMIT 1;", weatherDetailTable, cId);
	rows, err := runQuery(tableq)
	defer rows.Close()
	checkerr(err)

	dates := make([]time.Time, 0)
		for rows.Next() {
		var curDateString string
		err = rows.Scan(&curDateString) 
		checkerr(err)
		curDate, derr := time.Parse(dateLayout, curDateString)
		checkerr(derr)
		dates = append(dates, curDate)
	}

	 return dates[0];	

}