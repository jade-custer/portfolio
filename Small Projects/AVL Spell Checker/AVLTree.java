//Jade Custer
//ID: 20682049
//Github ID: jadecuster
//Lab06

//this is the code that was given by Alark Joshi as a started for this project refactored for this project using strings.
import java.util.ArrayList;

public class AVLTree {
    // Node class for the AVL Tree
    private class Node {
        String value;
        Node left;
        Node right;
        int height;

        Node(String value) {
            this.value = value;
            this.height = 1;
        }
    }

    // Root of the AVL Tree
    private Node root;

    // Get the height of a node
    private int height(Node node) {
        // If node is null, return 0 otherwise return the height of the node
        if (node == null) {
            return 0;
        } else {
            return node.height;
        }
    }

    // Update height of a node
    private void updateHeight(Node node) {
        if (node != null) {
            node.height = 1 + Math.max(height(node.left), height(node.right));
        }
    }

    // Get balance factor of a node
    private int getBalanceFactor(Node node) {
        // If node is null, return 0
        // else return the height diff between L and R nodes
        if (node == null) {
            return 0;
        } else {
            return height(node.left) - height(node.right);
        }
    }

    // Right rotation
    /*
     * Before Rotation:
     * 30
     * /
     * 20
     * /
     * 10
     * 
     * After Rotation:
     * 20
     * / \
     * 10 30
     */

    private Node rotateRight(Node y) {
        Node x = y.left;
        Node T2 = x.right;

        // Perform rotation
        x.right = y;
        y.left = T2;

        // Update heights
        updateHeight(y);
        updateHeight(x);

        return x;
    }

    // Left rotation
    /*
     * Before Rotation:
     * 20
     * \
     * 30
     * \
     * 40
     * 
     * After Rotation:
     * 30
     * / \
     * 20 40
     */
    private Node rotateLeft(Node x) {
        Node y = x.right;
        Node T2 = y.left;

        // Perform rotation
        y.left = x;
        x.right = T2;

        // Update heights
        updateHeight(x);
        updateHeight(y);

        return y;
    }

    // Insert a value into the AVL tree
    public void insert(String value) {
        root = insertRecursive(root, value);
    }

    // Recursive insert method
    private Node insertRecursive(Node node, String value) {
        // 1. Perform standard BST insertion
        if (node == null) {
            return new Node(value);
        }

        int compare = value.compareTo(node.value);

        if (compare < 0) {
            node.left = insertRecursive(node.left, value);
        } else if (compare > 0) {
            node.right = insertRecursive(node.right, value);
        } else {
            // Duplicate values are not allowed
            return node;
        }

        // 2. Update height of current node
        updateHeight(node);

        // 3. Get the balance factor
        int balance = getBalanceFactor(node);

        // 4. Perform rotations if needed (4 cases)

        // Left Left Case
        if (balance > 1 && value.compareTo(node.left.value) < 0) {
            return rotateRight(node);
        }

        // Right Right Case
        if (balance < -1 && value.compareTo(node.right.value) > 0) {
            return rotateLeft(node);
        }

        // Left Right Case
        if (balance > 1 && value.compareTo(node.left.value) > 0) {
            node.left = rotateLeft(node.left);
            return rotateRight(node);
        }

        // Right Left Case
        if (balance < -1 && value.compareTo(node.right.value) < 0) {
            node.right = rotateRight(node.right);
            return rotateLeft(node);
        }

        return node;
    }

    // Method to search for a word
    public boolean search(String word) {
        return searchRecursive(root, word);
    }

    // Method to search recursively for a word
    public boolean searchRecursive(Node node, String word) {
        // base case
        if (node == null) {
            return false;
        }

        //comparing the words 
        int compare = word.compareTo(node.value);

        // If found
        if (compare == 0) {
            return true;
        } else if (compare < 0) {
            // Searching left tree
           return searchRecursive(node.left, word);
        } else {
            // Searching right tree
            return searchRecursive(node.right, word);
        }
    }

    //Method to autocomplete words
    public ArrayList<String> autocomplete(String word){
        ArrayList<String> words = new ArrayList<>();

        //finding where the words that start with the same prefix is 
        Node startPoint = findPrefix(root,word);
        
        //collecting suggestions
        if (startPoint != null){
            collector(startPoint, word, words);
        }

        return words;
    }

    //Method to collect suggestions from tree
    public void collector (Node node, String prefix, ArrayList<String> suggestions){
        //if node is null nothing to collect
        if (node != null){
            //checking if the current node starts with the prefix and add it to the list of suggestions if it does
            if (node.value.startsWith(prefix)){
                suggestions.add(node.value);
            }

            //checking left subtree
            collector(node.left, prefix, suggestions);

            //checking right subtree
            collector(node.right, prefix,suggestions);
        }
    }

    //Method to find the given prefix so far
    public Node findPrefix (Node node,String word){
        //base case if node is null
        if (node  == null){
            return null;
        }

        //checking if the current word has the prefix that is wanted
        if (node.value.startsWith(word)){
            return node;
        }

        //comparing the words to see which subtree to go down
        int compare = word.compareTo(node.value);

        //left subtree
        if (compare < 0){
            return findPrefix(node.left, word);
        //right subtree
        } else {
            return findPrefix(node.right,word);
        } 
    }
}