package interface_is_nil

import (
"errors"
"fmt"
"reflect"
)

type Spinner interface {
	Spin()
}

type Wheels struct{}

func (wheels *Wheels) Spin() {
	fmt.Println("The wheels are spinning")
}

type Car struct {
	color string
	*Wheels
}

type Bicycle struct {
	gears int
	*Wheels
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
	defer func(){
		if r:=recover(); r!=nil {
			fmt.Println("Dude, where's my car?")
		}
	}()

	if vehicle != nil && !reflect.ValueOf(vehicle).IsNil() {
		//we just made sure vehicle is not nil, so should be safe, right?
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