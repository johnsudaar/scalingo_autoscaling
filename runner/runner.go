package runner

import (
	"log"
	"strconv"
	"time"

	"github.com/johnsudaar/scalingo_autoscaling/config"
)

type Runner struct {
	Scaler  Scaler
	App     string
	Refresh int
}

func NewRunner(scaler Scaler, app string, refresh int) *Runner {
	return &Runner{
		Scaler:  scaler,
		App:     app,
		Refresh: refresh,
	}
}

func (r *Runner) Start() {
	c := config.ScalingoClient()
	for {
		stats, err := c.AppsStats(r.App)
		if err != nil {
			log.Fatal(err.Error())
		}

		ram := formatRam(stats.Stats)
		cpu := formatCPU(stats.Stats)
		log.Println("[RUNNER] App " + r.App + " CPU: " + strconv.Itoa(cpu) + "% RAM: " + strconv.Itoa(ram) + "%")

		r.Scaler.ProcessStat(cpu, ram, r.App)

		time.Sleep(time.Duration(r.Refresh) * time.Second)
	}
}
