package entity

type ClientRawRepositoryInterface interface {
	Create(client []*ClientRaw) error
	GetClients(limit int, status string) ([]*ClientRaw, error)
	UpdateStatusClient(client []*Client) error
}

type ClientRepositoryInterface interface {
	Create(client []*Client) error
}
