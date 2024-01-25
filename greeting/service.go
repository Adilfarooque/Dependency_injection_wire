package greeting

import "fmt"

type GreetingService struct{}

func NewGreetingService() *GreetingService {
	return &GreetingService{}
}

func (s *GreetingService) Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
