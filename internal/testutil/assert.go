package testutil

import (
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) AssertPods(pods []v1.Pod, expected []map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (suite *TestSuite) AssertPod(pod v1.Pod, expected map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (suite *TestSuite) AssertLog(output *test.Hook, level log.Level, msg string, fields log.Fields) {
	_ = "STUB: not implemented"
	return
}
