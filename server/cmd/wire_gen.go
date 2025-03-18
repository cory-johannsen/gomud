// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/effect"
	"github.com/cory-johannsen/gomud/internal/engine"
	"github.com/cory-johannsen/gomud/internal/event"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
)

// Injectors from wire.go:

func InitializeEngine() (*engine.Engine, error) {
	configConfig, err := config.NewConfigFromEnv()
	if err != nil {
		return nil, err
	}
	database, err := storage.NewDatabase(configConfig)
	if err != nil {
		return nil, err
	}
	aimlessDrifter := effect.NewAimlessDrifter()
	alfredPfennigworth := effect.NewAlfredPfennigworth()
	arcaneMagick := effect.NewArcaneMagick()
	artfulDodger := effect.NewArtfulDodger()
	banterandJibe := effect.NewBanterandJibe()
	battenDowntheHatches := effect.NewBattenDowntheHatches()
	beastmaster := effect.NewBeastmaster()
	beggarsBowl := effect.NewBeggarsBowl()
	beguiler := effect.NewBeguiler()
	bewitching := effect.NewBewitching()
	beyondtheVeil := effect.NewBeyondtheVeil()
	blessinginDisguise := effect.NewBlessinginDisguise()
	bloodSand := effect.NewBloodSand()
	bookWorm := effect.NewBookWorm()
	broadBellied := effect.NewBroadBellied()
	castIronStomach := effect.NewCastIronStomach()
	catlikeReflexes := effect.NewCatlikeReflexes()
	cavesight := effect.NewCavesight()
	chatty := effect.NewChatty()
	childrenoftheEarth := effect.NewChildrenoftheEarth()
	clockworksofWar := effect.NewClockworksofWar()
	confidenceTrick := effect.NewConfidenceTrick()
	consumeAlcohol := effect.NewConsumeAlcohol()
	coordination := effect.NewCoordination()
	cragFighting := effect.NewCragFighting()
	craven := effect.NewCraven()
	cruisinforaBruisin := effect.NewCruisinforaBruisin()
	cultofPersonality := effect.NewCultofPersonality()
	dangerSense := effect.NewDangerSense()
	dauntless := effect.NewDauntless()
	deadlyAim := effect.NewDeadlyAim()
	denizenofStone := effect.NewDenizenofStone()
	dirtySecret := effect.NewDirtySecret()
	divineMagick := effect.NewDivineMagick()
	dogsofWar := effect.NewDogsofWar()
	dontKilltheMessenger := effect.NewDontKilltheMessenger()
	dungeonsDeep := effect.NewDungeonsDeep()
	eldritchSigns := effect.NewEldritchSigns()
	enduringMortality := effect.NewEnduringMortality()
	escapeArtist := effect.NewEscapeArtist()
	esotericMemory := effect.NewEsotericMemory()
	eureka := effect.NewEureka()
	fameFortune := effect.NewFameFortune()
	farsight := effect.NewFarsight()
	feastorFamine := effect.NewFeastorFamine()
	fetteredChaos := effect.NewFetteredChaos()
	fieldwarden := effect.NewFieldwarden()
	fireandBrimstone := effect.NewFireandBrimstone()
	firstborn := effect.NewFirstborn()
	fleetFooted := effect.NewFleetFooted()
	forAFewShillingsMore := effect.NewForAFewShillingsMore()
	fortunesWheel := effect.NewFortunesWheel()
	fourinHand := effect.NewFourinHand()
	frighteningBellow := effect.NewFrighteningBellow()
	goldbergian := effect.NewGoldbergian()
	greasetheWheels := effect.NewGreasetheWheels()
	grimResolve := effect.NewGrimResolve()
	grudgebearer := effect.NewGrudgebearer()
	guerillaWarfare := effect.NewGuerillaWarfare()
	guildPrivilege := effect.NewGuildPrivilege()
	gutplate := effect.NewGutplate()
	gutterspeak := effect.NewGutterspeak()
	handleYourDrugs := effect.NewHandleYourDrugs()
	hansShotFirst := effect.NewHansShotFirst()
	hedgewizardry := effect.NewHedgewizardry()
	hereComesTheCavalry := effect.NewHereComesTheCavalry()
	hijinks := effect.NewHijinks()
	hocusPocus := effect.NewHocusPocus()
	holyRoller := effect.NewHolyRoller()
	hueCry := effect.NewHueCry()
	hungerPangs := effect.NewHungerPangs()
	iAmTheLaw := effect.NewIAmTheLaw()
	iSelltheDead := effect.NewISelltheDead()
	inCrowdTreachery := effect.NewInCrowdTreachery()
	ironclad := effect.NewIronclad()
	issueChallenge := effect.NewIssueChallenge()
	itsaTrap := effect.NewItsaTrap()
	jackOfAllTrades := effect.NewJackOfAllTrades()
	kindredWarband := effect.NewKindredWarband()
	kleptomania := effect.NewKleptomania()
	lamentoftheAges := effect.NewLamentoftheAges()
	learnedDevotee := effect.NewLearnedDevotee()
	lowBlow := effect.NewLowBlow()
	luckoftheDraw := effect.NewLuckoftheDraw()
	machiavellianIntrigue := effect.NewMachiavellianIntrigue()
	magnificentBastard := effect.NewMagnificentBastard()
	manifestDestiny := effect.NewManifestDestiny()
	mastermind := effect.NewMastermind()
	meditativeHealing := effect.NewMeditativeHealing()
	memento := effect.NewMemento()
	metropolitan := effect.NewMetropolitan()
	mettleSteel := effect.NewMettleSteel()
	mightyThews := effect.NewMightyThews()
	mixedFamily := effect.NewMixedFamily()
	moreWork := effect.NewMoreWork()
	mortalCombat := effect.NewMortalCombat()
	mortalityWork := effect.NewMortalityWork()
	mountainAmongstMen := effect.NewMountainAmongstMen()
	mountainWarfare := effect.NewMountainWarfare()
	murdery := effect.NewMurdery()
	naturalSelection := effect.NewNaturalSelection()
	naturesOwn := effect.NewNaturesOwn()
	nighteyes := effect.NewNighteyes()
	nimbleFingers := effect.NewNimbleFingers()
	nobleSavage := effect.NewNobleSavage()
	nostrumRemedium := effect.NewNostrumRemedium()
	oathkeeper := effect.NewOathkeeper()
	oddCouple := effect.NewOddCouple()
	paythePiper := effect.NewPaythePiper()
	physicalProwess := effect.NewPhysicalProwess()
	pintsized := effect.NewPintsized()
	pliersScalpelBonesaw := effect.NewPliersScalpelBonesaw()
	rabbleRousing := effect.NewRabbleRousing()
	recidivist := effect.NewRecidivist()
	roadtoEnlightenment := effect.NewRoadtoEnlightenment()
	rotgutSpray := effect.NewRotgutSpray()
	runeMarkedGlory := effect.NewRuneMarkedGlory()
	sacredMantra := effect.NewSacredMantra()
	saltyDog := effect.NewSaltyDog()
	scrivenersSpeed := effect.NewScrivenersSpeed()
	seventhSense := effect.NewSeventhSense()
	shadowBroker := effect.NewShadowBroker()
	shakenNotStirred := effect.NewShakenNotStirred()
	shieldWall := effect.NewShieldWall()
	situationalAwareness := effect.NewSituationalAwareness()
	slamdance := effect.NewSlamdance()
	smart := effect.NewSmart()
	sneakAttack := effect.NewSneakAttack()
	standAndDeliver := effect.NewStandAndDeliver()
	stentorianVoice := effect.NewStentorianVoice()
	stoneheaded := effect.NewStoneheaded()
	strengthoftheMountain := effect.NewStrengthoftheMountain()
	swaggerWagon := effect.NewSwaggerWagon()
	theFullMonty := effect.NewTheFullMonty()
	theShowMustGoOn := effect.NewTheShowMustGoOn()
	theSlayersPath := effect.NewTheSlayersPath()
	thickLining := effect.NewThickLining()
	thievingWays := effect.NewThievingWays()
	tortuousInquisition := effect.NewTortuousInquisition()
	townGossip := effect.NewTownGossip()
	trueDetective := effect.NewTrueDetective()
	tunnelVision := effect.NewTunnelVision()
	turntheOtherCheek := effect.NewTurntheOtherCheek()
	underfoot := effect.NewUnderfoot()
	vimVigor := effect.NewVimVigor()
	violent := effect.NewViolent()
	wardenoftheWild := effect.NewWardenoftheWild()
	warpSpasm := effect.NewWarpSpasm()
	warriorsTattoo := effect.NewWarriorsTattoo()
	wendigo := effect.NewWendigo()
	whiteWitchery := effect.NewWhiteWitchery()
	wretchedPrankster := effect.NewWretchedPrankster()
	ambidexterity := effect.NewAmbidexterity()
	appallingMien := effect.NewAppallingMien()
	arbalestsSpeed := effect.NewArbalestsSpeed()
	azimuth := effect.NewAzimuth()
	badAxx := effect.NewBadAxx()
	battleMagick := effect.NewBattleMagick()
	beatdown := effect.NewBeatdown()
	bloodMagick := effect.NewBloodMagick()
	carousing := effect.NewCarousing()
	cheapShot := effect.NewCheapShot()
	clinchFighter := effect.NewClinchFighter()
	determination := effect.NewDetermination()
	dieHard := effect.NewDieHard()
	doppelganger := effect.NewDoppelganger()
	eagleEyes := effect.NewEagleEyes()
	electricalAlignment := effect.NewElectricalAlignment()
	fencersPanache := effect.NewFencersPanache()
	forkedTongue := effect.NewForkedTongue()
	fromTheHip := effect.NewFromTheHip()
	gallowsHumor := effect.NewGallowsHumor()
	gangsterGrip := effect.NewGangsterGrip()
	gatecrasher := effect.NewGatecrasher()
	groundPound := effect.NewGroundPound()
	gruesomeShot := effect.NewGruesomeShot()
	handspring := effect.NewHandspring()
	hardToKill := effect.NewHardToKill()
	higherMysteries := effect.NewHigherMysteries()
	holdout := effect.NewHoldout()
	housebreaker := effect.NewHousebreaker()
	impenetrableWall := effect.NewImpenetrableWall()
	imperviousMind := effect.NewImperviousMind()
	incredibleNumeration := effect.NewIncredibleNumeration()
	indifference := effect.NewIndifference()
	instincts := effect.NewInstincts()
	kidneyShot := effect.NewKidneyShot()
	knifework := effect.NewKnifework()
	larceny := effect.NewLarceny()
	lefthandedPath := effect.NewLefthandedPath()
	lightSleeper := effect.NewLightSleeper()
	lightningReaction := effect.NewLightningReaction()
	longwinded := effect.NewLongwinded()
	mariner := effect.NewMariner()
	meetingoftheMinds := effect.NewMeetingoftheMinds()
	menacingDemeanor := effect.NewMenacingDemeanor()
	militaryFormation := effect.NewMilitaryFormation()
	mineCraft := effect.NewMineCraft()
	mountedDefense := effect.NewMountedDefense()
	multilingual := effect.NewMultilingual()
	nervesofSteel := effect.NewNervesofSteel()
	noMercy := effect.NewNoMercy()
	overwhelmingForce := effect.NewOverwhelmingForce()
	runAmok := effect.NewRunAmok()
	ruralSensibility := effect.NewRuralSensibility()
	saddleborn := effect.NewSaddleborn()
	secondSkin := effect.NewSecondSkin()
	secretSigns := effect.NewSecretSigns()
	shieldSlam := effect.NewShieldSlam()
	shootFromTheHip := effect.NewShootFromTheHip()
	siegecraft := effect.NewSiegecraft()
	silverTongue := effect.NewSilverTongue()
	spiritedCharge := effect.NewSpiritedCharge()
	sprint := effect.NewSprint()
	stranglersUnion := effect.NewStranglersUnion()
	streetwise := effect.NewStreetwise()
	strongJaw := effect.NewStrongJaw()
	supernaturalParanoia := effect.NewSupernaturalParanoia()
	swordBoard := effect.NewSwordBoard()
	takeEmDown := effect.NewTakeEmDown()
	thereWillBeBlood := effect.NewThereWillBeBlood()
	toughAsNails := effect.NewToughAsNails()
	trueGrit := effect.NewTrueGrit()
	windsOfChange := effect.NewWindsOfChange()
	worldly := effect.NewWorldly()
	blackEye := effect.NewBlackEye()
	bruisedRibs := effect.NewBruisedRibs()
	dislocatedShoulder := effect.NewDislocatedShoulder()
	fortunesMercyModerate := effect.NewFortunesMercyModerate()
	hyperextendedElbow := effect.NewHyperextendedElbow()
	jammedFinger := effect.NewJammedFinger()
	misfortuneModerate := effect.NewMisfortuneModerate()
	pulledMuscle := effect.NewPulledMuscle()
	rattledBrain := effect.NewRattledBrain()
	sprainedWrist := effect.NewSprainedWrist()
	strainedGroin := effect.NewStrainedGroin()
	twistedAnkle := effect.NewTwistedAnkle()
	brokenRib := effect.NewBrokenRib()
	bustedKneecap := effect.NewBustedKneecap()
	fortunesMercySerious := effect.NewFortunesMercySerious()
	fracturedLarynx := effect.NewFracturedLarynx()
	headTraumaUntil := effect.NewHeadTraumaUntil()
	minorConcussion := effect.NewMinorConcussion()
	misfortuneSerious := effect.NewMisfortuneSerious()
	shellShockUntil := effect.NewShellShockUntil()
	skullFractureUntil := effect.NewSkullFractureUntil()
	stressFractureUntil := effect.NewStressFractureUntil()
	temporaryParalysis := effect.NewTemporaryParalysis()
	tornShoulder := effect.NewTornShoulder()
	butcheredLeg := effect.NewButcheredLeg()
	cerebralContusion := effect.NewCerebralContusion()
	detachedEar := effect.NewDetachedEar()
	fortunesMercyGrievous := effect.NewFortunesMercyGrievous()
	maimedFoot := effect.NewMaimedFoot()
	mangledOrgan := effect.NewMangledOrgan()
	mutilatedHand := effect.NewMutilatedHand()
	mutilatedNose := effect.NewMutilatedNose()
	puncturedLung := effect.NewPuncturedLung()
	severedArtery := effect.NewSeveredArtery()
	splinteredElbow := effect.NewSplinteredElbow()
	vitreousHemorrhage := effect.NewVitreousHemorrhage()
	badTicker := effect.NewBadTicker()
	blackCataract := effect.NewBlackCataract()
	bleeder := effect.NewBleeder()
	branded := effect.NewBranded()
	cholericTemperament := effect.NewCholericTemperament()
	cropEar := effect.NewCropEar()
	cursed := effect.NewCursed()
	dealwiththeDevil := effect.NewDealwiththeDevil()
	debtRidden := effect.NewDebtRidden()
	dunderhead := effect.NewDunderhead()
	eunuch := effect.NewEunuch()
	lilyLivered := effect.NewLilyLivered()
	melancholicTemperament := effect.NewMelancholicTemperament()
	neerDoWell := effect.NewNeerDoWell()
	nemesis := effect.NewNemesis()
	painkiller := effect.NewPainkiller()
	persecutionComplex := effect.NewPersecutionComplex()
	phlegmaticTemperament := effect.NewPhlegmaticTemperament()
	sanguineTemperament := effect.NewSanguineTemperament()
	sourStomach := effect.NewSourStomach()
	splitFace := effect.NewSplitFace()
	veteransBoot := effect.NewVeteransBoot()
	veteransEye := effect.NewVeteransEye()
	veteransHand := effect.NewVeteransHand()
	veteransLeg := effect.NewVeteransLeg()
	weakLungs := effect.NewWeakLungs()
	adaptable := effect.NewAdaptable()
	ammunition := effect.NewAmmunition()
	dangerous := effect.NewDangerous()
	defensive := effect.NewDefensive()
	entangling := effect.NewEntangling()
	fast := effect.NewFast()
	fiery := effect.NewFiery()
	finesse := effect.NewFinesse()
	gunpowder := effect.NewGunpowder()
	heavy := effect.NewHeavy()
	immolate := effect.NewImmolate()
	ineffective := effect.NewIneffective()
	light := effect.NewLight()
	natural := effect.NewNatural()
	powerful := effect.NewPowerful()
	protective := effect.NewProtective()
	pummeling := effect.NewPummeling()
	punishing := effect.NewPunishing()
	reach := effect.NewReach()
	reload := effect.NewReload()
	repeating := effect.NewRepeating()
	shrapnel := effect.NewShrapnel()
	slow := effect.NewSlow()
	throwing := effect.NewThrowing()
	vicious := effect.NewVicious()
	volatile := effect.NewVolatile()
	weak := effect.NewWeak()
	effects := effect.NewEffects(aimlessDrifter, alfredPfennigworth, arcaneMagick, artfulDodger, banterandJibe, battenDowntheHatches, beastmaster, beggarsBowl, beguiler, bewitching, beyondtheVeil, blessinginDisguise, bloodSand, bookWorm, broadBellied, castIronStomach, catlikeReflexes, cavesight, chatty, childrenoftheEarth, clockworksofWar, confidenceTrick, consumeAlcohol, coordination, cragFighting, craven, cruisinforaBruisin, cultofPersonality, dangerSense, dauntless, deadlyAim, denizenofStone, dirtySecret, divineMagick, dogsofWar, dontKilltheMessenger, dungeonsDeep, eldritchSigns, enduringMortality, escapeArtist, esotericMemory, eureka, fameFortune, farsight, feastorFamine, fetteredChaos, fieldwarden, fireandBrimstone, firstborn, fleetFooted, forAFewShillingsMore, fortunesWheel, fourinHand, frighteningBellow, goldbergian, greasetheWheels, grimResolve, grudgebearer, guerillaWarfare, guildPrivilege, gutplate, gutterspeak, handleYourDrugs, hansShotFirst, hedgewizardry, hereComesTheCavalry, hijinks, hocusPocus, holyRoller, hueCry, hungerPangs, iAmTheLaw, iSelltheDead, inCrowdTreachery, ironclad, issueChallenge, itsaTrap, jackOfAllTrades, kindredWarband, kleptomania, lamentoftheAges, learnedDevotee, lowBlow, luckoftheDraw, machiavellianIntrigue, magnificentBastard, manifestDestiny, mastermind, meditativeHealing, memento, metropolitan, mettleSteel, mightyThews, mixedFamily, moreWork, mortalCombat, mortalityWork, mountainAmongstMen, mountainWarfare, murdery, naturalSelection, naturesOwn, nighteyes, nimbleFingers, nobleSavage, nostrumRemedium, oathkeeper, oddCouple, paythePiper, physicalProwess, pintsized, pliersScalpelBonesaw, rabbleRousing, recidivist, roadtoEnlightenment, rotgutSpray, runeMarkedGlory, sacredMantra, saltyDog, scrivenersSpeed, seventhSense, shadowBroker, shakenNotStirred, shieldWall, situationalAwareness, slamdance, smart, sneakAttack, standAndDeliver, stentorianVoice, stoneheaded, strengthoftheMountain, swaggerWagon, theFullMonty, theShowMustGoOn, theSlayersPath, thickLining, thievingWays, tortuousInquisition, townGossip, trueDetective, tunnelVision, turntheOtherCheek, underfoot, vimVigor, violent, wardenoftheWild, warpSpasm, warriorsTattoo, wendigo, whiteWitchery, wretchedPrankster, ambidexterity, appallingMien, arbalestsSpeed, azimuth, badAxx, battleMagick, beatdown, bloodMagick, carousing, cheapShot, clinchFighter, determination, dieHard, doppelganger, eagleEyes, electricalAlignment, fencersPanache, forkedTongue, fromTheHip, gallowsHumor, gangsterGrip, gatecrasher, groundPound, gruesomeShot, handspring, hardToKill, higherMysteries, holdout, housebreaker, impenetrableWall, imperviousMind, incredibleNumeration, indifference, instincts, kidneyShot, knifework, larceny, lefthandedPath, lightSleeper, lightningReaction, longwinded, mariner, meetingoftheMinds, menacingDemeanor, militaryFormation, mineCraft, mountedDefense, multilingual, nervesofSteel, noMercy, overwhelmingForce, runAmok, ruralSensibility, saddleborn, secondSkin, secretSigns, shieldSlam, shootFromTheHip, siegecraft, silverTongue, spiritedCharge, sprint, stranglersUnion, streetwise, strongJaw, supernaturalParanoia, swordBoard, takeEmDown, thereWillBeBlood, toughAsNails, trueGrit, windsOfChange, worldly, blackEye, bruisedRibs, dislocatedShoulder, fortunesMercyModerate, hyperextendedElbow, jammedFinger, misfortuneModerate, pulledMuscle, rattledBrain, sprainedWrist, strainedGroin, twistedAnkle, brokenRib, bustedKneecap, fortunesMercySerious, fracturedLarynx, headTraumaUntil, minorConcussion, misfortuneSerious, shellShockUntil, skullFractureUntil, stressFractureUntil, temporaryParalysis, tornShoulder, butcheredLeg, cerebralContusion, detachedEar, fortunesMercyGrievous, maimedFoot, mangledOrgan, mutilatedHand, mutilatedNose, puncturedLung, severedArtery, splinteredElbow, vitreousHemorrhage, badTicker, blackCataract, bleeder, branded, cholericTemperament, cropEar, cursed, dealwiththeDevil, debtRidden, dunderhead, eunuch, lilyLivered, melancholicTemperament, neerDoWell, nemesis, painkiller, persecutionComplex, phlegmaticTemperament, sanguineTemperament, sourStomach, splitFace, veteransBoot, veteransEye, veteransHand, veteransLeg, weakLungs, adaptable, ammunition, dangerous, defensive, entangling, fast, fiery, finesse, gunpowder, heavy, immolate, ineffective, light, natural, powerful, protective, pummeling, punishing, reach, reload, repeating, shrapnel, slow, throwing, vicious, volatile, weak)
	effectLoader := loader.NewEffectLoader(configConfig, effects)
	appearanceLoader := loader.NewAppearanceLoader(configConfig, effectLoader)
	alignmentLoader := loader.NewAlignmentLoader(configConfig)
	traitLoader := loader.NewTraitLoader(configConfig, effectLoader)
	skillLoader := loader.NewSkillLoader(configConfig)
	qualityLoader := loader.NewQualityLoader(configConfig, effectLoader)
	equipmentLoader := loader.NewEquipmentLoader(configConfig, skillLoader, qualityLoader)
	archetypeLoader := loader.NewArchetypeLoader(configConfig, traitLoader, equipmentLoader)
	backgroundLoader := loader.NewBackgroundLoader(configConfig, traitLoader)
	disorderLoader := loader.NewDisorderLoader(configConfig)
	bus := EventBus.New()
	roomLoader := loader.NewRoomLoader(configConfig, bus)
	npcLoader := loader.NewNPCLoader(configConfig, alignmentLoader)
	generatorLoader := loader.NewGeneratorLoader(configConfig, roomLoader, npcLoader)
	injuryLoader := loader.NewInjuryLoader(configConfig, effectLoader)
	inventoryLoader := loader.NewInventoryLoader(configConfig)
	talentLoader := loader.NewTalentLoader(configConfig, effectLoader)
	jobLoader := loader.NewJobLoader(configConfig, archetypeLoader, skillLoader, talentLoader, traitLoader)
	teamLoader := loader.NewTeamLoader(configConfig)
	upbringingLoader := loader.NewUpbringingLoader(configConfig)
	actionLoader := loader.NewActionLoader(configConfig)
	sensorLoader := loader.NewSensorLoader()
	conditionLoader := loader.NewConditionLoader(configConfig)
	methodLoader := loader.NewMethodLoader(configConfig, conditionLoader)
	taskLoader := loader.NewTaskLoader(configConfig, actionLoader, conditionLoader, methodLoader)
	taskGraphLoader := loader.NewTaskGraphLoader(configConfig, taskLoader)
	loaders := loader.NewLoaders(appearanceLoader, alignmentLoader, archetypeLoader, backgroundLoader, disorderLoader, effectLoader, equipmentLoader, generatorLoader, injuryLoader, inventoryLoader, jobLoader, npcLoader, qualityLoader, roomLoader, skillLoader, talentLoader, traitLoader, teamLoader, upbringingLoader, actionLoader, sensorLoader, conditionLoader, methodLoader, taskLoader, taskGraphLoader)
	equipment := storage.NewEquipment(database, loaders)
	plannerGenerator := generator.NewPlannerGenerator()
	domainGenerator := generator.NewDomainGenerator()
	npCs := storage.NewNPCs(configConfig, database, loaders, equipment, plannerGenerator, domainGenerator, bus)
	players := storage.NewPlayers(database, npCs, loaders, equipment)
	playerGenerator := generator.NewPlayerGenerator(loaders)
	clock := event.NewClock(bus, configConfig)
	server := engine.NewServer(configConfig, database, players, npCs, loaders, playerGenerator, domainGenerator, plannerGenerator, bus, clock)
	engineEngine := engine.NewEngine(configConfig, server, bus)
	return engineEngine, nil
}
