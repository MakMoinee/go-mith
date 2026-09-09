package mithclientrest

type Operation struct {
	Method  string
	Path    string
	Headers map[string]string
}

// OperationResolver determines which REST operation should be
// performed for a given payload.
type OperationResolver interface {
	OperationFor(payload any) (Operation, error)
}
