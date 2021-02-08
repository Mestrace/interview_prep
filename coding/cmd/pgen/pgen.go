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

type implTplData struct {
	Pkg string
}

func main() {
	log.SetOutput(os.Stdout)

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

		// crate path if not exist
		syscall.Umask(0)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0700)
		} else if err != nil {
			log.Fatalf("Failed to create path because %s\n", err.Error())
			return err
		}

		// create `impl.go`
		impl, err := os.Create(filepath.Join(path, "impl.go"))
		if err != nil {
			log.Fatalf(failedToCreateFileStr, "impl.go")
			return err
		}
		defer impl.Close()

		// open template
		tpl, err := template.New("pgen").Parse(implTemplate)
		if err != nil {
			log.Fatalf("Failed to parse template because %s\n", err.Error())
			return err
		}

		// generate file
		if err := tpl.Execute(impl, implTplData{
			Pkg: pkg,
		}); err != nil {
			log.Fatalf("Failed to generate template file because %s\n", err.Error())
			return err
		}

		// > ginkgo bootstrap
		cmd := exec.Command("ginkgo", "bootstrap")
		cmd.Dir = path
		cmd.Dir = path
		sout, err := cmd.Output()
		fmt.Print(string(sout))
		if err != nil {
			log.Fatalf(ginkgoFailedErrStr, err.Error())
			return err
		}

		// > ginkgo generate impl.go
		cmd = exec.Command("ginkgo", "generate", "impl.go")
		cmd.Dir = path
		sout, err = cmd.Output()
		fmt.Print(string(sout))
		if err != nil {
			log.Fatalf("Unable to generate test for %s, because %s\n", "impl.go", err.Error())
			return err
		}

		return nil
	}

	app.Run(os.Args)
}

const (
	ginkgoFailedErrStr = `Execute ginkgo command failed because %s
Ginkgo might not have been installed.
Try run "go get -u github.com/onsi/ginkgo/ginkgo"`
	failedToCreateFileStr = `Failed to create file %s\n`
)
