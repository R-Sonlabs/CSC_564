/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 2, Implementation A - The Unisex Bathroom Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use C# and only locks.
*/

using System;
using System.Threading;

public class UnisexBathroomB1
{
    static int females;                 //Keeping track of the number of females in the washroom
    static int males;                   //Keeping track of the number of males in the washroom
    static int attendant;               //The attendant enforces fairness by only letting a max of three cycles of one gender in before forcing a gender change
    static int employees;               //For the number of employees in total.  Evenly distributed male/female
    static int totalMales;              //Tracked for the puposes of flagging
    static int totalFemales;            //Tracked for the puposes of flagging
    static Object door = new Object();  //Our lock, engaged for all incrementing and decrementing operations

    public static void Main(string[] args)
    {   //Parsing command line arguments
        employees = int.Parse(args[0]);
        totalFemales = employees;
        totalMales = employees;

        lock (door)         //Not strictly necessary to lock at this time, but just testing!
        {
            females = 0;    //Setting all variables to zero
            males = 0;
            attendant = 0;
        }

        for (int i = 0; i < employees; i++) //Spin up the threads.  One for each employee.
        {
            Thread female = new Thread(Female);
            female.Name = i.ToString();
            female.Start();
            Thread male = new Thread(Male);
            male.Name = i.ToString();
            male.Start();
        }

    }

    public static void Bathroom()  //Entered by each employee once they are cleared to do so
    {
        Console.Write("FLUSH!\n");        
    }


    public static void Female()
    {
        bool done = false;  //Just a sloppy way of breaking out of the loop later

        while (!done)
        {
            lock (door) //Get the lock!
            {
                if (attendant < 3)  //Ask permission from the attendant
                {
                    if (females < 3)    //Check to make sure that a) woman are in the washroom and b) that there aren't too many
                    {
                        males = 3;      //Lock the boys out
                        females++;      
                        attendant++;
                        done = true;    //Break the loop
                    }
                }

            }
        }
        Bathroom();     //Enter the washroom
        lock (door)     //Regain the lock on exiting
        {
            females--;  //Make way for the next gal
            totalFemales--;
            if (attendant >= 3) //See if it's time for a gender swap
            {
                if (females == 0)   //Last female out?
                {
                    females = 3;    //Lock out any more girls
                    males = 0;      //Let the boys back in
                    attendant = 0;  
                }
            }
            if (totalFemales <= 1)
            {
                attendant = 0;
                males = 0;
            }
        }
    }
                                    //Male code is identical to female code except with male in place of female
    public static void Male()
    {
        bool done = false;

        while (!done)
        {
            lock (door)
            {
                if (attendant < 3)
                {
                    if (males < 3)
                    {
                        females = 3;
                        males++;
                        attendant++;
                        done = true;
                    }
                }
            }
        }
        Bathroom();
        lock (door)
        {
            males--;
            totalMales--;
            if (attendant >= 2)
            {
                if (males == 0)
                {
                    males = 3;
                    females = 0;
                    attendant = 0;
                }
            }
            if (totalMales <= 1)
            {
                attendant = 0;
                females = 0;
            }
        }
    }
}