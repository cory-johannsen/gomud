package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
)


func NewEffects (
  aimlessDrifter *AimlessDrifter,
  alfredPfennigworth *AlfredPfennigworth,
  arcaneMagick *ArcaneMagick,
  artfulDodger *ArtfulDodger,
  banterandJibe *BanterandJibe,
  battenDowntheHatches *BattenDowntheHatches,
  beastmaster *Beastmaster,
  beggarsBowl *BeggarsBowl,
  beguiler *Beguiler,
  bewitching *Bewitching,
  beyondtheVeil *BeyondtheVeil,
  blessinginDisguise *BlessinginDisguise,
  bloodSand *BloodSand,
  bookWorm *BookWorm,
  broadBellied *BroadBellied,
  castIronStomach *CastIronStomach,
  catlikeReflexes *CatlikeReflexes,
  cavesight *Cavesight,
  chatty *Chatty,
  childrenoftheEarth *ChildrenoftheEarth,
  clockworksofWar *ClockworksofWar,
  confidenceTrick *ConfidenceTrick,
  consumeAlcohol *ConsumeAlcohol,
  coordination *Coordination,
  cragFighting *CragFighting,
  craven *Craven,
  cruisinforaBruisin *CruisinforaBruisin,
  cultofPersonality *CultofPersonality,
  dangerSense *DangerSense,
  dauntless *Dauntless,
  deadlyAim *DeadlyAim,
  denizenofStone *DenizenofStone,
  dirtySecret *DirtySecret,
  divineMagick *DivineMagick,
  dogsofWar *DogsofWar,
  dontKilltheMessenger *DontKilltheMessenger,
  dungeonsDeep *DungeonsDeep,
  eldritchSigns *EldritchSigns,
  enduringMortality *EnduringMortality,
  escapeArtist *EscapeArtist,
  esotericMemory *EsotericMemory,
  eureka *Eureka,
  fameFortune *FameFortune,
  farsight *Farsight,
  feastorFamine *FeastorFamine,
  fetteredChaos *FetteredChaos,
  fieldwarden *Fieldwarden,
  fireandBrimstone *FireandBrimstone,
  firstborn *Firstborn,
  fleetFooted *FleetFooted,
  forAFewShillingsMore *ForAFewShillingsMore,
  fortunesWheel *FortunesWheel,
  fourinHand *FourinHand,
  frighteningBellow *FrighteningBellow,
  goldbergian *Goldbergian,
  greasetheWheels *GreasetheWheels,
  grimResolve *GrimResolve,
  grudgebearer *Grudgebearer,
  guerillaWarfare *GuerillaWarfare,
  guildPrivilege *GuildPrivilege,
  gutplate *Gutplate,
  gutterspeak *Gutterspeak,
  handleYourDrugs *HandleYourDrugs,
  hansShotFirst *HansShotFirst,
  hedgewizardry *Hedgewizardry,
  hereComesTheCavalry *HereComesTheCavalry,
  hijinks *Hijinks,
  hocusPocus *HocusPocus,
  holyRoller *HolyRoller,
  hueCry *HueCry,
  hungerPangs *HungerPangs,
  iAmTheLaw *IAmTheLaw,
  iSelltheDead *ISelltheDead,
  inCrowdTreachery *InCrowdTreachery,
  ironclad *Ironclad,
  issueChallenge *IssueChallenge,
  itsaTrap *ItsaTrap,
  jackOfAllTrades *JackOfAllTrades,
  kindredWarband *KindredWarband,
  kleptomania *Kleptomania,
  lamentoftheAges *LamentoftheAges,
  learnedDevotee *LearnedDevotee,
  lowBlow *LowBlow,
  luckoftheDraw *LuckoftheDraw,
  machiavellianIntrigue *MachiavellianIntrigue,
  magnificentBastard *MagnificentBastard,
  manifestDestiny *ManifestDestiny,
  mastermind *Mastermind,
  meditativeHealing *MeditativeHealing,
  memento *Memento,
  metropolitan *Metropolitan,
  mettleSteel *MettleSteel,
  mightyThews *MightyThews,
  mixedFamily *MixedFamily,
  moreWork *MoreWork,
  mortalCombat *MortalCombat,
  mortalityWork *MortalityWork,
  mountainAmongstMen *MountainAmongstMen,
  mountainWarfare *MountainWarfare,
  murdery *Murdery,
  naturalSelection *NaturalSelection,
  naturesOwn *NaturesOwn,
  nighteyes *Nighteyes,
  nimbleFingers *NimbleFingers,
  nobleSavage *NobleSavage,
  nostrumRemedium *NostrumRemedium,
  oathkeeper *Oathkeeper,
  oddCouple *OddCouple,
  paythePiper *PaythePiper,
  physicalProwess *PhysicalProwess,
  pintsized *Pintsized,
  pliersScalpelBonesaw *PliersScalpelBonesaw,
  rabbleRousing *RabbleRousing,
  recidivist *Recidivist,
  roadtoEnlightenment *RoadtoEnlightenment,
  rotgutSpray *RotgutSpray,
  runeMarkedGlory *RuneMarkedGlory,
  sacredMantra *SacredMantra,
  saltyDog *SaltyDog,
  scrivenersSpeed *ScrivenersSpeed,
  seventhSense *SeventhSense,
  shadowBroker *ShadowBroker,
  shakenNotStirred *ShakenNotStirred,
  shieldWall *ShieldWall,
  situationalAwareness *SituationalAwareness,
  slamdance *Slamdance,
  smart *Smart,
  sneakAttack *SneakAttack,
  standAndDeliver *StandAndDeliver,
  stentorianVoice *StentorianVoice,
  stoneheaded *Stoneheaded,
  strengthoftheMountain *StrengthoftheMountain,
  swaggerWagon *SwaggerWagon,
  theFullMonty *TheFullMonty,
  theShowMustGoOn *TheShowMustGoOn,
  theSlayersPath *TheSlayersPath,
  thickLining *ThickLining,
  thievingWays *ThievingWays,
  tortuousInquisition *TortuousInquisition,
  townGossip *TownGossip,
  trueDetective *TrueDetective,
  tunnelVision *TunnelVision,
  turntheOtherCheek *TurntheOtherCheek,
  underfoot *Underfoot,
  vimVigor *VimVigor,
  violent *Violent,
  wardenoftheWild *WardenoftheWild,
  warpSpasm *WarpSpasm,
  warriorsTattoo *WarriorsTattoo,
  wendigo *Wendigo,
  whiteWitchery *WhiteWitchery,
  wretchedPrankster *WretchedPrankster,
  ambidexterity *Ambidexterity,
  appallingMien *AppallingMien,
  arbalestsSpeed *ArbalestsSpeed,
  azimuth *Azimuth,
  badAxx *BadAxx,
  battleMagick *BattleMagick,
  beatdown *Beatdown,
  bloodMagick *BloodMagick,
  carousing *Carousing,
  cheapShot *CheapShot,
  clinchFighter *ClinchFighter,
  determination *Determination,
  dieHard *DieHard,
  doppelganger *Doppelganger,
  eagleEyes *EagleEyes,
  electricalAlignment *ElectricalAlignment,
  fencersPanache *FencersPanache,
  forkedTongue *ForkedTongue,
  fromTheHip *FromTheHip,
  gallowsHumor *GallowsHumor,
  gangsterGrip *GangsterGrip,
  gatecrasher *Gatecrasher,
  groundPound *GroundPound,
  gruesomeShot *GruesomeShot,
  handspring *Handspring,
  hardToKill *HardToKill,
  higherMysteries *HigherMysteries,
  holdout *Holdout,
  housebreaker *Housebreaker,
  impenetrableWall *ImpenetrableWall,
  imperviousMind *ImperviousMind,
  incredibleNumeration *IncredibleNumeration,
  indifference *Indifference,
  instincts *Instincts,
  kidneyShot *KidneyShot,
  knifework *Knifework,
  larceny *Larceny,
  lefthandedPath *LefthandedPath,
  lightSleeper *LightSleeper,
  lightningReaction *LightningReaction,
  longwinded *Longwinded,
  mariner *Mariner,
  meetingoftheMinds *MeetingoftheMinds,
  menacingDemeanor *MenacingDemeanor,
  militaryFormation *MilitaryFormation,
  mineCraft *MineCraft,
  mountedDefense *MountedDefense,
  multilingual *Multilingual,
  nervesofSteel *NervesofSteel,
  noMercy *NoMercy,
  overwhelmingForce *OverwhelmingForce,
  runAmok *RunAmok,
  ruralSensibility *RuralSensibility,
  saddleborn *Saddleborn,
  secondSkin *SecondSkin,
  secretSigns *SecretSigns,
  shieldSlam *ShieldSlam,
  shootFromTheHip *ShootFromTheHip,
  siegecraft *Siegecraft,
  silverTongue *SilverTongue,
  spiritedCharge *SpiritedCharge,
  sprint *Sprint,
  stranglersUnion *StranglersUnion,
  streetwise *Streetwise,
  strongJaw *StrongJaw,
  supernaturalParanoia *SupernaturalParanoia,
  swordBoard *SwordBoard,
  takeEmDown *TakeEmDown,
  thereWillBeBlood *ThereWillBeBlood,
  toughAsNails *ToughAsNails,
  trueGrit *TrueGrit,
  windsOfChange *WindsOfChange,
  worldly *Worldly,
  blackEye *BlackEye,
  bruisedRibs *BruisedRibs,
  dislocatedShoulder *DislocatedShoulder,
  fortunesMercyModerate *FortunesMercyModerate,
  hyperextendedElbow *HyperextendedElbow,
  jammedFinger *JammedFinger,
  misfortuneModerate *MisfortuneModerate,
  pulledMuscle *PulledMuscle,
  rattledBrain *RattledBrain,
  sprainedWrist *SprainedWrist,
  strainedGroin *StrainedGroin,
  twistedAnkle *TwistedAnkle,
  brokenRib *BrokenRib,
  bustedKneecap *BustedKneecap,
  fortunesMercySerious *FortunesMercySerious,
  fracturedLarynx *FracturedLarynx,
  headTraumaUntil *HeadTraumaUntil,
  minorConcussion *MinorConcussion,
  misfortuneSerious *MisfortuneSerious,
  shellShockUntil *ShellShockUntil,
  skullFractureUntil *SkullFractureUntil,
  stressFractureUntil *StressFractureUntil,
  temporaryParalysis *TemporaryParalysis,
  tornShoulder *TornShoulder,
  butcheredLeg *ButcheredLeg,
  cerebralContusion *CerebralContusion,
  detachedEar *DetachedEar,
  fortunesMercyGrievous *FortunesMercyGrievous,
  maimedFoot *MaimedFoot,
  mangledOrgan *MangledOrgan,
  mutilatedHand *MutilatedHand,
  mutilatedNose *MutilatedNose,
  puncturedLung *PuncturedLung,
  severedArtery *SeveredArtery,
  splinteredElbow *SplinteredElbow,
  vitreousHemorrhage *VitreousHemorrhage,
  badTicker *BadTicker,
  blackCataract *BlackCataract,
  bleeder *Bleeder,
  branded *Branded,
  cholericTemperament *CholericTemperament,
  cropEar *CropEar,
  cursed *Cursed,
  dealwiththeDevil *DealwiththeDevil,
  debtRidden *DebtRidden,
  dunderhead *Dunderhead,
  eunuch *Eunuch,
  lilyLivered *LilyLivered,
  melancholicTemperament *MelancholicTemperament,
  neerDoWell *NeerDoWell,
  nemesis *Nemesis,
  painkiller *Painkiller,
  persecutionComplex *PersecutionComplex,
  phlegmaticTemperament *PhlegmaticTemperament,
  sanguineTemperament *SanguineTemperament,
  sourStomach *SourStomach,
  splitFace *SplitFace,
  veteransBoot *VeteransBoot,
  veteransEye *VeteransEye,
  veteransHand *VeteransHand,
  veteransLeg *VeteransLeg,
  weakLungs *WeakLungs,
  adaptable *Adaptable,
  ammunition *Ammunition,
  defensive *Defensive,
  entangling *Entangling,
  fast *Fast,
  fiery *Fiery,
  finesse *Finesse,
  gunpowder *Gunpowder,
  immolate *Immolate,
  ineffective *Ineffective,
  light *Light,
  powerful *Powerful,
  pummeling *Pummeling,
  punishing *Punishing,
  reach *Reach,
  reload *Reload,
  repeating *Repeating,
  shrapnel *Shrapnel,
  slow *Slow,
  throwing *Throwing,
  vicious *Vicious,
  volatile *Volatile,
  weak *Weak,
) domain.Effects {
  return domain.Effects{
    aimlessDrifter,
    alfredPfennigworth,
    arcaneMagick,
    artfulDodger,
    banterandJibe,
    battenDowntheHatches,
    beastmaster,
    beggarsBowl,
    beguiler,
    bewitching,
    beyondtheVeil,
    blessinginDisguise,
    bloodSand,
    bookWorm,
    broadBellied,
    castIronStomach,
    catlikeReflexes,
    cavesight,
    chatty,
    childrenoftheEarth,
    clockworksofWar,
    confidenceTrick,
    consumeAlcohol,
    coordination,
    cragFighting,
    craven,
    cruisinforaBruisin,
    cultofPersonality,
    dangerSense,
    dauntless,
    deadlyAim,
    denizenofStone,
    dirtySecret,
    divineMagick,
    dogsofWar,
    dontKilltheMessenger,
    dungeonsDeep,
    eldritchSigns,
    enduringMortality,
    escapeArtist,
    esotericMemory,
    eureka,
    fameFortune,
    farsight,
    feastorFamine,
    fetteredChaos,
    fieldwarden,
    fireandBrimstone,
    firstborn,
    fleetFooted,
    forAFewShillingsMore,
    fortunesWheel,
    fourinHand,
    frighteningBellow,
    goldbergian,
    greasetheWheels,
    grimResolve,
    grudgebearer,
    guerillaWarfare,
    guildPrivilege,
    gutplate,
    gutterspeak,
    handleYourDrugs,
    hansShotFirst,
    hedgewizardry,
    hereComesTheCavalry,
    hijinks,
    hocusPocus,
    holyRoller,
    hueCry,
    hungerPangs,
    iAmTheLaw,
    iSelltheDead,
    inCrowdTreachery,
    ironclad,
    issueChallenge,
    itsaTrap,
    jackOfAllTrades,
    kindredWarband,
    kleptomania,
    lamentoftheAges,
    learnedDevotee,
    lowBlow,
    luckoftheDraw,
    machiavellianIntrigue,
    magnificentBastard,
    manifestDestiny,
    mastermind,
    meditativeHealing,
    memento,
    metropolitan,
    mettleSteel,
    mightyThews,
    mixedFamily,
    moreWork,
    mortalCombat,
    mortalityWork,
    mountainAmongstMen,
    mountainWarfare,
    murdery,
    naturalSelection,
    naturesOwn,
    nighteyes,
    nimbleFingers,
    nobleSavage,
    nostrumRemedium,
    oathkeeper,
    oddCouple,
    paythePiper,
    physicalProwess,
    pintsized,
    pliersScalpelBonesaw,
    rabbleRousing,
    recidivist,
    roadtoEnlightenment,
    rotgutSpray,
    runeMarkedGlory,
    sacredMantra,
    saltyDog,
    scrivenersSpeed,
    seventhSense,
    shadowBroker,
    shakenNotStirred,
    shieldWall,
    situationalAwareness,
    slamdance,
    smart,
    sneakAttack,
    standAndDeliver,
    stentorianVoice,
    stoneheaded,
    strengthoftheMountain,
    swaggerWagon,
    theFullMonty,
    theShowMustGoOn,
    theSlayersPath,
    thickLining,
    thievingWays,
    tortuousInquisition,
    townGossip,
    trueDetective,
    tunnelVision,
    turntheOtherCheek,
    underfoot,
    vimVigor,
    violent,
    wardenoftheWild,
    warpSpasm,
    warriorsTattoo,
    wendigo,
    whiteWitchery,
    wretchedPrankster,
    ambidexterity,
    appallingMien,
    arbalestsSpeed,
    azimuth,
    badAxx,
    battleMagick,
    beatdown,
    bloodMagick,
    carousing,
    cheapShot,
    clinchFighter,
    determination,
    dieHard,
    doppelganger,
    eagleEyes,
    electricalAlignment,
    fencersPanache,
    forkedTongue,
    fromTheHip,
    gallowsHumor,
    gangsterGrip,
    gatecrasher,
    groundPound,
    gruesomeShot,
    handspring,
    hardToKill,
    higherMysteries,
    holdout,
    housebreaker,
    impenetrableWall,
    imperviousMind,
    incredibleNumeration,
    indifference,
    instincts,
    kidneyShot,
    knifework,
    larceny,
    lefthandedPath,
    lightSleeper,
    lightningReaction,
    longwinded,
    mariner,
    meetingoftheMinds,
    menacingDemeanor,
    militaryFormation,
    mineCraft,
    mountedDefense,
    multilingual,
    nervesofSteel,
    noMercy,
    overwhelmingForce,
    runAmok,
    ruralSensibility,
    saddleborn,
    secondSkin,
    secretSigns,
    shieldSlam,
    shootFromTheHip,
    siegecraft,
    silverTongue,
    spiritedCharge,
    sprint,
    stranglersUnion,
    streetwise,
    strongJaw,
    supernaturalParanoia,
    swordBoard,
    takeEmDown,
    thereWillBeBlood,
    toughAsNails,
    trueGrit,
    windsOfChange,
    worldly,
    blackEye,
    bruisedRibs,
    dislocatedShoulder,
    fortunesMercyModerate,
    hyperextendedElbow,
    jammedFinger,
    misfortuneModerate,
    pulledMuscle,
    rattledBrain,
    sprainedWrist,
    strainedGroin,
    twistedAnkle,
    brokenRib,
    bustedKneecap,
    fortunesMercySerious,
    fracturedLarynx,
    headTraumaUntil,
    minorConcussion,
    misfortuneSerious,
    shellShockUntil,
    skullFractureUntil,
    stressFractureUntil,
    temporaryParalysis,
    tornShoulder,
    butcheredLeg,
    cerebralContusion,
    detachedEar,
    fortunesMercyGrievous,
    maimedFoot,
    mangledOrgan,
    mutilatedHand,
    mutilatedNose,
    puncturedLung,
    severedArtery,
    splinteredElbow,
    vitreousHemorrhage,
    badTicker,
    blackCataract,
    bleeder,
    branded,
    cholericTemperament,
    cropEar,
    cursed,
    dealwiththeDevil,
    debtRidden,
    dunderhead,
    eunuch,
    lilyLivered,
    melancholicTemperament,
    neerDoWell,
    nemesis,
    painkiller,
    persecutionComplex,
    phlegmaticTemperament,
    sanguineTemperament,
    sourStomach,
    splitFace,
    veteransBoot,
    veteransEye,
    veteransHand,
    veteransLeg,
    weakLungs,
    adaptable,
    ammunition,
    defensive,
    entangling,
    fast,
    fiery,
    finesse,
    gunpowder,
    immolate,
    ineffective,
    light,
    powerful,
    pummeling,
    punishing,
    reach,
    reload,
    repeating,
    shrapnel,
    slow,
    throwing,
    vicious,
    volatile,
    weak,
  }
}
