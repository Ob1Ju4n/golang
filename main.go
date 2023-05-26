package main

import (
	"fmt"

	misc "github.com/ob1ju4n/golang/misc"
)

func main() {
	fmt.Println(misc.SayHello())
	v, err := misc.Sqrt(2500)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("Sqrt: %v\n", v)
	}
	// pic.Show(misc.GetPic) // picture is visible on a web page
	// misc.getPic(5,4) // local test
	// wc.Test(misc.WordCount)
	// misc.WordCount("Disney Disney Disney for the win") // local test
	// misc.PrintFibonacciNthTerm(15)
	hosts := map[string]misc.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
