/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 2, Implementation B - The Unisex Bathroom Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use C# and only semaphores.
*/

using System;
using System.Threading;

class UnisexBathroomA1
{
    static int iter;
    private static Semaphore ready;     //All employees must wait on this to enter the washroom
    private static Semaphore attendant; //To ensure gender equality
    private static Semaphore genF;      //A singular semaphore, more of a mutex, to hold while incrementing Fs.
    private static Semaphore genM;      //Male version of genF
    private static Semaphore femaleSem; //Initialized to 3 to ensure only 3 females enter
    private static Semaphore maleSem;   //Initialized to 3 to ensure only 3 males enter
    static int Fs;                      //Incremented to count the number of females in the washroom
    static int Ms;                      //Incremented to count the number of males in the washroom

    static void Main(string[] args)
    {   //Initialize all our primatives
        iter = int.Parse(args[0]);
        ready = new Semaphore(1, 1);
        attendant = new Semaphore(3, 3);
        genF = new Semaphore(1, 1);
        genM = new Semaphore(1, 1);
        femaleSem = new Semaphore(3, 3);
        maleSem = new Semaphore(3, 3);

        for (int i = 0; i < iter; i++)   //Spawning all out employees as threads.  Evenly divided between male and female
        {
            Thread female = new Thread(Female);
            female.Name = i.ToString();
            female.Start();
            Thread male = new Thread(Male);
            male.Name = i.ToString();
            male.Start();
        }


    }

    public static void Bathroom()
    {
        Console.Write("FLUSH!\n");      //What they all came here to do
    }

    static void Female()
    {
        attendant.WaitOne();    //Takes three waits before blocking
        genF.WaitOne();         //The mutex
        Fs++;                   //Used to check if they are the first or last person to enter
        if (Fs == 1)            //If first female to enter
        {
            ready.WaitOne();    //Lock the ready for the next person
        }
        genF.Release();         //Release the lock so the next female can increment or decrement the counter
        attendant.Release();    //Let the next female run
        femaleSem.WaitOne();    //Takes three wait before blocking
        Bathroom();             //Enter the bathroom
        femaleSem.Release();    //Released so that other females can proceed
        genF.WaitOne();         //Get the mutex again
        Fs--;                   //Decrement as leaving
        if (Fs == 0)            //If last female out
        {
            ready.Release();    //Unlock the ready for the next group
        }
        genF.Release();         //Unlock the mutex
    }

    static void Male()          //Identical to female code, except male in place of female
    {
        attendant.WaitOne();
        genM.WaitOne();
        Ms++;
        if (Ms == 1)
        {
            ready.WaitOne();
        }
        genM.Release();
        attendant.Release();
        maleSem.WaitOne();
        Bathroom();
        maleSem.Release();
        genM.WaitOne();
        Ms--;
        if (Ms == 0)
        {
            ready.Release();
        }
        genM.Release();
    }
}