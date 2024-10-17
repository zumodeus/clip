package clip

import "fmt"

func findEntity[T iEntity](entities []T, name, short, long string) T {
	for _, entity := range entities {
		if name != "" && entity.getName() == name {
			return entity
		}

		if short != "" && entity.getShort() == short {
			return entity
		}

		if long != "" && entity.getLong() == long {
			return entity
		}
	}
	var zero T
	return zero
}

type container struct {
	options  []*option
	commands []*command
}

func (c *container) findOneCommand(name, short, long string) *command {
	return findEntity(c.commands, name, short, long)
}

func (c *container) findOneOption(name, short, long string) *option {
	return findEntity(c.options, name, short, long)
}

func (c *container) AddNewEntity(newEntity interface{}) {
	if entity, ok := newEntity.(*command); ok {
		if c.findOneCommand(entity.name, entity.short, entity.long) != nil {
			NewException(
				DUPLICATE,
				fmt.Sprintf("command %s already exists.", entity.name),
				nil,
				false,
			)
		}
		c.commands = append(c.commands, entity)
		return
	}

	if entity, ok := newEntity.(*option); ok {
		if c.findOneOption(entity.name, entity.short, entity.long) != nil {
			NewException(
				DUPLICATE,
				fmt.Sprintf("option %s already exists.", entity.name),
				nil,
				false,
			)
		}
		c.options = append(c.options, entity)
		return
	}

	NewException(INVALID_TYPE, "newEntity type is not valid.", nil, false)
}
