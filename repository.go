package main

import "sync"

type ArticleRepository struct {
	data   []Article
	nextID int
	mutex  sync.Mutex
}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{
		data:   []Article{},
		nextID: 1,
	}
}

func (r *ArticleRepository) GetAll() []Article {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.data
}

func (r *ArticleRepository) GetByID(id int) (*Article, bool) {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, a := range r.data {
		if a.ID == id {
			return &a, true
		}
	}

	return nil, false
}

func (r *ArticleRepository) Create(a Article) Article {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	a.ID = r.nextID
	r.nextID++

	r.data = append(r.data, a)

	return a
}

func (r *ArticleRepository) Update(id int, a Article) (*Article, bool) {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, item := range r.data {
		if item.ID == id {

			a.ID = id
			r.data[i] = a

			return &a, true
		}
	}

	return nil, false
}

func (r *ArticleRepository) Delete(id int) bool {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, a := range r.data {
		if a.ID == id {

			r.data = append(r.data[:i], r.data[i+1:]...)
			return true
		}
	}

	return false
}
