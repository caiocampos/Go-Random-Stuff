// Package space provides ways to calculate how many orbital periods (years) a given time represents on other Planets
package space

// Planet a alias to string
type Planet string

// Enum of Planets
const (
	Earth   Planet = "Earth"
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// Enum of the duration of the orbital period (year) on each Planet
const (
	EarthYear   float64 = 31557600
	MercuryYear float64 = EarthYear * 0.2408467
	VenusYear   float64 = EarthYear * 0.61519726
	MarsYear    float64 = EarthYear * 1.8808158
	JupiterYear float64 = EarthYear * 11.862615
	SaturnYear  float64 = EarthYear * 29.447498
	UranusYear  float64 = EarthYear * 84.016846
	NeptuneYear float64 = EarthYear * 164.79132
)

// Age function calculate how old someone would be on a Planet
func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case Earth:
		return seconds / EarthYear
	case Mercury:
		return seconds / MercuryYear
	case Venus:
		return seconds / VenusYear
	case Mars:
		return seconds / MarsYear
	case Jupiter:
		return seconds / JupiterYear
	case Saturn:
		return seconds / SaturnYear
	case Uranus:
		return seconds / UranusYear
	case Neptune:
		return seconds / NeptuneYear
	}
	return 0
}
