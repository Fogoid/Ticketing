package station

type StationService interface {
	GetStations() ([]*Station, error)
	GetStation(string) (*Station, error)
	CreateStation(*Station) (*Station, error)
	UpdateStation(*Station) (*Station, error)
	DeleteStation(*Station) (int, error)
}

type Service struct {
	stationRepository StationRepository
}

func NewStationService(stationRepository StationRepository) StationService {
	return &Service{
		stationRepository: stationRepository,
	}
}

func (s *Service) GetStations() ([]*Station, error) {
	return s.stationRepository.GetStations()
}

func (s *Service) GetStation(id string) (*Station, error) {
    return s.stationRepository.GetStation(id)
}

func (s *Service) CreateStation(station *Station) (*Station, error) {
	return s.stationRepository.CreateStation(station)
}

func (s *Service) UpdateStation(station *Station) (*Station, error) {
	return s.stationRepository.UpdateStation(station)
}

func (s *Service) DeleteStation(station *Station) (int, error) {
    return s.stationRepository.DeleteStation(station)
}
