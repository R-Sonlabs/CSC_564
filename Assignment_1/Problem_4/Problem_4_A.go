/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 4, Implementation A - Building H2O
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use Go and any tools I wanted.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var hydrogenMutex sync.Mutex //For locking the hydrogenCounter, later for barring more hydrogen atoms from entry
var oxygenMutex sync.Mutex   //For locking out additional oxygen atoms
var bondMutex sync.Mutex     //For locking the bondCounter
var mainWait sync.WaitGroup  //To halt the main thread until all threads join
var hydrogenCounter int      //Incremented and decremented to track whether an atom is first or second in
var bondCounter int          //To check that all three atoms are present before bonding.
var bondNames string         //For debugging

func main() {
	bondNames = ""
	//Checking, getting, validating CL arguments for setting the number iterations the program will run for
	var iterS string
	if len(os.Args) > 1 {
		iterS = os.Args[1]
	} else {
		iterS = "1"
	}
	iter, err := strconv.Atoi(iterS)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	mainWaitCounter := iter * 3 //The mainWait will wait for ALL of the go routines to finish and call "Done" before allowing the program to exit.
	mainWait.Add(mainWaitCounter)

	bondWaitChan := make(chan bool)

	for i := 0; i < iter; i++ { //For each iteration specified by CL argument, spin up a go routine for one oxygen and two hydrogen
		go Oxygen(i, bondWaitChan)
		go Hydrogen(i, bondWaitChan)
		go Hydrogen(iter+i, bondWaitChan)
	}
	mainWait.Wait() //waiting for all routines to finish
}

func Oxygen(i int, bondWaitChan chan<- bool) {
	defer mainWait.Done()      //For joining later on the main thread
	defer oxygenMutex.Unlock() //This gets unlocked after the single oxygen thread returns from bonding
	oxygenMutex.Lock()         //Get the lock.  Bars other oxygen atoms from proceeding

	for hydrogenCounter != 2 { //Waits until there are two hydrogen atoms ready
	}
	bondWaitChan <- true //Signal the two waiting hydrogen atoms to enter Bond()
	bondWaitChan <- true
	Bond() //Proceed to Bond() with the freshly signalled hydrogen atoms.
}

func Hydrogen(i int, bondWaitChan <-chan bool) {
	defer mainWait.Done()     //For joining later on the main thread
	hydrogenMutex.Lock()      //Get the lock
	hydrogenCounter += 1      //Increment
	if hydrogenCounter == 1 { //Checking if first atom
		hydrogenMutex.Unlock() //Unlock the mutex so that the second atom can proceed
		<-bondWaitChan         //Wait on signal from Oxygen
		Bond()                 //Proceed to Bond()
	} else if hydrogenCounter == 2 { //Check if second atom
		<-bondWaitChan         //Wait on signal from Oxygen
		Bond()                 //Proceed to Bond()
		hydrogenCounter -= 2   //Reset the counter
		hydrogenMutex.Unlock() //Release the lock so that two more hydrogen can proceed
	}
}

func Bond() {
	//I thought more would happen here!  But not necessary for this example
}
