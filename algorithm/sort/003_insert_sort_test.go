package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{1, 2, 3, 5, 2, 6, 13, 3}
	fmt.Println(InsertSort(arr))
}
