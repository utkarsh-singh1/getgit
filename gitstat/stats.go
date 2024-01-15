package gitstat

func stats(email string) {

	commits := processRepositories(email)

	printCommitsStats(commits)
}


func processRepositories(email string) map[int]int {

	filepath := getDotFilePath()

	repos := parseFileLinesToSlice(filepath)

	daysInMap := daysInLastSixMonths

	commits := make(map[int]int, daysInMap)

	for i := daysInMap ; i > 0 ; i-- {
		commits[i] = 0
	}

	for _ , path := range repos {
		commits = fileCommits(email, path, commits)
	}

	return commits
}


