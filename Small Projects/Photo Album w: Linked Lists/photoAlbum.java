import javax.swing.*;
import java.awt.*;
import java.awt.image.BufferedImage;
import javax.imageio.ImageIO;
import java.io.File;
import java.io.FilenameFilter;
import java.io.IOException;

import java.util.logging.Level;
import java.util.logging.Logger;

public class photoAlbum extends JFrame {
    private static final Logger LOGGER = Logger.getLogger(photoAlbum.class.getName());
    private linkedList imageFiles;
    private JLabel imageLabel;
    private JLabel statusLabel;
    private Node cur = null;
    
    public photoAlbum (String directoryPath) {
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setTitle("Image Display");
        setLayout(new BorderLayout());

        imageLabel = new JLabel();
        add(imageLabel, BorderLayout.CENTER);

        //creating needed items
        JPanel buttonPanel = new JPanel();
        JButton prevButton = new JButton("Previous");
        JButton nextButton = new JButton("Next");
        JButton resetButton = new JButton("Reset");
        JButton firstButton = new JButton("First Picture");
        JButton lastButton = new JButton("Last Picture");
        
        //adding buttons into jpanel
        buttonPanel.add(prevButton);
        buttonPanel.add(nextButton);
        buttonPanel.add(resetButton);
        buttonPanel.add(firstButton);
        buttonPanel.add(lastButton);
        add(buttonPanel, BorderLayout.SOUTH);

        statusLabel = new JLabel("Status: ");
        add(statusLabel, BorderLayout.NORTH);

        //adding actionListeners
        prevButton.addActionListener(e -> showPreviousImage());
        nextButton.addActionListener(e -> showNextImage());
        resetButton.addActionListener(e -> showFirstImage() );
        firstButton.addActionListener(e -> showFirstPicture());
        lastButton.addActionListener(e-> showLastPicture());

        //loading images in 
        loadImagesFromDirectory(directoryPath);
        if (!imageFiles.isEmpty()) {
            displayImage(cur);
        } else {
            statusLabel.setText("Status: No images found in the specified directory.");
        }

        setSize(800, 600);
        setLocationRelativeTo(null);
    }

    private void loadImagesFromDirectory(String directoryPath) {
        File directory = new File(directoryPath);
        if (!directory.exists() || !directory.isDirectory()) {
            // Convey error message to the viewer 
            return;
        }

        //accepting in files with acceptable file types
        File[] files = directory.listFiles(new FilenameFilter() {
            @Override
            public boolean accept(File dir, String name) {
                String lowercaseName = name.toLowerCase();
                if (lowercaseName.endsWith(".jpg")){
                    return lowercaseName.endsWith(".jpg"); 
                } else if (lowercaseName.endsWith(".png")){
                    return lowercaseName.endsWith(".png");
                } else {
                    return lowercaseName.endsWith(".gif");
                }

            }
        });

        //creating a new linked list
        imageFiles = new linkedList();

        //adding each element to the linked list
        for (int i = 0; i < files.length;i++){
            try{
                String name = files[i].getName();
                String path = files[i].getCanonicalPath();
                long size = files[i].length();
    
                imageFiles.addNode(name, path, size);
            } catch (IOException e){
                System.err.println("IOException");
            }
        }
        //setting cur to the head of the list
        cur = imageFiles.getHead();

        LOGGER.log(Level.INFO, "Found {0} image files in directory: {1}", new Object[]{files.length, directoryPath});
        statusLabel.setText("Status: Found " + files.length + " image(s).");
    }

    private void displayImage(Node newCur) {
        //get the wanted image
        try {
            File imageFile = new File (newCur.getPath());
            System.out.println(imageFile.getPath());
            BufferedImage image = ImageIO.read(imageFile);
            if (image != null) {
                //scaling the image
                Image scaledImage = image.getScaledInstance(700, 500, DO_NOTHING_ON_CLOSE);
                imageLabel.setIcon(new ImageIcon(scaledImage));
            } else {
                LOGGER.log(Level.WARNING, "Failed to read image file: {0}", imageFile.getPath());
                statusLabel.setText("Status: Failed to read image file.");
            }
        } catch (IOException e) {
            LOGGER.log(Level.SEVERE, "Error loading image", e);
            statusLabel.setText("Status: Error loading image - " + e.getMessage());
        }
        
    }

    // Method to show the previous image 
    private void showPreviousImage() {
        displayImage(cur.prev);
        cur = cur.prev;
    }

    // Method to show the next image 
    private void showNextImage() {
        if (cur == imageFiles.getTail()){
            LOGGER.log(Level.INFO, "Reached end of photo album");
        }
       displayImage(cur.next);
       cur = cur.next;
       
    }

    //Method to show the first picture
    private void showFirstPicture(){
        cur = imageFiles.getHead();
        displayImage(cur);

    }

    //Method to show the last picture
    private void showLastPicture (){
        cur = imageFiles.getTail();
        displayImage(cur);

    }

    //Method to reset to first image 
    private void showFirstImage(){
        cur = imageFiles.getHead();
        displayImage(cur);
    }

    public static void main(String[] args) {
        
        SwingUtilities.invokeLater(() -> {
            String directoryPath = "images/"; // Replace with your image directory path
            photoAlbum display = new photoAlbum(directoryPath);
            display.setVisible(true);
        });
    }
}