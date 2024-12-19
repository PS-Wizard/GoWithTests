package selectRacer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(timeout time.Duration, w1, w2 string) (winner string, err error) {
	select {
	case <-ping(w1):
		return w1, nil
	case <-ping(w2):
		return w2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", w1, w2)
	}
}
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch

}
