package detection

import "runtime"

func CPUDetection() int {
	return runtime.NumCPU()
}
