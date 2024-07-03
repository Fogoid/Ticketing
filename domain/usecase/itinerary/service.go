package itinerary

import (
	"errors"
	"time"

	"github.com/fogoid/ticketing/domain/usecase/connections"
	"github.com/fogoid/ticketing/domain/usecase/station"
	"github.com/fogoid/ticketing/pkg/queue"
)

type ItineraryService interface {
	CreateItinerary(from string, to string)
}

type Service struct {
	stationService station.StationService
	connService    connections.ConnectionService
}

func (s *Service) CreateItinerary(from string, to string) (*Itinerary, error) {
	fromStation, err := s.stationService.GetStation(from)
	if err != nil {
		return nil, err
	}

	toStation, err := s.stationService.GetStation(to)
	if err != nil {
		return nil, err
	}

	graph, err := s.getStationsGraph(fromStation)
	if err != nil {
		return nil, err
	}

	exploredNodes := make(map[string]bool)
	currentNode := graph.initialStation
	queue := queue.New[*stationNode]()
	queue.Push(currentNode)
	for !queue.IsEmpty() {
		currentNode := queue.Pop().(*stationNode)
		if currentNode.station.Name == toStation.Name {
			return s.calculateItinerary(currentNode)
        }

		for _, childNode := range currentNode.connectedStations {
			if !exploredNodes[childNode.station.Name] {
				childNode.parent = currentNode
				queue.Push(childNode)
			}
		}
	}

	return nil, errors.New("Itinerary not found")
}

func (s *Service) getStationsGraph(initialStation *station.Station) (*StationsGraph, error) {
	expandedStations := make(map[string]bool)
	stationsToExpand := queue.New[*stationNode]()

	initialNode := NewStationNode(initialStation)
	stationsToExpand.Push(initialNode)
	for !stationsToExpand.IsEmpty() {
		currentStation := stationsToExpand.Pop().(*station.Station)

		if expandedStations[currentStation.Name] {
			continue
		}
		expandedStations[currentStation.Name] = true

		conn, err := s.connService.GetStationConnections(currentStation)
		if err != nil {
			return nil, err
		}

		for _, c := range conn {
			other, err := c.GetConnectedStation(currentStation.Name)
			if err != nil {
				return nil, err
			}

			otherConnections := NewStationNode(other)
			stationsToExpand.Push(otherConnections)
		}
	}

	return NewStationsGraph(initialNode), nil
}

func (s *Service) calculateItinerary(finalNode *stationNode) (*Itinerary, error) {
    stations := make([]*station.Station, 0)
    totalDuration := time.Second * 0

    currentNode := finalNode
    parentNode := finalNode.parent

    for parentNode != nil {
        conn, err := s.connService.GetStationsConnection(currentNode.station, parentNode.station)
        if err != nil {
            return nil, err
        }

        totalDuration += conn.Time
    }

    return NewItinerary(stations, totalDuration), nil
}

