package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	vecgen "tutorial/vec_gen"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the amount of loops done to fill the queue:")
	str, error := reader.ReadString('\n')
	if error != nil {
		panic("Error reading input, try again!")
	}
	loops, error := strconv.Atoi(str[:len(str)-2])
	if error != nil {
		panic("Error converting input to an int value!")
	}
	q := vecgen.NewQueue()
	for i := 0; i < loops; i++ {
		q.Append()
	}

	for i := 0; i < loops; i++ {
		fmt.Println(q.Step)
		q = *q.Node
	}
}
