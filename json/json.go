package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	//take in account that to export this structs the fields must start with uppercase
	location struct {
		X int `json:"coord_x"`
		Y int `json:"coord_y"`
	}
	building struct {
		Address string   `json:"address"`
		Loc     location `json:"location"`
	}
)

func main() {
	//create data to encode
	data := &building{
		Address: "sample street",
		Loc: location{
			X: 17,
			Y: 67,
		},
	}

	//encode data
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	//to display jsondata we need to convert it to string
	fmt.Println("\nSimple Marshal:\n", string(jsonData))

	//we have a pretty print version with indentation
	jsonDataPretty, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nPrettyfy:\n", string(jsonDataPretty))

	//decode json
	data2 := &building{}
	if err := json.Unmarshal(jsonDataPretty, data2); err != nil {
		panic(err)
	}
	fmt.Println("\ndata2:\n", data2)
	fmt.Println("\ndata2.Address:\n", data2.Address)

	//we can choose the stream which we use to send the encode, like http response
	enc := json.NewEncoder(os.Stdout)
	fmt.Println("\nos.Stdout encoder:")
	enc.Encode(data)
}
