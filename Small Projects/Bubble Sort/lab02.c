#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>

void bubbleSort(int* nums, int length) {
    //your code here does bubble sort
    for( int i = length - 1;i>0; i--){
    	for (int j =1 ; j<=i ; j++){
    		if( nums[j-1] > nums[j]){
    			int temp = nums[j-1];
    			nums[j-1]= nums[j];
    			nums[j] = temp;
    		}
    	}
    }
}

int main(int argc, char* argv[]) {
    // open the file 
    FILE* file = fopen(argv[1],"r");

	//throws an error if no file exists
    if (file == NULL){
    	printf("invalid file name \n");
    	exit(0);
    }
   
    // read the numbers in 
    int num;
    int len;
   	fscanf(file,"%d", &len); 

    //checking if the array is calid if not throws an error 
    if (len <= 0){
    	printf("invalid file format \n");
    	exit(0);
    }

    //setting up the array
    int nums[len];
    int i = 0;
    int r =0; //checks how many nums are put into the array
    
    //adding all the numbers from the file into the array 
    while (fscanf(file,"%d", &num) == 1){ 
   		nums[i] = num;
    	i++;
  		r++;
    }
    
    //checking if the amount of values in the array matches the array length
    if (len != r){
    	printf("invalid input format \n");
    	exit(0);
   	}

    //closing file 
    fclose(file);

    // pass the array to bubbleSort
    bubbleSort(nums,len);

    // print the array content
    for (int k = 0; k < len ; k++){
    	printf("%d", nums[k]);
    	printf (" ");
    }
    printf("\n");
    
}
