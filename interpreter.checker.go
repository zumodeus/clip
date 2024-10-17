package clip

func lookingForHelpFlag(
	holders []*Holder,
) bool {
	for _, holder := range holders {
		if holder.HasFlag("help") {
			return true
		}
	}
	return false
}

func lookingForRequiredNotSetted(
	holders []*Holder,
	options []*option,
) (e iEntity, f bool) {
	required := []*option{}
	for _, option := range options {
		if option.required && option.kind == PARAMETER {
			required = append(required, option)
		}
	}

	for _, holder := range holders {
		for _, option := range required {
			if _, exists := holder.GetParameterValue(option.name); !exists {
				return option, true
			}
		}
	}

	return nil, false
}

func lookingForMissingArguments(
	holders []*Holder,
	supports supports,
) bool {
	if supports == NULL {
		return false
	}

	for _, holder := range holders {
		length := len(holder.arguments)
		if length == 0 {
			return true
		}
	}

	return false
}
