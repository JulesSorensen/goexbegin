package exercises

// IsEven returns true if n is even, false otherwise
func IsEven(n int) bool {
	return n%2 == 0
}

// ReverseArray returns a new array with the elements of arr in reverse order
func ReverseArray(arr []int) []int {
	var newArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		newArr = append(newArr, arr[i])
	}
	return newArr
}

// ReverseString returns a new string with the characters of s in reverse order
// ℹ️ use the string() method to convert a byte to a string
func ReverseString(s string) string {
	var newStr string
	for i := len(s) - 1; i >= 0; i-- {
		newStr += string(s[i])
	}
	return newStr
}

// swap swaps the values of a and b, so that a becomes b and b becomes a
func Swap(a, b *int) {
	*a, *b = *b, *a
	// or
	// tmp := *a
	// *a = *b
	// *b = tmp
}
