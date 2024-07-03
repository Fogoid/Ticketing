package station

import (
	"context"
	"log"

	db "github.com/fogoid/ticketing/db/mysql"
)

type StationRepository interface {
	GetStations() ([]*Station, error)
	GetStation(string) (*Station, error)
	CreateStation(*Station) (*Station, error)
	UpdateStation(*Station) (*Station, error)
	DeleteStation(*Station) (int, error)
}

type MySqlRepository struct {
	queries db.Queries
}

func (r *MySqlRepository) GetStations() ([]*Station, error) {
	stations, err := r.queries.ListStations(context.Background())
	if err != nil {
		log.Printf("Error listing stations: %v", err)
		return nil, err
	}

	all := make([]*Station, len(stations))
	for _, s := range stations {
		domainS := &Station{
			Name: s.Name,
		}
		all = append(all, domainS)
	}

	return all, nil
}

func (r *MySqlRepository) GetStation(id string) (*Station, error) {
	var idInt int64 = 1
	station, err := r.queries.GetStation(context.Background(), idInt)
	if err != nil {
		log.Printf("Error obtaining station: %v", station)
		return nil, err
	}

	s := NewStation(station.Name)
	return s, nil
}

func (r *MySqlRepository) CreateStation(station *Station) (*Station, error) {
	_, err := r.queries.CreateStation(context.Background(), station.Name)
	if err != nil {
		return nil, err
	}

	return station, nil
}

func (r *MySqlRepository) UpdateStation(station *Station) (*Station, error) {
	params := db.UpdateStationParams{
		ID:   1,
		Name: station.Name,
	}

	_, err := r.queries.UpdateStation(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return station, nil
}

func (r *MySqlRepository) DeleteStation(*Station) (int, error) {
	err := r.queries.DeteteStation(context.Background(), 1)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
