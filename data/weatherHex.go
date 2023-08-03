package data


import (
	"fmt"
)
/* The structure of a Yoon Suin weather hex.
 * The numbers are pointer to the next hex ID for a given d6 roll.
 */
 
 
type WeatherHex struct {
	Id int
	weatherSystem int
	WeatherTypeId int 
	One int
	Two int
	Three int
	Four int
	Five int
	Six int	
}

 func GetWeatherHexesForSystem(id int) map[int]WeatherHex {

	tableq :=  fmt.Sprintf("SELECT * FROM weatherhexes WHERE weatherSystemId = %d;", id);
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	hexes := map[int]WeatherHex{}
	for rows.Next() {
		currentHex := WeatherHex{}
		err = rows.Scan(&currentHex.Id, &currentHex.weatherSystem, &currentHex.WeatherTypeId, &currentHex.One, &currentHex.Two, &currentHex.Three, &currentHex.Four, &currentHex.Five, &currentHex.Six) 
		checkerr(err, "none")
		hexes[currentHex.Id] = currentHex;
	}

	 return hexes;	
}
