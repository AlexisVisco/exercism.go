package space

import "fmt"

type Planet string

var planets = map[string]float64{
	"Earth":	31556952,
	"Mercury":	7600387.751258,
	"Venus":	19413750.404352,
	"Mars":		59352813.921442,
	"Jupiter":	374347972.14948,
	"Saturn":	929273280.906096,
	"Uranus":	2651315576.413392,
	"Neptune":	5200311775.25664,
}

func Age(seconds float64, planet string) float64 {
	secondsInYear := planets[planet]
	fmt.Println(secondsInYear)
	return seconds / secondsInYear
}