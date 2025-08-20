import java.io.BufferedReader;

import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Scanner;

/*
 * Main class of this program
 */
public class ereader {

    public static void main (String[] args){
         //creating arraylists
         purchasedBooks pB = new purchasedBooks();
         bookstore books = new bookstore();
         
         //creating scanner 
         Scanner keyboard = new Scanner(System.in);

         //viewing bookstore
         books.viewBooksForSale(pB, keyboard);;

         //viewing purchased books
         pB.viewPurchasedbooks(keyboard);

        //closing keyboard
        keyboard.close();
    }
}

/*
 * frequent class used to keep track of words in book and its frequency 
 */

class frequent {
    private String word;
    private int frequency;

    /*setWord sets a new word and sets the frequency to one
     * param newWord: word to be added 
    */
    public void setWord (String newWord){
        this.word = newWord;
        this.frequency = 1;
    }

    /*
     * updateFrequency updates the frequency of the word by 1
     */
    public void updateFrequency (){
        this.frequency += 1;
    }

    /*
     * getWord getter method
     * returns word attached to object
     */
    public String getWord() {
        return word;
    }

    /*getFrequency getter method
     * returns frequency attached to object 
     */
    public int getFrequency(){
        return frequency;
    }

}

/*
 * purchasedBooks is a class that keeps track of purchased books and various useful methods
 */
class  purchasedBooks {

    //creating new arrayList of purchased books
    public ArrayList purchasedBooks;
    
    //creating a new purchased books arraylist
    public purchasedBooks (){
        purchasedBooks = new ArrayList<>();
    }

    /*
     * viewPurchasedBooks allows the user to select what book they want to read 
     * param keyboard allows users to input user imputs
     */
    public  void viewPurchasedbooks (Scanner keyboard){
        //printing out menu items
        System.out.println("These are the books that you have purchased. Select the number of the one you want to read.");
        for (int i = 0; i < purchasedBooks.size();i++){
           System.out.println("(" + i + ")" + purchasedBooks.get(i));
        }

        //selecting book to read 
        int bookSelection = keyboard.nextInt();

        //reading book
        readBook(purchasedBooks.get(bookSelection).toString(), keyboard);
  
    }

    /*
     * readBook allows the user to go 20 lines forward and backwards and path into different methods
     */
    public void readBook (String txtName, Scanner keyboard){
        //setting up fileName
        String fileName = txtName.replaceAll(" ", "-");
        fileName += ".txt";

        //initializing variables
        ArrayList <String> wholeTxt = new ArrayList<>();
        int curIndex = 0;
        int curLine = 1;

        //used to read 20 lines at a time 
        //based on the example given by Alark Joshi in Campuswire
        try (BufferedReader reader = new BufferedReader(new FileReader(fileName))){
            String newLine;
            String curChunk = "";

            while((newLine = reader.readLine()) != null){
                curChunk += newLine + "\n";
                newLine = reader.readLine();
                curLine += 1;

                //every 20 lines becomes a new index in the arraylist
                if(curLine == 20){
                    wholeTxt.add(curChunk);
                    curIndex += 1;
                    curLine = 1;
                }
            }
        } catch (IOException e){
            System.err.println("Error reading file:" + e.getMessage());
        }


        //printing out user menu and setting up pages
        System.out.println("Now reading " + txtName + ". Options are \n(0)Forward \n(1)Backward \n(2)Search \n(3)Display frequent words \n");
        int curPage = -1;
        int newPrompt = -1; //set to -1 so the first page is 0

        //checking for forward or backward
        while(keyboard.hasNextInt()){
            //reading in next input
           newPrompt = keyboard.nextInt();
            
           //going forward
            if (newPrompt == 0){  
                if(curPage != curIndex){
                    curPage +=1;
                    System.out.println(wholeTxt.get(curPage));
                    System.out.println("current page =" + curPage);
                    System.out.println("(0)Forward \n (1)Backward \n");
                } else {
                    System.out.println("End of book."); //cant go forward if at end of book
                }
            //going backward
            } else if (newPrompt == 1){
                if (curPage != 0){
                    curPage -=1;
                    System.out.println(wholeTxt.get(curPage));
                    System.out.println("current page =" + curPage);
                    System.out.println("(0)Forward \n (1)Backward \n");
                } else {
                    System.out.println("Cannot go back."); //cant go back if at the beginning of the book
                }
            } else {
                break; // if not either option break out of the keyboard loop
            }
        }

        //if searching move to a different method
        if (newPrompt == 2){
            search(keyboard,fileName);
        //if trying to display frequent words (currently not working)
        } else if (newPrompt == 3){
            keyboard.close(); //scanner is used in frequentWords so keyboard needs to be closed 
            frequentWords(fileName);
        }
    }

