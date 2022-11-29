package mock

// DataStoreMock is a mock implementation of the datastore interface
type DataStoreMock struct {
}

// NewDataStoreMock returns a new instance of the mock datastore
func NewDataStoreMock() *DataStoreMock {
	return &DataStoreMock{}
}
