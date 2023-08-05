package data

import (
	"fmt"
)

const weatherTableTable = "weatherTableEntriesAndDetails"
type WeatherTable struct {
	tableName string
	diceNum int
	diceType int
	weatherType map[int]string
	weatherDetail map[int]string
	locations map[string]string //holds details for locations
}

func NewWeatherTable(n string, dn, dt int, wt, wd map[int]string, l map[string]string) (*WeatherTable) {
	return &WeatherTable{
		tableName: n,
		diceNum: dn,
		diceType: dt,
		weatherType: wt,
		weatherDetail: wd,
		locations: l,
	}
}

func GetWeatherTableForSystem(sysName string) (*WeatherTable, error) {
	id, ierr := getWeatherSystemId(sysName);
	if ierr != nil {
		return nil, ierr
	}
	tableq :=  fmt.Sprintf("SELECT name, diceType, diceNumber, weatherName, diceResult, location, details FROM %s WHERE weatherSystemId = ?", weatherTableTable);
	rows, err := runQuery(tableq, id)
	defer rows.Close()
	checkerr(err)

	tn := ""
	dt, dn := 0, 0
	types, details := map[int]string{}, map[int]string{}
	locations := map[string]string{}
		for rows.Next() {
			y, u, r := 0, 0, 0
			n, t, d, l := "", "", "", "" 
		err = rows.Scan(&n, &y, &u, &t, &r, &l, &d) 
		checkerr(err)
		tn = n
		dt = y
		dn = u
		types[r] = t
		details[r] = d
		locations[l] = d
	}
	 return NewWeatherTable(tn, dt, dn, types, details, locations), nil;	
}


func (wt WeatherTable) GetResultsForRoll(r int) (t, det string) {
	return wt.weatherType[r], wt.weatherDetail[r]
}

func (wt WeatherTable) GetDiceRollInformation() (t, n int) {
	return wt.diceType, wt.diceNum
}