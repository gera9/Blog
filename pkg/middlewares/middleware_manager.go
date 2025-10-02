package middlewares

type ContextKey struct {
	Name string
}

func (ck ContextKey) String() string {
	return "blog context key value " + ck.Name
}

type MiddlewareManager struct {
}
