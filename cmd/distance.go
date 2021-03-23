package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	pkg "github.com/fdaines/spm-go/utils/packages"
	"github.com/spf13/cobra"
	"math"
)

var (
	distanceCmd = &cobra.Command{
		Use:   "distance",
		Short: "Analyzes distance from the main sequence",
		Args:  ValidateArgs,
		Run:   analyzeDistance,
	}
)

func init() {
	rootCmd.AddCommand(distanceCmd)
}

func analyzeDistance(cmd *cobra.Command, args []string) {
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Distance from main sequence analysis started.")
		mainPackage, err := pkg.GetMainPackage()
		if err != nil {
			fmt.Printf("Error: %+v\n", err)
			return
		}
		var afferentMap = make(map[string][]string)
		pkgsInfo := pkg.GetBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkgsInfo[index] = dependencies.FillDependencies(pkgInfo, pkgsInfo)
			for _, current := range pkgsInfo[index].Dependencies.Internals {
				afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
			}
			abstractnessInfo, err := retrieveAbstractnessInfo(pkgInfo.PackageData, mainPackage)
			if err != nil {
				fmt.Printf("Error: %+v\n", err)
				return
			}
			pkgsInfo[index].AbstractnessDetails = abstractnessInfo
			pkgsInfo[index].AbstractionsCount = abstractnessInfo.StructsCount + abstractnessInfo.InterfacesCount
			pkgsInfo[index].ImplementationsCount = abstractnessInfo.MethodsCount + abstractnessInfo.FunctionsCount
			pkgsInfo[index].Abstractness = calculateAbstractness(pkgsInfo[index].AbstractionsCount, pkgsInfo[index].ImplementationsCount)
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
		printDistance(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "distance")
	})
}

func calculateDistance(instability float64, abstractness float64) float64 {
	return utils.RoundValue(math.Abs(instability + abstractness - 1))
}
