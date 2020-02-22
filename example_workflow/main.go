package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
)


func pythonCall(progName string, argument string, sendChannel chan <- string) {
	cmd := exec.Command("python3", progName, argument)
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

func sequenceConnector(receiveChannel <- chan string, sendChannel chan <- string ){
	msg := <- receiveChannel
	sendChannel <- msg
}

func parallelSplitConnector(receiveChannel <- chan string, sendChannel1 chan <- string, sendChannel2 chan <- string ){
	msg := <- receiveChannel

	sendChannel1 <- msg
	sendChannel2 <- msg

}

func synchronousConnector(receiveChannelA <- chan string, receiveChannelB <- chan  string, sendChannelC chan <- string) {
	msgA := <- receiveChannelA
	msgB := <- receiveChannelB

	if len(msgA) != 0 && len(msgB) != 0 {	
		sendChannelC <- "Trigger C"		
	}else {
		fmt.Println("Both A and B modules should send their outputs")
	}
}

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

//in==receive
//out==send

func main(){
	//Sequence connector Demonstration
	sendChannelModuleA := make(chan string, 1)
	receiveChannelModuleA := make(chan string, 1)
	go pythonCall("moduleA.py", "Start", sendChannelModuleA)
	go sequenceConnector(sendChannelModuleA, receiveChannelModuleA)
	outA := <- receiveChannelModuleA
	
	receiveChannelModuleB := make(chan string, 1)
	go pythonCall("moduleB.py", outA, receiveChannelModuleA)
	go sequenceConnector(receiveChannelModuleA, receiveChannelModuleB)
	

	//Parallel Split Connector Demonstration
	sendChannelModuleCin := make(chan string, 1)
	sendChannelModuleDin := make(chan string, 1)
	go parallelSplitConnector(receiveChannelModuleB, sendChannelModuleCin, 								sendChannelModuleDin)
	recChannelA1 := <- sendChannelModuleCin
	recChannelA2 := <- sendChannelModuleDin
	go pythonCall("moduleC.py", recChannelA1, sendChannelModuleCin)
	go pythonCall("moduleD.py", recChannelA2, sendChannelModuleDin)

	//Synchronous Connector Demonstration
	sendChannelModuleE := make(chan string, 1)
	synchronousConnector(sendChannelModuleCin, sendChannelModuleDin, 								sendChannelModuleE)
	sendChannelE := <-sendChannelModuleE
	if sendChannelE == "Trigger C"{
		receiveChannelModuleEin := make(chan string, 1)
		receiveChannelModuleEin <- sendChannelE
		execModule("moduleE.py")
		
	} else {
		fmt.Println("And didn't work")
	}
}





