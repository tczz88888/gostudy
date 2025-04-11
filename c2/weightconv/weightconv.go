package weightconv

import "fmt"

type Kg float64
type Bang float64

func (w Kg) String() string {
	return fmt.Sprintf("%gkg", w)
}

func (w Bang) String() string {
	return fmt.Sprintf("%gbang", w)
}

func Kg2Bang(kg Kg) Bang {
	return Bang(kg * 0.4536)
}

func Bang2Kg(bang Bang) Kg {
	return Kg(bang / 0.4536)
}
