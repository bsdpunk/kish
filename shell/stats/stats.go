package stats

import (
	"../commands"
	"fmt"
	"runtime"
)

var StatsSubs = commands.Commands{
	{
		Name:      "PrintMemUsage",
		ShortName: "mem",
		Usage:     "Print Memory Usage of the Shell",
		Action:    PrintMemUsage,
		Category:  "stats",
	},
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
