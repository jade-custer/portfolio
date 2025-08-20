public class Node {
    //iniitializing needed variables
    String fileName;
    String filePath;
    long fileSize;

    Node next;
    Node prev;
    
    public Node(String name, String path, long size){
        this.fileName = name;
        this.filePath = path;
        this.fileSize = size;
        this.next = null;
        this.prev = null;

    }

    //Method to get name of node
    public String getName (){
        return fileName;
    }

    //Method to get node path
    public String getPath(){
        return filePath;
    }

    //Method to get size of file
    public long getSize(){
        return fileSize;
    }


}
