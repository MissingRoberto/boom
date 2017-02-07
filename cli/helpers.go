package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/jszroberto/boom"
)

func exitWithError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
func diff(boom *boom.Boom, path string) {
	tmpFile, err := ioutil.TempFile("", "manifest.yml")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	writeFile(tmpFile.Name(), boom.String())
	cmd, _ := exec.Command("wdiff", "-n", "-w", "\033[30;41m", "-x", "\033[0m", "-y", "\033[30;42m", "-z", "\033[0m", tmpFile.Name(), path).Output()
	fmt.Printf("%s", cmd)
	os.Exit(0)
}

func writeFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
