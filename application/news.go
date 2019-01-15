package application

import (
	"github.com/jojoarianto/go-ddd-api/config"
	"github.com/jojoarianto/go-ddd-api/domain"
	"github.com/jojoarianto/go-ddd-api/infrastructure/persistence"
)

// GetNews returns domain.news by id
func GetNews(id int) (*domain.News, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetAllNews return all domain.news
func GetAllNews() ([]domain.News, error) {
	conn, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	return repo.GetAll()
}

// AddNews saves new news
func AddNews(title string, slug string) error {
	conn, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewNewsRepositoryWithRDB(conn)
	u := &domain.News{
		Title: title,
		Slug:  slug,
	}
	return repo.Save(u)
}
