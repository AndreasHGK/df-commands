package df_commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"strings"
)

type Gamemode struct {
	Gamemode GamemodeType
	Player   cmd.Optional[[]cmd.Target]
}

func (g Gamemode) Run(src cmd.Source, o *cmd.Output) {
	p, isPlayer := src.(*player.Player)

	targets, ok := g.Player.Load()
	if !isPlayer && !ok {
		// If the user did not provide any target AND is not a player, we cannot run the command.
		o.Errorf("You must run this command in-game.")
		return
	} else if !ok {
		// The user did not provide any targets, default to self.
		targets = []cmd.Target{p}
	}

	// Figure out what gamemode the user is trying to set.
	gamemode, ok := StringToGamemode(strings.ToLower(string(g.Gamemode)))
	if !ok {
		o.Errorf("Unknown gamemode: %s.", g.Gamemode)
		return
	}

	// Try to set the gamemode for each target, counting how many times it was sucessfully set, so it can be used in the
	// output message.
	count := 0
	for _, t := range targets {
		target, ok := t.(*player.Player)
		if !ok {
			continue
		}

		target.SetGameMode(gamemode)
		count++
	}

	// Send the success output message.
	if count == 1 {
		// We have a single target, so we display their name.
		if targets[0] == src {
			o.Printf("You changed your gamemode to %s.", GamemodeToString(gamemode))
			return
		}
		o.Printf("You changed %s's gamemode to %s.", targets[0].Name(), GamemodeToString(gamemode))
	} else {
		o.Printf("You set the gamemode to %s for %d people.", GamemodeToString(gamemode), count)
	}
}

type GamemodeType string

func (g GamemodeType) Type() string {
	return "gamemode"
}

func (g GamemodeType) Options(source cmd.Source) []string {
	return []string{
		"survival", "s", "0",
		"creative", "c", "1",
		"adventure", "a", "2",
		"spectator", "v", "3",
	}
}
