package main

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/packages"
)

func main() {
    // Load the package from the given path
    cfg := &packages.Config{
        Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
    }
    pkgs, err := packages.Load(cfg, "gomes/pkg/model/isa95")
    if err != nil {
        fmt.Println("error:", err)
        return
    }

    if len(pkgs) == 0 {
        fmt.Println("no packages found")
        return
    }

    // Iterate through the package and find struct declarations
    for _, pkg := range pkgs {
		fmt.Println("Package:", pkg.Types.Name())
        for _, syntax := range pkg.Syntax {
            for _, decl := range syntax.Decls {
                genDecl, ok := decl.(*ast.GenDecl)
                if !ok || genDecl.Tok != token.TYPE {
                    continue
                }
                for _, spec := range genDecl.Specs {
                    typeSpec, ok := spec.(*ast.TypeSpec)
                    if !ok {
                        continue
                    }
                    _, ok = typeSpec.Type.(*ast.StructType)
                    if !ok {
                        continue
                    }
                    fmt.Println("Struct:", typeSpec.Name.Name)
                }
            }
        }
    }
}
