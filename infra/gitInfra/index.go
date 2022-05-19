package gitInfra

import (
	"fmt"
	"os"
	"os/exec"
)

func Clone(project string) error {
	cmd := exec.Command("git", "clone", "https://github.com/fanama/next-secure-git.git", project)

	return cmd.Run()
}

func RemoveOrigin(project string) error {
	cmd := exec.Command("git", "remote", "rm", "origin")
	pathDir, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd.Dir = fmt.Sprintf("%s/%s", pathDir, project)

	return cmd.Run()
}
