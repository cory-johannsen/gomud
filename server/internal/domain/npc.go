package domain

import (
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

type NPC struct {
	mutex      sync.Mutex
	running    bool
	tickMillis int
	Character
	State          *htn.State
	Planner        *htn.Planner
	EventBus       eventbus.Bus
	playersEngaged map[int64]*Player
	Dialog         Dialog
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
			plan, err := n.Planner.Plan(n.State)
			if err != nil {
				log.Errorf("error planning NPC action: %v", err)
			}
			log.Debugf("NPC %s plan: %v", n.Name, plan)
			// Execute the plan
			if plan != nil {
				newState, err := htn.Execute(plan, n.State)
				if err != nil {
					log.Errorf("error executing NPC plan: %v", err)
				}
				if newState != nil {
					n.mutex.Lock()
					n.State = newState
					n.mutex.Unlock()
				}
			}
			// Sleep until the next tick
			time.Sleep(time.Duration(n.tickMillis) * time.Millisecond)
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

func NewNPC(character *Character, state *htn.State, planner *htn.Planner, dialog Dialog, eventBus eventbus.Bus, tickMillis int) *NPC {
	return &NPC{
		Character:      *character,
		State:          state,
		Planner:        planner,
		EventBus:       eventBus,
		running:        false,
		tickMillis:     tickMillis,
		playersEngaged: make(map[int64]*Player),
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