    /*
     * frequentWords iterates through the text and find the most frequent words and sorts out the stop words given
     * param fileName is the name of the text file
     */
    public void frequentWords (String fileName){
        ArrayList freqList = new ArrayList<>();

        //adding each word into the list and updating frequency 
        
            Scanner newWord = new Scanner(fileName);
            
            //checking to see if there are more text to go through
            while(newWord.hasNext()){
                String word = newWord.next();

                //if word doesn't exist yet then create and add a new object
                if (freqList.indexOf(word) == -1){
                    frequent newObj = new frequent();
                    newObj.setWord(word);
                    freqList.add(newObj);
                //else update the frequency 
                } else {
                    int index = freqList.indexOf(word);
                    frequent curWord = (frequent) freqList.get(index);
                    curWord.updateFrequency();
                }
            }
        

        //sorting the frequencies
        int n = freqList.size(); 
        for (int i = 0; i < n - 1; i++) {
            for (int j = 0; j < n - i - 1; j++) {
                frequent  first = (frequent) freqList.get(j);
                frequent second = (frequent) freqList.get(j+1);

                if (first.getFrequency() > second.getFrequency()) { 
                    // swap temp and arr[i] 
                    frequent temp = first; 
                    freqList.set(j, second);
                    freqList.set(j+1, temp);
                }

            }
        }
    }

    /*
     * search utlizes a given string to search through the text to find it
     */
    public void search (Scanner keyboard, String fileName){
        //voiding current keyboard input and prompting for user imput
        System.out.println("What string are you looking for? \n");
        keyboard.nextLine();

        while (keyboard.hasNext()){

            String wantedWord = keyboard.nextLine();
           
            try (BufferedReader reader = new BufferedReader(new FileReader(fileName))){
                String newLine;
    
                 while((newLine = reader.readLine()) != null){
                    if (newLine.contains(wantedWord)){
                        System.out.println(newLine);
                    }
                    
                    newLine = reader.readLine();
                }
                break;
            } catch (IOException e){
                System.err.println("Error reading file:" + e.getMessage());
            }
        }
    }

    //adding a new book
    public void addBook (String txtName){
        purchasedBooks.add(txtName);
    }

    public String get(int index){
        return purchasedBooks.get(index).toString();    
    }

    public int getSize(){
        return purchasedBooks.size();
    }
        
}

class bookstore {
     //creating new arrayList of bookstore books
     public ArrayList books;
    
     //creating a new purchased books arraylist
     public bookstore (){
         books = new ArrayList<String>();
         books.add("Crime and Punishment");
         books.add("Grimms Fairy Tales");
         books.add("Little Women");
         books.add("Romeo and Juliet");
         books.add("Winnie the Pooh");
     }

     public void  viewBooksForSale (purchasedBooks pB, Scanner keyboard){
        //formatting
        System.out.println("Welcome to the eReader. These are the current options available. Type the name of the book and exit bookstore with exit ");

        //printing out each book in the bookstore
        int size = books.size();

        for(int i = 0; i < size; i++){
            String title = books.get(i).toString();
            System.out.println("(" + i + ")" + title);
        }

       while (keyboard.hasNext()){
            String bookSelection = keyboard.nextLine();

            for (int i = 0; i < books.size(); i++){
                String curBook = books.get(i).toString();
                if (curBook.contains(bookSelection)){
                    pB.addBook(curBook);
                }
            }

            if (bookSelection.equals("exit")){
                break;
            }
       }
     }
}