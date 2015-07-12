//Package myTest are a bunch of simple function to make test examples
//the function that test the ones included here have to be in a file named
//like this one but with "_test.go" at the end
package myTest

//IntTest is a funtion to make a test by its returned value
func IntTest() int {
	return 5 + 6
}

//BoolTest is just other example for a test
func BoolTest(a, b int) bool {
	return a == b
}

//StringTest is a function to make a test bt the returned type
func StringTest() string {
	return "string"
}

//SumTest is a function to make a demo of how to test several parameters at once
func SumTest(a, b int) int {
	return a + b
}
