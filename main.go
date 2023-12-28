package main

import (
	"fmt"
	"os-detection/detection"
)

func main() {
	fmt.Printf("Operating System: %v", detection.OsDetection())
	fmt.Println()
	fmt.Printf("CPU Core: %v", detection.CPUDetection())
	fmt.Println()
	fmt.Printf("Memory Usage: %v GB of %v GB", detection.FreeRAMDetection(), detection.RAMDetection())
}
