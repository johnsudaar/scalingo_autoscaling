package neuralnetwork

import (
	"log"

	"github.com/NOX73/go-neural"
	"github.com/NOX73/go-neural/engine"
	"github.com/NOX73/go-neural/learn"
	"github.com/johnsudaar/scalingo_autoscaling/config"
	trainer "github.com/johnsudaar/scalingo_autoscaling/neural"
)

type NeuralNetworkScaler struct {
	UpNetwork   *neural.Network
	DownNetwork *neural.Network
	Size        int
}

func NewNeuralNetworkScaler(layers, size int) *NeuralNetworkScaler {
	sizes := make([]int, layers+1)
	for i := 0; i < layers; i++ {
		sizes[i] = size
	}
	sizes[layers] = 1

	return &NeuralNetworkScaler{
		UpNetwork:   neural.NewNetwork(size, sizes),
		DownNetwork: neural.NewNetwork(size, sizes),
		Size:        size,
	}
}

func (n *NeuralNetworkScaler) Learn() {
	n.UpNetwork.RandomizeSynapses()
	n.DownNetwork.RandomizeSynapses()
	upEngine := engine.New(n.UpNetwork)
	downEngine := engine.New(n.DownNetwork)

	upEngine.Start()
	downEngine.Start()
	i := 0

	for {
		predUp, predDown, values := trainer.GenerateTestWithPredicate(n.Size, 2)
		predUpF := []float64{-5}
		predDownF := []float64{-5}
		if predUp {
			predUpF[0] = 5
		}

		if predDown {
			predDownF[0] = 5
		}

		upEngine.Learn(values, predUpF, config.NEURAL_EPSILON)

		downEngine.Learn(values, predDownF, config.NEURAL_EPSILON)
		i = i + 1
		learn.Evaluation(n.UpNetwork, values, predUpF)
		learn.Evaluation(n.DownNetwork, values, predDownF)
		if i == 1000 {
			a, b := n.Evaluate()
			log.Println("-----------")
			log.Println(a)
			log.Println(b)
			i = 0
		}
	}
}

func (n *NeuralNetworkScaler) Evaluate() (float64, float64) {
	correctUp := 0
	correctDown := 0
	for i := 0; i < 10000; i++ {
		predUp, predDown, values := trainer.GenerateTestWithPredicate(n.Size, 2)
		resUp := n.UpNetwork.Calculate(values)
		resDown := n.DownNetwork.Calculate(values)

		if resUp[0] < 0.0 && !predUp {
			correctUp = correctUp + 1
		}

		if resUp[0] > 0.0 && predUp {
			correctUp = correctUp + 1
		}

		if resDown[0] < 0.0 && !predDown {
			correctDown = correctDown + 1
		}

		if resDown[0] > 0.0 && predDown {
			correctDown = correctDown + 1
		}
	}

	return float64(correctUp) / 10000.0, float64(correctDown) / 10000.0
}
