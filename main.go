package main

import "fmt"

func main() {

}


// Recursive search for folder contain .git folder and add them to goDotFilePath
func scan(folder string) {

	fmt.Println("Found Folders\n\n")

	repo := recursiveScanFolder(folder)

	filepath := goDotFilePath()

	addNewSliceToFile(filepath,repo)

	fmt.Println("Successfully Added \n\n")
	
}


