/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 4, Implementation B - Building H2O
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use Dlang and semaphores and barrier conditions only.
*/


import std.stdio;
import std.conv;
import std.concurrency;
import core.sync.semaphore;
import core.sync.barrier;
import core.thread;

int iter;
__gshared Semaphore hydmutex;           //The hydrogen mutex
__gshared Semaphore oxyMutex;           //The oxygen mutex
int hydrogenCounter;                    //For checking if first or second hydrogen atom
__gshared Barrier barrier;              //To halt until all atoms are ready to proceed
//Parsing command line arguments
void main(string[] args){
        if (args.length > 1){
        iter = to!int(args[1]);
    } else {
        iter = 10;        
    }

    hydMutex = new Semaphore(1);        //Initialized open
    oxyMutex = new Semaphore(1);        //Initialized open
    barrier = new Barrier(3);           //Initialized for 3 waits
    hydrogenCounter = 0;                //Ain't got no hydrogen yet!

    for (int i = 0; i < iter; i++){     //Spawn the atom threads.  1:2 ratio of oxygen to hydrogen
        auto oxygen = spawn(&Oxygen);
        auto hydrogen = spawn(&Hydrogen);
        hydrogen = spawn(&Hydrogen);
    }
}

void Oxygen(){
    oxyMutex.wait();                    //Decrement semaphore, bars other oxygen from proceeding until this one has finished bonding.
    barrier.wait();                     //Gets one of three waits on the barrier
    Bond();                             //Where the magic happens
    oxyMutex.notify;                    //Increment the semaphore for other oxygen atoms to proceed
}

void Hydrogen(){
    hydMutex.wait();                    //Decrement semaphore, bars other hydrogen from proceeding temporarily
    hydrogenCounter += 1;               //Increment for checking later
    if (hydrogenCounter == 2){          //If second hydrogen
        hydrogenCounter -= 2;           //Reset the counter
        barrier.wait();                 //Get one of three waits on this barrier
        Bond();                         //Zappo!
        hydMutex.notify();              //Release for next set of hydrogen
    } else {
        hydMutex.notify();              //If first hydrogen signal the second waiting atom to proceed
        barrier.wait();                 //Get one of three waits on this barrier
        Bond();                         //Bazinga!
    }
    
}

void Bond(){
    //Not much to do here in this example
}