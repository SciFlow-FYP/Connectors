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

func sequenceConnector(sendChannel <- chan string, receiveChannel chan <- string ){
	msg := <- sendChannel
	msg1 := msg + " sequence"
	receiveChannel <- msg1
}

//in==send
//out==recieve


func main(){
	sendChannelModuleA := make(chan string, 1)
	receiveChannelModuleA := make(chan string, 1)
	go pythonCall("moduleA.py", "", sendChannelModuleA)
	go sequenceConnector(sendChannelModuleA, receiveChannelModuleA)
	outA := <- receiveChannelModuleA
	fmt.Println(outA)

	receiveChannelModuleB := make(chan string, 1)
	//modB := "moduleB.py atmmoB"
	go pythonCall("moduleB.py", outA, receiveChannelModuleA)
	go sequenceConnector(receiveChannelModuleA, receiveChannelModuleB)
	outB := <- receiveChannelModuleB
	fmt.Println(outB)

/*

	receiveChannelModuleB := make(chan string, 1)
	modB := "moduleB.py " + outA
	go pythonCall(outA, receiveChannelModuleA)
	go sequenceConnector(receiveChannelModuleA, receiveChannelModuleB)
	fmt.Println(<- receiveChannelModuleB)
*/
}





