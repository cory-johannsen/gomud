package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
)

type ConditionLoader struct {
	config     *config.Config
	conditions htn.Conditions
}

func NewConditionLoader(cfg *config.Config) *ConditionLoader {
	return &ConditionLoader{
		config:     cfg,
		conditions: make(htn.Conditions),
	}
}

func (l *ConditionLoader) LoadConditions() (htn.Conditions, error) {
	if len(l.conditions) > 0 {
		return l.conditions, nil
	}

	flags, err := l.loadFlags()
	if err != nil {
		return nil, err
	}
	for k, v := range flags {
		l.conditions[k] = v
	}

	logical, err := l.loadLogical()
	if err != nil {
		return nil, err
	}
	for k, v := range logical {
		l.conditions[k] = v
	}

	notFlags, err := l.loadNotFlags()
	if err != nil {
		return nil, err
	}
	for k, v := range notFlags {
		l.conditions[k] = v
	}

	comparison, err := l.loadComparison()
	if err != nil {
		return nil, err
	}
	for k, v := range comparison {
		l.conditions[k] = v
	}

	property, err := l.loadPropertyComparison()
	if err != nil {
		return nil, err
	}
	for k, v := range property {
		l.conditions[k] = v
	}

	return l.conditions, nil
}

func (l *ConditionLoader) GetCondition(name string) (htn.Condition, error) {
	if len(l.conditions) == 0 {
		_, err := l.LoadConditions()
		if err != nil {
			return nil, err
		}
	}
	condition, ok := l.conditions[name]
	if !ok {
		return nil, nil
	}
	return condition, nil
}

func (l *ConditionLoader) SetCondition(name string, condition htn.Condition) {
	l.conditions[name] = condition
}

func (l *ConditionLoader) loadFlags() (htn.Conditions, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/conditions/flag")
	if err != nil {
		return nil, err
	}
	conditions := make(htn.Conditions)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		condition := &htn.FlagCondition{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/conditions/flag/" + item.Name())
		if err != nil {
			log.Errorf("error reading condition file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, condition)
		if err != nil {
			log.Errorf("error unmarshalling condition file %s: %v", item.Name(), err)
			continue
		}
		conditions[condition.Name()] = condition
	}
	return conditions, nil
}

func (l *ConditionLoader) loadNotFlags() (htn.Conditions, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/conditions/notflag")
	if err != nil {
		return nil, err
	}
	conditions := make(htn.Conditions)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		condition := &htn.FlagCondition{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/conditions/notflag/" + item.Name())
		if err != nil {
			log.Errorf("error reading condition file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, condition)
		if err != nil {
			log.Errorf("error unmarshalling condition file %s: %v", item.Name(), err)
			continue
		}

		conditions[condition.Name()] = &htn.NotFlagCondition{
			FlagCondition: *condition,
		}
	}
	return conditions, nil
}

func (l *ConditionLoader) loadLogical() (htn.Conditions, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/conditions/logical")
	if err != nil {
		return nil, err
	}
	conditions := make(htn.Conditions)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		condition := &htn.LogicalCondition{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/conditions/logical/" + item.Name())
		if err != nil {
			log.Errorf("error reading condition file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, condition)
		if err != nil {
			log.Errorf("error unmarshalling condition file %s: %v", item.Name(), err)
			continue
		}
		conditions[condition.Name()] = condition
	}
	return conditions, nil
}

func (l *ConditionLoader) loadComparison() (htn.Conditions, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/conditions/comparison")
	if err != nil {
		return nil, err
	}
	conditions := make(htn.Conditions)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		spec := &htn.ComparisonConditionSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/conditions/comparison/" + item.Name())
		if err != nil {
			log.Errorf("error reading condition file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling condition file %s: %v", item.Name(), err)
			continue
		}

		var condition htn.Condition
		switch spec.Type {
		case htn.Int64ConditionValueType:
			value, err := strconv.Atoi(spec.Value)
			if err != nil {
				log.Errorf("error converting comparison value to int64: %v", err)
				continue
			}
			newCondition := &htn.ComparisonCondition[int64]{
				ConditionName: spec.ConditionName,
				Comparison:    spec.Comparison,
				Value:         int64(value),
				Property:      spec.Property,
				Comparator:    htn.Int64Comparator,
			}
			condition = newCondition
		case htn.Float64ConditionValueType:
			value, err := strconv.ParseFloat(spec.Value, 64)
			if err != nil {
				log.Errorf("error converting comparison value to float64: %v", err)
				continue
			}
			newCondition := &htn.ComparisonCondition[float64]{
				ConditionName: spec.ConditionName,
				Comparison:    spec.Comparison,
				Value:         value,
				Property:      spec.Property,
				Comparator:    htn.Float64Comparator,
			}
			condition = newCondition
		}

		conditions[condition.Name()] = condition
	}
	return conditions, nil
}

func (l *ConditionLoader) loadPropertyComparison() (htn.Conditions, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/conditions/propertycomparison")
	if err != nil {
		return nil, err
	}
	conditions := make(htn.Conditions)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		condition := &htn.PropertyComparisonCondition{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/conditions/propertycomparison/" + item.Name())
		if err != nil {
			log.Errorf("error reading condition file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, condition)
		if err != nil {
			log.Errorf("error unmarshalling condition file %s: %v", item.Name(), err)
			continue
		}
		conditions[condition.Name()] = condition
	}
	return conditions, nil
}
