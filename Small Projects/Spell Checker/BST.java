
import java.util.ArrayList;

public class BST {
    private BSTNode root;

    // Method to insert
    public void insert(String word) {
        root = recursInsert(root, word);

    }

    // Recursive insert
    public BSTNode recursInsert(BSTNode root, String word) {
        // checking if root is null
        if (root == null) {
            root = new BSTNode(word);
            return root;
        }

        // comparing things to find where the root should go
        if (word.compareTo(root.word) < 0) {
            root.left = recursInsert(root.left, word);
        } else if (word.compareTo(root.word) > 0) {
            root.right = recursInsert(root.right, word);
        }

        return root;
    }

    // Method to search
    public boolean search(String word) {
        return recursSearch(root, word);
    }

    // Recursively searching
    public boolean recursSearch(BSTNode root, String word) {
        // if root = null
        if (root == null) {
            return false;
        }

        // if word is found return true
        if (word.equals(root.word)) {
            return true;
        }

        // if word is not found yet see which see of the tree you have to go down
        if (word.compareTo(root.word) < 0) {
            return recursSearch(root.left, word);
        } else {
            return recursSearch(root.right, word);
        }
    }

    // removing something from BST
    public BSTNode remove(String word) {
        BSTNode node = recursRemove(root, word);
        return node;
    }

    // Recursively removing
    public BSTNode recursRemove(BSTNode root, String word) {

        BSTNode parent = null;

        BSTNode current = root;

        //traversing to node
        while (current != null && !current.word.equals(word)) {
            parent = current;

            if (word.compareTo(current.word) < 0) {
                current = current.left;
            } else if (word.compareTo(current.word) > 0) {
                current = current.right;
            }
        }

        if (current == null) {
            return root;
        }

        // Case 1:
        if (current.left == null && current.right == null) {
        
            if (current != root){
                if (parent.left == current){
                        parent.left = null;
                }

                if (parent.right == current){
                    parent.right = null;
                }
            } else {
                root = null;
            }

        //Case 2:
        } else if (current.left != null && current.right != null){
            BSTNode successor = minimumVal(current.right);

            current.word = successor.word;

            current.right = recursRemove(current.right, successor.word);
        
        //Case 3:
        } else {
            BSTNode child;

            if (current.left!= null){
                child = current.left;
            } else {
                child = current.right;
            }

            if (current != root){
                if (current == parent.left){
                    parent.left = child;
                } else {
                    parent.right = child;
                }
            } else {
                root = child;
            }
        }

        //return root in case it was changed
        return root;
    }

    // Method to get minimum value in tree
    public BSTNode minimumVal(BSTNode root) {
        // traversing to the end of the left subtree until the minimum is gotten.
        while (root.left != null) {
            root = root.left;
        }

        return root;
    }

    // Method to colleact suggestions
    public ArrayList<String> collectSuggestions(String prefix, int minLength, int maxLength) {
        ArrayList<String> suggestions = new ArrayList<>();
        int max = 10;
        recursCollectSuggestions(root, prefix, minLength, maxLength, suggestions,max);

        return suggestions;
    }

    // collecting suggestions recursively
    public void recursCollectSuggestions(BSTNode root, String prefix, int minLength, int maxLength, ArrayList<String> suggestions,int max) {
        if (root == null) {
            return;
        }

        //checking if limit has been hit to prevent stack overflow
        if (suggestions.size() >= max){
            return;
        }

        // traversing left subtree
        recursCollectSuggestions(root.left, prefix, minLength, maxLength, suggestions,max);

        // checking the word of the current word
        String cur = root.word;

        if (cur.startsWith(prefix) && cur.length() >= minLength && cur.length() <= maxLength) {
            suggestions.add(cur);
        }

        // traversing right subtree
        recursCollectSuggestions(root.right, prefix, minLength, maxLength, suggestions, max);

    }

}
