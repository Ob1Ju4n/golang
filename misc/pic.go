package misc

import "fmt"

// GetPic - Creates a unit8 matrix that will be interpreted by the
// Show function from golang.org/x/tour/pic package
func GetPic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)

	for i := range y {
		x := make([]uint8, dx)
		for j := range x {
			x[j] = uint8((i + j) / 2)
		}
		y[i] = x
	}

	for _, v := range y {
		fmt.Printf("%d ", v)
	}
	return y

}
