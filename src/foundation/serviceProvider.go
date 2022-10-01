package foundation

type ServiceProvider interface {
	Register()
	Boot()
	Init(*Application)
}
