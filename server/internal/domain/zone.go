package domain

type Area struct{}

type Areas map[string]*Area

type ZoneSpec struct {
	ID          int64
	Name        string
	Description string
	Areas       []string
	Rooms       []string
}

type Zone struct {
	ID          int64
	Name        string
	Description string
	Areas       Areas
	Rooms       Rooms
}
