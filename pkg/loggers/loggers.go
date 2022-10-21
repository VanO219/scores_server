package loggers

import "VanO/scores_server/model/logger"

type LoggersWrapper struct {
	level int64 // Уровень логирования. 0 - логирование выключено; 1 - errors; 2 - errors, infos; 3 - errors, infos, debug.
	lg    []logger.Loggers
}

func (l *LoggersWrapper) Error(input ...interface{}) {
	if l.level < 1 {
		return
	}

	for _, lo := range l.lg {
		lo.Error(input)
	}
}

func (l *LoggersWrapper) Errorf(format string, args ...interface{}) {
	if l.level < 1 {
		return
	}

	for _, lo := range l.lg {
		lo.Errorf(format, args)
	}
}

func (l *LoggersWrapper) Info(input ...interface{}) {
	if l.level < 2 {
		return
	}

	for _, lo := range l.lg {
		lo.Info(input)
	}
}

func (l *LoggersWrapper) Infof(format string, args ...interface{}) {
	if l.level < 1 {
		return
	}

	for _, lo := range l.lg {
		lo.Infof(format, args)
	}
}

func (l *LoggersWrapper) Debug(input ...interface{}) {
	if l.level < 3 {
		return
	}

	for _, lo := range l.lg {
		lo.Debug(input)
	}
}

func (l *LoggersWrapper) Debugf(format string, args ...interface{}) {
	if l.level < 1 {
		return
	}

	for _, lo := range l.lg {
		lo.Debugf(format, args)
	}
}

// NewLoggersWrapper Инициализирует глобальную переменную Loggers
func NewLoggersWrapper(level int64, lg ...logger.Loggers) *LoggersWrapper {
	if level < 0 {
		level = 0
	}
	if level > 3 {
		level = 3
	}
	loggers := &LoggersWrapper{
		level: level,
		lg:    lg,
	}
	return loggers
}
