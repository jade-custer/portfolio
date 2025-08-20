//initial includes 
#include <stdio.h>

// returns i if nums[i]==target, -1 if target is not found
int linearSearch(int nums[], int length, int target) {
    //initializing variables 
    int i = 0;

    //while there is still numbers in the array check if the target exists 
    while (i < length){
    	//if found
    	if (nums[i] == target){
    		return i;
    	}
    	i++;
    }

	//if not found
    return -1;
}

int main(int argc, char* argv[]) {

    // create the nums array based on the numbers in the file
    // filename is given in argv[1]
    // the first number in the file is the size of the array (= the number of numbers to be in the array))
	FILE* file = fopen(argv[1],"r");
	int num;
	int length =fscanf(file,"%d", &num);
	int nums[length];
	int	i =0;

	while(fscanf(file,"%d", &num) == 1){
		nums[i] = num;
		i++;		
	}
	
    //get a target number from user
    int target;
    printf("What number are you looking for");
    scanf("%d",&target);

    if (linearSearch(nums, length, target) < 0) {
        printf("Not Found\n");
    } else {
        printf("Found\n");
    }
}
