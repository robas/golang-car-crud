package repository

import (
	"carProject/entity"
	"strings"
)

type inmem struct {
	m map[entity.ID]*entity.Car
}

func NewInmem() *inmem {
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

func (repo *inmem) Search(query string) ([]*entity.Car, error) {
	var carList []*entity.Car
	for _, v := range repo.m {
		if strings.Contains(strings.ToLower(v.Brand), query) ||
			strings.Contains(strings.ToLower(v.Model), query) {
			carList = append(carList, v)
		}
	}

	return carList, nil
}

func (repo *inmem) List() ([]*entity.Car, error) {
	var carList []*entity.Car
	for _, v := range repo.m {
		carList = append(carList, v)
	}
	return carList, nil
}

func (repo *inmem) Update(c *entity.Car) error {
	_, err := repo.Get(c.ID)
	if err != nil {
		return err
	}
	repo.m[c.ID] = c
	return nil
}

func (repo *inmem) Delete(id entity.ID) error {
	if repo.m[id] == nil {
		return entity.ErrNotFound
	}
	repo.m[id] = nil
	return nil
}
