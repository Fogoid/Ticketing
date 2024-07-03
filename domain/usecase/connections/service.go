package connections

import "github.com/fogoid/ticketing/domain/usecase/station"

type ConnectionService interface {
    GetStationsConnection(first *station.Station, second *station.Station) (*Connection, error)
    GetStationConnections(station *station.Station) ([]*Connection, error)
}
