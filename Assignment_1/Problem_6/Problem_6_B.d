/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 6, Implementation B - The Alarm Responders Problem
	Uvic, Fall 2018

	This problem was requested to be unique by Prof. Coady.
	I chose to fix a problem from a decade old piece of code I wrote.
	It is detailed on page 48 of my assignment paper.
*/


import std.stdio;
import std.concurrency;
import std.conv;
import std.random;
import core.sync.barrier;
import core.sync.mutex;
import core.thread;



Mutex mutex; //Our only mutex


void main(string[] args){
    //Parsing command line arguments
    int officersI;
    int keysI;
    int speedI;
    if (args.length > 3){
        officersI = to!int(args[1]);
        keysI = to!int(args[2]);
        speedI = to!int(args[3]);
    } else {
        officersI = 10;
        keysI = 100;
        speedI = 1;        
    }

    mutex = new Mutex();  //Initiallizing our mutex

    int[] keyDatabase = new int[keysI];   //Generating the officer's key "database"
    for (int i = 0; i < keysI; i++){
        keyDatabase[i] = i;
    }

    int[] keysCoded = new int[keysI];      //Generating the code pairs for the controller's "database"
    for (int i= 0; i <keysI; i++){
        int code = uniform(10000,99999);
        keysCoded[i] = code;
    }

    Tid controller = spawn(&Controller, cast(shared) keysCoded);  //Spawning the Controller thread and passing in the coded "database"

    for (int i = 0; i < officersI; i++){
        Tid officer = spawn(&Officer, cast(shared) keyDatabase);  //Spawning the officer threads and passing in the uncoded key "database"
        officer.send(controller);                   //Signalling the officer threads, and sending the ThreadID of the controller to help initialize communication
    }

    readln();    
}

void Controller(shared int[] keysCoded){
    while (true){  //Running in a perpetual loop
        int keyRequest = receiveOnly!int();  //Waiting on the keyRequest channel.  Each officer thread signals this with a key number to get a code returned.
        Tid requestingOfficer = receiveOnly!Tid();  //Getting the requesting officer's threadID.
        Tid coder = spawn(&Coder, keyRequest, requestingOfficer, keysCoded);  //Spawn a coder thread to handle the request with all the relevant information passed.  Leaving the controller available to process more requests.
    }   
}

void Coder(int keyRequest, Tid requestingOfficer, shared int[] keysCoded){
    int key = keyRequest;
    Tid officer = requestingOfficer;
    officer.send(keysCoded[key]);  //Return the code for the requested key
    int newCode = uniform(10000,99999);  //Generate a new random code
    mutex.lock();   //Locking the "database"
    keysCoded[key] = newCode;   //Writing the new code
    mutex.unlock(); //Unlocked the "Database"

}

void Officer(shared int[] keyDatabase){
    Tid controller = receiveOnly!Tid();  //Receive the controller's threadID
    while (true) {
        Thread.sleep(dur!("msecs")(500));  //zzzz...
        int key = uniform(0, keyDatabase.length);  //Select a random key from the uncoded "database"
        controller.send(key);  //Send the key request to the controller
        controller.send(thisTid);   //Send this officer's threadID so that it can be passed to the Coder thread to return the data
        int code = receiveOnly!int();   //Recieve the code from the Coder thread.
        writeln("Officer ",thisTid," requested Key",key," and got code ",code," back.");  //For debugging
    }

      
}


