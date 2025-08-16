import java.util.ArrayList;
import java.util.InputMismatchException;
import java.util.Scanner;

public class Tester {

    public static void main(String[] args) {
        // openning scanner
        Scanner keyboard = new Scanner(System.in);
        boolean active = true;

        // getting the name of the dictionary file
        System.out.println("Loading dictionary");
        String fileName = "mediumDictionary.txt";

        // populating the dictionary
        MyDictionary dict = new MyDictionary();
        dict.readFile(fileName);

        // opening the command line to ask for the word to spell check
        while (active) {
            boolean valid = false;
            int command = 0;

            while (!valid) {
                try {
                    System.out.println(
                            "\nWhat would you like to do? \n1)Check word\n2)Suggestion\n3)Add New Word\n4)Display Load Factor\n5)Remove Word\n 6)Exit");
                    command = keyboard.nextInt();
                    keyboard.nextLine();// consuming new line char
                    valid = true;
                } catch (InputMismatchException e) {
                    System.out.println("Input invalid, try again");
                    keyboard.nextLine();
                }

            }

            switch (command) {
                // check
                case 1 -> {
                    System.out.println("What word would you like to check?");
                    String word = keyboard.nextLine().trim();

                    // checking if the word exists
                    boolean exists = dict.contains(word);

                    // printing out if it exists or not
                    if (exists) {
                        System.out.println("Word is correct");
                    } else {
                        System.out.println("Word is not correct");
                    }
                }

                // suggestion
                case 2 -> {
                    System.out.println("What word would you like suggestions on?");
                    String word = keyboard.nextLine().trim();

                    ArrayList<String> suggestions = dict.suggestCorrections2(word);

                    System.out.println("Suggestions are " + suggestions);

                }

                // add
                case 3 -> {
                    System.out.println("What word would you like to add?");
                    String word = keyboard.nextLine().trim();

                    dict.addToDictionary(word);
                }

                // loadfactor
                case 4 -> {
                    dict.displayLoadFactor();
                }

                //remove word
                case 5 -> {
                    System.out.println("What word would you like to remove?");
                    String word = keyboard.nextLine().trim();

                    dict.removeFromDictionary(word);
                }

                // exit
                case 6 -> {
                    keyboard.close();
                    System.exit(0);
                }

                default -> {
                    System.out.println("Command unrecognized");
                }

            }
        }
    }
}