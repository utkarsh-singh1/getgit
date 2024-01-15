package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

func main() {

}


// Recursive search for folder contain .git folder and add them to goDotFilePath
func scan(folder string) {

	fmt.Println("Found Folders\n\n")

	repo := recursiveScanFolder(folder)

	filepath := getDotFilePath()

	addNewSliceToFile(filepath,repo)

	fmt.Println("Successfully Added \n\n")
	
}

func scanGitFolders( folders []string, folder string) []string {

	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)

	if err != nil {
		log.Fatal(err)
	}

	files , err := f.Readdir(-1)

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var path string

	
	for _, file := range files {

		if file.IsDir() {

			path = folder + "/" + file.Name()

			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}

			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string {

	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.getgit"

	return dotFile
}


func addNewSliceToFile(filepath string, newRepo []string) {
	existingRepos := parseFileLinesToSlice(filepath)

	repo := joinSlices(newRepo, existingRepos)

	dumpStringsSliceSliceToFile(repo, filepath)
} 

