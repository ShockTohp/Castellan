package data 

import ( 
	"time"
	"fmt"
	"log"
)
const markerTable = "weatherMarkers"
const dateLayout = "2006-01-02"
type WeatherMarker struct {
	Id int
	currentHex int
	date time.Time
	lastDiceRoll int
}

func NewWeatherMarker(id, ch, ldr int, date time.Time) *WeatherMarker {
	wm := &WeatherMarker {
		Id: id,
		currentHex: ch,
		date: date,
		lastDiceRoll: ldr,
	}
	log.Println(fmt.Sprintf("Created new marker\n %s", *wm))
	return wm
}

/*func GetWeatherMarkersForCampaign(cId int) []WeatherMarker {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE campaignId = %d ORDER BY weatherDate;", markerTable, cId);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)

	markers := make([]WeatherMarker, 0)
	for rows.Next() {
		currentMkr := WeatherMarker{}
		var dtStr = "" 
		err = rows.Scan(&currentMkr.Id, &currentMkr.CampaignId, &currentMkr.WeatherSystem, &currentMkr.CurrentHex, &dtStr, &currentMkr.LastDiceRoll) 
		currentMkr.Date, _ = time.Parse(dateLayout, dtStr)
		checkerr(err)
		markers = append(markers, currentMkr);
	}

	 return markers;	
}*/

func GetLastWeatherMarkerForCampaign(cId int) *WeatherMarker {
	tableq :=  fmt.Sprintf("SELECT id, hexId, weatherDate, lastDiceRoll FROM %s WHERE campaignId = %d ORDER BY weatherDate DESC LIMIT 1;", markerTable, cId);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)

	markers := make([]WeatherMarker, 0)
	for rows.Next() {
		currentMkr := WeatherMarker{}
		var dateString = ""
		err = rows.Scan(&currentMkr.Id, &currentMkr.currentHex, &dateString, &currentMkr.lastDiceRoll) 
		checkerr(err)
		currentMkr.date, _ = time.Parse(dateLayout, dateString)
		markers = append(markers, currentMkr);
	}

	if (len(markers) == 0) {
		return nil
	}
	log.Println(fmt.Sprintf("Loaded %d markers for campaing %d. The last marker is \n%s", len(markers), cId, markers[0]))
	 return NewWeatherMarker(markers[0].Id, markers[0].currentHex, markers[0].lastDiceRoll, markers[0].date);	
}

func (wm WeatherMarker) GetNextDate() time.Time {
	return wm.date.Add(24 * time.Hour)
}

func (wm WeatherMarker) WriteMarker(campId, ws int) {
	var dateString = wm.DateString() 
	registerQ := fmt.Sprintf("INSERT INTO %s (campaignId, weatherSystemId, hexId, WeatherDate, lastDiceRoll) VALUES (?, ?, ?, ?, ?);", markerTable)

	stmt, _ := db.Prepare(registerQ)
	stmt.Exec(campId, ws, wm.currentHex, dateString, wm.lastDiceRoll);
	defer stmt.Close()
}

func (wm WeatherMarker) String() string {
	return fmt.Sprintf("****\nId: %d \nCurrent Hex: %d\nDate: %s\nLast Dice Roll: %d", 
	wm.Id, 
 	wm.currentHex,
	wm.DateString(),
	wm.lastDiceRoll)
}

func (wm WeatherMarker) DateString() string {
	 return wm.date.Format(dateLayout)
}

func (wm WeatherMarker) Advance(hex, ldr int) (wmn *WeatherMarker) {
	date := wm.GetNextDate()
	return NewWeatherMarker(wm.Id, hex, ldr, date)
}

func (wm WeatherMarker) GetMarkedHex() int {
	return wm.currentHex
}

func (wm WeatherMarker) GetLastDiceRoll() int {
	return wm.lastDiceRoll
}
