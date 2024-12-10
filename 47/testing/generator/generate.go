package generator

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Module represents a single module configuration.
type Module struct {
	Name     string                      // Module name
	InitFunc func(*ServiceContext) error // Initialization function (optional)
	Path     string                      // Path for modules without InitFunc
}

// ServiceContext holds the instances of initialized modules.
type ServiceContext struct {
	Cfg  interface{} // Placeholder for config
	DB   interface{} // Placeholder for database
	Log  interface{} // Placeholder for logger
	Repo interface{} // Placeholder for repository
	// Add other fields as needed
}

// InfraGenerator handles the generation of infra files and initialization logic.
type InfraGenerator struct {
	Modules []Module
}

// NewInfraGenerator creates a new instance of InfraGenerator.
func NewInfraGenerator(modules []Module) *InfraGenerator {
	return &InfraGenerator{Modules: modules}
}

// RegisterModule adds a new module to the generator.
func (gen *InfraGenerator) RegisterModule(mod Module) {
	gen.Modules = append(gen.Modules, mod)
}

// BuildServiceContext initializes all registered modules.
// func (gen *InfraGenerator) BuildServiceContext() (*ServiceContext, error) {
// 	ctx := &ServiceContext{}

// 	for _, mod := range gen.Modules {
// 		if mod.InitFunc != nil {
// 			// Initialize using provided function
// 			if err := mod.InitFunc(ctx); err != nil {
// 				return nil, fmt.Errorf("failed to initialize module %s: %w", mod.Name, err)
// 			}
// 		} else if mod.Path != "" {
// 			// Placeholder for auto-generated interface
// 			fmt.Printf("Auto-generating interface for module: %s\n", mod.Name)
// 			// Replace with actual logic if necessary
// 		}
// 	}

// 	return ctx, nil
// }

// GenerateInfraFile creates the infra file with initialization logic.
func (gen *InfraGenerator) GenerateInfraFile(outputPath string) error {
	const infraTemplate = `package infra
	
import (
{{range .Imports}}
	"{{.}}"
{{end}}
)

type ServiceContext struct {
	{{range .Fields}}{{.}}
	{{end}}
}

func NewServiceContext() (*ServiceContext, error) {
	{{range .Initializations}}{{.}}
	{{end}}

	return &ServiceContext{
		{{range .ReturnFields}}{{.}},
		{{end}}
	}, nil
}
	`

	// Prepare template data
	imports := gen.collectImports()
	fields, initializations := gen.buildStructParts()

	data := struct {
		Imports         []string
		Fields          []string
		Initializations []string
	}{
		Imports:         imports,
		Fields:          fields,
		Initializations: initializations,
	}

	// Render template
	tmpl, err := template.New("infra").Parse(infraTemplate)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		return err
	}

	// Write the generated content to file
	if err := os.WriteFile(outputPath, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write infra file: %w", err)
	}

	return nil
}

func (gen *InfraGenerator) collectImports() []string {
	imports := map[string]bool{
		"go.uber.org/zap": true,
		"gorm.io/gorm":    true,
	}

	for _, mod := range gen.Modules {
		if mod.Path != "" {
			imports[mod.Path] = true
		}
	}

	result := []string{}
	for imp := range imports {
		result = append(result, imp)
	}
	return result
}

func (gen *InfraGenerator) buildStructParts() ([]string, []string) {
	var fields []string
	var initializations []string

	for _, mod := range gen.Modules {
		field := fmt.Sprintf("%s *%s", mod.Name, strings.Title(mod.Name))
		fields = append(fields, field)

		if mod.InitFunc != nil {
			init := fmt.Sprintf("// Initialize %s\nif err := %s(ctx); err != nil {\n\treturn nil, err\n}\n", mod.Name, mod.Name)
			initializations = append(initializations, init)
		} else if mod.Path != "" {
			init := fmt.Sprintf("// Auto-generate %s\nctx.%s = &%s{}", mod.Name, strings.Title(mod.Name), strings.Title(mod.Name))
			initializations = append(initializations, init)
		}
	}

	return fields, initializations
}

func (gen *InfraGenerator) SetupModules(initFunc map[string]interface{}, path map[string]string) {
	// Tambahkan modul berdasarkan fungsi inisialisasi
	for name, init := range initFunc {
		if fn, ok := init.(func(...interface{}) (interface{}, error)); ok {
			gen.RegisterModule(Module{
				Name: name,
				InitFunc: func(ctx *ServiceContext) error {
					// Panggil fungsi inisialisasi dengan parameter sesuai kebutuhan
					result, err := fn(ctx)
					if err != nil {
						return err
					}
					switch name {
					case "config":
						ctx.Cfg = result
					case "logger":
						ctx.Log = result
					case "database":
						ctx.DB = result
					default:
						return fmt.Errorf("unknown module: %s", name)
					}
					return nil
				},
			})
		}
	}

	// Tambahkan modul berdasarkan path
	for name, modulePath := range path {
		gen.RegisterModule(Module{
			Name: name,
			Path: modulePath,
		})
	}
}
