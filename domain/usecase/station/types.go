package station

type Station struct {
	Name string
}

func NewStation(name string) *Station {
    return &Station {
        name,
    }
}
