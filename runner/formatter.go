package runner

import "github.com/Scalingo/go-scalingo"

func formatCPU(stats []*scalingo.ContainerStat) int {
	cpuUsage := 0
	cpuCount := 0
	for _, stat := range stats {
		cpuUsage = cpuUsage + stat.CpuUsage
		cpuCount++
	}

	return int(float64(cpuUsage) / float64(cpuCount))
}

func formatRam(stats []*scalingo.ContainerStat) int {
	ramUsage := 0
	ramCount := 0

	for _, stat := range stats {
		ramUsage += int(float64(stat.MemoryUsage) / float64(stat.MemoryLimit) * 100.0)
		ramCount++
	}

	return int(float64(ramUsage) / float64(ramCount))
}
