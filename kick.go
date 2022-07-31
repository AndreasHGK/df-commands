package df_commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Kick struct {
	Player  []cmd.Target
	Message cmd.Optional[string]
}

func (k Kick) Run(src cmd.Source, o *cmd.Output) {
	message := k.Message.LoadOr("You have been kicked.")
	count := 0
	for _, t := range k.Player {
		p, ok := t.(*player.Player)
		if !ok {
			continue
		}

		p.Disconnect(message)
		count++
	}

	if count == 1 {
		o.Printf("You have kicked %s.", k.Player[0])
	} else {
		o.Printf("You have kicked %d players.", count)
	}
}
