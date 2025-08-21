package main

func Stop(word string) (exists bool) {
	//getting a map of the stop words
	idx := StopMap()

	//seeing if given word exists
	_, exists = idx[word]
	return
}
