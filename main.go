package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	CallApplication := "bin\\64bit\\obs64.exe"
	CallApplicationArguments := "--portable"

	AppLauncherExe, err := os.Executable()
	if err != nil {
		panic(err)
	}
	AppLauncherExePath := filepath.Dir(AppLauncherExe)

	if !filepath.IsAbs(CallApplication) {
		CallApplication = AppLauncherExePath + "\\" + CallApplication
	}

	if _, err := os.Stat(CallApplication); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("%v does not exists!\n", CallApplication)
		os.Exit(1)
	}

	var cmd *exec.Cmd
	if len(CallApplicationArguments) > 0 {
		args := strings.Fields(CallApplicationArguments)
		cmd = exec.Command(CallApplication, args[0:]...)
	} else {
		cmd = exec.Command(CallApplication)
	}
	cmd.Dir = filepath.Dir(CallApplication)
	_, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
