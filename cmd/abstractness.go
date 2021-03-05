package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	"github.com/spf13/cobra"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
		mainPackage, err := getMainPackage()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
		pkgsInfo := getBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				abstractnessInfo, err := retrieveAbstractnessInfo(pkg, mainPackage)
				if err != nil {
					fmt.Printf("Error: %+v\n", err)
					return
				}
				pkgsInfo[index].AbstractnessDetails = abstractnessInfo
				pkgsInfo[index].AbstractionsCount = abstractnessInfo.StructsCount + abstractnessInfo.InterfacesCount
				pkgsInfo[index].ImplementationsCount = abstractnessInfo.MethodsCount + abstractnessInfo.FunctionsCount
				pkgsInfo[index].Abstractness = calculateAbstractness(pkgsInfo[index].AbstractionsCount, pkgsInfo[index].ImplementationsCount)
			}
		}
		utils.PrintVerboseMessage("Done.")
		printAbstractness(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "abstractness")
	})
}

func retrieveAbstractnessInfo(pkg *build.Package, mainPackage string) (*model.AbstractnessDetails, error) {
	var methods, functions, interfaces, structs int
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	packageDir := strings.Replace(pkg.ImportPath, mainPackage, path, 1)

	for _, srcFile := range pkg.GoFiles {
		data, err := ioutil.ReadFile(filepath.Join(packageDir, srcFile))
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
	return utils.RoundValue(float64(abstractions) / float64(total))
}
