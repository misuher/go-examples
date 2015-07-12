package myTest

//the golang standard library suppot testing
import (
	"reflect"
	"testing"
)

//the testing function have to start with Testxxxx
//and have to receive "*testing.T" as a parameter
func TestIntTest(t *testing.T) {
	result := IntTest()
	if result != 11 {
		t.Fail()                                           //tells that fails with no extra information
		t.Errorf("Expected 11 but got %d instead", result) //custom error
	}
}

//now you can go to the CLI and type
//	$go test
//the first time you run it make sure that the IntTest implementation
//make the test fails so you can be sure that when you implement it
//you don't get a false positive result

func TestBoolTest(t *testing.T) {
	result := BoolTest()
	if result != true {
		t.Errorf("Expexted true but got %t instead", result)
	}
}

//now we are goint to check the types instead the return values taking advantage of reflect package
func TestStringTest(t *testing.T) {
	result := StringTest()
	resultType := reflect.ValueOf(result)
	if resultType.Kind() != reflect.String {
		t.Errorf("Expected string but got %s instead", resultType.Kind())
	}
}
