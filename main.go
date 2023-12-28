package main

import (
	"fmt"
	"os-detection/detection"
	"time"
)

func main() {
	for true {
		fmt.Printf("Operating System: %v", detection.OsDetection())
		fmt.Println()
		fmt.Printf("CPU Core: %v", detection.CPUDetection())
		fmt.Println()
		fmt.Printf("Total CPU Usage: %v", detection.TotalUsageCPU())
		fmt.Println()

		for index, value := range detection.TotalUsageEachCPU() {
			fmt.Printf("Core %v: %v", index, value)
			fmt.Println()
		}

		fmt.Println()
		fmt.Printf("Memory Usage: %v GB of %v GB", detection.FreeRAMDetection(), detection.RAMDetection())
		fmt.Println()
		fmt.Println("===========================================")
		time.Sleep(10000)
	}
}
