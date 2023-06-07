package kvs

import "time"

type StoreOptions struct {
	MaxRAMSize      uint64
	CleanupInterval time.Duration
}

type OptionSetter func(o *StoreOptions)

func LoadOptions(opts ...OptionSetter) *StoreOptions {
	newOptionsInstance := new(StoreOptions)

	for _, fn := range opts {
		fn(newOptionsInstance)
	}

	return newOptionsInstance
}

func WithMaxRAMSize(ramSize uint64) OptionSetter {
	return func(o *StoreOptions) {
		o.MaxRAMSize = ramSize
	}
}

func WithCleanupIntervalFrequency(freq time.Duration) OptionSetter {
	return func(o *StoreOptions) {
		o.CleanupInterval = freq
	}
}
