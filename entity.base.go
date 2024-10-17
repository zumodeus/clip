package clip

type entity struct {
	name        string
	short       string
	long        string
	description string
}

type iEntity interface {
	getName() string
	getShort() string
	getLong() string
}

func (e *entity) getName() string {
	return e.name
}

func (e *entity) getShort() string {
	return e.short
}

func (e *entity) getLong() string {
	return e.long
}
