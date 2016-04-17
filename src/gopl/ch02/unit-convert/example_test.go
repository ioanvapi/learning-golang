package unit_convert

import (
	"fmt"
	"os"
)

func ExampleConvert() {
	os.Args = []string{"1", "1.1", "1.2"}
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			process(arg)
		}
		return
	}

	fmt.Println("Input number. Ctrl-C to quit")

	for {
		var arg string
		_, err := fmt.Scanf("%s", &arg)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		process(arg)
	}
}
