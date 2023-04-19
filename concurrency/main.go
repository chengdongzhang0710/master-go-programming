package main

// func sayHello(n string, w *sync.WaitGroup) {
// 	fmt.Printf("Hello, %s!\n", n)
// 	w.Done()
// }

// func sum(a, b float64, w *sync.WaitGroup) {
// 	fmt.Printf("The sum of two float number is %.2f\n", a+b)
// 	w.Done()
// }

// func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	defer m.Unlock()
// 	*b += n
// 	wg.Done()
// }

// func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock()
// 	defer m.Unlock()
// 	*b -= n
// 	wg.Done()
// }

// func power(n int, ch chan int) {
// 	ch <- n * n
// }

func main() {
	// fmt.Println("No. of CPUs:", runtime.NumCPU())
	// fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	// fmt.Println("OS:", runtime.GOOS)
	// fmt.Println("Arch:", runtime.GOARCH)
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	// const gr = 100
	// var wg sync.WaitGroup
	// wg.Add(2 * gr)
	// n := 0
	// var m sync.Mutex
	// for i := 0; i < gr; i++ {
	// 	go func(wg *sync.WaitGroup) {
	// 		time.Sleep(time.Second / 10)
	// 		m.Lock()
	// 		defer m.Unlock()
	// 		n++
	// 		wg.Done()
	// 	}(&wg)
	// 	go func(wg *sync.WaitGroup) {
	// 		time.Sleep(time.Second / 10)
	// 		m.Lock()
	// 		defer m.Unlock()
	// 		n--
	// 		wg.Done()
	// 	}(&wg)
	// }
	// wg.Wait()
	// fmt.Println("The final value of var n is", n)

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go sayHello("Mr. Wick", &wg)
	// wg.Wait()

	// var wg sync.WaitGroup
	// wg.Add(3)
	// for i := 0; i < 3; i++ {
	// 	go sum(.1*float64(i), .2*float64(i), &wg)
	// }
	// wg.Wait()

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(x float64, wg *sync.WaitGroup) {
	// 	fmt.Println(math.Sqrt(x))
	// 	wg.Done()
	// }(.04, &wg)
	// wg.Wait()

	// var wg sync.WaitGroup
	// wg.Add(50)
	// start := time.Now().UnixNano() / 1_000_000
	// for i := 100; i <= 149; i++ {
	// 	go func(x float64, wg *sync.WaitGroup) {
	// 		time.Sleep(time.Second * 2)
	// 		fmt.Println(math.Sqrt(x))
	// 		wg.Done()
	// 	}(float64(i), &wg)
	// }
	// wg.Wait()
	// end := time.Now().UnixNano() / 1_000_000
	// fmt.Println(end - start)

	// var wg sync.WaitGroup
	// var m sync.Mutex
	// wg.Add(200)
	// balance := 100
	// for i := 0; i < 100; i++ {
	// 	go deposit(&balance, i, &wg, &m)
	// 	go withdraw(&balance, i, &wg, &m)
	// }
	// wg.Wait()
	// fmt.Println("Final balance value:", balance)

	// var c1 chan float64
	// c2 := make(<-chan rune)
	// c3 := make(chan<- rune)
	// c4 := make(chan int, 10)
	// fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)

	// ch := make(chan string)
	// defer close(ch)
	// go func(s string, ch chan string) {
	// 	ch <- s
	// }("Hello!", ch)
	// fmt.Println(<-ch)

	// c := make(chan int)
	// defer close(c)
	// go func(n int) {
	// 	c <- n
	// }(100)
	// fmt.Println(<-c)

	// ch := make(chan int)
	// defer close(ch)
	// for i := 1; i <= 50; i++ {
	// 	go power(i, ch)
	// 	fmt.Println(<-ch)
	// }

	// ch := make(chan int)
	// defer close(ch)
	// for i := 1; i <= 50; i++ {
	// 	go func(n int, ch chan int) {
	// 		ch <- n * n
	// 	}(i, ch)
	// 	fmt.Println(<-ch)
	// }
}
