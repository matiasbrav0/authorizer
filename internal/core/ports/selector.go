package ports

type Selector interface {
	OperationSelector(operation []byte) error
}
