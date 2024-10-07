package domain

type Difficulty string

func (d Difficulty) Value() interface{} {
	return d
}

func (d Difficulty) String() string {
	return string(d)
}

const (
	TrivialDifficulty     Difficulty = "Trivial"
	EasyDifficulty        Difficulty = "Easy"
	RoutineDifficulty     Difficulty = "Routine"
	StandardDifficulty    Difficulty = "Standard"
	ChallengingDifficulty Difficulty = "Challenging"
	HardDifficulty        Difficulty = "Hard"
	ArduousDifficulty     Difficulty = "Arduous"
)

type Difficulties []Difficulty

var RankedDifficulties = []Difficulty{
	TrivialDifficulty,
	EasyDifficulty,
	RoutineDifficulty,
	StandardDifficulty,
	ChallengingDifficulty,
	HardDifficulty,
	ArduousDifficulty,
}

func (d Difficulty) Rank() int {
	for i, v := range RankedDifficulties {
		if v == d {
			return i
		}
	}
	return -1
}

func (d Difficulty) EasierThan(that Difficulty) bool {
	return d.Rank() < that.Rank()
}

var _ Property = Difficulty("")
