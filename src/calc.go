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
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	//infinite loop using for
	for {
		//show menu
		menu()
	}
}

//func outside of main
func menu() {

	//print menu
	fmt.Println(".----------------------------------------------------------.")
	fmt.Println("| (╯°□ °）╯︵ ┻━┻ Meggu's Menu__̴ı̴̴̡̡̡ ̡͌l̡̡̡ ̡͌l̡*̡̡ ̴̡ı̴̴̡ ̡̡͡|̲̲̲͡͡͡ ̲▫̲͡ ̲̲̲͡͡π̲̲͡͡ ̲̲͡▫̲̲͡͡ ̲|̡̡̡ ̡ ̴̡ı̴̡̡ ̡͌l̡̡̡̡.___|")
	fmt.Println("'----------------------------------------------------------|")
	fmt.Println("' 1. Add --------------------------------------------------|")
	fmt.Println("' 2. Subtract----------------------------------------------|")
	fmt.Println("' 3. Multiply------------------------------------ ---------|")
	fmt.Println("' 4. Divide -----------------------------------------------|")
	fmt.Println("' 5. Remainder --------------------------------------------|")
	fmt.Println("' 6. Standard Deviation------------------------------------|")
	fmt.Println("' 7. Mean, Median, Mode -----------------------------------|")
	fmt.Println("' 8. Exit -------------------------------------------------|")
	fmt.Println("'----------------------------------------------------------'")

	//wait for output
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	//error statement
	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case '1':
		fmt.Println("Add")
		break
	case '2':
		fmt.Println("Subtract")
		break

	case '3':
		fmt.Println("Multiply")
		break

	case '4':
		fmt.Println("Divide")
		break

	case '5':
		fmt.Println("Remainder")
		break

	case '6':
		fmt.Println("Standard Deviation")
		break

	case '7':
		fmt.Println("Mean, Median, Mode")
		break

	case '8':
		os.Exit(3)
	}

}
