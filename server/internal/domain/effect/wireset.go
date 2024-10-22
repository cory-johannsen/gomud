// +build wireinject

package effect

import (
	"github.com/google/wire"
)

var EffectsSet = wire.NewSet(
  NewAimlessDrifter,
  NewAlfredPfennigworth,
  NewArcaneMagick,
  NewArtfulDodger,
  NewBanterandJibe,
  NewBattenDowntheHatches,
  NewBeastmaster,
  NewBeggarsBowl,
  NewBeguiler,
  NewBewitching,
  NewBeyondtheVeil,
  NewBlessinginDisguise,
  NewBloodSand,
  NewBookWorm,
  NewBroadBellied,
  NewCastIronStomach,
  NewCatlikeReflexes,
  NewCavesight,
  NewChatty,
  NewChildrenoftheEarth,
  NewClockworksofWar,
  NewConfidenceTrick,
  NewConsumeAlcohol,
  NewCoordination,
  NewCragFighting,
  NewCraven,
  NewCruisinforaBruisin,
  NewCultofPersonality,
  NewDangerSense,
  NewDauntless,
  NewDeadlyAim,
  NewDenizenofStone,
  NewDirtySecret,
  NewDivineMagick,
  NewDogsofWar,
  NewDontKilltheMessenger,
  NewDungeonsDeep,
  NewEldritchSigns,
  NewEnduringMortality,
  NewEscapeArtist,
  NewEsotericMemory,
  NewEureka,
  NewFameFortune,
  NewFarsight,
  NewFeastorFamine,
  NewFetteredChaos,
  NewFieldwarden,
  NewFireandBrimstone,
  NewFirstborn,
  NewForAFewShillingsMore,
  NewFortunesWheel,
  NewFourinHand,
  NewFrighteningBellow,
  NewGoldbergian,
  NewGreasetheWheels,
  NewGrimResolve,
  NewGrudgebearer,
  NewGuerillaWarfare,
  NewGuildPrivilege,
  NewGutplate,
  NewGutterspeak,
  NewHandleYourDrugs,
  NewHansShotFirst,
  NewHedgewizardry,
  NewHereComesTheCavalry,
  NewHijinks,
  NewHocusPocus,
  NewHolyRoller,
  NewHueCry,
  NewHungerPangs,
  NewIAmTheLaw,
  NewISelltheDead,
  NewInCrowdTreachery,
  NewIronclad,
  NewIssueChallenge,
  NewItsaTrap,
  NewJackOfAllTrades,
  NewKleptomania,
  NewLearnedDevotee,
  NewLowBlow,
  NewLuckoftheDraw,
  NewMachiavellianIntrigue,
  NewMagnificentBastard,
  NewManifestDestiny,
  NewMastermind,
  NewMeditativeHealing,
  NewMemento,
  NewMetropolitan,
  NewMettleSteel,
  NewMightyThews,
  NewMixedFamily,
  NewMoreWork,
  NewMortalCombat,
  NewMortalityWork,
  NewMountainAmongstMen,
  NewMountainWarfare,
  NewMurdery,
  NewNaturalSelection,
  NewNaturesOwn,
  NewNighteyes,
  NewNimbleFingers,
  NewNobleSavage,
  NewNostrumRemedium,
  NewOathkeeper,
  NewOddCouple,
  NewPaythePiper,
  NewPhysicalProwess,
  NewPintsized,
  NewPliersScalpelBonesaw,
  NewRabbleRousing,
  NewRecidivist,
  NewRoadtoEnlightenment,
  NewRotgutSpray,
  NewRuneMarkedGlory,
  NewSacredMantra,
  NewSaltyDog,
  NewScrivenersSpeed,
  NewSeventhSense,
  NewShadowBroker,
  NewShakenNotStirred,
  NewShieldWall,
  NewSituationalAwareness,
  NewSlamdance,
  NewSmart,
  NewSneakAttack,
  NewStandAndDeliver,
  NewStentorianVoice,
  NewStoneheaded,
  NewStrengthoftheMountain,
  NewSwaggerWagon,
  NewTheFullMonty,
  NewTheShowMustGoOn,
  NewTheSlayersPath,
  NewThickLining,
  NewThievingWays,
  NewTortuousInquisition,
  NewTownGossip,
  NewTrueDetective,
  NewTunnelVision,
  NewTurntheOtherCheek,
  NewUnderfoot,
  NewVimVigor,
  NewViolent,
  NewWardenoftheWild,
  NewWarpSpasm,
  NewWarriorsTattoo,
  NewWendigo,
  NewWhiteWitchery,
  NewWretchedPrankster,
  NewAmbidexterity,
  NewAppallingMien,
  NewArbalestsSpeed,
  NewAzimuth,
  NewBadAxx,
  NewBattleMagick,
  NewBeatdown,
  NewBloodMagick,
  NewCarousing,
  NewCheapShot,
  NewClinchFighter,
  NewDetermination,
  NewDieHard,
  NewDoppelganger,
  NewEagleEyes,
  NewElectricalAlignment,
  NewFencersPanache,
  NewForkedTongue,
  NewFromTheHip,
  NewGallowsHumor,
  NewGangsterGrip,
  NewGatecrasher,
  NewGroundPound,
  NewGruesomeShot,
  NewHandspring,
  NewHardToKill,
  NewHigherMysteries,
  NewHoldout,
  NewHousebreaker,
  NewImpenetrableWall,
  NewImperviousMind,
  NewIncredibleNumeration,
  NewIndifference,
  NewInstincts,
  NewKidneyShot,
  NewKnifework,
  NewLarceny,
  NewLefthandedPath,
  NewLightSleeper,
  NewLightningReaction,
  NewLongwinded,
  NewMariner,
  NewMeetingoftheMinds,
  NewMenacingDemeanor,
  NewMilitaryFormation,
  NewMineCraft,
  NewMountedDefense,
  NewMultilingual,
  NewNervesofSteel,
  NewNoMercy,
  NewOverwhelmingForce,
  NewRunAmok,
  NewRuralSensibility,
  NewSaddleborn,
  NewSecondSkin,
  NewSecretSigns,
  NewShieldSlam,
  NewShootFromTheHip,
  NewSiegecraft,
  NewSilverTongue,
  NewSpiritedCharge,
  NewSprint,
  NewStranglersUnion,
  NewStreetwise,
  NewStrongJaw,
  NewSupernaturalParanoia,
  NewSwordBoard,
  NewTakeEmDown,
  NewThereWillBeBlood,
  NewToughAsNails,
  NewTrueGrit,
  NewWindsOfChange,
  NewWorldly,
  NewBlackEye,
  NewBruisedRibs,
  NewDislocatedShoulder,
  NewFortunesMercyModerate,
  NewHyperextendedElbow,
  NewJammedFinger,
  NewMisfortuneModerate,
  NewPulledMuscle,
  NewRattledBrain,
  NewSprainedWrist,
  NewStrainedGroin,
  NewTwistedAnkle,
  NewBrokenRib,
  NewBustedKneecap,
  NewFortunesMercySerious,
  NewFracturedLarynx,
  NewHeadTraumaUntil,
  NewMinorConcussion,
  NewMisfortuneSerious,
  NewShellShockUntil,
  NewSkullFractureUntil,
  NewStressFractureUntil,
  NewTemporaryParalysis,
  NewTornShoulder,
  NewButcheredLeg,
  NewCerebralContusion,
  NewDetachedEar,
  NewFortunesMercyGrievous,
  NewMaimedFoot,
  NewMangledOrgan,
  NewMutilatedHand,
  NewMutilatedNose,
  NewPuncturedLung,
  NewSeveredArtery,
  NewSplinteredElbow,
  NewVitreousHemorrhage,
  NewBadTicker,
  NewBlackCataract,
  NewBleeder,
  NewBranded,
  NewCholericTemperament,
  NewCropEar,
  NewCursed,
  NewDealwiththeDevil,
  NewDebtRidden,
  NewDunderhead,
  NewEunuch,
  NewLilyLivered,
  NewMelancholicTemperament,
  NewNeerDoWell,
  NewNemesis,
  NewPainkiller,
  NewPersecutionComplex,
  NewPhlegmaticTemperament,
  NewSanguineTemperament,
  NewSourStomach,
  NewSplitFace,
  NewVeteransBoot,
  NewVeteransEye,
  NewVeteransHand,
  NewVeteransLeg,
  NewWeakLungs,
  NewAdaptable,
  NewAmmunition,
  NewDefensive,
  NewEntangling,
  NewFast,
  NewFiery,
  NewFinesse,
  NewGunpowder,
  NewImmolate,
  NewIneffective,
  NewLight,
  NewPowerful,
  NewPummeling,
  NewPunishing,
  NewReach,
  NewReload,
  NewRepeating,
  NewShrapnel,
  NewSlow,
  NewThrowing,
  NewVicious,
  NewVolatile,
  NewWeak,
)
