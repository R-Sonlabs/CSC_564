/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 6, Implementation A - The Alarm Responders Problem
	Uvic, Fall 2018

	This problem was requested to be unique by Prof. Coady.
	I chose to fix a problem from a decade old piece of code I wrote.
	It is detailed on page 48 of my assignment paper.
*/

package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var lock sync.Mutex //Our one and only mutex. It is only used to put the writing of the new code into a critical section.

func main() {
	//Parsing command line arguments
	var officersS string
	var keysS string
	var speedS string

	if len(os.Args) > 3 {
		officersS = os.Args[1]
		keysS = os.Args[2]
		speedS = os.Args[3]
	} else {
		officersS = "3"
		keysS = "100"
		speedS = "1"
	}
	officers, err := strconv.Atoi(officersS)
	if err != nil {
		println("Arguments incorrect")
		os.Exit(2)
	}
	keyNum, err := strconv.Atoi(keysS)
	if err != nil {
		println("Arguments incorrect")
		os.Exit(2)
	}
	speed, err := strconv.Atoi(speedS)
	if err != nil {
		println("Arguments incorrect")
		os.Exit(2)
	}

	//Generating our key "Database".  There are just the keys with their ID numbers that each officer has.
	keys := make([]string, keyNum)
	for i := 0; i < keyNum; i++ {
		keys[i] = "Key" + strconv.Itoa(i)
	}

	//Our channels!
	keyChan := make(chan string) //This is the channel that each officer uses to send their key request to the controller.
	memChan := make(chan *Code)  //This is the channel that each officer sends their memory pointer with so that the controller can populate it with the code.

	//Generating the CODED key "Database". The first codes are randomnly selected here. Only the controller's coder threads can access this.
	keysCoded := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		keysCoded[i] = keys[i] + ":" + strconv.Itoa(rand.Intn(99999-10000)+10000)
	}

	go Controller(keysCoded, keyChan, memChan) //Start our controller with the coded "Database" and the required channels.

	for i := 0; i < officers; i++ { //Spin up our officers, they get a unique int for their name, a sleep multiplier, their own copy of the uncoded "Database", and the channels
		go Officer(i, speed, keys, keyChan, memChan)
	}

	input := bufio.NewScanner(os.Stdin)
	input.Scan() //This gives us a keyboard interrupt.  Any key will halt the loops and exit.

}

type Code struct { //A struct for our memory block that each officer reserves and sends a pointer to along with their key request
	code int
}

func Controller(keysCoded []string, keyChan chan string, memChan chan *Code) { //The controller just listens for a key request, then farms it out to it's own thread for processing.
	for {
		key := <-keyChan

		go Coder(keysCoded, keyChan, memChan, key)

	}
}

func Coder(keysCoded []string, keyChan chan string, memChan chan *Code, key string) {
	rand.Seed(time.Now().UTC().UnixNano())       //Seed for geneting new code later.
	keyNumParsed := strings.Split(key, "Key")[1] //Parse the request for the actual key number
	memPointer := <-memChan                      //Get the address of the requestors memory block
	keyNum, err := strconv.Atoi(keyNumParsed)    //Further parsing of the key number
	if err != nil {                              //Exception catch
		println("Key parse failed")
		os.Exit(2)
	}
	keyCodeParsed := strings.Split(keysCoded[keyNum], ":")[1] //Get the coded pair for the requested key and parse out the code
	keyCode, err := strconv.Atoi(keyCodeParsed)               //Convert code string to int.  Not strictly neccessary for this demonstration.
	if err != nil {                                           //Exception catch
		println("Code parse failed")
		os.Exit(2)
	}
	memPointer.code = keyCode                                                                           //Write the code to the requestors memory
	lock.Lock()                                                                                         //Lock the coded key "Database" for writing of the new randomn code
	keysCoded[keyNum] = "Key" + strconv.Itoa(keyNum) + ":" + strconv.Itoa(rand.Intn(99999-10000)+10000) //Write the new random code back in to the "Database" for the next user
	lock.Unlock()                                                                                       //What you lock, you must unlock, naturally.
}

func Officer(name int, speed int, keys []string, keyChan chan string, memChan chan *Code) {
	rand.Seed(time.Now().UTC().UnixNano()) //Seed for random sleep
	codeMem := Code{}                      //Initialize this officer's memory block.
	for {
		rndSleep := rand.Intn(1000) * speed                                                           //Randomn sleep for creating randomn interleaving.  Each iteration of the loop has a new randomn sleep.  The multiplier 'speed' can be set from the command line (1 default if there are not arguments) using 0 as an argument results in no sleeps.
		time.Sleep(time.Duration(rndSleep) * time.Millisecond)                                        //Zzzz...
		rnd := rand.Intn(len(keys))                                                                   //Get a randomn key
		keyChan <- keys[rnd]                                                                          //Send it to the controller
		memChan <- &codeMem                                                                           //Send pointer to this officer's memory block
		println("Officer ", name, " requested ", keys[rnd], " and got code ", codeMem.code, " back.") //Let the humans know what happened!
	}
}
