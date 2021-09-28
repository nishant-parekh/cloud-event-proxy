package cloud_event_proxy_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCloudEventProxy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudEventProxy Suite")
}
