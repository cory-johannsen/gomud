package loader

import "github.com/cory-johannsen/gomud/internal/domain/htn"

type SensorLoader struct {
	sensors htn.Sensors
}

func NewSensorLoader() *SensorLoader {
	return &SensorLoader{
		sensors: make(htn.Sensors),
	}
}

func (l *SensorLoader) LoadSensors() (htn.Sensors, error) {
	return l.sensors, nil
}

func (l *SensorLoader) GetSensor(name string) (any, error) {
	return l.sensors[name], nil
}

func (l *SensorLoader) SetSensor(name string, sensor any) {
	l.sensors[name] = sensor
}
