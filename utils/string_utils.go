package utils

import (
	"fmt"
	"github.com/cleverua/tuna-timer-api/models"
	"net/http"
	"strings"
)

// NormalizeSlackCustomCommand will extract the subcommand form command (see scheme below_) and update the original command:
// Say the command was: Text = "start Add MongoDB service to docker-compose.yml"
// The method will do this:
//   SubCommand = "start"
//   Text = "Add MongoDB service to docker-compose.yml"
func NormalizeSlackCustomCommand(cmd models.SlackCustomCommand) models.SlackCustomCommand {

	text := strings.TrimSpace(cmd.Text)

	firstSpaceIndex := strings.Index(text, " ")
	if firstSpaceIndex != -1 {
		sub := text[:firstSpaceIndex]
		command := text[firstSpaceIndex+1:]

		cmd.SubCommand = strings.TrimSpace(sub)
		cmd.Text = strings.TrimSpace(command)
	} else {
		cmd.SubCommand = text
		cmd.Text = ""
	}

	return cmd
}

func GetSelfURLFromRequest(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)
}
