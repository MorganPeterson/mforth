package result

type Result[T any] struct {
	ok T
	err error
}

// produce a value with error == nil
func Ok[T any](val T) Result[T] {
	return Result[T]{ok: val, err: nil}
}

// produce an error with garbage value
func Error[T any](err error) Result[T] {
	return Result[T]{ok: *new(T), err: err}
}

// check for an error
func (r *Result[T]) IsOk() bool {
	return r.err == nil
}

// get value
func (r *Result[T]) UnwrapVal() T {
	return r.ok
}

// get error
func (r *Result[T]) UnwrapErr() error {
	return r.err
}
