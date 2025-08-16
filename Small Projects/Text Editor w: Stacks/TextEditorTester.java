
import java.util.ArrayList;
import java.util.Scanner;

public class TextEditorTester {

    // Creating the welcome menu
    public static void menu(TextEditor text) {
        // initiating keyboard
        Scanner keyboard =  new Scanner(System.in);

        //asking what action should be done
        System.out.println("");
        System.out.println(
                "What would you like to do? \nInsert at cursor \tDelete text \nMove the cursor \tUndo \nRedo \t\t\tSearch \nSave \t\t\tSave as \nCopy \t\t\tPaste \nFind and Replace \tExit\n");
        String input = keyboard.nextLine().strip().toLowerCase();

        // getting the right input using switch
        switch (input) {
            case "insert at cursor" -> {
                // getting the wanted text to input
                System.out.println("What would you like to insert?");
                String newText = keyboard.nextLine();

                // inserting text
                text.insert(newText);

                // printing content and resetting menu
                System.out.println(text.getContent());
                menu(text);
            }

            case "delete text" -> {
                // getting the amount wanted to delete
                System.out.println("What length do you want to delete?");
                int len = keyboard.nextInt();

                // deleting wanted text
                text.delete(len);

                // printing content and resetting menu
                System.out.println(text.getContent());
                menu(text);
            }

            case "move the cursor" -> {
                // getting where to move the cursor
                System.out.println("Where would you like to move the cursor?");
                int len = keyboard.nextInt();
                // moving cursor
                text.moveCursor(len);

                // resetting to start of menu
                menu(text);
            }

            case "exit" -> {
                // exiting and closing needed things
                System.out.println("Thank you for using text editor!");
                System.out.println(text.getContent());
                keyboard.close();
                System.exit(0);
            }

            case "undo" -> {
                // undoing action
                text.undoAction();

                // printing content and resetting menu
                System.out.println(text.getContent());
                menu(text);

            }

            case "redo" -> {
                // redoing action
                text.redoAction();

                // printing content and resetting menu
                System.out.println(text.getContent());
                menu(text);

            }

            case "search" -> {
                // getting where to move the cursor
                System.out.println("What word are you looking for?");
                String word = keyboard.nextLine();

                ArrayList<Integer> arr = text.searchText(word);

                // printing content and resetting menu
                System.out.println(arr);
                menu(text);

            }

            case "find and replace" -> {
                // getting where to move the cursor
                System.out.println("What word are you looking for?");
                String word = keyboard.nextLine();

                System.out.println("What would you like to replace it?");
                String replace = keyboard.nextLine();

                text.replaceText(word, replace);

                // printing content and resetting menu
                System.out.println(text.getContent());
                menu(text);

            }

            case "save" -> {
                text.saveFile();

                menu(text);
            }

            case "save as" -> {
                System.out.println("Where would you like this to be saved?");
                String file = keyboard.nextLine();

                text.saveAs(file);

                menu(text);   
            }

            case "copy" -> {
                //getting start index
                System.out.println("What is the start index?");
                int start = keyboard.nextInt();

                //getting end index
                System.out.println("What is the end index?");
                int end = keyboard.nextInt();

                //copying text
                text.copy(start, end);

                //back to menu
                menu(text);
            }

            case "paste" -> {
                System.out.println("Where would you like to paste?");
                int start = keyboard.nextInt();

                //pasting
                text.paste(start);
                
                //back to menu
                System.out.println(text.getContent());
                menu(text);
            }

            default -> {
                System.out.println("Not valid input");
                System.exit(0);
            }
        }

    }

    public static void main(String[] args) {
        // starting the welcome menu
        TextEditor text = new TextEditor();

        //asking if a file should be opened
        Scanner scan = new Scanner(System.in);
        System.out.println("Would you like to open a file? \nYes \nNo");
        String ans = scan.nextLine().strip().toLowerCase();
        if (ans.equals("yes")){
            System.out.println("What file would you like to open?");
            String file = scan.nextLine();
            text.openFile(file);
        }

        menu(text);
        scan.close();
    }

}
