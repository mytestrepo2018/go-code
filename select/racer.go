package main

import (
	"fmt"
	"net/http"
	"time"
)

func getDuration(a string) (duration time.Duration) {
	start := time.Now()
	http.Get(a)
	duration = time.Since(start)
	return
}

/* func Racer(a, b string) string {
	aDuration := getDuration(a)
	bDuration := getDuration(b)

	if aDuration < bDuration {
		return a
	}

	return b
} */

/* func Racer(a, b string) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(5 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
} */

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(a string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(a)
		close(ch)
	}()
	return ch
}
