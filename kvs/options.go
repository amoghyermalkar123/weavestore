package kvs

type StoreOptions struct {
	MaxSize uint64
}

type OptionSetter func(o *StoreOptions)

func LoadOptions(opts ...OptionSetter) *StoreOptions {
	newOptionsInstance := new(StoreOptions)

	for _, fn := range opts {
		fn(newOptionsInstance)
	}

	return newOptionsInstance
}

func WithMaxRAMSize(size uint64) OptionSetter {
	return func(o *StoreOptions) {
		o.MaxSize = size
	}
}
