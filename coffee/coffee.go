package coffee

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var completeBoilingWater = false
var delayForSwitchJob = time.Second * 3

func MakeASimpleCupOfCoffee() {
	var wg sync.WaitGroup
	wg.Add(2)

	/*
	**	Asynchronous
	**	Wait 3 jobs: "Boil water", "Prepare coffee into the filter", "Press coffee"
	**
	 */
	go func() {
		defer wg.Done()
		BoilingWaterJob(time.Second * 60)
		runtime.Gosched()
	}()

	go func() {
		defer wg.Done()
		DoAJob("Prepare coffee into the filter", time.Second*15)
		DoAJob("Press coffee", time.Second*10)
	}()
	wg.Wait()

	/* Synchronous */
	if completeBoilingWater {
		DoAJob("Put the boiled water into the filter", time.Second*5)
		DoAJob("Wait in some seconds ...", time.Second*60)
		fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), "Enjoy ...")
	}
}

func BoilingWaterJob(duration time.Duration) {
	fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), "Boil water")
	time.Sleep(duration)
	completeBoilingWater = true
	fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), "Ring! Ring! Ring! Completing boiling water ...")
}

func DoAJob(action string, duration time.Duration) {
	time.Sleep(delayForSwitchJob)
	fmt.Printf("%v: %v \n", time.Now().Format("02-Jan-2006 15:04:05"), action)
	time.Sleep(duration)
}
