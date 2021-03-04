package cmd

import (
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/cmd/packages"
	"github.com/fdaines/spm-go/utils"
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
		var afferentMap = make(map[string][]string)
		pkgsInfo := getBasicPackagesInfo()
		utils.PrintMessage("Gathering package metrics, please wait until the command is finished running.")
		mainPackage := pkgsInfo[0].Path
		for index, pkgInfo := range pkgsInfo {
			pkg, err := context.Import(pkgInfo.Path, "", 0)
			if err == nil {
				pkgsInfo[index] = packages.FillFiles(pkgInfo, pkg)
				pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
				for _, current := range pkgsInfo[index].Dependencies.Internals {
					afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
				}
				abstractnessInfo, _ := retrieveAbstractnessInfo(pkg, mainPackage)
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
		printAll(pkgsInfo)
	})
}
