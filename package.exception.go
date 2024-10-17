package clip

type typeof string

const (
	FATAL        typeof = "FatalError"
	MISSING      typeof = "MissingError"
	REQUIRED     typeof = "RequiredError"
	NOT_FOUND    typeof = "NotFoundError"
	NOT_MATCH    typeof = "NotMatchError"
	DUPLICATE    typeof = "DuplicateError"
	INVALID_TYPE typeof = "InvalidTypeError"
)

type exception struct {
	typeof          typeof
	message         string
	command         *command
	shouldPrintHelp bool
}

func NewException(
	typeof typeof,
	message string,
	command *command,
	shouldPrintHelp bool,
) {
	panic(exception{
		typeof,
		message,
		command,
		shouldPrintHelp,
	})
}
