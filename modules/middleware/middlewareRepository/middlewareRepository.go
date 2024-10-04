package middlewarerepository

type (
	MiddlewareRepositoryService interface{}

	middlewareRepository struct{
	}
)

func NewMiddlewareReposiroty() MiddlewareRepositoryService {
	return &middlewareRepository{}
}
