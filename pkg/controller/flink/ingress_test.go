package flink

import (
	"testing"

	config2 "github.com/lyft/flinkk8soperator/pkg/controller/config"
	"github.com/stretchr/testify/assert"
)

func TestReplaceJobUrl(t *testing.T) {
	assert.Equal(t,
		"ABC.lyft.xyz",
		ReplaceJobURL("{{$jobCluster}}.lyft.xyz", "ABC"))
}

func initTestConfigForIngress() error {
	return config2.ConfigSection.SetConfig(&config2.Config{
		FlinkIngressURLFormat: "{{$jobCluster}}.lyft.xyz",
	})
}
func TestGetFlinkUIIngressURL(t *testing.T) {
	err := initTestConfigForIngress()
	assert.Nil(t, err)
	assert.Equal(t,
		"ABC.lyft.xyz",
		GetFlinkUIIngressURL("ABC"))
}
