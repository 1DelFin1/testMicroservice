package domains

type IDGenerator interface {
	NewID() string
}

type IDGeneratorImpl struct {
}
