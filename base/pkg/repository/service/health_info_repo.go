package repo

type HealthInfoRepositoryService interface {
	BuildInfo() string
	HealthInfoDatabase() bool
	HealthInfoRedis() bool
}
