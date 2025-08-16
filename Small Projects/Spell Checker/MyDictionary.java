import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;

public class MyDictionary {
    private ArrayList<BST> table;
    private final int size = 50;
    private int elementCount;

    // Constructor to create basis of dictionary
    MyDictionary() {
        this.table = new ArrayList<>(size);
        this.elementCount = 0;

        for (int i = 0; i < size; i++) {
            table.add(new BST());
        }
    }

    // Method to read from dictionary
    public void readFile(String fileName) {
        try {
            BufferedReader reader = new BufferedReader(new FileReader(fileName));
            String line;

            while ((line = reader.readLine()) != null) {
                addToDictionary(line.trim());
            }

            reader.close();
        } catch (IOException e) {
            System.out.println("IOException");
            System.exit(0);
        }
    }

    // Method to add to dictionary
    public void addToDictionary(String newString) {
        // getting the hash value
        int index = polynomialHashFunction(newString);

        //getting the bucket
        BST curbucket = table.get(index);

        //adding to the bucket
        curbucket.insert(newString);
        elementCount++;
    }

    // Method to check if a word exists in the dictionary
    public boolean contains(String word) {
        //checking if word exists
        int index = polynomialHashFunction(word);
        return table.get(index).search(word);
    }

    // Method to use polynomial hash function(){
    public int polynomialHashFunction(String word) {
        int hashValue = 0;
        int base = 31; // chosen based on a recommendation from the internet
        int length = word.length();

        // getting the hash value by adding up each individual parts
        for (int i = 0; i < length; i++) {
            hashValue = (hashValue * base + (int) word.charAt(i)) % size;
        }

        return hashValue;
    }

    // Method to display load factor
    public void displayLoadFactor() {
        System.out.println((double) elementCount / size);
    }

    // Method to suggest corrections - first two starting letters
    public ArrayList<String> suggestCorrections2(String word) {
        ArrayList<String>suggestions = new ArrayList<>();

        //if nothing was given
        if ("".equals(word)){
            return suggestions;
        }
    
        //getting the first two character and length
        String prefix = word.substring(0,2);
        int minLength = word.length()-1;
        int maxLength = word.length()+1;

        //getting BST
       for (BST buckets:table){
        suggestions.addAll(buckets.collectSuggestions(prefix, minLength, maxLength));
       }
        return suggestions;
    }

    // Method to remove items from the dictionary
    public void removeFromDictionary(String word) {
        int index = polynomialHashFunction(word);

        //removing word from table
        BST tree = table.get(index);

        tree.remove(word);
    }

}
