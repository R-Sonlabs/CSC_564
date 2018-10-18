/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 5, Implementation A - Cigarette Smokers Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use Go and barrier conditions only.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var mainWait sync.WaitGroup     //For halting the main thread until all other threads have joined
var tobbaccoWait sync.WaitGroup //Tabbacco guy waits for this to signal
var matchesWait sync.WaitGroup  //Matches fella is waiting for this one
var paperWait sync.WaitGroup    //Paper lady is politely waiting and watching this one
var smokingWait sync.WaitGroup  //Everyone has to wait until the selfish smoker is done with this waitgroup
var startUp sync.WaitGroup      //Just pausing while all threads spin up

func main() {
	startUp.Add(3)      //To wait for all three smokers to finish spawning
	tobbaccoWait.Add(1) //This wait group is already halted
	go tobaccoHolder()  //Start goroutine
	matchesWait.Add(1)  //Halted from the start
	go matchesHolder()  //Start goroutine
	paperWait.Add(1)    //Immediately blocking
	go paperHolder()    //Start goroutine
	//Parsing command line arguments
	var iterS string
	var runsS string

	if len(os.Args) > 2 {
		iterS = os.Args[1]
		runsS = os.Args[2]
	} else {
		iterS = "30"
		runsS = "1"
	}
	iter, err := strconv.Atoi(iterS)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	runs, err := strconv.Atoi(runsS)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	mainWait.Add(1) //Create wait condition for main thread
	for r := 0; r < runs; r++ {

		go monitor(iter) //The monitor goroutine starts

	}
	mainWait.Wait()
}

func smoking() {
	defer smokingWait.Done() //To pause all smokers until this one has finished. I used a defer out of habit.
}

func monitor(iter int) {
	defer mainWait.Done()                          //For later joining on the main thread
	startUp.Wait()                                 //Wait for all the smokers to start up
	itemR := rand.NewSource(time.Now().UnixNano()) //Random seed for selecting which items to lay on the table
	for i := 0; i < iter; i++ {
		smokingWait.Wait()   //Wait for the smoker to finish and return the items
		var item1, item2 int //Get two item variables

		for item1 == item2 { //Make sure the items aren't the same
			itemN1 := rand.New(itemR)
			item1 = itemN1.Intn(3) //Get random item 1 as an int from 0-2
			itemN2 := rand.New(itemR)
			item2 = itemN2.Intn(3) //Get random item 2 as an int from 0-2
		}

		items := item1 + item2 //Add the two ints

		switch items { //Assess the conditions to see who smokes
		case 1:
			smokingWait.Add(1)  //Start the wait for other smokers to wait on
			tobbaccoWait.Done() //Signal the Tobbacco holder to proceed
		case 2:
			smokingWait.Add(1)
			matchesWait.Done() //Signal the Matches holder to proceed
		case 3:
			smokingWait.Add(1)
			paperWait.Done() //Signal the Paper holder to proceed
		}
	}
	smokingWait.Wait() //Wait until all the smoking has finished then restart the loop
}

func paperHolder() {
	startUp.Done() //Signal ready to start
	for {
		paperWait.Wait() //Wait for the signal to take the items and smoke away!
		paperWait.Add(1) //Reset waitgroup so that thread will halt again when the loop restarts
		go smoking()     //Don't know why I made this it's own goroutine.  Doesn't matter for this example
	}
}

//Similar to paperHolder() code
func matchesHolder() {
	startUp.Done() //Signal ready to start
	for {
		matchesWait.Wait()
		matchesWait.Add(1)
		go smoking()
	}
}

//Similar to paperHolder() code
func tobaccoHolder() {
	startUp.Done() //Signal ready to start
	for {
		tobbaccoWait.Wait()
		tobbaccoWait.Add(1)
		go smoking()
	}
}
