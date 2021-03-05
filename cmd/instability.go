package cmd

import (
	"fmt"
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
	"github.com/fdaines/spm-go/utils/output"
	"github.com/spf13/cobra"
)

var (
	instabilityCmd = &cobra.Command{
		Use:   "instability",
		Short: "Analyzes instability of packages",
		Args:  ValidateArgs,
		Run:   analyzeInstability,
	}
)

func init() {
	rootCmd.AddCommand(instabilityCmd)
}

func analyzeInstability(cmd *cobra.Command, args []string) {
	utils.ExecuteWithTimer(func() {
		utils.PrintMessage("Instability analysis started.")
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
				pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
				for _, current := range pkgsInfo[index].Dependencies.Internals {
					afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
				}
			}
		}
		for index, pkgInfo := range pkgsInfo {
			utils.PrintStep()
			pkgsInfo[index].Dependants = afferentMap[pkgInfo.Path]
			pkgsInfo[index].AfferentCoupling = len(pkgsInfo[index].Dependants)
			pkgsInfo[index].EfferentCoupling = pkgsInfo[index].Dependencies.InternalsCount
			pkgsInfo[index].Instability = calculateInstability(pkgsInfo[index])
		}
		utils.PrintVerboseMessage("Done.")
		printInstability(pkgsInfo)
		output.GenerateHtmlOutput(pkgsInfo, mainPackage, "instability")
	})
}

func calculateInstability(pksInfo *model.PackageInfo) float64 {
	if pksInfo.EfferentCoupling == 0 && pksInfo.AfferentCoupling == 0 {
		return 1
	}
	return utils.RoundValue(
		float64(pksInfo.EfferentCoupling) / float64(pksInfo.EfferentCoupling+pksInfo.AfferentCoupling))
}
