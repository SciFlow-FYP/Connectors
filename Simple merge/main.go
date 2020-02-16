package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
)


func pythonCall(progName string, sendChannel chan <- string) {
	cmd := exec.Command("python3", progName)
	out, err := cmd.CombinedOutput()
	log.Println(cmd.Run())

	if err != nil {
		fmt.Println(err)
		// Exit with status 3.
    		os.Exit(3)
	}
	
	fmt.Println(string(out))
	msg := string(out)[:len(out)-1]
	sendChannel <- msg
}


func sequenceConnector(sendChannel <- chan string, receiveChannel chan <- string ){
	msg := <- sendChannel
	msg1 := msg + " sequence"
	receiveChannel <- msg1
}

//in==send
//out==recieve

func execModule(progName string){
	cmd := exec.Command("python3", progName)
	out, err := cmd.CombinedOutput()
	log.Println(cmd.Run())

	if err != nil {
		fmt.Println(err)
		// Exit with status 3.
    		os.Exit(3)
	}

	fmt.Println(string(out))
}


func main(){

	program := "moduleA.py" 
	receiveChannelModuleA := make(chan string, 1)
	receiveChannelModuleB := make(chan string, 1)
	sendChannelModuleC := make(chan string, 1)
	
	
	
	if program == "moduleA.py"{
		pythonCall(program, receiveChannelModuleA)	
		sequenceConnector(receiveChannelModuleA, sendChannelModuleC)
		
	} else if program == "moduleB.py"{
		pythonCall(program, receiveChannelModuleB)
		sequenceConnector(receiveChannelModuleB, sendChannelModuleC)
	} else {
		fmt.Println("No such module!")
	}

	fmt.Println(<-sendChannelModuleC)
	
	receiveChannelModuleC := make(chan string, 1)
	go pythonCall("moduleC.py", sendChannelModuleC)
	go sequenceConnector(sendChannelModuleC, receiveChannelModuleC)
	outC := <- receiveChannelModuleC
	fmt.Println(outC)
}





