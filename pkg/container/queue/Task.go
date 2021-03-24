package queue

// Task you should implement
type Task interface {
	Run() error
}

// PriorityTask ..
type PriorityTask interface {
	Comparator
	Run() error
}

// Comparator like the java Comparator
type Comparator interface {
	Less(i interface{}) bool
}
