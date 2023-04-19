package main

import (
	"fmt"
)

type (
	Metric interface {
		Value() string
	}
	Prom interface {
		Send(metric Metric) string
	}

	metric  struct{}
	service struct{}
)

func (s *service) Send(m Metric) string {
	return m.Value()
}

func (m *metric) Value() string {
	return "some"
}

func NewService() *service {
	return &service{}
}

func NewMetric() *metric {
	return &metric{}
}

func main() {
	var s Prom
	m := NewMetric()
	s = NewService()
	fmt.Println(s.Send(m))
}
