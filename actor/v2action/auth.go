package v2action

import (
	"fmt"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/uaa/constant"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
)

// Authenticate authenticates the user in UAA and sets the returned tokens in
// the config.
//
// It unsets the currently targeted org and space whether authentication
// succeeds or not.
func (actor Actor) Authenticate(ID string, secret string, origin string, grantType constant.GrantType) error {
	if grantType == constant.GrantTypePassword && actor.Config.UAAGrantType() == string(constant.GrantTypeClientCredentials) {
		return actionerror.PasswordGrantTypeLogoutRequiredError{}
	}

	actor.Config.UnsetOrganizationAndSpaceInformation()

	accessToken, refreshToken, err := actor.UAAClient.Authenticate(ID, secret, origin, grantType)
	if err != nil {
		actor.Config.SetTokenInformation("", "", "")
		return err
	}

	accessToken = fmt.Sprintf("bearer %s", accessToken)
	actor.Config.SetTokenInformation(accessToken, refreshToken, "")

	if grantType != constant.GrantTypePassword {
		actor.Config.SetUAAGrantType(string(grantType))
		actor.Config.SetUAAClientCredentials(ID, secret)
	}

	return nil
}

func (actor Actor) GetLoginPromptsAndSaveUAAServerURL() (prompts map[string]coreconfig.AuthPrompt, apiErr error) {
	loginThing := actor.UAAClient.GetLogin()

	prompts = make(map[string]coreconfig.AuthPrompt)
	for key, val := range resource.Prompts {
		prompts[key] = coreconfig.AuthPrompt{
			Type:        knownAuthPromptTypes[val[0]],
			DisplayName: val[1],
		}
	}

	if resource.Links["uaa"] == "" {
		uaa.config.SetUaaEndpoint(uaa.config.AuthenticationEndpoint())
	} else {
		uaa.config.SetUaaEndpoint(resource.Links["uaa"])
	}
	return
}
