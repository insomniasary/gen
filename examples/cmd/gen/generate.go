package main

import (
	"github.com/insomniasary/gen"
	"github.com/insomniasary/gen/examples/conf"
	"github.com/insomniasary/gen/examples/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "/Users/linyanteng/Code/gen/examples/cmd/gen/query",
		ModelPkgPath:"/Users/linyanteng/Code/gen/examples/cmd/gen/model",
		FieldNullable: true,
		FieldWithIndexTag:true,
		FieldWithTypeTag:  true,
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
