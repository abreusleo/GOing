package user

import (
	"GOing/internal/db"
	"GOing/pkg/config"
	"fmt"
	"sync"
)

type Service struct {
	repo *Repository
}

func NewService(cfg config.Config) *Service {
	database := db.Connect()
	repo := NewRepository(database)
	return &Service{
		repo: repo,
	}
}

func (s *Service) QueryUsers(userIDs []int) []string {
	results := make(chan string)
	var wg sync.WaitGroup

	for _, id := range userIDs {
		wg.Add(1)
		go s.executeQuery(id, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var output []string
	for result := range results {
		output = append(output, result)
	}

	return output
}

func (s *Service) executeQuery(userID int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	name, err := s.repo.GetUserByID(userID)
	if err != nil {
		results <- err.Error()
		return
	}
	results <- fmt.Sprintf(name)
}
