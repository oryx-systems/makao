package datastore

import (
	"github.com/oryx-systems/makao/pkg/makao/application/common/helpers"
	"github.com/oryx-systems/makao/pkg/makao/infrastructure/datastore/psql"
	log "github.com/sirupsen/logrus"
)

// Repository holds a set of all the repository methods that interact with the database
type Repository interface {
	Create
	Query
	Update
}

// DbServiceImpl is an implementation of the database repository
// It is implementation agnostic i.e logic should be handled using
// the preferred database
type DbServiceImpl struct {
	Repository
}

// NewDbService creates a new database service
func NewDbService() *DbServiceImpl {
	// This implementation is database agnostic. It can be changed to use any database. e.g. Pg, Firebase, MongoDB, etc
	environment := helpers.MustGetEnvVar("REPOSITORY")

	switch environment {
	case "firebase":
		return nil

	case "postgres":
		pg, err := psql.NewPGInstance()
		if err != nil {
			log.Panicf("can't initialize postgres when setting up profile service: %s", err)
		}

		return &DbServiceImpl{
			Repository: pg,
		}

	default:
		log.Panicf("unknown repository: %s", environment)
	}

	return nil
}
