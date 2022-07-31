package df_commands

import "github.com/df-mc/dragonfly/server/world"

func StringToGamemode(s string) (world.GameMode, bool) {
	switch s {
	case "survival", "s", "0":
		return world.GameModeSurvival, true
	case "creative", "c", "1":
		return world.GameModeCreative, true
	case "adventure", "a", "2":
		return world.GameModeAdventure, true
	case "spectator", "v", "3":
		return world.GameModeSpectator, true
	default:
		return nil, false
	}
}

func GamemodeToString(gm world.GameMode) string {
	switch true {
	case gm == world.GameModeSurvival:
		return "Survival"
	case gm == world.GameModeCreative:
		return "Creative"
	case gm == world.GameModeAdventure:
		return "Adventure"
	case gm == world.GameModeSpectator:
		return "Spectator"
	default:
		return "Unknown"
	}
}
