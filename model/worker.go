package model

import (
	"math/rand"
	"time"
)

func generateRandomWorkTime(startTime time.Time) time.Time {

	// TODO: well...

	rnd := randomNumber(1000, 3000)
	randomWorkTime := startTime.Add(time.Millisecond * time.Duration(rnd))

	return randomWorkTime
}

func assignJobStage() string {
	stages := []string{"IN_QUEUE", "IN_PROGRESS", "COMPLETED", "FAILED"}

	item := randomNumber(0, len(stages))

	return stages[item]

}

func randomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	randNum := rand.Intn(max-min) + min

	return randNum
}
