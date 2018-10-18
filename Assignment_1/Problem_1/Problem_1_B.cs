/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 1, Implementation B - The Santa Claus Problem
	Uvic, Fall 2018

	This implementation is from the Little Book of Semaphores and it attributed to Allen Downey.
    The self-imposed restriction was to use C#.
*/

using System;
using System.Threading;
using static System.Console;

namespace NorthPoleA
{
    class SantasWorkshopA
    {
        static int elves;   //This counts the number of elves waiting to get help from Santa.  
        static int iter;    //Number of iterations to run for testing
        static int reindeer;    //Keeping track of the number of reindeer that have returned from holidays
        static Semaphore santaSem;  //Santa sleeps on this
        static Semaphore reindeerSem;   //The reindeer wait on this to enter GetHitched();
        static Semaphore elfTex;    //Protecting the elf incrementing
        static Semaphore mutex; //Protecting the reindeer incrementing

        static void Main(string[] args)
        {
            if (args.Length > 0)
            {
                iter = int.Parse(args[0]);
            }
            else
            {
                iter = 10;
            }

            elves = 0;
            reindeer = 0;
            santaSem = new Semaphore(0, 1);
            reindeerSem = new Semaphore(0, 1);
            elfTex = new Semaphore(1, 1);
            mutex = new Semaphore(1, 1);

            string[] reindeerNames = {
            "Dasher",
            "Dancer",
            "Prancer",
            "Vixen",
            "Comet",
            "Cupid",
            "Donner",
            "Blitzen",
            "Rudolph"};

            string[] elfNames = {
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
             "Cedar"
            };

            new Thread(Santa).Start();

            for (int i = 0; i < 9; i++)
            {
                Thread reindeerThread = new Thread(Reindeer);
                reindeerThread.Name = reindeerNames[i];
                reindeerThread.Start();
            }

            for (int i = 0; i < 42; i++)
            {
                Thread elfThread = new Thread(Elf);
                elfThread.Name = elfNames[i];
                elfThread.Start();
            }
        }

        static void Santa()
        {
            while (true)
            {
                santaSem.WaitOne();
                mutex.WaitOne();
                if (reindeer >= 9)
                {
                    PrepareSleigh();
                    reindeerSem.Release();
                    reindeer -= 9;
                } else if (elves == 3)
                {
                    HelpElves();
                }
                mutex.Release();
            }
        }

        static void Reindeer()
        {
            mutex.WaitOne();
            reindeer += 1;
            if (reindeer == 9)
            {
                santaSem.Release();
            }
            mutex.Release();
            reindeerSem.WaitOne();
            GetHitched();
        }

        static void PrepareSleigh()
        {
            WriteLine("Preparing Sleigh");
            new Thread(Christmas).Start(); ;
        }

        static void GetHitched()
        {
            WriteLine("Getting hitched");
            new Thread(Christmas).Start();
        }

        static void Elf()
        {
            elfTex.WaitOne();
            mutex.WaitOne();
            elves += 1;
            if (elves == 3)
            {
                santaSem.Release();
            } else
            {
                elfTex.Release();
            }
            mutex.Release();

            GetHelp();

            mutex.WaitOne();
            elves -= 1;
            if (elves == 0)
            {
                elfTex.Release();
            }
            mutex.Release();
        }

        static void GetHelp()
        {
            WriteLine("Elf waiting for help");
        }

        static void HelpElves()
        {
            WriteLine("Santa helped elves");
        }
    }
}
