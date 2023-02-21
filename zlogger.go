package zlogger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

const (
	ResetColor         = "\033[0m"
	BlackColor         = "\033[30m"
	RedColor           = "\033[31m"
	GreenColor         = "\033[32m"
	YellowColor        = "\033[33m"
	BlueColor          = "\033[34m"
	MagentaColor       = "\033[35m"
	CyanColor          = "\033[36m"
	WhiteColor         = "\033[37m"
	GreyColor          = "\033[90m"
	BrightRedColor     = "\033[91m"
	BrightGreenColor   = "\033[92m"
	BrightYellowColor  = "\033[93m"
	BrightBlueColor    = "\033[94m"
	BrightMagentaColor = "\033[95m"
	BrightCyanColor    = "\033[96m"
	BrightWhiteColor   = "\033[97m"
)

type ZLogger struct {
	// just a wrapper around zerolog, adding some helper functions
	zerolog.Logger
}

const (
	DefaultColor = WhiteColor
	TraceColor   = WhiteColor
	DebugColor   = BlueColor
	InfoColor    = GreenColor
	WarnColor    = YellowColor
	ErrorColor   = RedColor
	FatalColor   = BrightRedColor
	PanicColor   = BrightRedColor
)

var zlogger *ZLogger = nil

func GetLogger() *ZLogger {
	return zlogger
}

// Level defines log levels.
type LogLevel int8

const (
	// DebugLevel defines debug log level.
	DebugLevel LogLevel = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled

	// TraceLevel defines trace log level.
	TraceLevel LogLevel = -1
	// Values less than TraceLevel are handled as numbers.
)

type ZLoggerConfig struct {
	LogLevel `mapstructure:"level"`
}

func (l *ZLogger) SetLevel(ll LogLevel) {
	newlogger := l.Logger.Level(zerolog.Level(ll))
	l.Logger = newlogger
}

func (l *ZLogger) NewLogger(ll LogLevel) *ZLogger {
	newlogger := l.Logger.Level(zerolog.Level(ll))

	return &ZLogger{newlogger}
}

func (l *ZLogger) Tracef(format string, args ...interface{}) {
	l.Logger.Trace().Msgf(TraceColor+format+ResetColor, args...)
}

func (l *ZLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debug().Msgf(DebugColor+format+ResetColor, args...)
}

func (l *ZLogger) Infof(format string, args ...interface{}) {
	l.Logger.Info().Msgf(InfoColor+format+ResetColor, args...)
}

func (l *ZLogger) Printf(format string, args ...interface{}) {
	l.Logger.Info().Msgf(InfoColor+format+ResetColor, args...)
}

func (l *ZLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warn().Msgf(WarnColor+format+ResetColor, args...)
}

func (l *ZLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Error().Msgf(ErrorColor+format+ResetColor, args...)
}

func (l *ZLogger) Fatalf(format string, args ...interface{}) {
	l.Logger.Fatal().Msgf(FatalColor+format+ResetColor, args...)
}

func (l *ZLogger) Panicf(format string, args ...interface{}) {
	l.Logger.Panic().Msgf(PanicColor+format+ResetColor, args...)
}

