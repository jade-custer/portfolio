public class linkedList {
    Node head;
    Node tail;

    public void addNode(String name,String path, long size ){

        Node newNode = new Node(name,path,size); 

        //edge case if linkedlist doesnt exist 
        if(head == null){
            head = newNode; 
            tail = newNode;
            return; 
        }

        //iterating through list till the enc
        Node current = head;
        while(current != tail){
            current = current.next;

        }  
        //setting up the links 
        current.next = newNode; 
        newNode.prev = current;
        newNode.next = head;
        head.prev = newNode;
        tail = newNode;
       tail.next = head;
    }

    //Method to get head 
    public Node getHead(){
        return head;
    }

    //Method to get tail
    public Node getTail(){
        return tail;
    }

    //Method to check if list is empty
    public boolean isEmpty(){
        if(head == null){
            return true; 
        }
        else{
            return false; 
        }
    }
   
    //Method to get size of list
    public int size(){
        Node current = head;
        int size = 0;
       
        while (current != tail) {
            size++;
            current = current.next;
        }
        return size;
    }
}
