package main

import (
	"caesar"
	"flag"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"strings"
)

func main() {
	//BEGIN BANNER
	underline := strings.Repeat("*", 128)
	banner := figure.NewFigure("Caesar Cipher", "cosmic", true)
	banner.Print()
	fmt.Println(underline)
	//END BANNER
	encoded := flag.String("c", "", "The encrypted text")
	shift := flag.Int("s", 99, "The shift value. Will do brute force by default") //either do brute force bool or put in 99 as a shift option

	flag.Parse()

	subencoded := *encoded //This creates a new variable that is a copy of the flag values. Allowing them to be
	subshift := *shift        //manipulated

	if subencoded == "" {
		fmt.Println("Please run 'caesar -h' for usage information")
		os.Exit(0)
	}

	if subshift == 99 {
		bruteForce(subencoded)
	} else {
		ciphertext := caesar.DecryptCiphertext(subencoded, subshift)
		fmt.Println(ciphertext)
	}

}

func bruteForce(encoded string) string {
	for i := 1; i < 26; i++ {
		ciphertext := caesar.DecryptCiphertext(encoded, i)	//bruteforce := flag.Bool("brute", false, "Brute force caesar option")

		fmt.Println(ciphertext)
	}

	fmt.Println(encoded)
	return encoded
}
