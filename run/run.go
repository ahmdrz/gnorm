package run // import "gnorm.org/gnorm/run"

// Config holds the schema that is expected to exist in the gnorm.toml file.
type Config struct {
	// ConnStr is the connection string for the database.
	ConnStr string

	// Schemas holds the names of schemas to generate code for.
	Schemas []string

	// TemplateDir contains the relative path to the directory where gnorm
	// expects to find templates to render.  The default is the current
	// directory where gnorm is running.
	TemplateDir string

	// TablePath is a relative path for tables to be rendered.  The table
	// template will be rendered with each table in turn. If the path is empty,
	// tables will not be rendered.
	//
	// The table path may be a template, in which case the values .Schema and
	// .Table may be referenced, containing the name of the current schema and
	// table being rendered.  For example, "{{.Schema}}/{{.Table}}/{{.Table}}.go" would render
	// the "public.users" table to ./public/users/users.go.
	TablePath string

	// SchemaPath is a relative path for schemas to be rendered.  The schema
	// template will be rendered with each schema in turn. If the path is empty,
	// schema will not be rendered.
	//
	// The schema path may be a template, in which case the value .Schema may be
	// referenced, containing the name of the current schema being rendered. For
	// example, "schemas/{{.Schema}}/{{.Schema}}.go" would render the "public"
	// schema to ./schemas/public/public.go
	SchemaPath string
}

func Run(cfg Config) error {
	return nil
}