package cmd

import (
	"github.com/fdaines/spm-go/cmd/dependencies"
	"github.com/fdaines/spm-go/cmd/instability"
	"github.com/fdaines/spm-go/common"
	"github.com/fdaines/spm-go/model"
	"github.com/fdaines/spm-go/utils"
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
	utils.PrintMessage("Instability analysis started.")
	var afferentMap = make(map[string][]string)
	pkgsInfo := getBasicPackagesInfo()
	for index, pkgInfo := range pkgsInfo {
		pkg, err := context.Import(pkgInfo.Path, "", 0)
		if err == nil {
			pkgsInfo[index] = dependencies.FillDependencies(pkgsInfo[index], pkg, pkgsInfo)
			for _, current := range pkgsInfo[index].Dependencies.Internals {
				afferentMap[current] = append(afferentMap[pkgInfo.Path], current)
			}
		}
	}
	for index, pkgInfo := range pkgsInfo {
		pkgsInfo[index].Dependants = afferentMap[pkgInfo.Path]
		pkgsInfo[index].AfferentCoupling = len(pkgsInfo[index].Dependants)
		pkgsInfo[index].EfferentCoupling = pkgsInfo[index].Dependencies.InternalsCount
		pkgsInfo[index].Instability = calculateInstability(pkgsInfo[index])
	}
	instability.PrintPackages(pkgsInfo, common.OutputFormat)
}

func calculateInstability(pksInfo *model.PackageInfo) float64 {
	if pksInfo.EfferentCoupling == 0 && pksInfo.AfferentCoupling == 0 {
		return 1
	}
	return utils.RoundValue(
		float64(pksInfo.EfferentCoupling) / float64(pksInfo.EfferentCoupling + pksInfo.AfferentCoupling))
}