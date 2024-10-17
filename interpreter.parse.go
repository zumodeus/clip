package clip

import "fmt"

func parse(args []string, command *command) {
	handler := newHandler()

	for _, arg := range args {
		if matchLongOption(arg) {
			handler.appendOption("", arg[2:], command)
		} else if matchShortOption(arg) {
			for _, short := range arg[1:] {
				handler.appendOption(string(short), "", command)
			}
		} else if cmd := command.container.findOneCommand(
			"",
			arg,
			arg); cmd != nil && handler.length() == 0 {
			command = cmd
		} else {
			handler.appendArgument(arg, command.supports)
		}
	}

	if len(handler.holders) == 0 {
		handler.getLastHolder()
	}

	if lookingForHelpFlag(handler.holders) {
		printHelp(command)
		return
	}

	if required, found := lookingForRequiredNotSetted(
		handler.holders,
		command.container.options); found {
		NewException(
			REQUIRED,
			fmt.Sprintf(
				"Parameter -%s | --%s is required.",
				required.getShort(),
				required.getLong(),
			),
			command,
			true,
		)
	}

	if lookingForMissingArguments(
		handler.holders,
		command.supports,
	) {
		NewException(
			MISSING,
			"No arguments found for the command.",
			command,
			true,
		)
	}

	for _, holder := range handler.holders {
		command.handler(holder)
	}
}
