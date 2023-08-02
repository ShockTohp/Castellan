package castellancore
import (
	"castellan/data"
	"fmt"
)




func weather() (string) {
 weather := data.GetMarkedHex()	
 return fmt.Sprintf("%s", weather);
}

