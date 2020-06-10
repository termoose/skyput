package portal

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func Show() error {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		return fmt.Errorf("prompt failed: %v", err)
	}

	fmt.Printf("Selection: %s\n", result)

	return nil
}