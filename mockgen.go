package ApiGateWay

//go:generate mockery --dir ./internal/useCase/interfaces --name Store --output ./internal/useCase/mock/ --outpkg mockUseCase
//go:generate mockery --dir ./internal/useCase/adapters/storage --name Store --output ./internal/repository/store/storage/mock --outpkg mockStoreStorage
//go:generate mockery --dir ./internal/useCase/adapters/grpc --name Store --output ./internal/repository/store/grpc/mock --outpkg mockStoreGrpc
