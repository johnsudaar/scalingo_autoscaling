package threshold

import (
	"log"
	"time"

	"github.com/johnsudaar/scalingo_autoscaling/scaler"
)

type ThresholdScaler struct {
	CPUHighThreshold int
	CPULowThreshold  int
	RAMHighThreshold int
	RAMLowThreshold  int
	SleepTime        time.Duration
	LastScale        *time.Time
}

func NewThresholdScaler(cpuMin, cpuMax, ramMin, ramMax int, sleep time.Duration) *ThresholdScaler {
	return &ThresholdScaler{
		CPUHighThreshold: cpuMax,
		CPULowThreshold:  cpuMin,
		RAMHighThreshold: ramMax,
		RAMLowThreshold:  ramMin,
		LastScale:        nil,
		SleepTime:        sleep,
	}
}

func (t *ThresholdScaler) ProcessStat(cpu, ram int, app string) {
	// Time between scales
	if t.LastScale != nil && t.LastScale.Add(t.SleepTime).After(time.Now()) {
		return
	}
	now := time.Now()

	if cpu > t.CPUHighThreshold {
		scaled, err := scaler.ScaleUP(app)
		if err != nil {
			log.Println(err)
			return
		}
		if scaled {
			t.LastScale = &now
		}
		return
	}

	if ram > t.RAMHighThreshold {
		scaled, err := scaler.ScaleUP(app)
		if err != nil {
			log.Println(err)
			return
		}
		if scaled {
			t.LastScale = &now
		}
		return
	}

	if cpu < t.CPULowThreshold && ram < t.RAMLowThreshold {
		scaled, err := scaler.ScaleDown(app)
		if err != nil {
			log.Println(err)
			return
		}
		if scaled {
			t.LastScale = &now
		}

	}
}
