package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	scas "github.com/qiangmzsx/string-adapter/v2"
)

func main() {

	modelText := `
		[request_definition]
		r = sub, obj, act

		[policy_definition]
		p = sub, obj, act

		[role_definition]
		g = _, _
		g2 = _, _

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act
	`

	m := model.NewModel()
	m.LoadModelFromText(modelText)
	line := `
		p, alice, data1, read
		p, bob, data2, write
		p, data_group_admin, data_group, write

		g, alice, data_group_admin
		g2, data1, data_group
		g2, data2, data_group
	`
	sa := scas.NewAdapter(line)
	e, _ := casbin.NewEnforcer(m, sa)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	if res, _ := e.Enforce("alice", "data1", "read"); res {
		fmt.Println("permitted")
	} else {
		fmt.Println("rejected")
	}

}
