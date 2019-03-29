package main

import "fmt"

func main() {

	//const a, b, c int = 1, 2, 3

	var a = 1

	var b int16

	b = 2

	c := 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//fmt.Println(c2)
	//fmt.Println(*c2)

	var outId, outType, valid = Test()

	fmt.Println(fmt.Sprint("%v %v", outId, outType, valid))
}

//Test makes a test
func Test() (int, string, bool) {

	var id int = 1

	var id2 *int = &id

	var car = Automovel{Id: id, Type: "Mono motor"}

	var _, tipo = car.Run(2, "")

	fmt.Println(tipo)

	fmt.Printf("%s", car.Id)

	car.Id = id
	id = 3

	car.Id = *id2
	*id2 = 9

	//fmt.Printf("%s", car.Id)

	return car.Id, car.Type, true
}

// Automovel is a vehicle
type Automovel struct {
	Id   int
	Type string
}

//Run is running
func (vehicle *Automovel) Run(velocity int, tipo string) (int, string) {

	return (velocity * 3), "teste"
}

// private static int (this Automovel vehicle, int velocity) {

// }
