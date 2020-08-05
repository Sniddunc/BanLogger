package bot

import (
	"database/sql"

	"github.com/sniddunc/banlogger/internal/ban"
	"github.com/sniddunc/banlogger/internal/kick"
	"github.com/sniddunc/banlogger/internal/lookup"
	"github.com/sniddunc/banlogger/internal/warn"
	"github.com/sniddunc/gcmd"
)

var cmdBase gcmd.Base

// Setup initializes everything to do with gcmd
func Setup(db *sql.DB) {
	// Initialize the command base for use by the bot
	cmdBase = gcmd.New("!")
	cmdBase.UnknownCommandMessage = "Unknown command. Do /help for a list of commands."

	// Attach database
	cmdBase.Set("db", db)

	// Register warn command
	warnCommand := gcmd.Command{
		Name:    "warn",
		Usage:   "warn <profileURL> <reason>",
		Handler: warn.CommandHandler,
	}
	warnCommand.Use(warn.ValidateAndMapArgs)
	cmdBase.Register(warnCommand)

	// Register kick command
	kickCommand := gcmd.Command{
		Name:    "kick",
		Usage:   "kick <profileURL> <reason>",
		Handler: kick.CommandHandler,
	}
	kickCommand.Use(kick.ValidateAndMapArgs)
	cmdBase.Register(kickCommand)

	// Register ban command
	banCommand := gcmd.Command{
		Name:    "ban",
		Usage:   "ban <profileURL> <duration> <reason>\nDuration examples: 1min, 1h, 1d, 1m, 1y, perm",
		Handler: ban.CommandHandler,
	}
	banCommand.Use(ban.ValidateAndMapArgs)
	cmdBase.Register(banCommand)

	// Register lookup command
	lookupCommand := gcmd.Command{
		Name:    "lookup",
		Usage:   "ban <profileURL>",
		Handler: lookup.CommandHandler,
	}
	lookupCommand.Use(lookup.ValidateAndMapArgs)
	cmdBase.Register(lookupCommand)
}
