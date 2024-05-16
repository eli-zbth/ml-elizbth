package logger

import (
	"testing"
	logrus "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)


const customField = "custom field"

func setupLoggerTest() (Logger, *test.Hook){
    testLog, hook := test.NewNullLogger()
    logger := NewLogger()
    logger.log.Logger = testLog
    return logger,hook
}



func TestLogger(t*testing.T){

  t.Run("Should log info", func(t *testing.T) {

    logger,hook := setupLoggerTest()

    logger.Info("Info message")
    assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
    assert.Equal(t, "Info message", hook.LastEntry().Message)
  
    hook.Reset()
    assert.Nil(t, hook.LastEntry())
	})

  t.Run("Should log info with string format", func(t *testing.T) {

    logger,hook := setupLoggerTest()

    logger.Infof("Info message with field : %s",customField )
    assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
    assert.Equal(t, "Info message with field : custom field", hook.LastEntry().Message)
  
    hook.Reset()
    assert.Nil(t, hook.LastEntry())
	})



  t.Run("Should log error", func(t *testing.T) {

    logger,hook := setupLoggerTest()

    logger.Error("Error message")
    assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
    assert.Equal(t, "Error message", hook.LastEntry().Message)
  
    hook.Reset()
    assert.Nil(t, hook.LastEntry())
	})


  t.Run("Should log error with string format", func(t *testing.T) {

    logger,hook := setupLoggerTest()

    logger.Errorf("Error message with field : %s",customField )
    assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
    assert.Equal(t, "Error message with field : custom field", hook.LastEntry().Message)
  
    hook.Reset()
    assert.Nil(t, hook.LastEntry())
	})

}
