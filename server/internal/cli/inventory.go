package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/fatih/color"
)

type InventoryHandler struct {
	stateProvider domain.StateProvider
}

func NewInventoryHandler(stateProvider domain.StateProvider) *InventoryHandler {
	return &InventoryHandler{
		stateProvider: stateProvider,
	}
}

func (i *InventoryHandler) Handle(ctx context.Context, args []string) (string, error) {
	cyan := color.New(color.FgCyan).SprintFunc()
	state := i.stateProvider()
	player := state.Player()
	inventory := player.Inventory()
	msg := fmt.Sprintf("  %s\n\tMain Hand: ", cyan("Inventory"))
	if inventory.MainHand() == nil {
		msg += "empty"
	} else {
		msg += inventory.MainHand().Name()
	}
	msg += fmt.Sprintf("\n\tOff Hand: ")
	if inventory.OffHand() == nil {
		msg += "empty"
	} else {
		msg += inventory.OffHand().Name()
	}
	msg += fmt.Sprintf("\n\tArmor: ")
	if inventory.Armor() == nil {
		msg += "empty"
	} else {
		msg += inventory.Armor().Name()
	}
	msg += fmt.Sprintf("\n\tCash: %d\n", inventory.Cash())
	msg += "\tPack:\n"
	if len(inventory.Pack().Items()) == 0 {
		msg += "\t\tempty\n"
	} else {
		for _, item := range inventory.Pack().Items() {
			itemMsg := fmt.Sprintf("\t\t%s", item.Name())
			if item.Description() != "" {
				itemMsg += fmt.Sprintf(" - %s", item.Description())
			}
			msg += fmt.Sprintf("%s, %.2f kg\n", itemMsg, float64(item.MassInGrams())/1000.0)
		}
	}
	return msg, nil
}

func (i *InventoryHandler) Help(args []string) string {
	return "view your inventory"
}

func (i *InventoryHandler) State() domain.GameState {
	return i.stateProvider()
}

var _ Handler = &InventoryHandler{}
