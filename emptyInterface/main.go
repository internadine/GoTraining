package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{
		vehicle: vehicle{
			Seats:    4,
			MaxSpeed: 210,
			Color:    "red",
		},
		Wheels: 4,
		Doors:  2,
	}
	tacoma := car{}
	bmw528 := car{}
	boeing747 := plane{
		vehicle: vehicle{
			Seats:    20,
			MaxSpeed: 4,
			Color:    "green",
		},
		Jet: true,
	}
	boeing757 := plane{}
	boeing767 := plane{}
	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	v := vehicle{4, 210, "red"}
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu, v}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
