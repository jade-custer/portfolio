import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.Stack;

public class TextEditor {
    StringBuilder str;
    int cursorPosition;
    String copy;
    String curFile;
    Stack<EditAction> undo;
    Stack<EditAction> redo;

    // constructor
    public TextEditor() {
        this.str = new StringBuilder();
        this.cursorPosition = 0;
        this.copy = "";
        this.curFile = "";
        this.undo = new Stack<>();
        this.redo = new Stack<>();
    }

    // Method to insert text at cursor position
    public void insert(String text) {
        // adding what it looked like before into undo in case the user wants to refresh
        // to this
        EditAction prevAction = new EditAction("undo", cursorPosition, str.toString());
        undo.push(prevAction);

        // inserting text
        str.insert(cursorPosition, text);
    }

    // Method to delete a substring for a certain length
    public void delete(int length) {
        // adding what it looked like before into undo in case the user wants to refresh
        // to this
        EditAction prevAction = new EditAction("undo", cursorPosition, str.toString());
        undo.push(prevAction);

        str.delete(cursorPosition, length + cursorPosition);
    }

    // Method to move the cursor to a wanted point
    public void moveCursor(int position) {
        cursorPosition = position;
    }

    // Method to get all content of the string
    public String getContent() {
        return "\nString:" + str.toString() + "\n";
    }

    // Method to get the position of cursor
    public int getCursor() {
        return cursorPosition;
    }

    // Method to undo action
    public void undoAction() {
        // getting what the string looked like before the last action
        EditAction action = undo.pop();

        // putting what the text currently looks like in the redo list
        EditAction current = new EditAction("redo", cursorPosition, str.toString());
        redo.push(current);

        // getting back the last action
        str.setLength(0);
        str.append(action.text);
        cursorPosition = action.position;

    }

    // Method to redo action
    public void redoAction() {
        // check if there is anything to redo
        if (redo.isEmpty()) {
            System.out.println("Nothing has been undone");
        } else {
            EditAction element = redo.pop();
            str.setLength(0);
            str.append(element.text);
            cursorPosition = element.position;
        }
    }

    // Method to search for occurences
    public ArrayList<Integer> searchText(String searchTerm) {
        ArrayList<Integer> positions = new ArrayList<>();
        StringBuilder worker = new StringBuilder();
        worker.append(str.toString());
        int lastOccurence = 0;

        // go through string and find each occurence and the first character position
        // starting at 0
        while (worker.length() > 0) {
            int index = worker.indexOf(searchTerm) + lastOccurence;

            // if present
            if (index != -1) {
                positions.add(index);
                worker.delete(0, index + searchTerm.length());
                lastOccurence += index;
            } else {
                break;
            }
        }

        return positions;
    }

    // Method to replace text
    public void replaceText(String searchTerm, String replacement) {
        // adding this into the undo stack in case this needs to be undone
        EditAction newAction = new EditAction("undo", cursorPosition, str.toString());
        undo.push(newAction);

        StringBuilder worker = new StringBuilder();
        worker.append(str.toString());

        // find each iteration of the wanted search term
        while (worker.length() > 0) {
            int index = worker.indexOf(searchTerm);

            // if present
            if (index != -1) {
                worker.replace(index, index+searchTerm.length(), replacement);
            } else {
                break;
            }
        }

        //setting str
        str.setLength(0);
        str.append(worker.toString());
    }

    // Method to open file
    public void openFile(String fileName) {
        try {
            // opening file and getting the contents
            curFile = fileName;
            File newFile = new File(fileName);
            Scanner reader = new Scanner(newFile);
            while (reader.hasNextLine()) {
                str.append(reader.nextLine());
            }
            // close reader
            reader.close();

        } catch (IOException e) {
            System.err.println("IOException");
        }
    }

    // Method to save file
    public void saveFile() {
        // checking if a file is currently open
        if (curFile.equals("")) {
            System.err.println("No file open");
            return;
        }

        try {
            // printing to the current file that is open
            FileWriter fw = new FileWriter(curFile);
            BufferedWriter writer = new BufferedWriter(fw);

            writer.write(str.toString());

            writer.close();

        } catch (IOException e) {
            System.err.println("IOException");
        }

    }

    // Method for save as
    public void saveAs(String fileName) {
        // saving current text to a given fileName
        try {
            FileWriter fw = new FileWriter(fileName);
            BufferedWriter writer = new BufferedWriter(fw);

            writer.write(str.toString());

            writer.close();

        } catch (IOException e) {
            System.err.println("IOException");
        }

    }

    // Method to copy substrings
    public void copy(int start, int end) {
        // getting the substring at the given point;
        String subsString = str.substring(start, end);

        // setting copy to the wanted substring
        copy = subsString;

    }

    // Method to paste
    public void paste(int position) {
        // saving the text before the paste is completed into the undo stack
        EditAction action = new EditAction("undo", cursorPosition, str.toString());
        undo.push(action);

        // System.out.println(copy);
        // inserting copied words into position
        str.insert(position, copy);
    }
}