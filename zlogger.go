package zlogger

// A wrapper around zerolog

// playing with zerolog
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
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Debug().Msg(TraceColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
	l.Logger.Trace().Msg(TraceColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Debug(args ...interface{}) {
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Debug().Msg(DebugColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
	l.Logger.Debug().Msg(DebugColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Info(args ...interface{}) {
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Info().Msg(InfoColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
	l.Logger.Info().Msg(InfoColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Warn(args ...interface{}) {
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Warn().Msg(WarnColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
	l.Logger.Warn().Msg(WarnColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Error(args ...interface{}) {
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Error().Msg(ErrorColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
	l.Logger.Error().Msg(ErrorColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Fatal(args ...interface{}) {
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Fatal().Msg(FatalColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
	l.Logger.Fatal().Msg(FatalColor + fmt.Sprint(args...) + ResetColor)
}

func (l *ZLogger) Panic(args ...interface{}) {
	/*
		if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
			l.Logger.Panic().Msg(PanicColor + fmt.Sprintf(args[0].(string), args[1:]...) + ResetColor)
			return
		}
	*/
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
		/*
			output.FormatTimestamp = func(i interface{}) string {
				// Note: all layouts should use the standard reference time: Mon Jan 2 15:04:05 MST 2006
				t, _ := time.Parse("2006-1-2 15:04:05", i.(string))
				return fmt.Sprintf("[%s%s%s]", BrightCyanColor, t.Format("2006-01-02 15:04:05"), ResetColor)
			}
		*/

		zlogger = &ZLogger{}
		zlogger.Logger = zerolog.New(output).With().Timestamp().Logger()
	}

}

/*
// playing with logrus
import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	ResetColor         = "\033[0m"
	RedColor           = "\033[31m"
	GreenColor         = "\033[32m"
	YellowColor        = "\033[33m"
	BlueColor          = "\033[34m"
	MagentaColor       = "\033[35m"
	CyanColor          = "\033[36m"
	WhiteColor         = "\033[37m"
	BrightRedColor     = "\033[91m"
	BrightGreenColor   = "\033[92m"
	BrightYellowColor  = "\033[93m"
	BrightBlueColor    = "\033[94m"
	BrightMagentaColor = "\033[95m"
	BrightCyanColor    = "\033[96m"
	BrightWhiteColor   = "\033[97m"
)

type logWrapper struct {
	*logrus.Logger
}

var logger *logWrapper

func GetLogger() *logWrapper {
	return logger
}

func (l *logWrapper) Debug(args ...interface{}) {
	// rjk
	if f, ok := args[0].(string); ok {
		if strings.Contains(f, "%") {
			l.Logger.Debugf(f, args[1:]...)
			return
		}
	}
	l.Logger.Debug(args)
}

func (l *logWrapper) Info(args ...interface{}) {
	// rjk
	if (len(args) > 0) && (strings.Contains(args[0].(string), "%")) {
		l.Logger.Infof(args[0].(string), args[1:]...)
		return
	}

	l.Logger.Info(args)
}

func (l *logWrapper) Warn(args ...interface{}) {
	// rjk
	if f, ok := args[0].(string); ok {
		if strings.Contains(f, "%") {
			l.Logger.Warnf(f, args[1:]...)
			return
		}
	}
	l.Logger.Warn(args)
}

func (l *logWrapper) Error(args ...interface{}) {
	// rjk
	if f, ok := args[0].(string); ok {
		if strings.Contains(f, "%") {
			l.Logger.Errorf(f, args[1:]...)
			return
		}
	}
	l.Logger.Error(args)
}

func (l *logWrapper) Fatal(args ...interface{}) {
	// rjk
	if f, ok := args[0].(string); ok {
		if strings.Contains(f, "%") {
			l.Logger.Fatalf(f, args[1:]...)
			return
		}
	}
	l.Logger.Fatal(args)
}

type customFormatter struct {
	logrus.TextFormatter
}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
	var levelColor string
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = BrightWhiteColor // white
	case logrus.WarnLevel:
		levelColor = YellowColor // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = BrightRedColor // red
	case logrus.InfoLevel:
		levelColor = BrightBlueColor
	default:
		levelColor = BrightBlueColor // blue
	}
	//return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
	return []byte(fmt.Sprintf("%s[%s] %s%s\n", levelColor, strings.ToUpper(entry.Level.String()), entry.Message, ResetColor)), nil
}

func init() {
	if logger == nil {
		logger = &logWrapper{&logrus.Logger{}}
	}
	logger.Out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&customFormatter{
		logrus.TextFormatter{
			ForceColors:            true,
			DisableTimestamp:       true,
			DisableLevelTruncation: true,
		},
	})
}
*/
/////////////////////////////////////////////////////////////////////

// playing with zap logger
/*
import (
	"encoding/json"

	"go.uber.org/zap"
)

// var log = logger.GetLogger()
var log *zap.SugaredLogger

func init() {

	zapConfig := []byte(`{
		"level" : "info",
		"encoding": "json",
		"outputPaths":["stdout"],
		"errorOutputPaths":["stderr"],
		"encoderConfig": {
			"messageKey":"message",
			"levelKey":"level",
			"levelEncoder":"capital"
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

	log = logger.Sugar()

	defer logger.Sync()
}
*/
