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

// QueryUsers realiza consultas aos usuários usando goroutines e channels
func (s *Service) QueryUsers(userIDs []int) []string {
	// Channel para coletar os resultados
	results := make(chan string)
	// WaitGroup para esperar todas as goroutines terminarem
	var wg sync.WaitGroup

	for _, id := range userIDs {
		wg.Add(1)
		go s.executeQuery(id, results, &wg)
	}

	// Goroutine para fechar o channel após todas as consultas serem concluídas
	go func() {
		wg.Wait()
		close(results)
	}()

	// Coletar os resultados do channel
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
