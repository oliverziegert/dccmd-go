package alias

import (
	"errors"
	"fmt"
	"github.com/oliverziegert/dccmd-go/config"

	"github.com/asaskevich/govalidator"
	"github.com/manifoldco/promptui"
	"strconv"
)

func GetDomain(target string) (*string, error) {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("Invalid token variable. (Domain can't be less then 1 char.)")
		}
		if !govalidator.IsDNSName(input) {
			return errors.New("Invalid token variable. (Domain can't be less then 1 char.)")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Domain for target %s", target),
		Default:  "dracoon.team",
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	return &result, nil
}

func GetClientId(target string) (*string, error) {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("Invalid ClientID variable. (ClientID can't be less then 1 char.)")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("ClientID for target %s", target),
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	return &result, nil
}

func GetClientSecret(target string) (*string, error) {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("Invalid Secret variable. (ClientID can't be less then 1 char.)")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("ClientSecret for target %s", target),
		Mask:     '*',
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	return &result, nil
}

func GetReturnFlow(target string) (*config.ReturnFlow, error) {
	prompt := promptui.Select{
		Label: fmt.Sprintf("Select return flow for target %s", target),
		Items: []string{string(config.ReturnFlowBrowser), string(config.ReturnFlowCli)},
		Size:  2,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	returnFlow := config.ReturnFlow(result)
	return &returnFlow, nil
}

func GetBindAddress(target string) (*string, error) {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("Invalid Secret variable. (BindAddress can't be less then 1 char.)")
		}
		if !govalidator.IsIP(input) {
			return errors.New("Invalid Secret variable. (BindAddress must ne a IP address.)")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("Bind address for target %s", target),
		Default:  "127.0.0.1",
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	return &result, nil
}
func GetBindPort(target string) (*uint16, error) {
	validate := func(input string) error {
		if len(input) <= 1 {
			return errors.New("Invalid Secret variable. (BindPort can't be less then 1 char.)")
		}
		if !govalidator.IsInt(input) {
			return errors.New("Invalid Secret variable. (BindPort must be a number.)")
		}
		port, err := strconv.Atoi(input)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid Secret variable. %v", err))
		}
		if port <= 0 || port > 65535 {
			return errors.New("Invalid Secret variable. (BindPort must be between 0 and 65535.)")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    fmt.Sprintf("ClientSecret for target %s", target),
		Default:  "1337",
		Validate: validate,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}
	portS, _ := strconv.Atoi(result)
	port := uint16(portS)
	return &port, nil
}
