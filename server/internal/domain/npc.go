package domain

import (
	"context"
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type SkillRankSpec struct {
	Job   string `yaml:"job"`
	Skill string `yaml:"skill"`
}

type SkillRankSpecs []SkillRankSpec

type NPCInventorySpec struct {
	MainHand string   `yaml:"mainHand"`
	OffHand  string   `yaml:"offHand"`
	Armor    string   `yaml:"armor"`
	Pack     []string `yaml:"pack"`
	Cash     int      `yaml:"cash"`
}

type TaskDialog struct {
	Task string   `yaml:"task"`
	Text []string `yaml:"text"`
}

type Dialog map[string]TaskDialog

type NPCSpec struct {
	Name                string              `yaml:"name"`
	Age                 int                 `yaml:"age"`
	Alignment           AlignmentSpec       `yaml:"alignment"`
	Archetype           string              `yaml:"archetype"`
	Background          string              `yaml:"background"`
	BackgroundTrait     string              `yaml:"backgroundTrait"`
	BirthSeason         Season              `yaml:"birthSeason"`
	Condition           Condition           `yaml:"condition"`
	ConsumedAdvances    ConsumedAdvances    `yaml:"consumedAdvances"`
	DistinguishingMarks DistinguishingMarks `yaml:"distinguishingMarks"`
	Drawback            string              `yaml:"drawback"`
	FatePoints          int                 `yaml:"fatePoints"`
	Inventory           NPCInventorySpec    `yaml:"inventory"`
	Job                 string              `yaml:"job"`
	Peril               string              `yaml:"peril"`
	Poorness            Poorness            `yaml:"poorness"`
	Room                string              `yaml:"room"`
	Team                string              `yaml:"team"`
	Tattoo              Tattoo              `yaml:"tattoo"`
	SkillRanks          SkillRankSpecs      `yaml:"skillRanks"`
	Stats               Stats               `yaml:"stats"`
	Talents             []string            `yaml:"talents"`
	Upbringing          string              `yaml:"upbringing"`
	Dialog              Dialog              `yaml:"dialog"`
}

type NPCSpecs map[string]*NPCSpec

type NPCResolver interface {
	FetchNPCById(ctx context.Context, id int) (*NPC, error)
}

type NPC struct {
	mutex      sync.Mutex
	running    bool
	tickMillis int
	Character
	Domain         *htn.Domain
	Planner        *htn.Planner
	EventBus       eventbus.Bus
	Dialog         Dialog
	playersEngaged map[int64]*Player
	playersGreeted map[*Player]time.Time
}

func (n *NPC) IsPlayer() bool {
	return false
}

func (n *NPC) IsNPC() bool {
	return true
}

func (n *NPC) Start() error {
	if n.running {
		return nil
	}
	tickDuration := time.Duration(n.tickMillis) * time.Millisecond
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.running = true
	log.Printf("starting NPC %s", n.Name)
	go func() {
		for {
			if !n.running {
				log.Printf("NPC %s is not running", n.Name)
				break
			}
			// Plan the next action
			startTime := time.Now()
			plan, err := n.Planner.Plan(n.Domain)
			if err != nil {
				log.Errorf("error planning NPC action: %v", err)
			}
			log.Debugf("NPC %s plan: %v", n.Name, plan)
			// Execute the plan
			if plan != nil {
				newDomain, err := htn.Execute(plan, n.Domain, tickDuration)
				if err != nil {
					log.Errorf("error executing NPC plan: %v", err)
				}
				if newDomain != nil {
					n.mutex.Lock()
					n.Domain = newDomain
					n.mutex.Unlock()
				}
			}
			endTime := time.Now()
			elapsedTime := endTime.Sub(startTime)
			// Sleep until the next tick
			sleepTime := tickDuration - elapsedTime
			if sleepTime > 0 {
				time.Sleep(sleepTime)
			}
		}
	}()
	return nil
}

func (n *NPC) Stop() error {
	if !n.running {
		return nil
	}
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.running = false
	log.Printf("stopping NPC %s", n.Name)
	return nil
}

func (n *NPC) PlayersEngaged() int {
	return len(n.playersEngaged)
}

func (n *NPC) PlayerLastGreeted(player *Player) time.Time {
	last, ok := n.playersGreeted[player]
	if !ok {
		return time.Time{}
	}
	return last
}

func (n *NPC) SetPlayerLastGreeted(player *Player, t time.Time) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.playersGreeted[player] = t
}

func (n *NPC) Sobriety(substance string) float64 {
	return 1.0
}

func NewNPC(character *Character, domain *htn.Domain, planner *htn.Planner, dialog Dialog, eventBus eventbus.Bus, tickMillis int) *NPC {
	return &NPC{
		Character:      *character,
		Domain:         domain,
		Planner:        planner,
		EventBus:       eventBus,
		running:        false,
		tickMillis:     tickMillis,
		playersEngaged: make(map[int64]*Player),
		playersGreeted: make(map[*Player]time.Time),
		Dialog:         dialog,
	}
}

// PlayersEngagedSensor contains a NPC that can queried to calculate the number of engaged customers that NPC has
type PlayersEngagedSensor struct {
	NPC *NPC
}

func (s *PlayersEngagedSensor) Get() (int, error) {
	return s.NPC.PlayersEngaged(), nil
}

func (s *PlayersEngagedSensor) Name() string {
	return "PlayersEngaged"
}

func (s *PlayersEngagedSensor) String() string {
	value, _ := s.Get()
	return fmt.Sprintf("PlayersEngaged: %d", value)
}

var _ htn.Sensor[int] = &PlayersEngagedSensor{}

type PlayersInRangeSensor struct {
	NPC *NPC
}

func (s *PlayersInRangeSensor) Get() (int, error) {
	return s.NPC.Room().PlayerCount(), nil
}

func (s *PlayersInRangeSensor) Name() string {
	return "PlayersInRange"
}

func (s *PlayersInRangeSensor) String() string {
	value, _ := s.Get()
	return fmt.Sprintf("PlayersInRange: %d", value)
}

var _ htn.Sensor[int] = &PlayersInRangeSensor{}

type SobrietySensor struct {
	NPC       *NPC
	Substance string
}

func (s *SobrietySensor) String() string {
	val, err := s.Get()
	if err != nil {
		return fmt.Sprintf("Sobriety: %s: error", s.Substance)
	}
	return fmt.Sprintf("Sobriety: %s: %f", s.Substance, val)
}

func (s *SobrietySensor) Name() string {
	return fmt.Sprintf("Sobriety: %s", s.Substance)
}

func (s *SobrietySensor) Get() (float64, error) {
	return s.NPC.Sobriety(s.Substance), nil
}

var _ htn.Sensor[float64] = &SobrietySensor{}

type IntoxicationSensor struct {
	SobrietySensor
}

func (s *IntoxicationSensor) Get() (float64, error) {
	sobriety, err := s.SobrietySensor.Get()
	if err != nil {
		return 0, err
	}
	return 1.0 - sobriety, nil
}

var _ htn.Sensor[float64] = &IntoxicationSensor{}
