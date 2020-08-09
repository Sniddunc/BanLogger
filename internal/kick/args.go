package kick

import (
	"fmt"
	"strings"

	"github.com/sniddunc/banlogger/internal/steam"
	"github.com/sniddunc/banlogger/pkg/config"
	"github.com/sniddunc/banlogger/pkg/logging"
	"github.com/sniddunc/banlogger/pkg/validation"
	"github.com/sniddunc/gcmd"
)

// ValidateAndMapArgs validates the provided arguments and maps them out to the store on context
// so that they can be easily used by the command handler.
func ValidateAndMapArgs(next gcmd.HandlerFunc) gcmd.HandlerFunc {
	return func(c gcmd.Context) error {
		args := c.Args

		// Check amount of arguments provided
		if len(args) < 2 {
			return fmt.Errorf("Invalid arguments")
		}

		profileURL := args[0]

		// Validate profileURL
		ok, isVanity := validation.ValidateProfileURL(profileURL)
		if !ok {
			return fmt.Errorf("Invalid profile URL provided")
		}

		reason := strings.Join(args[1:], " ")

		if len(reason) < config.ReasonMinLen || len(reason) > config.ReasonMaxLen {
			return fmt.Errorf("Reason must be between %d and %d in length", config.ReasonMinLen, config.ReasonMaxLen)
		}

		// Get the user's SteamID
		steamID, err := steam.GetSteamID(isVanity, profileURL)
		if err != nil {
			return fmt.Errorf("Could not resolve SteamID")
		}

		// Map data to context store
		c.Set("reason", reason)
		c.Set("steamID", steamID)

		logging.Info("kick/args.go", fmt.Sprintf("Kick command passed arguments check. steamID: %s | reason: %s", steamID, reason))

		return next(c)
	}
}
