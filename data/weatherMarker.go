package data 

import ( 
	"time"
	"fmt"
)
const markerTable = "weatherMarkers"
const dateLayout = "2006-01-02"
type WeatherMarker struct {
	Id int
	CampaignId int
	WeatherSystem int
	CurrentHex int
	Date time.Time
	LastDiceRoll int
}

func GetWeatherMarkersForCampaign(cId int) []WeatherMarker {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE campaignId = %d ORDER BY weatherDate;", markerTable, cId);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	markers := make([]WeatherMarker, 0)
	for rows.Next() {
		currentMkr := WeatherMarker{}
		var dtStr = "" 
		err = rows.Scan(&currentMkr.Id, &currentMkr.CampaignId, &currentMkr.WeatherSystem, &currentMkr.CurrentHex, &dtStr, &currentMkr.LastDiceRoll) 
		currentMkr.Date, _ = time.Parse(dateLayout, dtStr)
		checkerr(err, "none")
		markers = append(markers, currentMkr);
	}

	 return markers;	
}

func GetLastWeatherMarkerForCampaign(cId int) *WeatherMarker {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE campaignId = %d ORDER BY weatherDate;", markerTable, cId);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	markers := make([]WeatherMarker, 0)
	for rows.Next() {
		currentMkr := WeatherMarker{}
		var dateString = ""
		err = rows.Scan(&currentMkr.Id, &currentMkr.CampaignId, &currentMkr.WeatherSystem, &currentMkr.CurrentHex, &dateString, &currentMkr.LastDiceRoll) 
		currentMkr.Date, _ = time.Parse(dateLayout, dateString)
		checkerr(err, "none")
		markers = append(markers, currentMkr);
	}

	if (len(markers) == 0) {
		return nil
	}
	 return &markers[len(markers) - 1];	
}

func (wm WeatherMarker) WriteMarker() {
	var dateString = wm.DateString() 
	registerQ := fmt.Sprintf("INSERT INTO %s (campaignId, weatherSystemId, hexId, WeatherDate, lastDiceRoll) VALUES (?, ?, ?, ?, ?);", markerTable)

	stmt, _ := db.Prepare(registerQ)
	stmt.Exec(wm.CampaignId, wm.WeatherSystem, wm.CurrentHex, dateString, wm.LastDiceRoll);
	defer stmt.Close()
}

func (wm WeatherMarker) String() string {
	return fmt.Sprintf("****\nId: %d\nCampaign: %d\nWeather System Id: %d\nCurrent Hex: %d\nDate: %s\nLast Dice Roll: %d", 
	wm.Id, 
	wm.CampaignId, 
	wm.WeatherSystem,
	wm.DateString,
 	wm.CurrentHex,
	wm.LastDiceRoll)
}

func (wm WeatherMarker) DateString() string {
	 return wm.Date.Format(dateLayout)
}
