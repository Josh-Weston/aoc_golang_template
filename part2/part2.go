package part2

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func Run(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if scanner.Err() != nil {
			fmt.Println(scanner.Err())
			os.Exit(1)
		}
		line := scanner.Text()
		fmt.Println(line)
	}
	return -1, errors.New("some error")
}
