package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//tasks.Mutable()
	//tasks.MultiDimensional()
	//tasks.Matrix()
	//tasks.Slice()
	//tasks.WorkRange()
	//sayGood("good")
	/*var a string = "shakh"

	fmt.Println(a)
	a = "Salim"
	fmt.Println(a)

	slice := []string{"salim", "yaxyo", "Qodir"}
	slice[1] = "salim"
	fmt.Println(slice)*/

	//tasks.SayGoodMorning("Salim")
	/*tasks.Function([]string{"Ixti", "Shakh", "Abi"}, tasks.SayGdmorning)
	tasks.Function([]string{"Ixti", "Shakh", "Abi"}, tasks.SayBye)

	a1 := tasks.CycleArea(12.4)
	a2 := tasks.CycleArea(2)
	fmt.Println(a1, a2)
	fmt.Printf("%0.2f and %0.2f", a1, a2)*/
	/*a1, b1 := tasks.GetInitials("shakh abdul")
	fmt.Println(a1, b1)

	a2, b2 := tasks.GetInitials("salim lohix")
	fmt.Println(a2, b2)

	a3, b3 := tasks.GetInitials("shakh")
	fmt.Println(a3, b3)
	*/
	//tasks.MapSructure()
	//	phoneNumber := tasks.CreatePhoneNumber2([10]int{2, 4, 5, 6, 7, 8, 9, 7, 9, 0})
	//	fmt.Println(phoneNumber)

	//name := "shakha"

	//name = tasks.UpdateName(name)

	//fmt.Println(name)
	//profes := tasks.Salim("shakh")
	//fmt.Println(profes)
	/*tasks.CountSheeps([]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	})*/

	/*tasks.CountSheeps([]bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	})*/
	//InputFunc()

	mybill := CreateBill()
	fmt.Println(mybill)
	PromptOpt(mybill)

	//	res := IsUpper("SsSsSSSS")
	//	fmt.Print(res)

	/*s := "salom sevgi"
	fmt.Println(len(s))
	fmt.Print(s[0:2])*/
}

/*
func InputFunc() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type something: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %q\n", input)

}*/

type bill struct {
	name  string
	items map[string]float32
	tips  float64
}

func (b bill) UpdateTip(tip float64) {
	b.tips = float64(tip)
}

func (b bill) AddItem(name string, price float64) {
	b.items[name] = float32(price)

}

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func CreateBill() bill {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Create a new bill name: ")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	name, _ := GetInput("create new bill name: ", reader)
	b := NewBill(name)
	fmt.Println("Created new bill - ", b.name)
	return b
}

func NewBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float32{},
		tips:  0,
	}
	return b
}

func exitoption(b bill) {
	r := bufio.NewReader(os.Stdin)
	ExC, _ := GetInput("Do you want to continue y/n:", r)
	switch ExC {
	case "y":

		PromptOpt(b)
	case "n":
		os.Exit(0)
	}
}

func PromptOpt(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := GetInput("Choose option (a - add item, s - save bill, t - add tip):", reader)
	switch opt {
	case "a":
		name, _ := GetInput("Item name: ", reader)
		price, _ := GetInput("Item price($): ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be number")
			PromptOpt(b)
		}
		b.AddItem(name, p)

		fmt.Println("Irem are added - ", name, price)
		fmt.Println(b)
	case "t":
		tip, _ := GetInput("Enter tip amount ($)", reader)
		fmt.Println(tip)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be number")
			PromptOpt(b)
		}
		b.UpdateTip(t)
		fmt.Println("Tip added - ", tip)
		PromptOpt(b)

	case "s":
		fmt.Println("You chose s to save the bill: ", b)
	default:
		fmt.Println("that was not a valid option")
		PromptOpt(b)
	}
	exitoption(b)

}

//func PromptOptions(b bill) {
//reader := bufio.NewReader(os.Stdin)
//opt, _ := get
//}

/*func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}*/
