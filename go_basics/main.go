package main

func main() {
	// typed and untyped const

	// const x = 5
	// const y = 2.2 * x
	// var z float64 = x
	// fmt.Println(x, y, z)
	// fmt.Printf("const x is of type %T, const y is of type %T, var z is of type %T\n", x, y, z)

	// map definition

	// map1 := map[string]float64{
	// 	"USD": 152.,
	// 	"ERO": .254,
	// }
	// _ = map1

	// default types

	// a := 5
	// b := 3.4
	// fmt.Printf("var a is of type %T, var b is of type %T\n", a, b)

	// type convertion

	// s := string(99)
	// fmt.Printf("var s is of type %T and of value %v", s, s)

	// a := 223344
	// var str = fmt.Sprintf("%d", a)
	// fmt.Printf("var str is of type %T\n", str)

	// var result, err = strconv.ParseFloat("3.142", 64)
	// _ = err
	// fmt.Printf("result is of type %T and of value %v\n", result, result)

	// i, err := strconv.Atoi("-50")
	// s := strconv.Itoa(20)
	// _ = err
	// fmt.Printf("i Type is %T, i value is %v\n", i, i)
	// fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// Declare Variables Exercise

	// var (
	// 	a int
	// 	b float64
	// 	c bool
	// 	d string
	// )
	// _, _, _, _ = a, b, c, d

	// x, y, z := 20, 15.5, "Gopher!"
	// fmt.Println(x, y, z)
	// fmt.Printf("var y is of type %T\n", y)

	// var a float64 = 7.1
	// x, y := true, 3.7
	// a, x = 5.5, false
	// _, _, _ = a, x, y

	// name := "Golang"
	// fmt.Println(name)

	// Constant Exercise

	// const daysWeek = 7
	// const lightSpeed = 299792458
	// const pi = 3.14159
	// fmt.Printf("const daysWeek is of type %T, const lightSpeed is of type %T, const pi is of type %T\n", daysWeek, lightSpeed, pi)

	// const (
	// 	daysWeek   = 7
	// 	lightSpeed = 299792458
	// 	pi         = 3.14159
	// )

	// const daysWeek, lightSpeed, pi = 7, 299792458, 3.14159
	// fmt.Println(daysWeek, lightSpeed, pi)

	// const (
	// 	secPerDay int = 60 * 60 * 24
	// 	daysYear  int = 365
	// )
	// fmt.Printf("There are %d seconds in a year.\n", secPerDay*daysYear)

	// m := []int{1: 3, 4: 5, 6: 8}
	// fmt.Printf("var m is of type %T\n", m)
	// fmt.Println(m[1], m[4], m[6])

	// const (
	// 	Jun = iota + 6
	// 	Jul
	// 	Aug
	// )
	// fmt.Println(Jun, Jul, Aug)

	// Package fmt Exercise

	// x, y, z := 10, 15.5, "Gophers"
	// score := []int{10, 20, 30}
	// fmt.Printf("var x is of value %d, var y is of value %f, var z is of value %s\n", x, y, z)
	// fmt.Printf("var score is of value %#v\n", score)
	// fmt.Printf("var z is of value %q\n", z)
	// fmt.Printf("var x is of value %v, var y is of value %v, var z is of value %v\n", x, y, z)
	// fmt.Printf("var x is of value %#v, var y is of value %#v, var z is of value %#v\n", x, y, z)
	// fmt.Printf("var score is of value %v\n", score)
	// fmt.Printf("var y is of type %T, var score is of type %T\n", y, score)

	// const x float64 = 1.422349587101
	// fmt.Printf("const x is of value %.4f\n", x)

	// shape := "circle"
	// radius := 3.2
	// const pi float64 = 3.14159
	// circumference := 2 * pi * radius
	// fmt.Printf("Shape: %q\n", shape)
	// fmt.Printf("Circle's circumference with radius %v is %v\n", radius, circumference)

	// Operators and Conversions Exercise

	// var i int = 3
	// var f float64 = 3.2
	// fmt.Printf("The type of var x is %T\n", float64(i))
	// fmt.Printf("The value of var x is %f\n", float64(i))
	// fmt.Printf("The type of var f is %T\n", int(f))
	// fmt.Printf("The type of var f is %d\n", int(f))

	// x, y := 4, 5.1
	// z := float64(x) * y
	// fmt.Println(z)
	// a := 5
	// b := 6.2 * float64(a)
	// fmt.Println(b)

	// distance := 149.6 * 1000000 * 1000
	// speed := 299_792_458
	// time := distance / float64(speed)
	// fmt.Println(time)

	// x, y := 0.1, 5
	// var z float64
	// result1 := x < z && int(x) != int(z)
	// fmt.Println(result1)
	// result2 := y == 1*5 && int(z) == 0
	// fmt.Println(result2)

	// Named Types and Aliases Exercise

	// type duration int
	// var hour duration
	// fmt.Printf("var hour is of value %v and of type %T\n", hour, hour)
	// hour = 3600
	// fmt.Printf("var hour is now of value %v\n", hour)

	// type duration int
	// var hour duration = 3600
	// minute := 60
	// fmt.Println(int(hour) != minute)

	// type mile float64
	// type kilometer float64
	// const m2km = 1.609
	// var mileBerlinToParis mile = 655.3
	// var kmBerlinToParis kilometer
	// kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	// fmt.Println(kmBerlinToParis)
}
