package main
import (
	"fmt"
)

type WEEKDAYS int8
const (
	SUNDAY WEEKDAYS = 0
	MONDAY WEEKDAYS = 1
	TUESDAY WEEKDAYS = 2
	WEDNESDAY WEEKDAYS = 3
	THURSDAY WEEKDAYS = 4
	FRIDAY WEEKDAYS = 5
	SATURDAY WEEKDAYS = 6
)
func main(){
	fmt.Println("Enumeration")
	day := SUNDAY
	fmt.Println("Day code for SUNDAY:",day)
	day = TUESDAY
	fmt.Println("Day code for TUESDAY:",day)
}
