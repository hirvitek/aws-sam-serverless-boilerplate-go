package mocks

type EventDispatcher struct{}

func (e EventDispatcher) Dispatch(message interface{}) error {
	panic("implement me")
}

func (e EventDispatcher) DispatchWithAttributes(message interface{}, messageAttributes map[string]string) error {
	panic("implement me")
}
