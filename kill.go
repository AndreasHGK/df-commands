package df_commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Kill struct {
	Player cmd.Optional[[]cmd.Target]
}

func (k Kill) Run(src cmd.Source, o *cmd.Output) {
	p, isPlayer := src.(*player.Player)

	targets, ok := k.Player.Load()
	if !isPlayer && !ok {
		// If the user did not provide any target AND is not a player, we cannot run the command.
		o.Errorf("You must run this command in-game.")
		return
	} else if !ok {
		// The user did not provide any targets, default to self.
		targets = []cmd.Target{p}
	}

	// Try to set the gamemode for each target, counting how many times it was sucessfully set, so it can be used in the
	// output message.
	count := 0
	for _, t := range targets {
		target, ok := t.(*player.Player)
		if !ok {
			continue
		}

		target.Hurt(target.MaxHealth()*10, SourceCommand{})
		count++
	}

	// Send the success output message.
	if count == 1 {
		// We have a single target, so we display their name.
		if targets[0] == src {
			o.Printf("You killed yourself.")
			return
		}
		o.Printf("You killed %s.", targets[0].Name())
	} else {
		o.Printf("You killed %d players.", count)
	}
}

type SourceCommand struct{}

func (s SourceCommand) ReducedByArmour() bool {
	return false
}

func (s SourceCommand) ReducedByResistance() bool {
	return false
}

func (SourceCommand) Fire() bool {
	return false
}
