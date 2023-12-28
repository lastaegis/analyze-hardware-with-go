package detection

import (
	"github.com/shirou/gopsutil/cpu"
	"log"
	"runtime"
	"time"
)

func CPUDetection() int {
	return runtime.NumCPU()
}

func TotalUsageCPU() float64 {
	percentageTotalCPU, err := cpu.Percent(time.Second, false)

	if err != nil {
		log.Fatal(err)
	}

	return percentageTotalCPU[0]
}

func TotalUsageEachCPU() []float64 {
	percentageAllCPU, err := cpu.Percent(time.Second, true)

	if err != nil {
		log.Fatal(err)
	}

	return percentageAllCPU

}
