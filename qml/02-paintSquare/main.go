package main

import "gopkg.in/qml.v1"

func run() error{
	//create the engine
	engine := qml.NewEngine()
	//load the qml description file, return the element as a qml.Object
    window, err := engine.LoadFile("main.qml")
    if err != nil {
            return err
    }
	square, err := engine.LoadFile("square.qml")
    if err != nil {
            return err
    }
	//draw them in the scene
    win := window.CreateWindow(nil)  //createWindow method for root window
	sq := square.Create(nil)	//create method for components
	//describe the relationship between qml.Objects
	sq.Set("parent", win.Root() )
	//show the scene painting the root qml.Object
    win.Show()
	//wait for events
    win.Wait()
    return nil
}

func main() {
	//execute the qml engine
	err := qml.Run(run)
	if err!=nil{
		panic(err)
	}
}
