package detection

import (
	"github.com/pbnjay/memory"
)

var (
	ramUsage            uint64
	humanReadableMemory float64
)

func RAMDetection() float64 {
	return humanReadable(memory.TotalMemory())
}

func FreeRAMDetection() float64 {
	ramUsage = memory.TotalMemory() - memory.TotalMemory()
	return humanReadable(memory.FreeMemory())
}

func humanReadable(totalMemory uint64) float64 {
	humanReadableMemory = float64(totalMemory) / 1000000000
	return humanReadableMemory
}
