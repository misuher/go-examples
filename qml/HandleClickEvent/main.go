package main

import "gopkg.in/qml.v1"

type Arg struct{
	Data string
}

func (a *Arg)SetArg(){
	if(a.Data=="Click Detected!!! Click Me!!!"){
		a.Data = "Click Me again!!"
	}else{
		a.Data = "Click Detected!!!"
	}
    qml.Changed(a, &a.Data)
}

func run() error{
	//create the engine
	engine := qml.NewEngine()
	//subtitute vble values
	golang := Arg{Data: "Click me!!"}
	context := engine.Context()
	context.SetVar("golang", &golang)
	//load the qml description file, return the element as a qml.Object
    window, err := engine.LoadFile("main.qml")
    if err != nil {
            return err
    }
	text, err := engine.LoadFile("text.qml")
    if err != nil {
            return err
    }
	//draw them in the scene
    win := window.CreateWindow(nil)  //createWindow method for root window
	txt := text.Create(nil)	//create method for components
	//describe the relationship between qml.Objects
	txt.Set("parent", win.Root() )
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
