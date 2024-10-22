//go:build wireinject
// +build wireinject

package loader

import (
	"github.com/google/wire"
)

var LoaderSet = wire.NewSet(
	NewAppearanceLoader, NewAlignmentLoader, NewArchetypeLoader, NewBackgroundLoader,
	NewEffectLoader, NewEquipmentLoader, NewInjuryLoader, NewInventoryLoader, NewJobLoader,
	NewQualityLoader, NewRoomLoader, NewSkillLoader,
	NewTalentLoader, NewTeamLoader, NewTraitLoader, NewUpbringingLoader)
