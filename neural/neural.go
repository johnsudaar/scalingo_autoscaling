package neural

import (
	"log"
	"time"

	"github.com/johnsudaar/scalingo_autoscaling/scaler"
)

type NeuralScaler struct {
	CPUValues []float64
	RAMValues []float64
	Size      int
	Up        *Neuron
	Down      *Neuron
	SleepTime time.Duration
	LastScale *time.Time
}

func makeArray(size int) []float64 {
	values := make([]float64, size)
	values[0] = 1.0
	for i := 1; i < size; i++ {
		values[i] = -1.0
	}

	return values
}

func NewNeuralScaler(size int, sleep time.Duration) *NeuralScaler {
	return &NeuralScaler{
		CPUValues: makeArray(size),
		RAMValues: makeArray(size),
		Size:      size,
		Up:        NewNeuron(size),
		Down:      NewNeuron(size),
		LastScale: nil,
		SleepTime: sleep,
	}
}

func NewTrainerNeuralScaler(sleep time.Duration) *NeuralScaler {

	n := NewNeuralScaler(11, sleep)
	n.Up.Weights = []float64{-1.808803281925636, -0.12332832192460394, -0.003511034160217701, -0.00021938544015336243, 0.001397166811978586, -0.007971006831471871, 0.009197114164396794, 0.027941779613391556, 0.05891383981789711, 1.7892363005629277, 0.5081301372528442}
	n.Down.Weights = []float64{0.29599708354150445, 0.41105983706390065, 0.026670820385796722, 0.008506771195584961, 0.016459138419652703, -0.026935551950602297, 0.005314572154632299, -0.17530761260368588, -0.0381626996644667, -0.4616092957156408, -0.37798198642016756}
	return n
}

func (n *NeuralScaler) ProcessStat(cpu, ram int, app string) {
	cpus := append(n.CPUValues, float64(cpu)/100.0)[1:]
	rams := append(n.RAMValues, float64(ram)/100.0)[1:]
	cpus[0] = 1.0
	rams[0] = 1.0
	log.Println(cpus)
	log.Println(rams)
	n.CPUValues = cpus
	n.RAMValues = rams
	if cpus[1] == -1 {
		return
	}

	if n.LastScale != nil && n.LastScale.Add(n.SleepTime).After(time.Now()) {
		return
	}
	now := time.Now()

	cpuUpF, errCpuUp := n.Up.Evaluate(n.CPUValues)
	cpuDownF, errCpuDown := n.Down.Evaluate(n.CPUValues)
	ramUpF, errRamUp := n.Up.Evaluate(n.CPUValues)
	ramDownF, errRamDown := n.Down.Evaluate(n.CPUValues)

	cpuUp := cpuUpF > 0
	cpuDown := cpuDownF > 0
	ramUp := ramUpF > 0
	ramDown := ramDownF > 0

	if errCpuUp != nil {
		log.Println(errCpuUp)
		return
	}

	if errCpuDown != nil {
		log.Println(errCpuDown)
		return
	}
	if errRamUp != nil {
		log.Println(errRamUp)
		return
	}
	if errRamDown != nil {
		log.Println(errRamDown)
		return
	}

	if ramUp || cpuUp {
		scaler.ScaleUP(app)
		n.LastScale = &now
	}

	if ramDown && cpuDown {
		scaler.ScaleDown(app)
		n.LastScale = &now
	}

}
