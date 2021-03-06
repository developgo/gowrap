package templatestests

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using ../templates/prometheus template

//go:generate gowrap gen -p github.com/hexdigest/gowrap/templates_tests -i TestInterface -t ../templates/prometheus -o interface_with_prometheus.go

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// TestInterfaceWithPrometheus implements TestInterface interface with all methods wrapped
// with Prometheus metrics
type TestInterfaceWithPrometheus struct {
	base         TestInterface
	instanceName string
}

var testinterfaceDurationSummaryVec = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:   "testinterface_duration_seconds",
		Help:   "testinterface runtime duration and result",
		MaxAge: time.Minute,
	},
	[]string{"instance_name", "method", "result"})

// NewTestInterfaceWithPrometheus returns an instance of the TestInterface decorated with prometheus metrics
func NewTestInterfaceWithPrometheus(base TestInterface, instanceName string) TestInterfaceWithPrometheus {
	return TestInterfaceWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// F implements TestInterface
func (_d TestInterfaceWithPrometheus) F(ctx context.Context, a1 string, a2 ...string) (result1 string, result2 string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		testinterfaceDurationSummaryVec.WithLabelValues(_d.instanceName, "F", result).Observe(time.Since(_since).Seconds())
	}()

	return _d.base.F(ctx, a1, a2...)
}

// NoError implements TestInterface
func (_d TestInterfaceWithPrometheus) NoError(s1 string) (s2 string) {
	_since := time.Now()
	defer testinterfaceDurationSummaryVec.WithLabelValues(_d.instanceName, "NoError", "ok").Observe(time.Since(_since).Seconds())

	return _d.base.NoError(s1)
}

// NoParamsOrResults implements TestInterface
func (_d TestInterfaceWithPrometheus) NoParamsOrResults() {
	_since := time.Now()
	defer testinterfaceDurationSummaryVec.WithLabelValues(_d.instanceName, "NoParamsOrResults", "ok").Observe(time.Since(_since).Seconds())

	_d.base.NoParamsOrResults()
	return
}
