package langs

import (
	"github.com/denizgursoy/gotouch/internal/logger"
)

func NewEmptySetupChecker() LanguageChecker {
	return &emptySetupChecker{}
}

func (e *emptySetupChecker) CompletePreTask() error {
	return nil
}

func (e *emptySetupChecker) CheckDependency(dependency interface{}) error {
	return nil
}

type emptySetupChecker struct {
	Logger logger.Logger
}

func (e *emptySetupChecker) GetDependency(dependency interface{}) error {
	return nil
}

func (e *emptySetupChecker) CheckSetup() error {
	return nil
}
