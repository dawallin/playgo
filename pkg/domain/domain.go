package domain

type domain struct {
	repository Repository
}

func NewDomain(repository Repository) *domain {
	d := domain{repository}
	return &d
}

type Repository interface {
	GetItems(count int) ([]string, error)
	GetItem(id string) (string, error)
	SetItem(s string) (string, error)
}

func (d *domain) GetItems(count int) ([]string, error) {
	return d.repository.GetItems(count)
}

func (d *domain) GetItem(id string) (string, error) {
	return d.repository.GetItem(id)
}

func (d *domain) SetItem(s string) (string, error) {
	return d.repository.SetItem(s)
}
