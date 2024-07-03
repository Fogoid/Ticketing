package connections

import (
	"errors"
	"time"

	"github.com/fogoid/ticketing/domain/usecase/station"
)

type Connection struct {
    First *station.Station
    Second *station.Station
    Time time.Duration
}

func (c Connection) GetConnectedStation(name string) (*station.Station, error) {
    switch name {
    case c.First.Name:
        return c.Second, nil
    case c.Second.Name:
        return c.First, nil
    default:
        return nil, errors.New("Could not find connected station")
    }
}

