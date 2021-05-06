package dataStructures

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCreation(t *testing.T) {
	merkle := InitMerkle()
	var list []string
	for i := 0; i < 5; i++ {
		list = append(list, strconv.Itoa(i))
	}
	merkle.FullTree(list)
	fmt.Println(merkle.Root.izquierda.izquierda.izquierda)
}
