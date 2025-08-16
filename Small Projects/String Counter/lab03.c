#include <stdio.h>
#include <stdlib.h>
#include <string.h>

/* substringcounter() counts the number of instances of a substring in a string
 * @param string: a string that is used by substring to find number of instances
 * @param substring: a string that may or may not occur in the string
 * @return int: number of instances found
 *
 */

int substringCounter(char* string, char* substring){
	//initializing variables
	int lengthString = strlen(string);
	int lengthSubstring = strlen(substring);
	int instanceCount = 0;
	int i = 0;
	int j = 0;
	int count = 0;

	//while loop to go through the main string and find instances of substring
	while (i < lengthString){
		j = 0;
		count = 0;

		//counts up when the char of string matches the substring
		while (string[i] == substring[j] && count < lengthSubstring){
			count++;
			i++;
			j++;
		}
		//counts up if the string matches the substring based on count
		if (count == lengthSubstring){
			instanceCount++;
			count= 0;
		} else {
			i++;
		}
	}

	//returns number of instances
	return instanceCount;
}

int main(int argc, char* argv[]) {
	//checking if there is enough things in argv
	if (argc != 3){
		printf("invalid arguments  \n");
		exit(0);
	}

	//calling substringCounter to find how many instances there are of the substring
	int numOfInstances = substringCounter(argv[1], argv[2]);

	printf("%d  \n", numOfInstances);

}
