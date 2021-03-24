package sort

import (
	"fmt"
	"testing"
)

func test(i [5]int) {
	i[0] = 19
}

func TestInsertSort(t *testing.T) {
	list := [5]int{1, 2, 3, 4, 5}
	arr := list[:]
	test(list)
	fmt.Println(arr)
	fmt.Println(list)
}
