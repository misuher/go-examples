package interfaceType

type customType struct {
	i int
	s string
}

func example(a interface{}) {
	//first way:
	switch a {
	case a.(string):
		//do something
	case a.(int32):
		//do something
	case a.(customType):
		//do something
	}

	//second way
	switch t := a.(type) {
	case string:
		//do something
	case int32:
		//do something
	case customType:
		//do something
	}
}
