package main

// type car struct {
// 	brand string
// 	price int
// }

// func (c *car) changeCar(newBrand string, newPrice int) {
// 	c.brand = newBrand
// 	c.price = newPrice
// }

// type money float64

// func (m money) print() {
// 	fmt.Printf("%.2f\n", m)
// }

// func (m money) printStr() string {
// 	return fmt.Sprintf("%.2f", m)
// }

// type book struct {
// 	price float64
// 	title string
// }

// func (b book) vat() float64 {
// 	return b.price * .09
// }

// func (b *book) discount() {
// 	b.price *= .9
// }

// func (b *book) changePrice(p float64) {
// 	b.price = p
// }

// type vehicle interface {
// 	License() string
// 	Name() string
// }

// type car struct {
// 	licenseNo string
// 	brand     string
// }

// func (c car) License() string {
// 	return fmt.Sprintln("The license of the car:", c.licenseNo)
// }

// func (c car) Name() string {
// 	return fmt.Sprintln("The name of the car:", c.brand)
// }

// type cube struct {
// 	edge float64
// }

// func volume(c cube) float64 {
// 	return math.Pow(c.edge, 3)
// }

func main() {
	// myCar := car{brand: "Audi", price: 40000}
	// fmt.Println(myCar)
	// myCar.changeCar("BMW", 45000)
	// fmt.Println(myCar)

	// var m map[int]float64
	// fmt.Printf("var m is of type %T and is of value %#v\n", m, m)
	// fmt.Println(m)

	// var empty interface{}
	// empty = 5
	// fmt.Println(empty)
	// empty = "Go"
	// fmt.Println(empty)
	// empty = []int{1, 2, 3}
	// fmt.Println(empty)
	// fmt.Println(len(empty.([]int)))

	// var myMoney money = 3.14159
	// myMoney.print()
	// fmt.Println(myMoney.printStr())

	// myBook := book{price: 15.4, title: "The Go Programming Language"}
	// myBook.discount()
	// fmt.Println("The vat value of the book afer applying discount is", myBook.vat())

	// bestBook := book{title: "The trial by Kafka", price: 9.9}
	// bestBook.changePrice(11.99)
	// fmt.Printf("Book's price: %.2f\n", bestBook.price)

	// var myVehicle vehicle = car{licenseNo: "NO001", brand: "BMW"}
	// fmt.Printf(myVehicle.License())
	// fmt.Printf(myVehicle.Name())

	// var empty interface{}
	// fmt.Printf("var empty is of type %T and is of value %#v\n", empty, empty)
	// empty = 5
	// fmt.Printf("var empty is now of type %T\n", empty)
	// empty = .5
	// fmt.Printf("var empty is now of type %T\n", empty)
	// empty = []int{1, 2, 3}
	// fmt.Printf("var empty is now of type %T\n", empty)
	// empty = append(empty.([]int), 4)
	// fmt.Println(empty)

	// var x interface{}
	// x = cube{edge: 5}
	// v := volume(x.(cube))
	// fmt.Printf("Cube Volume: %v\n", v)
}
