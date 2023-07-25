package environment

const (
	EnvLocal      Environment = "local"
	EnvProduction Environment = "production"
)

type Environment string

func (env Environment) IsLocal() bool {
	return env == EnvLocal
}

func (env Environment) IsProduction() bool {
	return env == EnvProduction
}
