import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.NoSuchElementException;
import java.util.Scanner;
import java.util.logging.Logger;

public class Traverse {
    private static final Logger LOGGER = Logger.getLogger(Traverse.class.getName());

    public static void main(String[] args) throws IOException {
        try (
            //initiating scanner
            Scanner keyboard = new Scanner(System.in)) {

            // asking user what they would like to do
            System.out.println("What would you like to do? \nFind a file\nDisplay summary of a folder \n");
            String input = keyboard.nextLine().toLowerCase().strip();
            
            // initializing needed traversal items and getting to the wanted folder
            String home = System.getProperty("user.home");
            File directory = new File(home + "/Desktop");
            String fileName;
            File[] baseFiles = directory.listFiles();
            if (baseFiles != null) {
                for (File file : baseFiles) {
                    if (file.isDirectory() && file.getName().equals(args[0])) {
                        System.out.println(file.getAbsolutePath());
                        directory = new File(file.getAbsolutePath());
                    }
                }
            }

            //going to the correct mneu item 
            switch (input) {
                case ("find a file") -> {
                    //creating arrayList
                    ArrayList<String> files = new ArrayList<>();

                    // asking user for wanted file
                    System.out.println("What are you looking for?");
                    fileName = keyboard.nextLine();

                    // searching for file
                    findFile(fileName, files, directory);

                    // writing to a result file
                    try (FileWriter writer = new FileWriter("results_search.txt")) {
                        for (String path : files) {
                            writer.write(path + "\n");
                        }
                    }
                }

                case ("display summary of a folder") -> {
                    // creating queue
                    LinkedQueue<File> queue = new LinkedQueue<>();

                    // searching directory and returning information
                    findTextFile(queue, directory);

                    //calculating smallest, biggest, and the count and avg disk usage 
                    File smallest = null;
                    File biggest = null;
                    int count = 0;
                    long diskUsage = 0;
                    LinkedQueue<TextFile> fileQueue = new LinkedQueue<>();

                    //iterating through the queue till empty
                    while (!queue.isEmpty()) {
                        //get first item and remove it 
                        File item = queue.dequeue();

                        //incrementing count and diskUsage
                        count++;
                        diskUsage += item.length();

                        //if this is the first element set smallest and biggest to this file
                        if (smallest == null) {
                            smallest = item;
                            biggest = item;
                        }

                        // checking if smallest
                        if (item.length() <= smallest.length()) {
                            smallest = item;
                        }

                        // checking if biggest
                        if (item.length() >= biggest.length()) {
                            biggest = item;
                        }
                        TextFile newFile = new TextFile(item.getName(), item.length(), item.getAbsolutePath());
                        fileQueue.enqueue(newFile);
                    }

                    //printing out wanted things to console
                    System.out.println("Total is: " + count + "\nAverage DiskUsage: " + (diskUsage / count));
                    System.out.println("Smallest File is: " + smallest.getName() + "\nBiggest File is: " + biggest.getName());

                    //if no files exist then the given directory doesn't exist
                    if (fileQueue.isEmpty()) {
                        LOGGER.warning("not a directory");
                    } else {
                        // writing to a result file
                        try (FileWriter writer = new FileWriter("results.txt")) {
                            while (!fileQueue.isEmpty()) {
                                TextFile file = fileQueue.dequeue();
                                writer.write(file.fileName + ", " + file.path + ", " + file.fileSize + "\n");
                            }
                        }

                    }
                }
                default -> LOGGER.warning("not an option");
            }
        } catch (NoSuchElementException e2) {
            LOGGER.severe("input mismatch");
        }
    }

    // Method to traverse and find a file
    public static void findFile(String fileName, ArrayList<String> arr, File directory) {
        // checking if the current file is a directory
        if (directory.isDirectory()) {
            // getting each file and directory out of this directory
            File[] files = directory.listFiles();

            // checking through the files
            if (files != null) { 
                for (File each : files) {// checking through each file
                    if (each.canRead()) {// checking that it is readable and not needing super access
                        // if directory access directory
                        if (each.isDirectory()) {
                            findFile(fileName, arr, each);
                        } else {// if a file check if it is the file that we are looking for
                            if (each.getName().equals(fileName)) {
                                LOGGER.info(each.getAbsolutePath());
                                arr.add(each.getAbsolutePath());
                            }
                        }
                    }
                }
            }
        }
    }

    // Method to traverse and find a text file
    public static void findTextFile(LinkedQueue<File> arr1, File directory) {
        //iterating through directory 
        if (directory.isDirectory() && directory.exists()) {
            //access all files in directory
            File[] files = directory.listFiles();

            //if thre are files then access the, 
            if (files != null) {
                //go through each file and add them to the queue if they are a text file
                for (File newFile : files) {
                    if (newFile.isFile() && newFile.getName().strip().endsWith(".txt")) {
                        arr1.enqueue(newFile);
                    } else {// else directory and recursion happens
                        findTextFile(arr1, newFile);
                    }
                }
            }
        }
    }
}