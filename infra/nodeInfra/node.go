package nodeinfra

import (
	"encoding/json"
	"fanama/init-project/domain"
	"io/ioutil"
	"os/exec"
)

func InitInfra(port string) error {
	script := domain.NodeScript{Start: "next start -p " + port, Dev: "next dev", Build: "next build", Lint: "next lint", Prepare: "husky install .husky"}
	pack := domain.NodePackage{Name: "test", Version: "1.0", Script: script}

	file, err := json.MarshalIndent(pack, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("package.json", file, 0664)
	if err != nil {
		return err
	}
	return nil
}

func InstallInfraDependencies() error {
	cmd := exec.Command("npm", "install", "next", "react", "react-dom", "sass")
	return cmd.Run()
}

func InstallInfraDevDependencies() error {
	cmd := exec.Command("npm", "install", "-D", "@commitlint/config-conventional", "@types/node", "@types/react", "@types/react-dom", "@typescript-eslint/eslint-plugin", "commitlint", "eslint", "eslint-config-next", "eslint-config-prettier", "eslint-plugin-prettier", "husky", "typescript")
	return cmd.Run()
}
