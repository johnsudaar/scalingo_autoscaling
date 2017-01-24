package scaler

import (
	"log"

	"github.com/Scalingo/go-scalingo"
	"github.com/johnsudaar/scalingo_autoscaling/config"
	"gopkg.in/errgo.v1"
)

func getWebContainers(c *scalingo.Client, app string) (int, error) {
	containers, err := c.AppsPs(app)
	if err != nil {
		return -1, errgo.Mask(err)
	}

	for _, container := range containers {
		if container.Name == "web" {
			return container.Amount, nil
		}
	}

	return -1, errgo.New("Cannot find web container")
}

func Scale(c *scalingo.Client, app string, amount int) error {
	container := scalingo.Container{
		Name:   "web",
		Amount: amount,
	}

	scaleParams := &scalingo.AppsScaleParams{
		Containers: []scalingo.Container{container},
	}
	_, err := c.AppsScale(app, scaleParams)

	if err != nil {
		return errgo.Mask(err)
	} else {
		return nil
	}
}

func ScaleUP(app string) (bool, error) {
	log.Println("[SCALER] Scale up order received for: " + app)

	c := config.ScalingoClient()
	web, err := getWebContainers(c, app)
	if err != nil {
		return false, errgo.Mask(err)
	}

	if web < config.MAX_CONTAINER {
		log.Println("[SCALER] Scaling up: " + app)

		if err := Scale(c, app, web+1); err != nil {
			return false, errgo.Mask(err)
		}
		return true, nil
	}

	return false, nil
}

func ScaleDown(app string) (bool, error) {
	log.Println("[SCALER] Scale down order received for: " + app)
	c := config.ScalingoClient()
	web, err := getWebContainers(c, app)
	if err != nil {
		return false, errgo.Mask(err)
	}

	if web > config.MIN_CONTAINER {
		log.Println("[SCALER] Scaling down: " + app)
		if err := Scale(c, app, web-1); err != nil {
			return false, errgo.Mask(err)
		}
		return true, nil
	}
	return false, nil
}
