package clip

import "fmt"

type kind string

const (
	FLAG            kind = "FLAG"
	PARAMETER       kind = "PARAMETER"
	POSITIONAL_FLAG kind = "POSITIONAL_FLAG"
)

type option struct {
	entity
	kind     kind
	unique   bool
	required bool
}

func NewOption(
	name, short, long, description string,
	kind kind,
	unique, required bool,
) *option {
	if !matchNameString(name) {
		NewException(
			NOT_MATCH,
			fmt.Sprintf("name %s does not match the pattern.", name),
			nil,
			false,
		)
	}

	if !matchShortString(short) {
		NewException(
			NOT_MATCH,
			fmt.Sprintf("short %s does not match the pattern.", short),
			nil,
			false,
		)
	}

	if !matchLongString(long) {
		NewException(
			NOT_MATCH,
			fmt.Sprintf("long %s does not match the pattern.", long),
			nil,
			false,
		)
	}

	return &option{
		entity{
			name,
			short,
			long,
			description,
		},
		kind,
		unique,
		required,
	}
}
