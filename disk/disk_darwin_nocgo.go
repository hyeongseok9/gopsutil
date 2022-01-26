//go:build darwin && !cgo
// +build darwin,!cgo

package disk

import (
	"context"

	"github.com/hyeongseok9/gopsutil/v3/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
