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

func messagePassing(sendChannel <- chan string, receiveChannel chan <- string ){
	msg := <- sendChannel
	msg1 := msg + " sequence"
	receiveChannel <- msg1
}

//in==send
//out==recieve

func sequenceConnector(program1 string, program2 string){
	sendChannelModuleA := make(chan string, 1)
	receiveChannelModuleA := make(chan string, 1)
	go pythonCall(program1, "", sendChannelModuleA)
	go messagePassing(sendChannelModuleA, receiveChannelModuleA)
	outA := <- receiveChannelModuleA
	fmt.Println(outA)

	sendChannelModuleB := make(chan string, 1)
	go pythonCall(program2, outA, receiveChannelModuleA)
	go messagePassing(receiveChannelModuleA, sendChannelModuleB)
	outB := <- sendChannelModuleB
	fmt.Println(outB)
}


func main(){

	sequenceConnector("moduleA.py", "moduleB.py")

}





