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

func exclusiveChoiceConnector(sendChannel <- chan string) (string, string){
	msg := <- sendChannel
	//msg1 := msg + " multi choice"
	receiveChannel1 := make(chan string, 1)
	receiveChannel2 := make(chan string, 1)

	switch msg {
		case "Module B":
			receiveChannel1 <- msg
			receiveChannel2 <- "NULL"
        		//return <-receiveChannel1, <-receiveChannel2


		case "Module C":
			receiveChannel1 <- "NULL"
			receiveChannel2 <- msg
			//return <-receiveChannel1, <-receiveChannel2
		
		default:
			fmt.Printf("No matching path\n")
			break    
		

    	}
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


func main(){
	sendChannelModuleA := make(chan string, 1)
	go pythonCall("moduleA.py", sendChannelModuleA)
	
	recChannelA1,recChannelA2 := exclusiveChoiceConnector(sendChannelModuleA)

	if recChannelA2 == "NULL"{
		fmt.Println("test1", recChannelA1)
		sendChannelModuleBin := make(chan string, 1)
		sendChannelModuleBin <- recChannelA1
		execModule("moduleB.py")
		
	} else if recChannelA1 == "NULL"{
		fmt.Println("test2", recChannelA2)
		sendChannelModuleCin := make(chan string, 1)
		sendChannelModuleCin <- recChannelA1
		execModule("moduleC.py")
	}

	
	
/*

	outA1 := <- receiveChannelModuleA1
	outA2 := <- receiveChannelModuleA2
	fmt.Println(outA1)
	fmt.Println(outA2)


	sendChannelModuleB := make(chan string, 1)
	//modB := "moduleB.py atmmoB"
	go pythonCall("moduleB.py", outA, receiveChannelModuleA)
	go sequenceConnector(receiveChannelModuleA, sendChannelModuleB)
	outB := <- sendChannelModuleB
	fmt.Println(outB)



	receiveChannelModuleB := make(chan string, 1)
	modB := "moduleB.py " + outA
	go pythonCall(outA, receiveChannelModuleA)
	go sequenceConnector(receiveChannelModuleA, receiveChannelModuleB)
	fmt.Println(<- receiveChannelModuleB)
*/
}





