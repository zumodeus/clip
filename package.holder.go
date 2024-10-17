package clip

type parameter struct {
	key   string
	value string
}

type Holder struct {
	flags      []string
	arguments  []string
	parameters []parameter
}

func (h *Holder) addFlag(newFlag string) bool {
	for _, flag := range h.flags {
		if flag == newFlag {
			return false
		}
	}

	h.flags = append(h.flags, newFlag)
	return true
}

func (h *Holder) addParameter(newParameter string, unique bool) bool {
	if unique {
		for _, parameter := range h.parameters {
			if parameter.key == newParameter {
				return false
			}
		}
	}

	h.parameters = append(h.parameters, parameter{newParameter, ""})
	return true
}

func (h *Holder) addArgument(newArgument string, supports supports) bool {
	for i := range h.parameters {
		if h.parameters[i].value == "" {
			h.parameters[i].value = newArgument
			return true
		}
	}

	if supports == STRINGS || (len(h.arguments) == 0 && supports == STRING) {
		h.arguments = append(h.arguments, newArgument)
		return true
	}

	return false
}

func (h *Holder) GetFlags() []string {
	return h.flags
}

func (h *Holder) HasFlag(flag string) bool {
	for _, f := range h.flags {
		if f == flag {
			return true
		}
	}
	return false
}

func (h *Holder) GetParameterValue(key string) (string, bool) {
	for _, parameter := range h.parameters {
		if parameter.key == key {
			return parameter.value, true
		}
	}
	return "", false
}

func (h *Holder) GetParameterValues(key string) []string {
	var values []string
	for _, parameter := range h.parameters {
		if parameter.key == key {
			values = append(values, parameter.value)
		}
	}
	return values
}

func (h *Holder) GetArguments() []string {
	return h.arguments
}
