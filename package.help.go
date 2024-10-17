package clip

import "fmt"

func cleanCommand(command *command) string {
	if command.short == "" && command.long == "" {
		return command.name
	}
	return fmt.Sprintf(
		"[%s | %s]",
		command.short,
		command.long,
	)
}

func cleanArguments(supports supports) string {
	if supports == STRING {
		return " <arg>"
	} else if supports == STRINGS {
		return " [<arg>]"
	} else {
		return ""
	}
}

func printOptions(options []*option) {
	for _, opt := range options {
		fmt.Printf(
			"  -%s | --%-16s %s",
			opt.getShort(),
			opt.getLong(),
			opt.description,
		)
		if opt.unique {
			fmt.Print(" \033[1m[once]\033[0m")
		}
		if opt.kind == PARAMETER && opt.required {
			fmt.Print(" \033[1m[required]\033[0m")
		}
		fmt.Println()
	}
}

func printHelp(command *command) {
	positional_flags := []*option{}
	parameters := []*option{}
	flags := []*option{}

	for _, opt := range command.container.options {
		if opt.kind == PARAMETER {
			parameters = append(parameters, opt)
		} else if opt.kind == FLAG {
			flags = append(flags, opt)
		} else if opt.kind == POSITIONAL_FLAG {
			positional_flags = append(positional_flags, opt)
		}
	}

	fmt.Printf("%s\n", command.description)

	fmt.Printf("\nUsage: %s", cleanCommand(command))
	for _, opt := range append(positional_flags, append(flags, parameters...)...) {
		fmt.Printf(" [-%s | --%s]", opt.getShort(), opt.getLong())
		if opt.kind == PARAMETER {
			fmt.Printf(" <%s>", opt.getName())
		}
	}

	if len(command.container.commands) > 0 {
		fmt.Print(" <subcommands>")
	}
	fmt.Printf("%s", cleanArguments(command.supports))
	fmt.Println()

	fmt.Println("\nHere is a list of sub commands and options of this command:")
	fmt.Println()

	if len(command.container.commands) > 0 {
		fmt.Println("Subcommands:")
		for _, cmd := range command.container.commands {
			fmt.Printf(
				"  %s | %-19s %s\n",
				cmd.getShort(),
				cmd.getLong(),
				cmd.description,
			)
		}
		fmt.Println()
	}

	if len(positional_flags) > 0 {
		fmt.Println("Positional Flags:")
		printOptions(positional_flags)
		fmt.Println()
	}

	if len(flags) > 0 {
		fmt.Println("Flags:")
		printOptions(flags)
		fmt.Println()
	}

	if len(parameters) > 0 {
		fmt.Println("Parameters:")
		printOptions(parameters)
	}
}
