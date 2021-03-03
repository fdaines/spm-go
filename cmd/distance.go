package cmd

import (
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/cmd/distance"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/utils"
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
	utils.PrintMessage("Distance from main sequence analysis started.")
	var afferentMap = make(map[string][]string)
	pkgsInfo := getBasicPackagesInfo()
	for index, pkgInfo := range pkgsInfo {
		pkg, err := context.Import(pkgInfo.Path, "", 0)
		if err == nil {
			pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
			for _, current := range pkgsInfo[index].Dependencies.Internals {
				afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
			}
			abstractnessInfo, _ := retrieveAbstractnessInfo(pkg)
			pkgsInfo[index].AbstractnessDetails = abstractnessInfo
			pkgsInfo[index].AbstractionsCount = abstractnessInfo.StructsCount + abstractnessInfo.InterfacesCount
			pkgsInfo[index].ImplementationsCount = abstractnessInfo.MethodsCount + abstractnessInfo.FunctionsCount
			pkgsInfo[index].Abstractness = calculateAbstractness(pkgsInfo[index].AbstractionsCount, pkgsInfo[index].ImplementationsCount)
		}
	}
	for index, pkgInfo := range pkgsInfo {
		pkgsInfo[index].Dependants = afferentMap[pkgInfo.Path]
		pkgsInfo[index].AfferentCoupling = len(pkgsInfo[index].Dependants)
		pkgsInfo[index].EfferentCoupling = pkgsInfo[index].Dependencies.InternalsCount
		pkgsInfo[index].Instability = calculateInstability(pkgsInfo[index])
		pkgsInfo[index].Distance = calculateDistance(pkgsInfo[index].Instability, pkgsInfo[index].Abstractness)
	}
	distance.PrintPackages(pkgsInfo, common.OutputFormat)
}

func calculateDistance(instability float64, abstractness float64) float64 {
	return utils.RoundValue(math.Abs(instability + abstractness - 1))
}