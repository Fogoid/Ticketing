package itinerary

import (
	"time"

	"github.com/fogoid/ticketing/domain/usecase/station"
)

type Itinerary struct {
	stations []*station.Station
    duration time.Duration
}

func NewItinerary(stations []*station.Station, duration time.Duration) *Itinerary {
    return &Itinerary {
        stations,
        duration,
    }
}

type connection struct {
	from *station.Station
	to   *station.Station
	time time.Duration
}

type StationsGraph struct {
	initialStation *stationNode
}

func NewStationsGraph(initialStation *stationNode) *StationsGraph {
	return &StationsGraph{
		initialStation,
	}
}

type stationNode struct {
	station           *station.Station
	connectedStations []*stationNode
	parent            *stationNode
}

func NewStationNode(station *station.Station) *stationNode {
	return &stationNode{
		station,
		make([]*stationNode, 0),
		nil,
	}
}
