package function1

import (
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/version"
	"go.uber.org/zap"
)

func DoFunction1(name string) error {
	logger.Info("hello "+name, zap.String("version", version.Get()))
	test := healthpb.StatusResponse{}
	logger.Infof("pb: %v", test)

	return nil
}
