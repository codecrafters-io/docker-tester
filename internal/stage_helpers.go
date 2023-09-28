package internal

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func initRandom() {
	if seed := os.Getenv("CODECRAFTERS_RANDOM_SEED"); seed != "" {
		seedInt, err := strconv.Atoi(seed)

		if err != nil {
			panic(err)
		}

		rand.Seed(int64(seedInt))
	} else {
		rand.Seed(time.Now().UnixNano())
	}
}

func randomInt(n int64) int64 {
	return int64(rand.Intn(int(n)))
}

func randomString(n int64) string {
	return strconv.FormatInt(randomInt(n), 10)
}