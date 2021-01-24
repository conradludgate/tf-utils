package example

type Foo struct {
	// Some value
	A string `tf:"default=foo"`

	// Some other value
	B int `tf:"required"`

	// Simple lists can be implemented
	List []string

	// Complex lists can too, as long as Baz
	// implements the 'Schema' interface
	AnotherList []Baz

	// Maps can only be over simple types (terraform limitation)
	Map map[string]int

	// map[int]... represents a Set.
	// If Bar implements the `Set` interface,
	// then that will be the Set function
	Set map[int]Bar
}

type Bar struct {
	C string
}

type Baz struct {
	D float32
}
