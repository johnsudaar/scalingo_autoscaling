package neural

import (
	"log"
	"strconv"

	"gopkg.in/errgo.v1"
)

type NeuralNetwork struct {
	Layers     [][]*Neuron
	LastNeuron *Neuron
	LayerCount int
	LayerSize  int
}

func NewNeuralNetwork(neuronSize, layerCount int) *NeuralNetwork {
	layers := make([][]*Neuron, layerCount)
	for i := 0; i < layerCount; i++ {
		layers[i] = make([]*Neuron, neuronSize)
		for j := 0; j < neuronSize; j++ {
			layers[i][j] = NewNeuron(neuronSize)
		}
	}

	return &NeuralNetwork{
		Layers:     layers,
		LayerCount: layerCount,
		LayerSize:  neuronSize,
		LastNeuron: NewNeuron(neuronSize),
	}
}

func (n *NeuralNetwork) Evaluate(values []float64) (float64, error) {
	layerValues := make([]float64, n.LayerSize)
	oldLayerValues := values
	var err error

	for i := 0; i < n.LayerCount; i++ {
		for j := 0; j < n.LayerSize; j++ {
			layerValues[j], err = n.Layers[i][j].Evaluate(oldLayerValues)
			if err != nil {
				return 0.0, errgo.Mask(err)
			}
		}
		oldLayerValues = layerValues
	}

	result, err := n.LastNeuron.Evaluate(oldLayerValues)

	if err != nil {
		return 0.0, errgo.Mask(err)
	}
	return result, err
}

func (n *NeuralNetwork) Print() {
	for i := 0; i < n.LayerCount; i++ {
		log.Println(" ### LAYER " + strconv.Itoa(i) + " ### ")
		for j := 0; j < n.LayerSize; j++ {
			log.Println("Neuron " + strconv.Itoa(j) + ":  ")
			log.Println(n.Layers[i][j].Weights)
		}
	}
	log.Println("Last neuron : ")
	log.Println(n.LastNeuron.Weights)
}
