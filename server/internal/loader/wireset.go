//go:build wireinject
// +build wireinject

package loader

import (
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	"github.com/google/wire"
)

var LoaderSet = wire.NewSet(
	NewAppearanceLoader, NewAlignmentLoader, NewArchetypeLoader, NewBackgroundLoader, NewDisorderLoader,
	NewEffectLoader, NewEquipmentLoader, NewInjuryLoader, NewInteractiveObjectLoader, NewInventoryLoader, NewJobLoader,
	NewQualityLoader, NewRoomLoader, NewSkillLoader,
	NewTalentLoader, NewTeamLoader, NewTraitLoader, NewUpbringingLoader, NewGeneratorLoader, NewNPCLoader,
	NewActionLoader, wire.Bind(new(ActionResolver), new(*ActionLoader)),
	NewConditionLoader, NewMethodLoader, NewSensorLoader, NewTaskLoader, NewTaskGraphLoader)
