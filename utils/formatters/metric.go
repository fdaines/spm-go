package formatters

import "github.com/fdaines/spm-go/model"

type getValue func(*model.PackageInfo) interface{}

type MetricOutput struct {
	Title string
	Value getValue
}

var FilesCount = MetricOutput{
	Title: "Files",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.FilesCount
	},
}
var InternalDependencies = MetricOutput{
	Title: "Internals",
	Value: func (pkg *model.PackageInfo) interface{} {
		if pkg.Dependencies == nil {
			return "-"
		}
		return pkg.Dependencies.InternalsCount
	},
}
var ExternalDependencies = MetricOutput{
	Title: "Externals",
	Value: func (pkg *model.PackageInfo) interface{} {
		if pkg.Dependencies == nil {
			return "-"
		}
		return pkg.Dependencies.ExternalsCount
	},
}
var StandardDependencies = MetricOutput{
	Title: "Standard",
	Value: func (pkg *model.PackageInfo) interface{} {
		if pkg.Dependencies == nil {
			return "-"
		}
		return pkg.Dependencies.StandardCount
	},
}
var TotalDependencies = MetricOutput{
	Title: "Total",
	Value: func (pkg *model.PackageInfo) interface{} {
		if pkg.Dependencies == nil {
			return "-"
		}
		return pkg.Dependencies.TotalCount
	},
}
var AfferentCoupling = MetricOutput{
	Title: "Afferent",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.AfferentCoupling
	},
}
var EfferentCoupling = MetricOutput{
	Title: "Efferent",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.EfferentCoupling
	},
}
var Instability = MetricOutput{
	Title: "Instability",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.Instability
	},
}
var Abstractions = MetricOutput{
	Title: "Abstractions",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.AbstractionsCount
	},
}
var Implementations = MetricOutput{
	Title: "Implementations",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.ImplementationsCount
	},
}
var Abstractness = MetricOutput{
	Title: "Abstractness",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.Abstractness
	},
}
var Distance = MetricOutput{
	Title: "Distance",
	Value: func (pkg *model.PackageInfo) interface{} {
		return pkg.Distance
	},
}

