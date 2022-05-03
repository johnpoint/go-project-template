package bootstrap

import (
	"context"
	"go.uber.org/zap"
	"math/rand"
	"reflect"
	"time"
)

var globalComponent = make([]Component, 0)

func AddGlobalComponent(components ...Component) {
	globalComponent = append(globalComponent, components...)
}

type Helper struct {
	logger     log
	components []Component
}

func (i *Helper) Init(ctx context.Context) error {
	i.logger.Info("[Bootstrap] Start")
	rand.Seed(time.Now().UnixNano())
	for j := range globalComponent {
		i.logger.Info("[Bootstrap]", zap.String("Component", reflect.TypeOf(globalComponent[j]).String()))
		err := globalComponent[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	for j := range i.components {
		i.logger.Info("[Bootstrap]", zap.String("Component", reflect.TypeOf(i.components[j]).String()))
		err := i.components[j].Init(ctx)
		if err != nil {
			return err
		}
	}
	i.logger.Info("[Bootstrap] Finish")
	return nil
}

func (i *Helper) AddComponent(components ...Component) *Helper {
	for j := range components {
		i.components = append(i.components, components[j])
	}
	return i
}

func (i *Helper) SetLogger(logger log) *Helper {
	i.logger = logger
	return i
}
