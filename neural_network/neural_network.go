package neuralnetwork

import (
	"github.com/NOX73/go-neural"
	trainer "github.com/johnsudaar/scalingo_autoscaling/neural"
)

type NeuralNetworkScaler struct {
	UpEngine   *neural.Network
	DownEngine *neural.Network
	Size       int
}

func NewNeuralNetworkScaler(layers, size int) *NeuralNetworkScaler {
	sizes := make([]int, layers+1)
	for i := 0; i < layers; i++ {
		sizes[i] = size
	}
	sizes[layers] = 1

	return &NeuralNetworkScaler{
		UpEngine:   neural.NewNetwork(layers+1, sizes),
		DownEngine: neural.NewNetwork(layers+1, sizes),
		Size:       size,
	}
}

func (n *NeuralNetworkScaler) Learn() {
	trainer.GenerateTestWithPredicate(n.Size, 7)
}
