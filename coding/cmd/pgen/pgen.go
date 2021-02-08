package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"text/template"

	"github.com/urfave/cli/v2"
)

type ImplTplData struct {
	Pkg string
}

func main() {
	app := cli.NewApp()
	app.Name = "pgen"
	app.Usage = "Generates code for interview prep"
	app.UsageText = "pgen PACKAGE_NAME"
	app.Version = "0.0.1"
	app.Authors = []*cli.Author{
		{
			Name:  "Yida Liu",
			Email: "yida.liu@case.edu",
		},
	}
	app.Writer = os.Stdout
	app.Action = func(c *cli.Context) error {
		if c.NArg() <= 0 {
			return fmt.Errorf("Must provide PACKAGE_NAME")
		}

		pkg := c.Args().First()

		path := filepath.Join(".", pkg)

		// crate path
		syscall.Umask(0)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		} else if err != nil {
			return err
		}

		// create `impl.go`
		impl, err := os.Create(filepath.Join(path, "impl.go"))
		if err != nil {
			return err
		}

		// open template
		tpl, err := template.New("pgen").Parse(implTemplate)
		if err != nil {
			return err
		}

		// generate file
		if err := tpl.Execute(impl, ImplTplData{
			Pkg: pkg,
		}); err != nil {
			return err
		}

		// > ginkgo bootstrap
		cmd := exec.Command("ginkgo", "bootstrap")
		cmd.Dir = path
		if err := cmd.Run(); err != nil {
			return err
		}

		// > ginkgo generate impl.go
		cmd = exec.Command("ginkgo", "generate", "impl.go")
		cmd.Dir = path
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
