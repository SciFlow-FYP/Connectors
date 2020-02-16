package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"strconv"
)


func pythonCall(progName string, sendChannel chan <- string, itr string) {
	cmd := exec.Command("python3", progName, itr)
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


func loop(program string){
	for i:=1; i<=10; i++ {
		sendChannel := make(chan string, 1)
		receiveChannel := make(chan string, 1)
		pythonCall(program, sendChannel, strconv.Itoa(i))
		msg := <- sendChannel 	
		receiveChannel <- msg
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
	
	loop("module.py")

}





