package data

import (
	"fmt"
)

const weatherSystemTable = "weatherSystems"

type WeatherSystem struct {
	Id int
	systemName string
	ResolutionType string
	StartingHex int
}


func GetWeatherSystemByName(name string) WeatherSystem {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE systemName LIKE \"%s\";", weatherSystemTable, name);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	systems := make([]WeatherSystem, 0)
	for rows.Next() {
		currentSys := WeatherSystem{}
		err = rows.Scan(&currentSys.Id, &currentSys.systemName, &currentSys.ResolutionType, &currentSys.StartingHex) 
		checkerr(err, "none")
		systems = append(systems, currentSys);
	}

	 return systems[0];	
	
}

func getWeatherSystemById(id int) WeatherSystem {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE systemName = %d;", weatherSystemTable, id);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	systems := make([]WeatherSystem, 0)
	for rows.Next() {
		currentSys := WeatherSystem{}
		err = rows.Scan(&currentSys.Id, &currentSys.systemName, &currentSys.ResolutionType, &currentSys.StartingHex) 
		checkerr(err, "none")
		systems = append(systems, currentSys);
	}

	 return systems[0];	
	
}