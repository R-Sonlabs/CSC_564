/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 3, Implementation A - Readers-Writers Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use Go and only barrier conditions.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var roomBarrier sync.WaitGroup   //Controlled by the writer to bar readers from entering
var writerBarrier sync.WaitGroup //Controlled by readers to bar writers from entering when they are in the critical section still.
var dogBone int                  //Our critical section variable

func main() {

	var mainWait sync.WaitGroup //To halt the main thread until all threads have completed
	mainWait.Add(1)
	//Parsing command line arguments
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

	readersIter := iter * 10 //Ten times as many readers as writers
	writersIter := iter

	roomBarrier.Add(0)   //Initialize and open the roomBarrier
	writerBarrier.Add(0) //Initialize and open the writerBarrier

	for i := 0; i < readersIter; i++ { //Spawn reader threads
		go readers(i)
	}

	for i := 0; i < writersIter; i++ { //Spawn writer threads
		go writers(i)
	}

	mainWait.Wait() //Halting main until completion
}

func readers(i int) {
	id := strconv.Itoa(i)
	readerName := "Reader " + id //For debugging purposes
	for {
		roomBarrier.Wait()                                  //Decrement the room barrier.  Will wait here if a writer shows up.
		writerBarrier.Add(1)                                //Notify any writers that show up that the room is occupied
		dogBoneS := strconv.Itoa(dogBone)                   //Read the variable
		fmt.Printf(readerName + " sees " + dogBoneS + "\n") //For debugging
		writerBarrier.Done()                                //Signal leaving the critical section
	}

}

func writers(i int) {
	id := strconv.Itoa(i)
	writerName := "***WRITER " + id //For debugging
	for {
		roomBarrier.Add(1)   //Increment the roomBarrier barring further readers from entering
		writerBarrier.Wait() //Wait for last reader to leave
		dogBone += 1         //Write to the variable
		dogBoneI := strconv.Itoa(dogBone)
		fmt.Printf(writerName + " wrote " + dogBoneI + "\n") //For debugging
		roomBarrier.Done()                                   //Signal that the writer has left so that readers may enter again
	}
}
