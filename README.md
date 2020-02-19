# StopWatch

A simple stop watch, allowing for timing of a number of tasks.exposing total running time and running time for each named task.
Note that [Stopwatch](https://github.com/sanwishe/Stopwatch) is not designed for mutli-thread use.

# Usage

```
go get github.com/sanwishe/Stopwatch
```

example usage

```go
import (
    "time"
    "fmt"
    "github.com/sanwishe/Stopwatch"
)

func yourFunc() {
	stopWatch := stopwatch.New("test")

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

```

The example will expose print such as:

```
StopWatch 'test': running time (ms) = 114
-----------------------------------------
ms        %         Task name
-----------------------------------------
103       90%       [<step 1>]
2         2%        [<step 2>]
3         2%        [<step 3>]
5         4%        [<step 4>]
```
