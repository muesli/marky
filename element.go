package marky

// Element is implemented by any markdown element we can render.
type Element interface {
	String() string
}
