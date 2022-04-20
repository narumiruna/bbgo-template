package template

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/c9s/bbgo/pkg/bbgo"
	"github.com/c9s/bbgo/pkg/types"
)

const ID = "template"

var log = logrus.WithField("strategy", ID)

func init() {
	bbgo.RegisterStrategy(ID, &Strategy{})
}

type Strategy struct {
	Notifiability *bbgo.Notifiability

	Symbol   string         `json:"symbol"`
	Interval types.Interval `json:"interval"`
}

func (s *Strategy) ID() string {
	return ID
}

func (s *Strategy) Initialize() error {
	return nil
}

func (s *Strategy) Validate() error {
	return nil
}

func (s *Strategy) Subscribe(session *bbgo.ExchangeSession) {
	session.Subscribe(types.KLineChannel, s.Symbol, types.SubscribeOptions{Interval: s.Interval.String()})
}

func (s *Strategy) Run(ctx context.Context, orderExecutor bbgo.OrderExecutor, session *bbgo.ExchangeSession) error {
	session.MarketDataStream.OnKLineClosed(func(kline types.KLine) {
		log.Infof("kline: %v", kline)
	})
	return nil
}