func (l *ZLogger) Trace(args ...interface{}) {
	l.Logger.Trace().Msg(TraceColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Debug(args ...interface{}) {
	l.Logger.Debug().Msg(DebugColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Info(args ...interface{}) {
	l.Logger.Info().Msg(InfoColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Print(args ...interface{}) {
	l.Logger.Info().Msg(InfoColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Warn(args ...interface{}) {
	l.Logger.Warn().Msg(WarnColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Error(args ...interface{}) {
	l.Logger.Error().Msg(ErrorColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal().Msg(FatalColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Panic(args ...interface{}) {
	l.Logger.Panic().Msg(PanicColor + fmt.Sprint(args...) + ResetColor)
}

func init() {
	if zlogger == nil {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		output := zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false}
		output.TimeFormat = GreyColor + "[2006-1-2 15:04:05]" + ResetColor
		output.FormatLevel = func(i interface{}) string {
			var formatLevelColor string = ""
			switch i.(string) {
			case "trace":
				formatLevelColor = TraceColor
			case "debug":
				formatLevelColor = DebugColor
			case "info":
				formatLevelColor = InfoColor
			case "warn":
				formatLevelColor = WarnColor
			case "error":
				formatLevelColor = ErrorColor
			case "fatal":
				formatLevelColor = FatalColor
			case "panic":
				formatLevelColor = PanicColor
			default:
				formatLevelColor = DefaultColor
			}
			return ResetColor + formatLevelColor + strings.ToUpper(fmt.Sprintf("[%s]", i)) + ResetColor
		}
		output.FormatMessage = func(i interface{}) string {
			return ResetColor + fmt.Sprintf("%s", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return ResetColor + fmt.Sprintf("(%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return ResetColor + fmt.Sprintf("%s)", i)
		}

		zlogger = &ZLogger{}
		zlogger.Logger = zerolog.New(output).With().Timestamp().Logger()

		ResetDefault(zlogger)
	}

}

// expose functions (like export in js) for the static log
var (
	Info     = zlogger.Info
	Infof    = zlogger.Infof
	Print     = zlogger.Print
	Printf    = zlogger.Printf
	Warn     = zlogger.Warn
	Warnf    = zlogger.Warnf
	Error    = zlogger.Error
	Errorf   = zlogger.Errorf
	Panic    = zlogger.Panic
	Panicf   = zlogger.Panicf
	Fatal    = zlogger.Fatal
	Fatalf   = zlogger.Fatalf
	Debug    = zlogger.Debug
	Debugf   = zlogger.Debugf
	SetLevel = zlogger.SetLevel
)

func ResetDefault(l *ZLogger) {
	var zLogger = l
	Info = zLogger.Info
	Infof = zLogger.Infof
	Print = zLogger.Print
	Printf = zLogger.Printf
	Warn = zLogger.Warn
	Warnf = zLogger.Warnf
	Error = zLogger.Error
	Errorf = zLogger.Errorf
	Panic = zLogger.Panic
	Panicf = zLogger.Panicf
	Fatal = zLogger.Fatal
	Fatalf = zLogger.Fatalf
	Debug = zLogger.Debug
	Debugf = zLogger.Debugf
	SetLevel = zLogger.SetLevel
}

/// zap logger test
/*

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogLevel = zapcore.Level

type ZLogger struct {
	l     *zap.SugaredLogger // zap ensure that zap.Logger is safe for concurrent use
	level LogLevel
}

type ZLoggerConfig struct {
	LogLevel `mapstructure:"level"`
}

var stdlog *ZLogger

func GetLogger() *ZLogger {
	return stdlog
}

func (l *ZLogger) SetLevel(ll LogLevel) {
	stdlog.level = ll
	//newlogger := l.Logger.Level(zerolog.Level(ll))
	//l.Logger = newlogger
}

func init() {
	zapConfig := []byte(`{
		"level" : "info",
		"encoding": "console",
		"outputPaths":["stdout"],
		"errorOutputPaths":["stderr"],
		"encoderConfig": {
			"messageKey":"message",
			"levelKey":"level",
			"levelEncoder":"capitalColor"
		}
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(zapConfig, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()

	if err != nil {
		panic(err)
	}

	stdlog = &ZLogger{
		l:     logger.Sugar(),
		level: cfg.Level.Level(),
	}
	ResetDefault(stdlog)
	//log = logger.Sugar()

	defer logger.Sync()
}

func (l *ZLogger) Debug(args ...interface{}) {
	l.l.Debug(args...)
}
func (l *ZLogger) Info(args ...interface{}) {
	l.l.Info(args...)
}
func (l *ZLogger) Warn(args ...interface{}) {
	l.l.Warn(args...)
}
func (l *ZLogger) Error(args ...interface{}) {
	l.l.Error(args...)
}
func (l *ZLogger) Fatal(args ...interface{}) {
	l.l.Fatal(args...)
}
func (l *ZLogger) DPanic(args ...interface{}) {
	l.l.DPanic(args...)
}
func (l *ZLogger) Panic(args ...interface{}) {
	l.l.Panic(args...)
}
func (l *ZLogger) Debugf(msg string, args ...interface{}) {
	l.l.Debugf(msg, args...)
}
func (l *ZLogger) Infof(msg string, args ...interface{}) {
	l.l.Infof(msg, args...)
}
func (l *ZLogger) Warnf(msg string, args ...interface{}) {
	l.l.Warnf(msg, args...)
}
func (l *ZLogger) Errorf(msg string, args ...interface{}) {
	l.l.Errorf(msg, args...)
}
func (l *ZLogger) Fatalf(msg string, args ...interface{}) {
	l.l.Fatalf(msg, args...)
}
func (l *ZLogger) DPanicf(msg string, args ...interface{}) {
	l.l.DPanicf(msg, args...)
}
func (l *ZLogger) Panicf(msg string, args ...interface{}) {
	l.l.Panicf(msg, args...)
}

// expose functions (like export in js) for the static log
var (
	Info     = stdlog.Info
	Infof    = stdlog.Infof
	Warn     = stdlog.Warn
	Warnf    = stdlog.Warnf
	Error    = stdlog.Error
	Errorf   = stdlog.Errorf
	DPanic   = stdlog.DPanic
	DPanicf  = stdlog.DPanicf
	Panic    = stdlog.Panic
	Panicf   = stdlog.Panicf
	Fatal    = stdlog.Fatal
	Fatalf   = stdlog.Fatalf
	Debug    = stdlog.Debug
	Debugf   = stdlog.Debugf
	SetLevel = stdlog.SetLevel
)

func ResetDefault(l *ZLogger) {
	stdlog = l
	Info = stdlog.Info
	Infof = stdlog.Infof
	Warn = stdlog.Warn
	Warnf = stdlog.Warnf
	Error = stdlog.Error
	Errorf = stdlog.Errorf
	DPanic = stdlog.DPanic
	DPanicf = stdlog.DPanicf
	Panic = stdlog.Panic
	Panicf = stdlog.Panicf
	Fatal = stdlog.Fatal
	Fatalf = stdlog.Fatalf
	Debug = stdlog.Debug
	Debugf = stdlog.Debugf
	SetLevel = stdlog.SetLevel
}

*/
