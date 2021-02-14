package http

import "github.com/manifoldco/promptui"

// RunPrompt is the CLI Entry point for the PromptUI
func RunPrompt() error {
	pattern, err := getPattern()
	if err != nil {
		return err
	}

	packageName, err := getPackageName()
	if err != nil {
		return err
	}

	output, err := getOutputName()
	if err != nil {
		return err
	}

	prefix, err := getPrefix()
	if err != nil {
		return err
	}

	HttpHandle(pattern, packageName, output, prefix)

	return nil
}

func getPrefix() (string, error) {
	httpPrompt := promptui.Prompt{
		Label:   "prefix",
		Default: "files",
	}
	prefix, err := httpPrompt.Run()
	if err != nil {
		return "", err
	}

	return prefix, nil
}

func getPattern() (string, error) {
	httpPrompt := promptui.Prompt{
		Label:   "input",
		Default: defaultPattern,
	}
	pattern, err := httpPrompt.Run()
	if err != nil {
		return "", err
	}

	return pattern, nil
}

func getPackageName() (string, error) {
	httpPrompt := promptui.Prompt{
		Label:   "Name of the Package",
		Default: packageName,
	}
	output, err := httpPrompt.Run()
	if err != nil {
		return "", err
	}

	return output, nil
}

func getOutputName() (string, error) {
	httpPrompt := promptui.Prompt{
		Label:   "Filename of the generated File",
		Default: outputName,
	}
	output, err := httpPrompt.Run()
	if err != nil {
		return "", err
	}

	return output, nil
}
