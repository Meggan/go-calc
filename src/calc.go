package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println(".----------------------------------------------------------------.")
	fmt.Println("| (╯°□ °）╯︵ ┻━┻ Meggu's Calculator__̴ı̴̴̡̡̡ ̡͌l̡̡̡ ̡͌l̡*̡̡ ̴̡ı̴̴̡ ̡̡͡|̲̲̲͡͡͡ ̲▫̲͡ ̲̲̲͡͡π̲̲͡͡ ̲̲͡▫̲̲͡͡ ̲|̡̡̡ ̡ ̴̡ı̴̡̡ ̡͌l̡̡̡̡.___|")
	fmt.Println("'----------------------------------------------------------------'")

	//infinite loop using for
	for {
		//show menu
		menu()
	}
}

//func outside of main
func menu() {

	//print menu
	fmt.Println("   .----------------------------------------------------------.")
	fmt.Println("   | (╯°□ °）╯︵ ┻━┻ Meggu's Menu__̴ı̴̴̡̡̡ ̡͌l̡̡̡ ̡͌l̡*̡̡ ̴̡ı̴̴̡ ̡̡͡|̲̲̲͡͡͡ ̲▫̲͡ ̲̲̲͡͡π̲̲͡͡ ̲̲͡▫̲̲͡͡ ̲|̡̡̡ ̡ ̴̡ı̴̡̡ ̡͌l̡̡̡̡.___|")
	fmt.Println("   '----------------------------------------------------------|")
	fmt.Println("   ' 1 Add ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 2. Subtract ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 3. Multiply ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 4. Divide ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 5. Remainder -⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 6. Standard Deviation ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 7. Mean, Median, Mode ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 8. Exit ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   '----------------------------------------------------------'")

	fmt.Println()
	fmt.Println("Pick an option, desu")
	//wait for output
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	//error statement
	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case '1':
		add()
		break
	case '2':
		subtract()
		break

	case '3':
		fmt.Println("Multiply")
		multiply()
		break

	case '4':
		fmt.Println("Divide")
		divide()
		break

	case '5':
		fmt.Println("Remainder")
		remainder()
		break

	case '6':
		fmt.Println("Standard Deviation")
		// standardDeviation()
		break

	case '7':
		fmt.Println("Mean, Median, Mode")
		// mmm()
		break

	case '8':
		os.Exit(3)
	}

}

func add() {

	var x float64
	var y float64

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Sum : %f\n\n\n", x+y)
}

func subtract() {

	var x float64
	var y float64

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Result : %f\n\n\n", x-y)
}

func multiply() {

	var x float64
	var y float64

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Result : %f\n\n\n", x*y)
}

func divide() {

	var x float64
	var y float64

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Division : %f\n\n\n", x/y)
}

func remainder() {

	var x int
	var y int

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Remainder : %d\n\n\n", x%y)
}

func standardDeviation() {

	var x int
	var y int

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Sum : %d\n\n\n", x-y)
}

func mmm() {

	var x int
	var y int

	fmt.Println("First Number: ")
	fmt.Scan(&x)

	fmt.Println("Second Number: ")
	fmt.Scan(&y)
	fmt.Println()
	fmt.Printf("Sum : %d\n\n\n", x-y)
}
