package v6

import (
	"errors"
	"net/url"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/uaa/constant"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/command/v6/shared"
)

//go:generate counterfeiter . LoginActor

type LoginActor interface {
	BetterAuthenticate(creds map[string]string, origin string, grantType constant.GrantType) error
	SetTarget(settings v2action.TargetSettings) (v2action.Warnings, error)
	GetLoginPrompts() map[string]coreconfig.AuthPrompt
}

type LoginCommand struct {
	APIEndpoint       string      `short:"a" description:"API endpoint (e.g. https://api.example.com)"`
	Organization      string      `short:"o" description:"Org"`
	Password          string      `short:"p" description:"Password"`
	Space             string      `short:"s" description:"Space"`
	SkipSSLValidation bool        `long:"skip-ssl-validation" description:"Skip verification of the API endpoint. Not recommended!"`
	SSO               bool        `long:"sso" description:"Prompt for a one-time passcode to login"`
	SSOPasscode       string      `long:"sso-passcode" description:"One-time passcode"`
	Username          string      `short:"u" description:"Username"`
	usage             interface{} `usage:"CF_NAME login [-a API_URL] [-u USERNAME] [-p PASSWORD] [-o ORG] [-s SPACE] [--sso | --sso-passcode PASSCODE]\n\nWARNING:\n   Providing your password as a command line option is highly discouraged\n   Your password may be visible to others and may be recorded in your shell history\n\nEXAMPLES:\n   CF_NAME login (omit username and password to login interactively -- CF_NAME will prompt for both)\n   CF_NAME login -u name@example.com -p pa55woRD (specify username and password as arguments)\n   CF_NAME login -u name@example.com -p \"my password\" (use quotes for passwords with a space)\n   CF_NAME login -u name@example.com -p \"\\\"password\\\"\" (escape quotes if used in password)\n   CF_NAME login --sso (CF_NAME will provide a url to obtain a one-time passcode to login)"`
	relatedCommands   interface{} `related_commands:"api, auth, target"`

	UI     command.UI
	Actor  LoginActor
	Config command.Config
}

func (cmd *LoginCommand) Setup(config command.Config, ui command.UI) error {
	ccClient, _, err := shared.NewClients(config, ui, false)
	if err != nil {
		return err
	}
	cmd.Actor = v2action.NewActor(ccClient, nil, config)
	cmd.UI = ui
	cmd.Config = config
	return nil
}

func (cmd *LoginCommand) doStuff(args []string) error {
	if cmd.APIEndpoint == "" {
		apiEndpoint, err := cmd.UI.DisplayTextPrompt("API endpoint")
		if err != nil {
			return err
		}
		cmd.APIEndpoint = apiEndpoint
	} else {
		cmd.UI.DisplayText("API endpoint: {{.APIEndpoint}}", map[string]interface{}{
			"APIEndpoint": cmd.APIEndpoint,
		})
	}
	endpoint, _ := url.Parse(cmd.APIEndpoint)
	if endpoint.Scheme == "" {
		endpoint.Scheme = "https"
	}
	settings := v2action.TargetSettings{
		URL:               endpoint.String(),
		SkipSSLValidation: true,
	}
	_, err := cmd.Actor.SetTarget(settings)
	if err != nil {
		return err
	}

	// Recreating the actor because we now have an auth endpoint we can use
	ccClient, uaaClient, err := shared.NewClients(cmd.Config, cmd.UI, true)
	if err != nil {
		return err
	}
	cmd.Actor = v2action.NewActor(ccClient, uaaClient, cmd.Config)

	cmd.authenticate("", "")

	cmd.UI.DisplayText("API endpoint: {{.APIEndpoint}} (API Version: {{.APIVersion}})", map[string]interface{}{
		"APIEndpoint": cmd.APIEndpoint,
		"APIVersion":  cmd.Config.APIVersion(),
	})
	return nil
}

func (cmd *LoginCommand) Execute(args []string) error {
	if cmd.Config.Experimental() {
		return cmd.doStuff(args)
	}
	return translatableerror.UnrefactoredCommandError{}
}

func (cmd *LoginCommand) authenticate(usernameFlagValue, passwordFlagValue string) error {
	prompts := cmd.Actor.GetLoginPrompts()
	passwordKeys := []string{}
	credentials := make(map[string]string)
	var err error

	if value, ok := prompts["username"]; ok {
		if prompts["username"].Type == coreconfig.AuthPromptTypeText && usernameFlagValue != "" {
			credentials["username"] = usernameFlagValue
		} else {
			credentials["username"], err = cmd.UI.DisplayTextPrompt(value.DisplayName)
			if err != nil {
				return err
			}
		}
	}

	for key, prompt := range prompts {
		if prompt.Type == coreconfig.AuthPromptTypePassword {
			if key == "passcode" || key == "password" {
				continue
			}

			passwordKeys = append(passwordKeys, key)
		} else if key == "username" {
			continue
		} else {
			credentials[key], err = cmd.UI.DisplayTextPrompt(prompt.DisplayName)
			if err != nil {
				return err
			}
		}
	}

	maxLoginTries := 3
	for i := 0; i < maxLoginTries; i++ {

		// ensure that password gets prompted before other codes (eg. mfa code)
		if passPrompt, ok := prompts["password"]; ok {
			if passwordFlagValue != "" {
				credentials["password"] = passwordFlagValue
				passwordFlagValue = ""
			} else {
				credentials["password"], err = cmd.UI.DisplayPasswordPrompt(passPrompt.DisplayName)
				if err != nil {
					return err
				}
			}
		}

		for _, key := range passwordKeys {
			credentials[key], err = cmd.UI.DisplayPasswordPrompt(prompts[key].DisplayName)
			if err != nil {
				return err
			}
		}

		credentialsCopy := make(map[string]string, len(credentials))
		for k, v := range credentials {
			credentialsCopy[k] = v
		}

		cmd.UI.DisplayText("Authenticating...")
		err = cmd.Actor.BetterAuthenticate(credentials, "", constant.GrantTypePassword)

		if err == nil {
			cmd.UI.DisplayOK()
			cmd.UI.DisplayNewline()
			break
		}

		cmd.UI.DisplayError(err)
	}

	if err != nil {
		// In non-spike, use translatableerror
		return errors.New(("Unable to authenticate."))
	}

	return nil
}
