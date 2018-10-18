/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *  The Yvonne-A-Tron, 2018!!                                               *
 *  Smashed out by Matt Richardson, CSC564, UVIC. Spring, 2018.             *
 *                                                                          *
 *  Now with EXTRA colors and exclamation points!!!!!!!!                    *
 *                                                                          *
 *  With much admiration to Prof. Yvonne Coady,                             *
 *          My teacher, mentor, and friend.                                 *
 *                                                                          *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

using System;
using static System.Console;

class Yvonneatron
{

    private static Random rnd = new Random();
    static int mode;

    static void Main(string[] args)
    {
        if (args.Length > 0)
        {
            mode = int.Parse(args[0]);
        }
        else
        {
            mode = 1;
        }

        string[] problems = {
            "4.1 Producer-consumer problem",
            "4.2 Readers-writers problem",
            "4.3 No-starve mutex",
            "4.4 Dining philosophers",
            "4.5 Cigarette smokers problem",
            "5.1 The dining savages problem",
            "5.2 The barbershop problem",
            "5.3 The FIFO barbershop",
            "5.4 Hilzer's Barbershop problem",
            "5.5 The Santa Claus problem",
            "5.6 Building H20",
            "5.7 River crossing problem",
            "5.8 The roller coaster problem",
            "6.1 The search-insert-delete problem",
            "6.2 The unisex bathroom problem",
            "6.3 Baboon crossing problem",
            "6.4 The Modus Hall Problem",
            "7.1 The sushi bar problem",
            "7.2 The child care problem",
            "7.3 The room part problem",
            "7.4 The Senate Bus problem",
            "7.5 The Faneuil Hall problem",
            "7.6 Dining Hall problem"
        };

        string[] languages = {
            "C/C#/C++",            
            "D",
            "Go",
            "Rust"
        };

        string[] methods = {
            "Semaphores",
            "Mutex",
            "Barriers",
            "Channels",
            "locks",
            "whatever tools you want",
            "Someone elses implementation (Whew!)"
        };

        string[] luckArray = {
            "YOU CAN DO IT!!!!!!!",
            "GOOD LUCK!!!!!!!!",
            "THIS IS GOING TO BE AWESOME!!!!!!!!",
            "RAH RAH RAH!!!!!!!!",
            "GO GO GO GO GO!!!!!  Not the language, the verb. Duh.",
            "HOW FREAKING COOL!!!!!!!!!"
        };

        
        string problem = problems[rnd.Next(0,23)];
        string language = languages[rnd.Next(0,4)];
        int methodsCounter = rnd.Next(1,3);
        string method;
        string methodA = "";
        string methodB = "";
        string luckString = luckArray[rnd.Next(0, 6)];
        string luck = "";
        foreach (char c in luckString)
            {
                ForegroundColor = GetRandomConsoleColor();
                luck += c;

            }
        WriteLine();
        ForegroundColor = ConsoleColor.Yellow;
        WriteLine("---------------------------------------------------------------------------------------");
        WriteLine();
        ForegroundColor = ConsoleColor.Cyan;
        if (methodsCounter == 1) {

            method = methods[rnd.Next(0,6)]+" ONLY!";

        } else {

            if (methodsCounter == 2){
                methodA = methods[rnd.Next(0, 6)];
                methodB = methods[rnd.Next(0, 6)];
                while (methodA == methodB)
                {
                    methodA = methods[rnd.Next(0, 6)];
                }
            }
            if (methodA == "Someone elses implementation (Whew!)" || methodB == "Someone elses implementation (Whew!)")
            {
                method = "someone elses implementation (Whew!).";
            } else if (methodA == "whatever tools you want" || methodB == "whatever tools you want" ){
                method = methods[6];
            }
            {
                method = methodA + " and " + methodB + ".";
            }            
        }  

        if (mode == 1){
            Write("   For this problem, do "+problem+" in "+language+" using "+method);
        } else if (mode == 2){
            Write("   Ok, for the second implementation use "+language+" using "+method+" "+luck);  
        } else if (mode == 3){
            Write("   "+luck);
        }
      
        WriteLine();
        WriteLine();
        ForegroundColor = ConsoleColor.Yellow;
        WriteLine("---------------------------------------------------------------------------------------");
        WriteLine();
        ForegroundColor = ConsoleColor.Gray;
    }

    private static ConsoleColor GetRandomConsoleColor()
    {
        var consoleColors = Enum.GetValues(typeof(ConsoleColor));
        return (ConsoleColor)consoleColors.GetValue(rnd.Next(consoleColors.Length));
    }

}