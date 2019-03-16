package generator

import (
	"errors"
	"os"
	"path"
	"path/filepath"

	"github.com/maximelamure/restli-go/common"
)

const (
	PDSC = iota
	IDL
)

type Generator interface {
	Generate(input string) error
}

type generatorFactory struct {
	generator map[int]Generator
}

func newGeneratorFactory() *generatorFactory {
	return &generatorFactory{generator: make(map[int]Generator)}
}

func (g *generatorFactory) Add(t int, gen Generator) {
	g.generator[t] = gen
}

func (g *generatorFactory) Genertor(t int) (Generator, error) {
	gen := g.generator[t]
	if gen == nil {
		return nil, errors.New("No generator supported for this format")
	}

	return gen, nil
}

// Generate will generate the client stub and the go structs based on the
// PDSC and IDL files.
func Generate(outputPath string, inputPath []string) error {
	gen := newGeneratorFactory()
	gen.Add(PDSC, NewSchemaGenerator(outputPath))
	gen.Add(IDL, NewResourceGenerator(outputPath))

	for _, root := range inputPath {
		err := filepath.Walk(root, visit(gen))
		if err != nil {
			return common.Wrap(err)
		}
	}

	return nil
}

func visit(gen *generatorFactory) filepath.WalkFunc {
	return func(input string, info os.FileInfo, err error) error {
		if err != nil {
			return common.Wrap(err)
		}

		if info.IsDir() {
			return nil
		}

		ext := path.Ext(info.Name())
		var g Generator
		if ext == SchemaExtension {
			g, err = gen.Genertor(PDSC)
		} else if ext == IDLExtension {
			g, err = gen.Genertor(IDL)
		} else {
			return nil // file extension not covered
		}

		if err != nil {
			return common.Wrap(err)
		}

		return g.Generate(input)
	}
}
