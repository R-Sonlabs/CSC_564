/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 5, Implementation B - Cigarette Smokers Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use Dlang and Semaphores only.
*/

import std.stdio;
import std.conv;
import std.concurrency;
import std.random;
import std.range;
import core.stdc.stdlib;
import core.sync.semaphore;
import core.thread;

    int iter;
    int runs;
    __gshared Semaphore mainSem;            //To halt the main thread until other threads have completed
    __gshared Semaphore tobbaccoSem;        //Signal for the tobbaccco holder
    __gshared Semaphore matchesSem;         //Signal for the matches holder
    __gshared Semaphore paperSem;           //Signal for the paper holder
    __gshared Semaphore smokingSem;         //Signal for the others to wait until a smoker is finished smoking
    __gshared Semaphore startUp;            //Signal that all threads have started and are ready
    
void main(string[] args)
{   //Initialize those little semaphores, baby!
    startUp = new Semaphore();
    smokingSem = new Semaphore();
    smokingSem.notify();            //Release the smoking semaphore as no one is smoking yet
    //Parsing command line arguments
    if (args.length > 2){
        iter = to!int(args[1]);
        runs = to!int(args[2]);
    } else {
        iter = 10;
        runs = 1;
    }
    //Even more semaphores to initialize
    tobbaccoSem = new Semaphore();
    matchesSem = new Semaphore();
    paperSem = new Semaphore();
    auto tobbaccoThread = spawn(&tobbaccoHolder);       //Spawn tobbacco holder thread
    startUp.wait();                                     
    auto matchesThread = spawn(&matchesHolder);         //Spawn matches holder thread
    startUp.wait();
    auto paperThread = spawn(&paperHolder);             //Spawn paper holder thread
    startUp.wait();
    //Wait, another semaphore to initialize?  Oh right, gotta make the main thread wait until the others are finished
    mainSem = new Semaphore();                          
    mainSem.notify();
    for (int i = 0; i < runs; i++){
        mainSem.wait();
        auto monitorThread = spawn(&monitor, iter);      //Spawn the monitor thread
    }
}

void smoking() {
    smokingSem.notify();                                //Signal that smoking has finished
}

void monitor(int iter){
    
    for (int r = 0; r < iter; r++){
        smokingSem.wait();                              //Wait on any smokers smoking to finish before starting again
        int item1;
        int item2;

        while (item1 == item2) {                        //Make sure our items are not the same
        auto rnd1 = Random(unpredictableSeed);          //I love that seed name.
        item1 = uniform(0,3, rnd1);                     //Randomnly select item1 as an int 0-2
        auto rnd2 = Random(unpredictableSeed);          //Seriously, this should be a band name
        item2 = uniform(0,3, rnd2);                     //Randomnly select item2 as an int 0-2
        }
        int items = item1 + item2;                      //Add the item ints together.  Saves having two switch-cases
    
        switch (items) {                                //Let's evaluate!
            default:
                throw new Exception("bad switch");      //Also a band name?
            case 1:
                //tobbacco signal
                tobbaccoSem.notify();
                goto case;
            case 2:
                //matches signal
                matchesSem.notify();
                goto case;
            case 3:
                //paper signal
                paperSem.notify();
                
        }       
    }
    mainSem.notify();     //Signal main thread to end
}

void tobbaccoHolder() {
    startUp.notify();           //Signal ready to proceed
    while (true) {
        tobbaccoSem.wait();     //Wait for the signal.  Like Batman...
        smoking();              //Do the thing!
    }    
}
//Similar to tobbaccoHolder code
void matchesHolder() {
    startUp.notify();           //Signal ready to proceed
    while (true) {
        matchesSem.wait();
        smoking();
    }
}
//Similar to tobbaccoHolder code
void paperHolder() {
    startUp.notify();           //Signal ready to proceed
    while (true) {
        paperSem.wait();
        smoking();      
    }
}
