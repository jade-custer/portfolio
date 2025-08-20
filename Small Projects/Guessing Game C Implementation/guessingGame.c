#include <stdlib.h>
#include <stdio.h>
#include <time.h>

int main (int argc,char* argv[]){
	//initialize the random number
	srand(time(0));

	//generate rand number
	int i=rand();

	//find the mod to get the rand number 
	i = i % 10;
	
	if (i == 0){
		i = 10;
	}

	//initializing variables 
	int guess = 0;
	int userGuess = 0;

	while (guess == 0){
		//ask the user what their guess is 
		printf ("What is your guess \n");
		scanf("%d",&userGuess);

		//if higher print higher
		if (userGuess > i ){
			printf("Too high\n");
			
		}

		//if low print low
		else if (userGuess< i ){
			printf("Too low\n");
			
		}

		//if correct
		else{
			guess = 1;
			printf("You guessed it \n");
		}
	}
}
