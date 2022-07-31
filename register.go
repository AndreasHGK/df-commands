package df_commands

import "github.com/df-mc/dragonfly/server/cmd"

func RegisterAll() {
	cmd.Register(cmd.New("gamemode", "change someone's gamemode", []string{"gm"},
		Gamemode{},
	))
	cmd.Register(cmd.New("kill", "kill a player", []string{},
		Kill{},
	))
	cmd.Register(cmd.New("teleport", "teleport to a destination", []string{"tp"},
		TeleportToCoords{},
		TeleportToPlayer{},
		TeleportPlayerToCoords{},
		TeleportPlayerToPlayer{},
	))
	cmd.Register(cmd.New("kick", "disconnect a player", []string{},
		Kick{},
	))
}
