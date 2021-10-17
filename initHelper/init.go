package initHelper

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
	"{{project_name}}/initHelper/depend"
	"{{project_name}}/config"
)

type Helper struct {
	*log.Logger
	Depends []depend.Depend
}

func (i *Helper) Init(ctx context.Context) error {
	startTime := time.Now()
	rand.Seed(time.Now().UnixNano())
	if err := config.Config.ReadConfig(); err != nil {
		panic(err)
	}
	for j := range i.Depends {
		startSubTime := time.Now()
		if i.Depends[j].GetEnable() {
			err := i.Depends[j].Init(ctx)
			if err != nil {
				return err
			}
			fmt.Printf("[init] \u001B[1;32;40m%s\u001B[0m at %d ms\n", i.Depends[j].GetName(), time.Now().UnixMilli()-startSubTime.UnixMilli())
		}
	}
	fmt.Println(fmt.Sprintf("[init] Finish at %d ms", time.Now().UnixMilli()-startTime.UnixMilli()))
	return nil
}

func (i *Helper) AddDepend(depend ...depend.Depend) *Helper {
	for j := range depend {
		depend[j].SetEnable(true)
		i.Depends = append(i.Depends, depend[j])
	}
	return i
}
