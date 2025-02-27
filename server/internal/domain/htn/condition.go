package htn

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Condition interface {
	Name() string
	IsMet(state *State) bool
	String() string
}

type Conditions map[string]Condition

// FlagCondition is a simple condition that is gated by a boolean Value that can be set
type FlagCondition struct {
	FlagName string `yaml:"name"`
	Value    bool   `yaml:"value"`
}

func (f *FlagCondition) Name() string {
	return f.FlagName
}

func (f *FlagCondition) IsMet(state *State) bool {
	return f.Value
}

func (f *FlagCondition) Set(value bool) {
	f.Value = value
}

func (f *FlagCondition) String() string {
	return fmt.Sprintf("FlagCondition: %t", f.Value)
}

// NotFlagCondition embeds a FlagCondition and inverts the behavior
type NotFlagCondition struct {
	FlagCondition
}

func (n *NotFlagCondition) IsMet(state *State) bool {
	return !n.FlagCondition.IsMet(state)
}

func (n *NotFlagCondition) String() string {
	return fmt.Sprintf("NotFlagCondition %s: %t", n.FlagName, n.FlagCondition.Value)
}

type Comparison string

const (
	EQ  Comparison = "=="
	NEQ Comparison = "!="
	LT  Comparison = "<"
	LTE Comparison = "<="
	GT  Comparison = ">"
	GTE Comparison = ">="
)

type Comparator[T any] func(value T, property T, comparison Comparison) bool

var IntComparator = func(value int, property int, comparison Comparison) bool {
	switch comparison {
	case EQ:
		return value == property
	case NEQ:
		return value != property
	case LT:
		return value < property
	case LTE:
		return value <= property
	case GT:
		return value > property
	case GTE:
		return value >= property
	}
	return false
}

var Int64Comparator = func(value int64, property int64, comparison Comparison) bool {
	switch comparison {
	case EQ:
		return value == property
	case NEQ:
		return value != property
	case LT:
		return value < property
	case LTE:
		return value <= property
	case GT:
		return value > property
	case GTE:
		return value >= property
	}
	return false
}

// ComparisonCondition is a condition that is met if the given Property compares to the specified Value using the given Comparison function
type ComparisonCondition[T any] struct {
	ConditionName string
	Comparison    Comparison
	Value         T
	Property      string
	Comparator    Comparator[T]
}

func (c *ComparisonCondition[T]) Name() string {
	return c.ConditionName
}

func (c *ComparisonCondition[T]) IsMet(state *State) bool {
	property, err := state.Property(c.Property)
	if err != nil {
		return false
	}
	value := property.(*Property[T]).Value(state)
	log.Debugf("ComparisonCondition %s comparing %s(%v) %s %v", c.ConditionName, c.Property, value, c.Comparison, c.Value)
	return c.Comparator(c.Value, value, c.Comparison)
}

func (c *ComparisonCondition[T]) String() string {
	return fmt.Sprintf("ComparisonCondition %s: property %s %s value %v", c.ConditionName, c.Property, c.Comparison, c.Value)
}

// PropertyComparisonCondition is a condition that compares to Property values
type PropertyComparisonCondition struct {
	ConditionName string     `yaml:"name"`
	Comparison    Comparison `yaml:"comparison"`
	LHS           string     `yaml:"lhs"`
	RHS           string     `yaml:"rhs"`
}

func (p *PropertyComparisonCondition) Name() string {
	return p.ConditionName
}

func (p *PropertyComparisonCondition) IsMet(state *State) bool {
	lhsProperty, err := state.Property(p.LHS)
	if err != nil {
		return false
	}
	lhs := lhsProperty.(Property[any]).Value(state).(float64)
	rhsProperty, err := state.Property(p.RHS)
	if err != nil {
		return false
	}
	rhs := rhsProperty.(Property[any]).Value(state).(float64)
	log.Printf("PropertyComparisonCondition %s comparing %s(%f) %s %s(%f)", p.ConditionName, p.LHS, lhs, p.Comparison, p.RHS, rhs)
	switch p.Comparison {
	case EQ:
		return lhs == rhs
	case NEQ:
		return lhs != rhs
	case GT:
		return lhs > rhs
	case GTE:
		return lhs >= rhs
	case LT:
		return lhs < rhs
	case LTE:
		return lhs <= rhs
	}
	return false
}

func (p *PropertyComparisonCondition) String() string {
	return fmt.Sprintf("PropertyComparisonCondition %s:  %s %s %s", p.ConditionName, p.LHS, p.Comparison, p.RHS)
}

type LogicalOperator string

const (
	AND LogicalOperator = "AND"
	OR  LogicalOperator = "OR"
	NOT LogicalOperator = "NOT"
	XOR LogicalOperator = "XOR"
)

type LogicalCondition struct {
	ConditionName string          `yaml:"name"`
	Operator      LogicalOperator `yaml:"operator"`
	LHSProperty   string          `yaml:"lhs"`
	RHSProperty   string          `yaml:"rhs"`
}

func (l *LogicalCondition) Name() string {
	return l.ConditionName
}

func (l *LogicalCondition) IsMet(state *State) bool {
	lhsProperty, err := state.Property(l.LHSProperty)
	if err != nil {
		return false
	}
	lhsFloat := lhsProperty.(Property[any]).Value(state).(float64)
	lhs := lhsFloat > 0
	rhsProperty, err := state.Property(l.RHSProperty)
	rhsFloat := float64(0)
	if err != nil {
		if l.Operator == NOT {
			rhsFloat = rhsProperty.(Property[any]).Value(state).(float64)
		} else {
			return false
		}
	}
	rhs := rhsFloat > 0
	log.Printf("LogicalCondition %s comparing %s(%f) %s %s(%f)", l.ConditionName, l.LHSProperty, lhsFloat, l.Operator, l.RHSProperty, rhsFloat)

	switch l.Operator {
	case AND:
		return lhs && rhs
	case OR:
		return lhs || rhs
	case NOT:
		return !lhs
	case XOR:
		return (lhs || rhs) && !(lhs && rhs)
	}
	return false
}

func (l *LogicalCondition) String() string {
	return fmt.Sprintf("LogicalCondition %s: %s %s %s", l.ConditionName, l.LHSProperty, l.Operator, l.RHSProperty)
}

// TaskCondition is a condition that is met when the given Task is complete
type TaskCondition struct {
	Task Task `yaml:"task"`
}

func (t *TaskCondition) IsMet(state *State) bool {
	return t.Task.IsComplete()
}

func (t *TaskCondition) String() string {
	return fmt.Sprintf("TaskCondition: %s, complete: %t", t.Task.Name(), t.Task.IsComplete())
}

type Evaluator func(state *State) bool

type FuncCondition struct {
	ConditionName string    `yaml:"name"`
	Evaluator     Evaluator `yaml:"evaluator"`
}

func (f *FuncCondition) Name() string {
	return f.ConditionName
}

func (f *FuncCondition) IsMet(state *State) bool {
	return f.Evaluator(state)
}

func (f *FuncCondition) String() string {
	return fmt.Sprintf("FuncCondition: %s", f.ConditionName)
}
