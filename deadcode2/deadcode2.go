package deadcode2

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"path/filepath"

	"github.com/soider/d"
)

func Deadcode2(path string) ([]string, error) {

	fset := token.NewFileSet()
	pkgs, err := ParseDir(fset, path, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	d.D(pkgs)

	res := make([]string, 0, len(pkgs))
	for name, pkg := range pkgs {
		res = append(res, fmt.Sprintf("%v %v", name, pkg))
	}

	return res, nil
}

func ParseDir(fset *token.FileSet, path string, filter func(os.FileInfo) bool, mode parser.Mode) (pkgs map[string]*ast.Package, first error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	list, err := fd.Readdir(-1)
	if err != nil {
		return nil, err
	}

	pkgs = make(map[string]*ast.Package)
	for _, d := range list {
		if strings.HasSuffix(d.Name(), ".go") && (filter == nil || filter(d)) {
			filename := filepath.Join(path, d.Name())
			if src, err := parser.ParseFile(fset, filename, nil, mode); err == nil {
				name := src.Name.Name
				pkg, found := pkgs[name]
				if !found {
					pkg = &ast.Package{
						Name:  name,
						Files: make(map[string]*ast.File),
					}
					pkgs[name] = pkg
				}
				pkg.Files[filename] = src
			} else if first == nil {
				first = err
			}
		}
	}

	return
}
