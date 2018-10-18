/*
	CSC564 - Assignment 1
	Matt Richardson, V00905254
	Problem 3, Implementation B - Readers-Writers Problem
	Uvic, Fall 2018

	This implementation is original in that it was not derived from any other solutions known
	to the author.  The self-imposed restrictions were to use C# and any tools I wanted.
*/

using System;
using System.Threading;
using static System.Console;

class ReadersWriters
{
    static ReaderWriterLock lockRW;     //Our only tool.
    static int dogBone;                 //The variable for the critical section
    static int iter;                    //For testing runs

    static void Main(string[] args)
    {   //Parsing command line arguments
        if (args.Length > 0)
        {
            iter = int.Parse(args[0]);
        }
        else
        {
            iter = 1;
        }

        lockRW = new ReaderWriterLock();    //Initialze the lock

        int readersIter = iter * 10;        //Time times as many readers as writers
        int writersIter = iter;

        for (int i = 0; i < readersIter; i++)   //Spawn reader threads
        {
            Thread readerThread = new Thread(reader);
            readerThread.Name = "Reader " + i;
            readerThread.Start();
        }

        for (int i = 0; i < writersIter; i++)   //Spawn writer threads
        {
            Thread writerThread = new Thread(writer);
            writerThread.Name = "****WRITER " + i;
            writerThread.Start();
        }
    }

    static void reader()
    {
        while (true)
        {
            lockRW.AcquireReaderLock(Timeout.Infinite);                 //Any number of readers make get this lock.  Disabled when a writer gets the WriterLock
            WriteLine(Thread.CurrentThread.Name + " sees " + dogBone);  //Read variable
            lockRW.ReleaseReaderLock();                                 //Release the reader lock.  When 0 the writer know that no readers are in the critical section
        }
    }

    static void writer()
    {
        while (true)
        {
            lockRW.AcquireWriterLock(Timeout.Infinite);                 //Only one writer may get this lock.  When acquired, all readers are barred from entering the critical section
            dogBone++;                                                  //Write to the variable
            WriteLine(Thread.CurrentThread.Name + " wrote " + dogBone); //For debugging
            lockRW.ReleaseWriterLock();                                 //Let any readers or writers waiting proceed
        }
    }
}