// To be placed in the output Go repo at cmd/go.

package main

import (
	"errors"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
)

var cmdServe = &Command{
	UsageLine: "serve [serve flags] [application_dir | package | yaml_files...]",
	Short:     "starts a local development App Engine server",
	Long: `
Serve launches your application on a local development App Engine server.

The argument to this command should be your application's root directory or a
single package which contains an app.yaml file. If you are using the Modules
feature, then you should pass multiple YAML files to serve, rather than a
directory, to specify which modules to serve. If no arguments are provided,
serve looks in your current directory for an app.yaml file.

The -host flag controls the host name to which application modules should bind
(default localhost).

The -port flag controls the lowest port to which application modules should
bind (default 8080).

The -use_mtime_file_watcher flag causes the development server to use mtime
polling for detecting source code changes, as opposed to inotify watches.

The -admin_port flag controls the port to which the admin server should bind
(default 8000).

The -clear_datastore flag clears the local datastore on startup.

This command wraps the dev_appserver.py command provided as part of the
App Engine SDK. For help using that command directly, run:
  ./dev_appserver.py --help
  `,
}

var (
	serveHost       string // serve -host flag
	servePort       int    // serve -port flag
	serveUseModTime bool   // serve -use_mtime_file_watcher flag
	serveAdminPort  int    // serve -admin_port flag
	clearDatastore  bool   // serve -clear_datastore flag
)

func init() {
	// break init cycle
	cmdServe.Run = runServe

	cmdServe.Flag.StringVar(&serveHost, "host", "localhost", "")
	cmdServe.Flag.IntVar(&servePort, "port", 8080, "")
	cmdServe.Flag.BoolVar(&serveUseModTime, "use_mtime_file_watcher", false, "")
	cmdServe.Flag.IntVar(&serveAdminPort, "admin_port", 8000, "")
	cmdServe.Flag.BoolVar(&clearDatastore, "clear_datastore", false, "")
}

func runServe(cmd *Command, args []string) {
	devAppserver, err := findDevAppserver()
	if err != nil {
		fatalf("goapp serve: %v", err)
	}
	toolArgs := []string{
		"--host", serveHost,
		"--port", strconv.Itoa(servePort),
		"--admin_port", strconv.Itoa(serveAdminPort),
		"--skip_sdk_update_check", "yes",
	}
	if serveUseModTime {
		toolArgs = append(toolArgs, "--use_mtime_file_watcher", "yes")
	}
	if clearDatastore {
		toolArgs = append(toolArgs, "--clear_datastore", "yes")
	}
	files, err := resolveAppFiles(args)
	if err != nil {
		fatalf("goapp serve: %v", err)
	}
	runSDKTool(devAppserver, append(toolArgs, files...))
}

func runSDKTool(tool string, args []string) {
	python, err := findPython()
	if err != nil {
		fatalf("could not find python interpreter: %v", err)
	}

	toolName := filepath.Base(tool)

	cmd := exec.Command(python, tool)
	cmd.Args = append(cmd.Args, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err = cmd.Start(); err != nil {
		fatalf("error starting %s: %v", toolName, err)
	}

	// Swallow SIGINT. The tool will catch it and shut down cleanly.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for s := range sig {
			logf("goapp: caught SIGINT, waiting for %s to shut down", toolName)
			cmd.Process.Signal(s)
		}
	}()

	if err = cmd.Wait(); err != nil {
		errorf("error while running %s: %v", toolName, err)
	}
}

func findPython() (path string, err error) {
	for _, name := range []string{"python2.7", "python"} {
		path, err = exec.LookPath(name)
		if err == nil {
			return
		}
	}
	return
}

func findDevAppserver() (string, error) {
	if p := os.Getenv("APPENGINE_DEV_APPSERVER"); p != "" {
		return p, nil
	}
	return "", fmt.Errorf("unable to find dev_appserver.py")
}

// resolveAppFiles returns a list of arguments suitable for passing appcfg.py
// or dev_appserver.py corresponding to the user-provided args.
func resolveAppFiles(args []string) ([]string, error) {
	if len(args) == 0 {
		if fileExists("app.yaml") {
			return []string{"./"}, nil
		}
		return nil, errors.New("no app.yaml file in current directory")
	}

	if len(args) == 1 && !strings.HasSuffix(args[0], ".yaml") {
		if fileExists(filepath.Join(args[0], "app.yaml")) {
			return args, nil
		}
		// Try to resolve this arg as a package.
		if build.IsLocalImport(args[0]) {
			return nil, fmt.Errorf("unable to find app.yaml at %s", args[0])
		}
		pkgs := packages(args)
		if len(pkgs) > 1 {
			return nil, errors.New("only a single package may be provided")
		}
		if len(pkgs) == 0 {
			return nil, fmt.Errorf("unable to find app.yaml at %s (unable to resolve package)", args[0])
		}
		dir := pkgs[0].Dir
		if !fileExists(filepath.Join(dir, "app.yaml")) {
			return nil, fmt.Errorf("unable to find app.yaml at %s", dir)
		}
		return []string{dir}, nil
	}

	// The 1 or more args must all end with .yaml at this point.
	for _, a := range args {
		if !strings.HasSuffix(a, ".yaml") {
			return nil, fmt.Errorf("%s is not a YAML file", a)
		}
		if !fileExists(a) {
			return nil, fmt.Errorf("%s does not exist", a)
		}
	}
	return args, nil
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}
