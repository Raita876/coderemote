package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var (
	version string
	name    string
)

func Exec(cmd ...string) error {
	fmt.Println(strings.Join(cmd[:], " "))

	var c *exec.Cmd

	if len(cmd) < 2 {
		c = exec.Command(cmd[0])
	} else {
		c = exec.Command(cmd[0], cmd[1:]...)
	}

	rStdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}

	rStderr, err := c.StderrPipe()
	if err != nil {
		return err
	}

	out := bytes.NewBuffer(nil)

	wStdout := io.MultiWriter(out, os.Stdout)
	wStderr := io.MultiWriter(out, os.Stderr)

	err = c.Start()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		io.Copy(wStdout, rStdout)
		rStdout.Close()
		wg.Done()
	}()

	go func() {
		io.Copy(wStderr, rStderr)
		rStderr.Close()
		wg.Done()
	}()

	wg.Wait()

	return c.Wait()
}

func main() {
	app := &cli.App{
		Version: version,
		Name:    name,
		Usage:   "code command for remote host.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "workdir",
				Aliases: []string{"w"},
				Value:   "/",
				EnvVars: []string{"CODE_WORKDIR"},
			},
			&cli.StringFlag{
				Name:    "remote-host",
				Aliases: []string{"r"},
				Value:   "remote",
				EnvVars: []string{"CODE_HOST"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.Args().Len() != 1 {
				return xerrors.New("Incorrect number of arguments.")
			}

			path := c.Args().First()
			host := c.String("remote-host")
			workdir := c.String("workdir")

			folderURI := fmt.Sprintf("vscode-remote://ssh-remote+%s%s", host, workdir+"/"+path)

			err := Exec("code", "--folder-uri", folderURI)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
