package logger

type ILogger interface {
	With(k, v string) ILogger
	Errorf(fmt string, p ...interface{})
	Warnf(fmt string, p ...interface{})
	Infof(fmt string, p ...interface{})
	Debugf(fmt string, p ...interface{})
}

type LogConfig struct {
	Level    string `env:"LOG_LEVEL" envDefault:"debug"`
	IsPretty bool   `env:"LOG_PRETTY" envDefault:""`
}

type noLogger struct {
}

func (l *noLogger) With(k, v string) ILogger {
	return l
}
func (l *noLogger) Errorf(fmt string, p ...interface{}) {}
func (l *noLogger) Warnf(fmt string, p ...interface{})  {}
func (l *noLogger) Infof(fmt string, p ...interface{})  {}
func (l *noLogger) Debugf(fmt string, p ...interface{}) {}

var NoLogger = &noLogger{}
