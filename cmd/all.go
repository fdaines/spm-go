package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/cmd/packages"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	"github.com/spf13/cobra"
)

var (
	allCmd = &cobra.Command{
		Use:   "all",
		Short: "Displays all metric for packages",
		Args:  ValidateArgs,
		Run:   listAllMetrics,
	}
)

func init() {
	rootCmd.AddCommand(allCmd)
}

func listAllMetrics(cmd *cobra.Command, args []string) {
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Full analysis started.")
		mainPackage, err := getMainPackage()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
		var afferentMap = make(map[string][]string)
		pkgsInfo := getBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				pkgsInfo[index] = packages.FillFiles(pkgInfo, pkg)
				pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
				for _, current := range pkgsInfo[index].Dependencies.Internals {
					afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
				}
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
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkgsInfo[index].Dependants = afferentMap[pkgInfo.Path]
			pkgsInfo[index].AfferentCoupling = len(pkgsInfo[index].Dependants)
			pkgsInfo[index].EfferentCoupling = pkgsInfo[index].Dependencies.InternalsCount
			pkgsInfo[index].Instability = calculateInstability(pkgsInfo[index])
			pkgsInfo[index].Distance = calculateDistance(pkgsInfo[index].Instability, pkgsInfo[index].Abstractness)
		}
		utils.PrintVerboseMessage("Done.")
		printAll(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage)
	})
}
