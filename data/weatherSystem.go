package data

import (
	"fmt"
	"errors"
)

const weatherSystemTable = "weatherSystems"

var validWeatherSystems = map[string]int{}
type WeatherSystem struct {
	Id int
	systemName string
	ResolutionType *ResolutionType
	StartingHex int
}

func IsValidWeatherSystem(n string) bool {
	_, yes := validWeatherSystems[n]
	return yes;
}

func getWeatherSystemId(n string) (int, error) {
	if !IsValidWeatherSystem(n) {
		return -1, errors.New("Invalid Weather System")
	}
	return validWeatherSystems[n], nil
}
func NewWeatherSystem(id, sh int, name string, rt *ResolutionType) * WeatherSystem {
	return &WeatherSystem{
		Id: id,
		systemName: name,
		ResolutionType: rt,
		StartingHex: sh,
	}
}

func GetWeatherSystemByName(name string) (* WeatherSystem, error) {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE systemName LIKE \"%s\";", weatherSystemTable, name);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)

	systems := make([]WeatherSystem, 0)
	rIds := make([]int, 0)
	for rows.Next() {
		currentSys := WeatherSystem{}
		var crId int;
		err = rows.Scan(&currentSys.Id, &currentSys.systemName, &crId, &currentSys.StartingHex) 
		checkerr(err)
		systems = append(systems, currentSys);
		rIds = append(rIds, crId);
	}

	 if len(systems) > 1 {
		return nil, errors.New(fmt.Sprintf("Too many systems with that name, try being more specific."))
	 } else if len(systems) < 1 {
			return nil, errors.New(fmt.Sprintf("No systems with that name, please ensure it is a valid system."))
		}
	 
	 rt, err := getResolutionTypeById(rIds[0]); 
	 if err != nil {
		return nil, err
	 }
	 return NewWeatherSystem(systems[0].Id, systems[0].StartingHex, systems[0].systemName, rt), nil;	
	
}

func getWeatherSystemById(id int) (* WeatherSystem, error) {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE id = %d;", weatherSystemTable, id);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)

	systems := make([]WeatherSystem, 0)
	rIds := make([]int, 0)
	for rows.Next() {
		currentSys := WeatherSystem{}
		var rId int
		err = rows.Scan(&currentSys.Id, &currentSys.systemName, &rId, &currentSys.StartingHex) 
		checkerr(err)
		systems = append(systems, currentSys);
		rIds = append(rIds, rId)
	}
	rt, err := getResolutionTypeById(rIds[0]); 
	 if err != nil {
		return nil, err
	 }
	 return NewWeatherSystem(systems[0].Id, systems[0].StartingHex, systems[0].systemName, rt), nil;	

}