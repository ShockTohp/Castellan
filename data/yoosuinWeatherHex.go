package data

import (
	"fmt"
)
/* The structure of a Yoon Suin weather hex.
 * The numbers are pointer to the next hex ID for a given d6 roll.
 */
type WeatherHex struct {
	id int
	weather_id WeatherType
	marked bool
	one int
	two int
	three int
	four int
	five int
	six int	
}

type WeatherType int64
const (
	ClearSkies    WeatherType = 1
	HeatWave      WeatherType = 2
	ColdFront     WeatherType = 3
	StrongWinds   WeatherType = 4
	TropicalStorm WeatherType = 5
	Typhoon       WeatherType = 6
)
var WeatherNames = []string {
	"None",
	"Clear Skies",
	"Heat Wave",
	"Cold Front",
	"Strong Winds",
	"Tropical Storm",
	"Typhoon",
} 
 func GetMarkedHex() WeatherHex {

	tableq :=  "SELECT * FROM weatherhexes WHERE marked = 1;";
	rows, err := runQuery(tableq)//"SELECT name FROM sqlite_schema WHERE type = 'table' AND name NOT LIKE 'sqlite_%';")
	defer rows.Close()
	checkerr(err, "tableq")

	hexes := make([]WeatherHex, 0)
	for rows.Next() {
		currentHex := WeatherHex{}
		err = rows.Scan(&currentHex.id, &currentHex.weather_id, &currentHex.marked, &currentHex.one, &currentHex.two, &currentHex.three, &currentHex.four, &currentHex.five, &currentHex.six) 
		checkerr(err, "none")
		hexes = append(hexes, currentHex);
	}

	 return hexes[0];	
}

func  (w WeatherHex) String() string {
	return fmt.Sprintf("The Hex id is: %d\n the current weather pattern is %s.", w.id, WeatherNames[w.weather_id])
}