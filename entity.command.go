package clip

import "fmt"

type supports string

const (
	NULL    supports = "NULL"
	STRING  supports = "STRING"
	STRINGS supports = "STRINGS"
)

type command struct {
	entity
	supports  supports
	container *container
	handler   func(*Holder)
}

func (c *command) GetContainer() *container {
	return c.container
}

func NewCommand(
	name, short, long, description string,
	supports supports,
	handler func(*Holder),
) *command {
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

	return &command{
		entity{
			name,
			short,
			long,
			description,
		},
		supports,
		&container{
			[]*option{
				NewOption(
					"help",
					"h",
					"help",
					"Print helpful information about the current command.",
					FLAG,
					true,
					false,
				),
			},
			make([]*command, 0),
		},
		handler,
	}
}
