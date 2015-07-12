//First of all notice that the test file is named "xxxxxx_test.go"
//so this is not where the test take place
package myTest

//we are going to define some weird function
//in order to test them based on the types they return
func IntTest() int {
	return 5 + 6
}

func BoolTest() bool {
	return true
}

func StringTest() string {
	return "string"
}
