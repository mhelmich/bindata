package bindata

// Type -
type Type int

const (
	// TarBz -
	TarBz Type = iota
)

var (
	typeToTemplate = map[Type][]string{
		TarBz: {"templates/root.tmpl", "templates/tree.tmpl", "templates/tarbz.tmpl"},
	}
)
