package clip

import (
	"fmt"
	"os"
	"sync"
)

type clip struct {
	command *command
}

var instance *clip
var once sync.Once

func (c *clip) Parse(args []string) {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	parse(args, c.command)
}

func (c *clip) ErrorCatcher() {
	if err, ok := recover().(exception); ok {
		fmt.Printf("%s: %s\n\n", err.typeof, err.message)
		if err.shouldPrintHelp {
			printHelp(err.command)
		}
		os.Exit(1)
	}
}

func (c *clip) GetContainer() *container {
	return c.command.container
}

func GetCLIP(
	name, description string,
	defaultSupports supports,
	defaultHandler func(*Holder),
) *clip {
	once.Do(func() {
		cmd := &command{
			entity{name, "", "", description},
			defaultSupports,
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
			defaultHandler,
		}
		instance = &clip{
			cmd,
		}
	})
	return instance
}
