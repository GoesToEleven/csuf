// To be placed in the output Go repo at cmd/go.

package main

import (
	"path/filepath"
)

var cmdDeploy = &Command{
	UsageLine: "deploy [deploy flags] [ application_dir | package | yaml_files...]",
	Short:     "deploys your application to App Engine",
	Long: `
Deploy uploads your application files to Google App Engine, and then compiles
and lauches your application.

The argument to this command should be your application's root directory or a
single package which contains an app.yaml file. If you are using the Modules
feature, then you should pass multiple YAML files to deploy, rather than a
directory, to specify which modules to update. If no arguments are provided,
deploy looks in your current directory for an app.yaml file.

The -application flag sets the application ID, overriding the application value
from the app.yaml file.

The -version flag sets the major version, overriding the version value from the
app.yaml file.

The -oauth flag causes authentication to be done using OAuth2, instead of
interactive password auth.

This command wraps the appcfg.py command provided as part of the App Engine
SDK. For help using that command directly, run:
  ./appcfg.py help update
  `,
}

var (
	deployApp   string // deploy -application flag
	deployVer   string // deploy -version flag
	deployOAuth bool   // deploy -oauth flag
)

func init() {
	// break init cycle
	cmdDeploy.Run = runDeploy

	cmdDeploy.Flag.StringVar(&deployApp, "application", "", "")
	cmdDeploy.Flag.StringVar(&deployVer, "version", "", "")
	cmdDeploy.Flag.BoolVar(&deployOAuth, "oauth", false, "")
}

func runDeploy(cmd *Command, args []string) {
	appcfg, err := findAppcfg()
	if err != nil {
		fatalf("goapp serve: %v", err)
	}
	toolArgs := []string{"update"}
	if deployApp != "" {
		toolArgs = append(toolArgs, "--application", deployApp)
	}
	if deployVer != "" {
		toolArgs = append(toolArgs, "--version", deployVer)
	}
	if deployOAuth {
		toolArgs = append(toolArgs, "--oauth2")
	}
	files, err := resolveAppFiles(args)
	if err != nil {
		fatalf("goapp deploy: %v", err)
	}
	runSDKTool(appcfg, append(toolArgs, files...))
}

func findAppcfg() (string, error) {
	devAppserver, err := findDevAppserver()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(devAppserver), "appcfg.py"), nil
}
