package car

import "carProject/entity"

type inmem struct {
	m map[entity.ID]*entity.Car
}

func newInmem() *inmem {
	var newMap = map[entity.ID]*entity.Car{}
	return &inmem{m: newMap}
}

func (repo *inmem) Create(c *entity.Car) (entity.ID, error) {
	repo.m[c.ID] = c
	return c.ID, nil
}

func (repo *inmem) Get(id entity.ID) (*entity.Car, error) {
	if repo.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return repo.m[id], nil
}
