package main

import "fmt"

func main() {
	// sliceExample()
	// appendExample()
	copyExample()	
}

func sliceExample() {
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	println("Slice x:", x)
	fmt.Println("Length of slice x:", len(x))
	fmt.Println("Capacity of slice x:", cap(x))

	y := arr[1:4]
	println("Slice y:", y)
	fmt.Println("Length of slice y:", len(y))
	fmt.Println("Capacity of slice y:", cap(y))

	z := arr[:3]
	println("Slice z:", z)
	fmt.Println("Length of slice z:", len(z))
	fmt.Println("Capacity of slice z:", cap(z))

	w := arr[2:]
	println("Slice w:", w)
	fmt.Println("Length of slice w:", len(w))
	fmt.Println("Capacity of slice w:", cap(w))

	var yArr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array yArr:", yArr)
	fmt.Println("Length of array yArr:", len(yArr))

	zArr := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array zArr:", zArr)
	fmt.Println("Length of array zArr:", len(zArr))

}

func appendExample() {
	slice1 := []int{1, 2, 3}
	fmt.Println("Slice1 before append:", slice1)

	slice2 := append(slice1, 4, 5)
	fmt.Println("Slice2 after append:", slice2)

	slice3 := append(slice2, []int{6, 7, 8}...)
	fmt.Println("Slice3 after appending another slice:", slice3)

}

func copyExample() {
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, 3)

	numCopied := copy(destination, source)
	fmt.Println("Number of elements copied:", numCopied)
	fmt.Println("Destination slice after copy:", destination)
}
