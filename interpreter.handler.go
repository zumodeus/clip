package clip

import "fmt"

type handler struct {
	holders []*Holder
}

func convertOptionString(short, long string) string {
	if short != "" {
		return "-" + short
	}
	return "--" + long
}

func (h *handler) getLastHolder() *Holder {
	if len(h.holders) == 0 {
		h.holders = append(h.holders, &Holder{
			make([]string, 0),
			make([]string, 0),
			make([]parameter, 0),
		})
	}

	return h.holders[len(h.holders)-1]
}

func (h *handler) optionAlreadyExistsInAnotherContext(
	name string,
) bool {
	for _, holder := range h.holders {
		if holder.HasFlag(name) {
			return true
		}
	}
	return false
}

func (h *handler) length() int {
	return len(h.holders)
}

func (h *handler) appendOption(
	short, long string,
	command *command,
) {
	o := command.container.findOneOption("", short, long)

	if o == nil {
		NewException(
			NOT_FOUND,
			fmt.Sprintf(
				"option %s was not found.",
				convertOptionString(short, long),
			),
			command,
			true,
		)
	}

	switch o.kind {
	case FLAG:
		if h.optionAlreadyExistsInAnotherContext(o.name) && o.unique {
			NewException(
				DUPLICATE,
				fmt.Sprintf(
					"Flag -%s | --%s must be used once.",
					o.short,
					o.long,
				),
				command,
				true,
			)
		}
		h.getLastHolder().addFlag(o.name)
	case PARAMETER:
		if !h.getLastHolder().addParameter(o.name, o.unique) {
			NewException(
				DUPLICATE,
				fmt.Sprintf("Parameter %s already exists.", o.name),
				command,
				true,
			)
		}
	case POSITIONAL_FLAG:
		if len(h.holders) == 0 || len(h.holders) > 0 && !(h.getLastHolder().HasFlag(o.name) &&
			o.unique) {
			newHolder := &Holder{
				make([]string, 0),
				make([]string, 0),
				make([]parameter, 0),
			}
			newHolder.addFlag(o.name)
			h.holders = append(h.holders, newHolder)
		} else {
			NewException(
				DUPLICATE,
				fmt.Sprintf(
					"Flag -%s | --%s must be used once.",
					o.short,
					o.long,
				),
				command,
				true,
			)
		}
	default:
		NewException(
			INVALID_TYPE,
			"Unknown option kind.",
			command,
			true,
		)
	}
}

func (h *handler) appendArgument(
	argument string,
	supports supports,
) {
	if len(h.holders) == 0 {
		h.holders = append(h.holders, &Holder{
			make([]string, 0),
			make([]string, 0),
			make([]parameter, 0),
		})
	}

	for _, holder := range h.holders {
		if holder.addArgument(argument, supports) {
			break
		}
	}
}

func newHandler() *handler {
	return &handler{
		make([]*Holder, 0),
	}
}
