package main

import "os"
import "fmt"

//@Author Bartosz Banasik
// <file path> 'talking' 'waiting'

func start() {
	//var model_ptr *Simulation_Model
	if(len(os.Args)>1){
		fmt.Println("init path existed")
	//	model_ptr = getModel(os.Args[1]);
		getModel(os.Args[1])
	} else {
		fmt.Println("init file not existed")
	}
}