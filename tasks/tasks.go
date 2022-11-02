package tasks

import "fmt"

func TaskN1() {

	a := "Aziza"
	b := "Umida"
	d := "Nargiza"
	c := "Imrona"

	arr := [4]string{a, b, c, d}

	fmt.Println("Number of women in our family is: ", len(arr))
	fmt.Println("\n they are : ", arr)

}

func MultiDimensional() {

	fmt.Println("\n\n\n\t**** multidimensional arrays ****")
	fmt.Println("")
	fmt.Println("")

	l1 := [2][4]int{
		{2, 4, 5, 6},
		{3, 7, 8, 9},
	}

	fmt.Println(l1, len(l1))
	fmt.Println("")

	//	fmt.Println("\t\n***Matrix***")
	fmt.Println("")

	l1[1] = [4]int{21, 34, 4, 5}
	//l1[0] = [4]int{}
	fmt.Println(l1, len(l1))
	fmt.Println(l1[0][0])
}

func Matrix() {
	a := 0
	b := 0
	c := 0
	d := 0
	g := 0
	f := 0

	fmt.Println("\t\n***Enter 2x3 matrix elements here ***")

	fmt.Println("\nEnter 1-number:")
	fmt.Scanln(&a)

	fmt.Println("\nEnter 2-number:")
	fmt.Scanln(&b)

	fmt.Println("\nEnter 3-number:")
	fmt.Scanln(&c)

	fmt.Println("\nEnter 4-number:")
	fmt.Scanln(&d)

	fmt.Println("\nEnter 5-number:")
	fmt.Scanln(&g)

	fmt.Println("\nEnter 6-number:")
	fmt.Scanln(&f)

	matrix_1 := [2][3]int{
		{a, b, g},
		{c, d, f},
	}

	fmt.Println(matrix_1, len(matrix_1))
	//sum := [][]int{}
	/*for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {

			sum += matrix_1[i][j]

		}
	}
	fmt.Println("Sum is: ", sum)*/
}

func Slice() {

	arr := [4]string{
		"shakhzod", "Ali", "Abdullokh", "Salom",
	}

	scores := []int{
		21, 21, 4, 5,
	}

	scores[3] = 32

	scores = append(scores, 32)

	fmt.Println(scores, len(scores))
	rangeOne := arr[1:2]
	rangeOne2 := scores[2:4]
	fmt.Println(rangeOne, rangeOne2)

}

func WorkRange() {
	/*slice := []int{21, 22, 23}
	sum := 0
	for _, numSlice := range slice {
		sum += numSlice
	}
	fmt.Println(sum)
	*/
	Slices := []int{21, 44, 55, 6}
	for i := 0; i < len(Slices); i++ {
		fmt.Println(Slices[i])
	}
	fmt.Println("*************")
	for i, v := range Slices {
		fmt.Printf("\n %d: %d", i, v)
	}
}
