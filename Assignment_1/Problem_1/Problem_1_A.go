/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 1, Implementation A - The Santa Claus Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use Go and only Waitgroups.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var mainWait sync.WaitGroup      //Waits for all threads to complete their iterations before ending the program.
var santaSleep sync.WaitGroup    //This is what Santa waits on (sleeps) and what the pesky elves and reindeer signal to wake him.
var reindeerWait sync.WaitGroup  //The reindeer wait on this 9 times before it signals, causing the last reindeer to wake the boss with santaSleep.Done()
var hitchWait sync.WaitGroup     //All reindeer will then wait on this, to be signalled by Santa after he wakes up.
var elfWait sync.WaitGroup       //The elf version of reindeerWait. Takes 3 signals.
var elfHelp sync.WaitGroup       //The elf version of hitchWait.  Also takes 3 signals.
var christmasWait sync.WaitGroup //Used to muster all the parties, 9 reindeer and Santa, and cause the elves to hold all their problems until they get back.
var elves bool;						//True if there are enough to elves to make it worth Santas time.
var reindeer bool;					//True if all the reindeer are actually back.

func main() {
	mainWait.Add(1)       //Only one required as Santa calls this when he breaks out of his last loop.
	reindeerWait.Add(9)   //9 for 9 reindeer.
	santaSleep.Add(1)     //Only one elf or reindeer actually wakes Santa.
	hitchWait.Add(9)      //Also 9 for 9 reindeer.
	elfWait.Add(3)        //It takes 3 elves to wake Santa.
	christmasWait.Add(10) //All the reindeer AND Santa signal this when Christmas is over, keeps the elves and their little problems at bay.

	//Code to parse command line arguments.  Allows user to add more iterations.
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

	//Name each thread for debugging purposes.
	reindeerNames := [9]string{
		"Dasher",
		"Dancer",
		"Prancer",
		"Vixen",
		"Comet",
		"Cupid",
		"Donner",
		"Blitzen",
		"Rudolph"}
	//Ok, this might be a little overkill.
	elfNames := [42]string{
		"Dash",
		"Evergreen",
		"Sugarplum",
		"Pixie",
		"Pudding",
		"Perky",
		"Candycane",
		"Glitter-toes",
		"Happy",
		"Angel-Eyes",
		"Sugar-Socks",
		"McJingles",
		"Frost",
		"Tinsel",
		"Twinkle",
		"Jingle",
		"Ginger",
		"Joy",
		"Merry",
		"Pepper",
		"Sparkle",
		"Tinsel",
		"Winter",
		"Trinket",
		"Buddy",
		"Noel",
		"Snowball",
		"Tiny",
		"Elfin",
		"Candy",
		"Carol",
		"Angel",
		"Nick",
		"Plum",
		"Holly",
		"Snow",
		"Pine",
		"Garland",
		"Joseph",
		"Gabriel",
		"Hope",
		"Cedar"}

	//Start goroutines
	go Santa(iter) //One of him.  The one and only.
	for i := 0; i < 9; i++ {
		go Reindeer(reindeerNames[i]) //The 9 reindeer goroutines.
	}
	for i := 0; i < 42; i++ {
		go Elf(elfNames[1]) //Thats a lot of elves.
	}

	mainWait.Wait() //To join all the threads upon completion.  Called by Santa as he dies.

}

func Santa(iter int) {
	defer mainWait.Done()	//Signal to end the program after all iterations
	for i := 0; i < iter; i++ {
		christmasWait.Add(1) //Thinking about next year (resets waitgroup)
		santaSleep.Wait()    //zzzzz....
		if (reindeer == true){  //Was it the reindeer who woke him?	
			PrepareSleigh();	//Get the show on the road.
		}
		else if (elves == true){ //Or was it the elves begging for help?
			HelpElves();	//Because somebody has to do it.
		}
	}
}

func Reindeer(name string) {
	for {
		reindeerWait.Done()	//Each reindeer signals the waitgroup.
		WarmingHut()	//It's the northpole, right?  Who doesn't need a warming hut, I don't care how much fur you have.
	}
}

func PrepareSleigh() {
	santaSleep.Add(9) //Fluffs up pillows so that he can go back to sleep when he returns to his bed (reseting waitgroup)
	christmasWait.Done()  //Signal to the pesky elves that they can start asking for help again and reindeer can go back on vacation
	Christmas()	//The big day!
}

func WarmingHut() {
	reindeerWait.Wait() //waiting for all 9 reindeer (probably playing cards)
	reindeerWait.Add(1) // asks an elf to book tickets so that the reindeer can go back on vacation
	reindeer = true;	//There are enough reindeer! (this lets santa know if they were the ones that woke him.)
	santaSleep.Done()   //wake santa... gently
	GetHitched()	//Yeehaw!
}

func GetHitched() {
	christmasWait.Wait() //don't you hate waiting for christmas?
	reindeer = false
}

func Elf(name string) {
	for {
		santaSleep.Done();	//Elves have to wake santa 3 times each. Santa waits for 9 wakes in total, so 3/3/3 9.
		santaSleep.Done();
		santaSleep.Done();
		elves = true;	//It was an elf who woke you, Santa!
		elfWait.Wait();
		elfWait.Add(1);
	}
}

func GetHelp() {
	elves = false;	//Back to work for the elves
}

func HelpElves() {
	elfWait.Done();
	elfWait.Done();
	elfWait.Done();
}

func Christmas() {
	fmt.Println("Christmas")
	//Finally!
}
