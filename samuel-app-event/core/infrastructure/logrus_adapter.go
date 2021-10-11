package infrastructure

import (
	"github.com/ahmetb/go-linq/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
)

type LogrusAdapter struct {
	echo.Logger
	Instance *logrus.Logger
}

func NewLogrusAdapter(instance *logrus.Logger) *LogrusAdapter {
	adapter := new(LogrusAdapter)
	adapter.Instance = instance

	return adapter
}

var levelMapping = map[logrus.Level]log.Lvl{
	logrus.DebugLevel: log.DEBUG,
	logrus.InfoLevel:  log.INFO,
	logrus.WarnLevel:  log.WARN,
	logrus.ErrorLevel: log.ERROR,
}

func (adapter LogrusAdapter) Output() io.Writer {
	return adapter.Instance.Out
}

func (adapter *LogrusAdapter) SetOutput(writer io.Writer) {
	adapter.Instance.Out = writer
}

func (adapter LogrusAdapter) Prefix() string {
	// TODO: Has to be decided later ...
	return ""
}

func (adapter LogrusAdapter) SetPrefix(_ string) {
	// TODO: Has to be decided later ...
}

func (adapter LogrusAdapter) Level() log.Lvl {
	return levelMapping[adapter.Instance.Level]
}

func (adapter LogrusAdapter) SetLevel(level log.Lvl) {
	logrusLevel := linq.From(levelMapping).
		FirstWithT(func(pair linq.KeyValue) bool {
			return pair.Value == level
		}).(linq.KeyValue).Key.(logrus.Level)

	adapter.Instance.SetLevel(logrusLevel)
}

func (adapter LogrusAdapter) SetHeader(_ string) {
	// TODO: Has to be decided later ...
}

func (adapter LogrusAdapter) Print(args ...interface{}) {
	adapter.Instance.Print(args...)
}

func (adapter LogrusAdapter) Printf(format string, args ...interface{}) {
	adapter.Instance.Printf(format, args...)
}

func (adapter LogrusAdapter) Printj(json log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(json)).Print()
}

func (adapter LogrusAdapter) Debug(args ...interface{}) {
	adapter.Instance.Debug(args...)
}

func (adapter LogrusAdapter) Debugf(format string, args ...interface{}) {
	adapter.Instance.Debugf(format, args...)
}

func (adapter LogrusAdapter) Debugj(json log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(json)).Debug()
}

func (adapter LogrusAdapter) Info(args ...interface{}) {
	adapter.Instance.Info(args...)
}

func (adapter LogrusAdapter) Infof(format string, args ...interface{}) {
	adapter.Instance.Infof(format, args...)
}

func (adapter LogrusAdapter) Infoj(json log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(json)).Info()
}

func (adapter LogrusAdapter) Warn(args ...interface{}) {
	adapter.Instance.Warn(args...)
}

func (adapter LogrusAdapter) Warnf(format string, args ...interface{}) {
	adapter.Instance.Warnf(format, args...)
}

func (adapter LogrusAdapter) Warnj(json log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(json)).Warn()
}

func (adapter LogrusAdapter) Error(args ...interface{}) {
	adapter.Instance.Error(args...)
}

func (adapter LogrusAdapter) Errorf(format string, args ...interface{}) {
	adapter.Instance.Errorf(format, args...)
}

func (adapter LogrusAdapter) Errorj(json log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(json)).Error()
}

func (adapter LogrusAdapter) Fatal(args ...interface{}) {
	adapter.Instance.Fatal(args...)
}

func (adapter LogrusAdapter) Fatalf(format string, args ...interface{}) {
	adapter.Instance.Fatalf(format, args...)
}

func (adapter LogrusAdapter) Fatalj(json log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(json)).Fatal()
}

func (adapter LogrusAdapter) Panic(args ...interface{}) {
	adapter.Instance.Panic(args...)
}

func (adapter LogrusAdapter) Panicf(format string, args ...interface{}) {
	adapter.Instance.Panicf(format, args...)
}

func (adapter LogrusAdapter) Panicj(j log.JSON) {
	adapter.Instance.WithFields(logrus.Fields(j)).Panic()
}
