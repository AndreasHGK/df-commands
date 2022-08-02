package df_commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

type Weather struct {
	Weather  WeatherType        `cmd:"weather"`
	Duration cmd.Optional[uint] `cmd:"duration"`
}

func (c Weather) Run(src cmd.Source, o *cmd.Output) {
	runner, ok := src.(interface {
		World() *world.World
	})
	if !ok {
		o.Errorf("Please run this command in-game")
		return
	}
	w := runner.World()

	t := time.Duration(c.Duration.LoadOr(60)) * time.Second
	switch c.Weather {
	case "clear":
		w.StopRaining()
		w.StopThundering()
		o.Printf("Changing to clear weather")
	case "rain":
		w.StartRaining(t)
		o.Printf("Changing to rainy weather")
	case "thunder":
		w.StartRaining(t)
		w.StartThundering(t)
		o.Printf("Changing to rain and thunder")
	default:
		o.Errorf("Unknown weather type: %s", c.Weather)
		return
	}
}

type WeatherType string

func (w WeatherType) Type() string {
	return "weather"
}

func (w WeatherType) Options(source cmd.Source) []string {
	return []string{
		"clear", "rain", "thunder",
	}
}
