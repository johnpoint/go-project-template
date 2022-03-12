package depend

import (
	"PROJECT_NAME/config"
	"PROJECT_NAME/dao/mongoDao"
	"PROJECT_NAME/pkg/bootstrap"
	"context"
)

type MongoDB struct{}

var _ bootstrap.Component = (*MongoDB)(nil)

func (r *MongoDB) Init(ctx context.Context) error {
	mongoDao.InitMongoClient(config.Config.MongoDBConfig)
	return nil
}
