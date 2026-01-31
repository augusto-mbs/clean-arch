package graph

import "github.com/augusto-mbs/clean-arch/internal/usecase"

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}
