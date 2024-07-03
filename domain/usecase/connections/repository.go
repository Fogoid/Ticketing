package connections

type ConnectionsRepository interface {
    GetStationConnections(string) []*Connection
}
