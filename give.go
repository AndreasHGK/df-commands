package df_commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"golang.org/x/text/language"
	"strings"
)

type Give struct {
	Player []cmd.Target        `cmd:"player"`
	Item   Item                `cmd:"itemName"`
	Count  cmd.Optional[uint]  `cmd:"amount"`
	Data   cmd.Optional[int16] `cmd:"data"`
	// todo: NBT data
}

func (g Give) Run(src cmd.Source, o *cmd.Output) {
	itemCount := g.Count.LoadOr(1)
	if itemCount == 0 {
		o.Errorf("You must give at least one item.")
		return
	}
	itemType := string(g.Item)
	if !strings.ContainsRune(itemType, ':') {
		itemType = "minecraft:" + itemType
	}
	iType, ok := world.ItemByName(itemType, g.Data.LoadOr(0))
	if !ok {
		o.Errorf("Unknown item: %s.", g.Item)
		return
	}
	i := item.NewStack(iType, int(itemCount))

	count := 0
	for _, t := range g.Player {
		p, ok := t.(*player.Player)
		if !ok {
			continue
		}

		_, _ = p.Inventory().AddItem(i)
		count++
	}

	itemName := item.DisplayName(i.Item(), language.BritishEnglish)
	if count == 1 {
		o.Printf("You have given %s * %d to %s.", itemName, itemCount, g.Player[0].Name())
	} else {
		o.Printf("You have given %s * %d to %d players.", itemName, itemCount, count)
	}
}

type Item string

func (Item) Type() string {
	return "Item"
}

func (Item) Options(source cmd.Source) (items []string) {
	for _, i := range world.Items() {
		n, d := i.EncodeItem()
		if d != 0 {
			continue
		}
		items = append(items, strings.TrimPrefix(n, "minecraft:"))
	}
	for _, i := range world.CustomItems() {
		n, d := i.EncodeItem()
		if d != 0 {
			continue
		}
		items = append(items, n)
	}
	return items
}
