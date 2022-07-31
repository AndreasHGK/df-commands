package df_commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

type TeleportToCoords struct {
	To mgl64.Vec3
}

func (t TeleportToCoords) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		o.Errorf("Please run this command in-game.")
		return
	}

	p.Teleport(t.To)
	o.Printf("You have been teleported to %v.", t.To)
}

type TeleportToPlayer struct {
	To []cmd.Target
}

func (t TeleportToPlayer) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		o.Errorf("Please run this command in-game.")
		return
	}

	if len(t.To) != 1 {
		o.Errorf("Must provide exactly one player to teleport to.")
		return
	}
	destination := t.To[0].Position()
	p.Teleport(destination)
	o.Printf("You have been teleported to %s.", t.To[0].Name())
}

type TeleportPlayerToCoords struct {
	Player []cmd.Target
	To     mgl64.Vec3
}

func (t TeleportPlayerToCoords) Run(src cmd.Source, o *cmd.Output) {
	count := 0
	for _, ta := range t.Player {
		target, ok := ta.(*player.Player)
		if !ok {
			continue
		}

		target.Teleport(t.To)
		count++
	}

	if count == 1 {
		o.Printf("You teleported %s to %v.", t.Player[0].Name(), t.To)
	} else {
		o.Printf("You teleported %d players to %v.", count, t.To)
	}
}

type TeleportPlayerToPlayer struct {
	Player []cmd.Target
	To     []cmd.Target
}

func (t TeleportPlayerToPlayer) Run(src cmd.Source, o *cmd.Output) {
	if len(t.To) != 1 {
		o.Errorf("Must provide exactly one player as destination.")
		return
	}
	destination := t.To[0].Position()

	count := 0
	for _, ta := range t.Player {
		p, ok := ta.(*player.Player)
		if !ok {
			continue
		}

		p.Teleport(destination)
		count++
	}

	if count == 1 {
		o.Printf("You teleported %s to %s.", t.Player[0].Name(), t.To[0].Name())
	} else {
		o.Printf("You teleported %d players to %s.", count, t.To[0].Name())
	}
}
