package main

import (
	"bufio"
	"fmt"
	"os"

	"gonum.org/v1/gonum/stat"
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
	fmt.Println("   ' 1  Add -⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 2. Subtract ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 3. Multiply ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 4. Divide ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 5. Remainder -⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 6. Standard Deviation ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
	fmt.Println("   ' 7. Variance ⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆-⋆|")
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
		multiply()
		break

	case '4':
		divide()
		break

	case '5':
		remainder()
		break

	case '6':
		standardDeviation()
		break

	case '7':
		variance()
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
	var y float64
	i := 0
	var sd []float64

	fmt.Println("How Many Numbers: ")
	fmt.Scan(&x)

	for i < x {

		fmt.Println("Number: ")
		fmt.Scan(&y)
		fmt.Println()
		sd = append(sd, y)
		i++
	}

	stdev := stat.StdDev(sd, nil)
	fmt.Printf("Standard Deviation: %.4f \n", stdev)
}

func variance() {

	var x int
	var y float64
	i := 0
	var va []float64

	fmt.Println("How Many Numbers: ")
	fmt.Scan(&x)

	for i < x {

		fmt.Println("Number: ")
		fmt.Scan(&y)
		fmt.Println()
		va = append(va, y)
		i++
	}

	vari := stat.Variance(va, nil)
	fmt.Printf("Variance: %.4f \n", vari)

}
