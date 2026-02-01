package main

type ArticleService struct {
	repo *ArticleRepository
	log  chan string
}

func NewArticleService(r *ArticleRepository, log chan string) *ArticleService {
	return &ArticleService{
		repo: r,
		log:  log,
	}
}

func (s *ArticleService) GetAll() []Article {
	return s.repo.GetAll()
}

func (s *ArticleService) GetByID(id int) (*Article, bool) {
	return s.repo.GetByID(id)
}

func (s *ArticleService) Create(a Article) Article {

	result := s.repo.Create(a)

	s.log <- "Created: " + result.Title

	return result
}

func (s *ArticleService) Update(id int, a Article) (*Article, bool) {

	result, ok := s.repo.Update(id, a)

	if ok {
		s.log <- "Updated: " + a.Title
	}

	return result, ok
}

func (s *ArticleService) Delete(id int) bool {

	ok := s.repo.Delete(id)

	if ok {
		s.log <- "Deleted article"
	}

	return ok
}
