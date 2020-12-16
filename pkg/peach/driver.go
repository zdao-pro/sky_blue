package peach

// Driver defined paladin remote client impl
// each remote config center driver must do
// 1. implements `New` method
type Driver interface {
	New(args ...interface{}) (Client, error)
}
