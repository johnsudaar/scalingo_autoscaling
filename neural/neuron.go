package neural

import (
	"log"

	"github.com/johnsudaar/scalingo_autoscaling/config"
	"gopkg.in/errgo.v1"
)

type Neuron struct {
	Weights []float64
	Size    int
}

func NewNeuron(size int) *Neuron {
	n := &Neuron{
		Weights: make([]float64, size+1),
		Size:    size + 1,
	}

	for i := 0; i < size+1; i++ {
		n.Weights[i] = rng.Float64()
	}
	return n
}

func (n *Neuron) Evaluate(values []float64) (float64, error) {
	if len(values) != n.Size-1 {
		return 0.0, errgo.New("Values size do not correspond with Neuron size")
	}

	values = append(values, 1.0)
	value := float64(0)

	for i := 0; i < n.Size; i++ {
		value += values[i] * n.Weights[i]
	}
	return value, nil
}

func (n *Neuron) Learn(values []float64, predicate bool) (bool, error) {
	resultF, err := n.Evaluate(values)
	if err != nil {
		return false, errgo.Mask(err)
	}

	result := resultF > 0

	if result == predicate {
		return true, nil
	}

	oc := 1.0

	if result && !predicate {
		oc = -1.0
	}
	for i := 0; i < n.Size; i++ {
		if i == n.Size-1 {
			// Last value is always 1
			n.Weights[i] = n.Weights[i] + config.NEURAL_EPSILON*oc
		} else {
			n.Weights[i] = n.Weights[i] + config.NEURAL_EPSILON*oc*values[i]
		}
	}

	return false, nil
}

func LearnAll() {
	up := NewNeuron(11)
	down := NewNeuron(11)
	passed := 0
	max := 0
	for passed < 1*1000*1000 {
		resUp, resDown, err := GenerateTest(10, up, down)
		if err != nil {
			log.Fatal(err)
		}

		if resUp && resDown {
			passed++
		} else {
			if passed > max {
				log.Println(passed)
				max = passed
			}
			passed = 0
		}
	}

	log.Println(passed)
	log.Println("UP WEIGHTS : ")
	log.Println(up.Weights)
	log.Println("DOWN WEIGHTS : ")
	log.Println(down.Weights)
}
