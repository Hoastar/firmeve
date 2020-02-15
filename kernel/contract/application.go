package contract

type (
	Application interface {
		Container
		SetMode(mode uint8)
		Mode() uint8
		IsDevelopment() bool
		IsProduction() bool
		IsTesting() bool
		Resolve(abstract interface{}, params ...interface{}) interface{}
		Boot()
		Register(provider Provider, force bool)
		RegisterMultiple(providers []Provider, force bool)
		HasProvider(name string) bool
		GetProvider(name string) Provider
		Reset()
	}
)