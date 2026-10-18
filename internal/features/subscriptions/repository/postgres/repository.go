package subscription_repository_postgres

import core_postgres_pool "github.com/Andromeeeda/Subtracker/internal/core/repository/postgres/pool"

type SubPostgresRep struct {
	pool core_postgres_pool.Pool
}

func NewSubPostgresRep(pool core_postgres_pool.Pool) *SubPostgresRep {
	return &SubPostgresRep{
		pool: pool,
	}
}


