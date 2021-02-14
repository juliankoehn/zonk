package tmpl

type (
	tmplParams struct {
		Encode  bool
		Package string
		Format  string
		Funcs   string
		Files   []*tmplFile
	}
	tmplFile struct {
		Base string
		Name string
		Path string
		Ext  string
		Data string
	}
)
