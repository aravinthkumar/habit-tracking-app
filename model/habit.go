package model

type habit struct {
	date                   string
	completedItems         int
	plannedItems           int
	percentageOfCompletion float32
}

type items struct {
	id          int
	name        string
	date        string
	startTime   int
	endTime     int
	isCompleted bool
	comment     string
}
