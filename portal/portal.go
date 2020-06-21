package portal

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/termoose/skyput/config"
)

func Show(portals config.PortalList) (error, string) {
	prompt := promptui.Select{
		Label: "Select Portal",
		Items: portals,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return fmt.Errorf("prompt failed: %v", err), ""
	}

	return nil, result
}