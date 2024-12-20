package picker

import (
	"context"
	"errors"
	"gchat/pkg/logger"
	"os"

	"go.uber.org/zap"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

const AddrPickerName = "addr"

type addrKey struct{}

var ErrNoSubConnSelect = errors.New("no sub conn select")

func init() {
	balancer.Register(newBuilder())
}

func ContextWithAddr(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, addrKey{}, addr)
}

type addrPickerBuilder struct{}

func newBuilder() balancer.Builder {
	return base.NewBalancerBuilder(AddrPickerName, &addrPickerBuilder{}, base.Config{HealthCheck: true})
}

func (*addrPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	subConns := make(map[string]balancer.SubConn, len(info.ReadySCs))
	for k, sc := range info.ReadySCs {
		subConns[sc.Address.Addr] = k
	}
	return &addrPicker{
		subConnes: subConns,
	}
}

type addrPicker struct {
	subConnes map[string]balancer.SubConn
}

func (p *addrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	pr := balancer.PickResult{}

	address := info.Ctx.Value(addrKey{}).(string)
	if os.Getenv("GCHAT_ENV") == "gcp" {
		address = "34.44.91.248:8000"
	}
	sc, ok := p.subConnes[address]
	if !ok {
		logger.Logger.Error("Pick error", zap.String("address", address), zap.Any("subConnes", p.subConnes))
		return pr, ErrNoSubConnSelect
	}
	pr.SubConn = sc
	return pr, nil
}
