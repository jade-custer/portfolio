//Jade Custer
//ID: 20682049
//Github ID: jadecuster
//Lab06

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.InputMismatchException;
import java.util.Scanner;

public class Tester {

    //Method to read in the file
    public static ArrayList<String> readFile(String fileName) {
        //initalilizing the arraylist
        ArrayList<String> arr = new ArrayList<>();

        //try catch for bufferedreader
        try {
            BufferedReader reader = new BufferedReader(new FileReader(fileName));
            String line;

            //adding each line into the array to be processed into the tree
            while ((line = reader.readLine()) != null) {
                arr.add(line);
            }

            reader.close();
        } catch (IOException e) {
            System.out.println("IOException");
            System.exit(0);
        }

        return arr;
    }

    public static void main(String[] args) {
        //creating the tree
        AVLTree avlTree = new AVLTree();
        boolean active = true;

        //getting the wanted words from the dictionary
        ArrayList<String> arr = readFile("smallDictionary.txt");

        String[] newArr = new String[arr.size()];
        newArr = arr.toArray(newArr);

        //inserting each value into the tree
        for (String value : newArr) {
            avlTree.insert(value);
        }

        // Setting up to do a menu
        Scanner keyboard = new Scanner(System.in);

        //Keeping the user in the menu while they arent ready to be exited
        while (active) {
            boolean valid = false;
            int command = 0;

            //getting the user to type in something that matches whats in the menu 
            while (!valid) {
                try {
                    System.out.println(
                            "\nWhat would you like to do? \n1)Check word\n2)Autocomplete\n3)Exit");
                    command = keyboard.nextInt();
                    keyboard.nextLine();// consuming new line char
                    valid = true;
                } catch (InputMismatchException e) {
                    System.out.println("Input invalid, try again");
                    keyboard.nextLine();
                }
            }

            //switch to use to process menu items
            switch (command) {
                //Checking word
                case 1 -> {
                    //Getting user input
                    System.out.println("What word are you searching for?");
                    String word = keyboard.nextLine();

                    //Searching for word
                    boolean check = avlTree.search(word);

                    System.out.println(check);
                }

                //Autocompleting word
                case 2 -> {
                    //Getting user input
                    System.out.println("What word are you looking to autocomplete?");
                    String word = keyboard.nextLine();

                    //Getting wanted suggestions
                    ArrayList<String> suggestions = avlTree.autocomplete(word);

                    System.out.println(suggestions);

                }

                //Exiting
                case 3 -> {
                    keyboard.close();
                    System.exit(0);
                }
            }
        }
    }
}
