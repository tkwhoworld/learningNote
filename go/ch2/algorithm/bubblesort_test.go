package algorithm

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	values := []int{3,4,5,3,1}
	BubbleSort(values)
	if values[0] !=1 || values[1]!=3 || values[2] !=3 || values[3]!=4 || values[4]!=5 {
		t.Error("BubbleSort() failed.Got",values,"Excepted 1 3 3 4 5")
	}
}