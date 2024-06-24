package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func setEnvironmentStrings(env Environment) {
	for envName, envValue := range env {
		if envValue.NeedRemove {
			os.Unsetenv(envName)
		} else {
			os.Setenv(envName, envValue.Value)
		}
	}
}

// func getEnvironmentStrings(env Environment) []string {
// 	environments := make([]string, len(env))
// 	for envName, envValue := range env {
// 		if envValue.NeedRemove {
// 			continue
// 		}
// 		envString := envName + "=" + envValue.Value
// 		environments = append(environments, envString)
// 	}
// 	return environments
// }

func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) < 1 {
		fmt.Println("Недостаточно аргументов")
		return 1
	}

	// #nosec G204
	command := exec.Command(cmd[0], cmd[1:]...)
	setEnvironmentStrings(env)
	// command.Env = getEnvironmentStrings(env)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		fmt.Printf("Команда завершилась с ошибкой: %v\n", err)

		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		return 1
	}
	return 0
}
