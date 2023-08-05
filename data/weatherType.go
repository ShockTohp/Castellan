package data 
import (
	"fmt"
)

const weatherTypeTable = "weatherTypes"

type WeatherType struct {
	Id int
	name string
}

func NewWeatherType(id int, n string) *WeatherType {
	return &WeatherType{
		Id: id,
		name: n,
	}
}

func GetWeatherTypesForSystem(id int) map[int]*WeatherType {
	tableq :=  fmt.Sprintf("SELECT id, weatherName FROM %s WHERE weatherSystemId = %d;", weatherTypeTable, id);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err)

	types := map[int]*WeatherType{}
	for rows.Next() {
		currentType := WeatherType{}
		err = rows.Scan(&currentType.Id, &currentType.name) 
		checkerr(err)
		types[currentType.Id] = NewWeatherType(currentType.Id, currentType.name);
	}

	 return types;	
}

func getWeatherTypeById(id int) *WeatherType {
	tableq :=  fmt.Sprintf("SELECT weatherName FROM %s WHERE id = %d;", weatherTypeTable, id);
	var name string
	err := db.QueryRow(tableq).Scan(&name)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	checkerr(err)

	 return NewWeatherType(id, name)	
}

func (wt WeatherType) GetWeatherName() string {
	return wt.name
}

func (wt WeatherType) String() string {
	return fmt.Sprintf("*****\n WEATHER TYPE\n %s", wt.name)
}
