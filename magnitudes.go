package points

import (
	"fmt"
)

type Magnitudes []Magnitude

func (magnitudes Magnitudes) Len() int {
	return len(magnitudes)
}

func (magnitudes Magnitudes) Less(i, j int) bool {
	return magnitudes[i].Abs() < magnitudes[j].Abs()
}

func (magnitudes Magnitudes) Swap(i, j int) {
	magnitudes[i], magnitudes[j] = magnitudes[j], magnitudes[i]
}

func (magnitudes Magnitudes) String() string {
	str := ""
	for _, p := range magnitudes {
		str += p.String() + "\t" + fmt.Sprintf("%f", p.Abs()) + "\n"
	}
	return str
}
