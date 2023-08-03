package data 
import (
	"fmt"
)

const weatherTypeTable = "weatherTypes"

type WeatherType struct {
	Id int
	weatherSystemId int
	Name string
}

func GetWeatherTypesForSystem(id int) map[int]WeatherType {
	tableq :=  fmt.Sprintf("SELECT * FROM %s WHERE weatherSystemId = %d;", weatherTypeTable, id);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	types := map[int]WeatherType{}
	for rows.Next() {
		currentType := WeatherType{}
		err = rows.Scan(&currentType.Id, &currentType.weatherSystemId, &currentType.Name) 
		checkerr(err, "none")
		types[currentType.Id] = currentType;
	}

	 return types;	
}
