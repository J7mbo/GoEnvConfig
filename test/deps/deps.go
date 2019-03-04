package deps

/* These exist here because they should be in a separate package to show private properties being modified.  */

type Dep struct {
	i  int    `env:"i" default:"1337"`
	s  string `env:"s"`
	ip int    `env:"ip"`
}

func (d *Dep) GetInt() int       { return d.i }
func (d *Dep) GetString() string { return d.s }
func (d *Dep) GetPublicInt() int { return d.ip }
