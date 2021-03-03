package cmd

import (
	"github.com/fdaines/spm-go/cmd/abstractness"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/spf13/cobra"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
)

var (
	abstractnessCmd = &cobra.Command{
		Use:   "abstractness",
		Short: "Analyzes abstractness of packages",
		Args:  ValidateArgs,
		Run:   analyzeAbstractness,
	}
)

func init() {
	rootCmd.AddCommand(abstractnessCmd)
}

func analyzeAbstractness(cmd *cobra.Command, args []string) {
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Abstractness analysis started.")
		pkgsInfo := getBasicPackagesInfo()
		for index, pkgInfo := range pkgsInfo {
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				abstractnessInfo, _ := retrieveAbstractnessInfo(pkg)
				pkgsInfo[index].AbstractnessDetails = abstractnessInfo
				pkgsInfo[index].AbstractionsCount = abstractnessInfo.StructsCount + abstractnessInfo.InterfacesCount
				pkgsInfo[index].ImplementationsCount = abstractnessInfo.MethodsCount + abstractnessInfo.FunctionsCount
				pkgsInfo[index].Abstractness = calculateAbstractness(pkgsInfo[index].AbstractionsCount, pkgsInfo[index].ImplementationsCount)
			}
		}
		abstractness.PrintPackages(pkgsInfo, common.OutputFormat)
	})
}

func retrieveAbstractnessInfo(pkg *build.Package) (*model.AbstractnessDetails, error) {
	var methods, functions, interfaces, structs int

	for _, srcFile := range pkg.GoFiles {
		data, err := ioutil.ReadFile(filepath.Join(pkg.SrcRoot, pkg.ImportPath, srcFile))
		if err != nil {
			return nil, err
		}
		fileset := token.NewFileSet()
		node, err := parser.ParseFile(fileset, srcFile, data, 0)
		if err != nil {
			return nil, err
		}
		ast.Inspect(node, func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.FuncDecl:
				if t.Recv != nil {
					methods++
				} else {
					functions++
				}
			case *ast.InterfaceType:
				interfaces++
			case *ast.StructType:
				structs++
			}
			return true
		})
	}

	return &model.AbstractnessDetails{
		MethodsCount:    methods,
		FunctionsCount:  functions,
		InterfacesCount: interfaces,
		StructsCount:    structs,
	}, nil
}

func calculateAbstractness(abstractions int, implementations int) float64 {
	total := implementations + abstractions
	if total == 0 {
		total = 1
	}
	return utils.RoundValue(float64(abstractions)/float64(total))
}