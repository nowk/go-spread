package spread

import . "github.com/nowk/go-wrker"

// Spread spins up a set number of workers and throws the set of jobs at them
// then shuts the workers down
func Spread(num int, jobs ...Job) <-chan error {
	var n = len(jobs)
	jobschan := make(chan Job, n)

	pool := NewPool(num)
	er := pool.Dispatch(jobschan)

	for _, job := range jobs {
		jobschan <- job
	}

	go pool.Drain()

	return er
}
