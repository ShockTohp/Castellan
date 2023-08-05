package data


import (
	"fmt"
	"log"
)
/* The structure of a Yoon Suin weather hex.
 * The numbers are pointer to the next hex ID for a given d6 roll.
 */
 
 
type WeatherHex struct {
	Id int
	WeatherType *WeatherType
	One int
	Two int
	Three int
	Four int
	Five int
	Six int	
}

func NewWeatherHex(id, one, two, three, four, five, six int, wt *WeatherType) * WeatherHex {
	return &WeatherHex{
		Id: id,
	    WeatherType: wt,
		One: one,
		Two: two,
		Three: three,
		Four: four,
		Five: five,
		Six: six,	
	}
}

 func GetWeatherHexesForSystem(id int) map[int]*WeatherHex {

	tableq := "SELECT id, weatherTypeId, one, two, three, four, five, six FROM weatherhexes WHERE weatherSystemId = ?;"
	rows, err := runQuery(tableq, id)
	defer rows.Close()
	checkerr(err)

	hexes := map[int]*WeatherHex{}
	for rows.Next() {
		ch := WeatherHex{}
		var wtId int;
		err = rows.Scan(&ch.Id, &wtId, &ch.One, &ch.Two, &ch.Three, &ch.Four, &ch.Five, &ch.Six) 
		checkerr(err)
		wt := getWeatherTypeById(wtId)
		hexes[ch.Id] = NewWeatherHex(ch.Id, ch.One, ch.Two, ch.Three, ch.Four, ch.Five, ch.Six, wt);
	}
	 if len(hexes) < 1 {
		log.Fatal(fmt.Sprintf("No hexes retrieved for system %d", id))
	 }
	 return hexes;	
}

func (wh WeatherHex) GetWeatherName() string {
	return wh.WeatherType.GetWeatherName()
}

func (wh WeatherHex) String() string {
	return fmt.Sprintf("******\n WEATHER HEX:\n %d\n%s\n%d\n%d\n%d\n%d\n%d\n%d\n",
	wh.Id,
	wh.WeatherType,
	wh.One,
	wh.Two,
	wh.Three,
	wh.Four,
	wh.Five,
	wh.Six)
}