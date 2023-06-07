package kvs

// StoreOptions defines the configuration settings for the KV engine
type StoreOptions struct {
	MaxSize int
}

type OptionSetter func(o *StoreOptions)

func LoadOptions(opts ...OptionSetter) *StoreOptions {
	newOptionsInstance := new(StoreOptions)

	for _, fn := range opts {
		fn(newOptionsInstance)
	}

	return newOptionsInstance
}

// WithMaxRAMSize defines the max capacity of the in-memory store
func WithMaxRAMSize(size int) OptionSetter {
	return func(o *StoreOptions) {
		o.MaxSize = size
	}
}
