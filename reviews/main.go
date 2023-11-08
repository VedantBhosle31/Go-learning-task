package reviews

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Getreview() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Thanks for rating us %v stars", numrating)
	}

	fmt.Printf("Thanks for rating us %v", input)

}
