package portal

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/termoose/skyput/config"
)

func Show(config *config.Config) error {
	prompt := promptui.Select{
		Label: "Select Portal",
		Items: config.GetPortals(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		return fmt.Errorf("prompt failed: %v", err)
	}

	config.SetDefaultPortal(result)
	return nil
}