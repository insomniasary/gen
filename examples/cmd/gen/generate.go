package main

import (
	"github.com/insomniasary/gen"
	"github.com/insomniasary/gen/examples/conf"
	"github.com/insomniasary/gen/examples/dal"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN)

	prepare(dal.DB) // prepare table for generate
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
	})

	g.UseDB(dal.DB)

	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
