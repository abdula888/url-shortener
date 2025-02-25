package log

type Logger interface {
	Info(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}
