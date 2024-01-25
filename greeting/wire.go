package greeting

import "github.com/google/wire"

func InitializeGreetingService() *GreetingService {
	wire.Build(NewGreetingService)
	return &GreetingService{}
}

