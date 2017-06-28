package launch

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Exec struct {
	ExampleDir string
}

func (self *Exec) Launch(app *App) error {
	return self.goApp(app)
}

func (self *Exec) goApp(app *App) error {
	appDir := fmt.Sprintf("%s/%s/cmd", self.ExampleDir, app.Type)
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	startup := fmt.Sprintf("%s/%s.json", cwd, app.Id)
	configFile, err := os.Create(startup)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(configFile).Encode(app.Startup); err != nil {
		return err
	}
	configFile.Close()
	cmd := exec.Command("go", "run", "main.go", "-startup", startup)
	log := fmt.Sprintf("%s.log", app.Id)
	cmd.Dir = appDir
	cmd.Stdout, err = os.Create(log)
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return err
	}
	return cmd.Start()
}
