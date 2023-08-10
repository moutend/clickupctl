package constant

type contextKey struct {
	name string
}

var (
	ClientContextKey    = &contextKey{"client"}
	TransportContextKey = &contextKey{"responseCache"}
)
