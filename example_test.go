package stopwatch

import (
	"fmt"
	"testing"
	"time"
)

func TestStopWatch(t *testing.T) {
	stopWatch := New("test")

	stopWatch.Start("[<step 1>]")
	time.Sleep(time.Millisecond * 100)
	stopWatch.Stop()

	stopWatch.Start("[<step 2>]")
	time.Sleep(time.Millisecond * 2)
	stopWatch.Stop()

	stopWatch.Start("[<step 3>]")
	time.Sleep(time.Millisecond * 3)
	stopWatch.Stop()

	stopWatch.Start("[<step 4>]")
	time.Sleep(time.Millisecond * 4)
	stopWatch.Stop()

	fmt.Println(stopWatch.PrettyPrint())
}
