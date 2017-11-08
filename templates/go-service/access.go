package {{.Name}}

// I{{.Name | Title}}Access represents an interface that needs to be satisfied
// for any data access used by the Service
type IAccess interface {

}

// {{.Name | Title}}Access represents a structure that implements data access
// for the {{.Name | Title}}
type Access struct {

}

// New{{.Name | Title}}Access creates a new access object and returns it
func NewAccess() *Access {
    return &Access{}
}