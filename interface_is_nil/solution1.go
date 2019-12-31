package interface_is_nil

import (
"errors"
"fmt"
)

type Spinner interface {
	Spin()
	IsNil() bool
}

type Wheels struct{}

func (wheels *Wheels) Spin() {
	fmt.Println("The wheels are spinning")
}

func (car *Car) IsNil() bool{
	if car == (*Car)(nil) {
		return false
	}
	return true
}

type Car struct {
	color string
	*Wheels
}

type Bicycle struct {
	color string
	*Wheels
}

func main() {
	Volkswagen, _ := BuyCar(300000, "black")
	fmt.Println("\nTaking our Volkswagen for a ride:")
	Ride(Volkswagen)

	AstonMartin, _ := BuyCar(12, "red")
	fmt.Println("\nTaking our AstonMartin for a ride:")
	Ride(AstonMartin)
}

func Ride(vehicle Spinner) {
	defer func(){
		if r:=recover(); r!=nil {
			fmt.Println("Dude, where's my car?")
		}
	}()

	fmt.Printf("%T %v \n", vehicle, vehicle)
	if vehicle.IsNil() {
		vehicle.Spin()
	}

	// #1 Approach
	//switch v := vehicle.(type) {
	//	//case *Car:
	//	//	fmt.Printf("%T %v \n", v, v)
	//	//	if v != nil {
	//	//		//we just made sure vehicle is not nil, so should be safe, right?
	//	//		vehicle.Spin()
	//	//	}
	//	//default:
	//	//	fmt.Printf("Unsupported type %T!\n", v)
	//	//}

	//#2 Approach
	//if vehicle != (*Car)(nil) {
	//	//we just made sure vehicle is not nil, so should be safe, right?
	//	vehicle.Spin()
	//}

	//(*interface{})(nil)

	//if _, ok := vehicle.(interface{}); ok {
	//	vehicle.Spin()
	//}
}

func BuyCar(cash int, color string) (*Car, error) {
	if cash > 100500 {
		return &Car{color: color}, nil
	} else {
		return nil, errors.New("go find a job, pal")
	}
}
