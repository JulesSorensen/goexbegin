package exercises

import (
	"fmt"
	"testing"
)

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		fmt.Println("🔴 2 should be even")
		t.Fail()
	}
	if IsEven(3) {
		fmt.Println("🔴 3 should not be even")
		t.Fail()
	}
}

func TestReverseArray(t *testing.T) {
	arr := []int{1, 2, 3}
	newArr := ReverseArray(arr)
	if len(newArr) < 3 || newArr[0] != 3 || newArr[1] != 2 || newArr[2] != 1 {
		fmt.Println("🔴 ReverseArray should reverse the array")
		t.Fail()
	}
}

func TestReverseString(t *testing.T) {
	str := "hello"
	newStr := ReverseString(str)
	if newStr != "olleh" {
		fmt.Println("🔴 ReverseString should reverse the string")
		t.Fail()
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		fmt.Println("🔴 Swap should swap the values of a and b")
		t.Fail()
	}
}
