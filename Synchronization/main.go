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


func sync(receiveChannelA <- chan string, receiveChannelB <- chan  string, sendChannelC chan <- string) {
	msgA := <- receiveChannelA
	msgB := <- receiveChannelB

	if len(msgA) != 0 && len(msgB) != 0 {	
		sendChannelC <- "Trigger C"		
	}else {
		fmt.Println("Both A and B modules should send their outputs")
	}
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
	receiveChannelModuleA := make(chan string, 1)
	receiveChannelModuleB := make(chan string, 1)
	sendChannelModuleC := make(chan string, 1)
	go pythonCall("moduleA.py", receiveChannelModuleA)
	go pythonCall("moduleB.py", receiveChannelModuleB)
	sync(receiveChannelModuleA, receiveChannelModuleB, sendChannelModuleC)
	sendChannelC := <-sendChannelModuleC
	if sendChannelC == "Trigger C"{
		//fmt.Println("test1", sendChannelC)
		receiveChannelModuleCin := make(chan string, 1)
		receiveChannelModuleCin <- sendChannelC
		execModule("moduleC.py")
		
	} else {
		fmt.Println("And didn't work")
	}	
	
}





