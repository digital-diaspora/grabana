package encoder

import (
	"github.com/K-Phoen/sdk"
	"github.com/lueurxax/grabana/encoder/golang"
	"go.uber.org/zap"
)

func ToGolang(logger *zap.Logger, dashboard sdk.Board) (string, error) {
	golangEncoder := golang.NewEncoder(logger)

	return golangEncoder.EncodeDashboard(dashboard)
}
