package storage

type Access interface {
	Create(dest interface{}) (interface{}, error)
	Fetch(dest interface{}) (interface{}, error)
	FetchParams(dest interface{}, params map[string]any) (interface{}, error)
	Update(dest interface{}) (interface{}, error)
	Delete(dest interface{}) error
}

type accessDriver struct {
	storage Access
}

func NewDataAccess(driver Access) Access {
	return &accessDriver{storage: driver}
}

func (d accessDriver) Create(dest interface{}) (interface{}, error) {
	return d.storage.Create(dest)
}

func (d accessDriver) Fetch(dest interface{}) (interface{}, error) {
	return d.storage.Fetch(dest)
}

func (d accessDriver) FetchParams(dest interface{}, params map[string]any) (interface{}, error) {
	return d.storage.FetchParams(dest, params)
}

func (d accessDriver) Update(dest interface{}) (interface{}, error) {
	return d.storage.Update(dest)
}

func (d accessDriver) Delete(dest interface{}) error {
	return d.storage.Delete(dest)
}
