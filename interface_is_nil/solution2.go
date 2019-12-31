package interface_is_nil

import (
	"errors"
	"fmt"
)

type InterfaceNilChecker interface {
	IsNilInterface() bool
}

type Spinner interface {
	Spin()
	InterfaceNilChecker
}

type Wheels struct{}

func (wheels *Wheels) Spin() {
	fmt.Println("The wheels are spinning")
}

type Car struct {
	color string
	*Wheels
}

func (car *Car) IsNilInterface() bool {
	if car == (*Car)(nil) {
		return false
	}
	return true
}

type Bicycle struct {
	gears int
	*Wheels
}

func (bike *Bicycle) IsNilInterface() bool {
	if bike == (*Bicycle)(nil) {
		return false
	}
	return true
}

func main() {
	Volkswagen, _ := BuyCar(300000, "black")
	fmt.Println("\nTaking our Volkswagen for a ride:")
	Ride(Volkswagen)

	AstonMartin, _ := BuyCar(12, "red")
	fmt.Println("\nTaking our AstonMartin for a ride:")
	Ride(AstonMartin)

	bike, _ := BuyBicycle(2000, 17)
	fmt.Println("\nTaking our bike for a ride:")
	Ride(bike)

	cheapBike, _ := BuyBicycle(100, 1)
	fmt.Println("\nTaking our cheapBike for a ride:")
	Ride(cheapBike)

}

func Ride(vehicle Spinner) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Dude, where's my car?")
		}
	}()

	fmt.Printf("%T %v \n", vehicle, vehicle)
	if vehicle.IsNilInterface() {
		vehicle.Spin()
	}
}

func BuyCar(cash int, color string) (*Car, error) {
	if cash > 100500 {
		return &Car{color: color}, nil
	} else {
		return nil, errors.New("go find a job, pal")
	}
}

func BuyBicycle(cash int, gears int) (*Bicycle, error) {
	if cash > 1050 {
		return &Bicycle{gears: gears}, nil
	} else {
		return nil, errors.New("go find a job, pal")
	}
}