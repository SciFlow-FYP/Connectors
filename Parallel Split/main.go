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

func messagePassing(sendChannel <- chan string) (string, string){
	msg := <- sendChannel
	//msg1 := msg + " multi choice"
	receiveChannel1 := make(chan string, 1)
	receiveChannel2 := make(chan string, 1)

	receiveChannel1 <- msg
	receiveChannel2 <- msg

	return <-receiveChannel1, <-receiveChannel2
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


func parallelSplitConnector(program1 string, program2 string, program3 string){
	sendChannelModuleA := make(chan string, 1)
	go pythonCall(program1, sendChannelModuleA)
	
	recChannelA1,recChannelA2 := messagePassing(sendChannelModuleA)

	fmt.Println("test3", recChannelA1)
	
	sendChannelModuleBin := make(chan string, 1)
	sendChannelModuleBin <- recChannelA1
	execModule(program2)
	
	fmt.Println("test4", recChannelA2)

	sendChannelModuleCin := make(chan string, 1)
	sendChannelModuleCin <- recChannelA2
	execModule(program3)
}

func main(){

	parallelSplitConnector("moduleA.py","moduleB.py","moduleC.py")

}





