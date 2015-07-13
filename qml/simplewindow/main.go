package main

import "gopkg.in/qml.v1"

func run() error{
	//create the engine
	engine := qml.NewEngine()
	//load the qml description file, return the element as a qml.Object
    component, err := engine.LoadFile("main.qml")
    if err != nil {
            return err
    }
	//draw it in the scene
    win := component.CreateWindow(nil)
	//show the scene
    win.Show()
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
