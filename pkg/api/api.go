// Copyright 2019 bodhi369. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Echo Gorm Casbin restful API
//
// API Doc for echoAtom v1
//
// Terms Of Service:NA
//
//     Schemes: http, https
//     Host: localhost:8080
//     BasePath:
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: bodhi369 <bodhi369@example.com> http://bodhi369.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer:
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//          description: Bearer xxx
//
// swagger:meta
package api

import (
	"crypto/sha1"
	"fmt"

	casbinmw "github.com/bodhi369/echoatom/pkg/utl/middleware/casbin"

	"github.com/bodhi369/echoatom/pkg/utl/config"
	"github.com/bodhi369/echoatom/pkg/utl/gormplug"
	"github.com/bodhi369/echoatom/pkg/utl/middleware/jwt"
	"github.com/bodhi369/echoatom/pkg/utl/secure"
	"github.com/bodhi369/echoatom/pkg/utl/server"
	"github.com/bodhi369/echoatom/pkg/utl/zlog"
	"github.com/casbin/casbin"

	"github.com/bodhi369/echoatom/pkg/api/auth"
	al "github.com/bodhi369/echoatom/pkg/api/auth/logging"
	at "github.com/bodhi369/echoatom/pkg/api/auth/transport"
	"github.com/bodhi369/echoatom/pkg/api/menu"
	ml "github.com/bodhi369/echoatom/pkg/api/menu/logging"
	mt "github.com/bodhi369/echoatom/pkg/api/menu/transport"
	"github.com/bodhi369/echoatom/pkg/api/role"
	rl "github.com/bodhi369/echoatom/pkg/api/role/logging"
	rt "github.com/bodhi369/echoatom/pkg/api/role/transport"
	"github.com/bodhi369/echoatom/pkg/api/user"
	ul "github.com/bodhi369/echoatom/pkg/api/user/logging"
	ut "github.com/bodhi369/echoatom/pkg/api/user/transport"
)

// Start starts the API service
func Start(cfg *config.Configuration) error {

	//连接数据库
	db, err := gormplug.New(cfg)
	if err != nil {
		return err
	}

	//casbin 配置
	md := casbin.NewModel(cfg.Casbinmode.ModeText)
	ad := casbinmw.NewAdapter(db)
	ce := casbin.NewEnforcer(md, ad)

	rbac := casbinmw.New(ce)

	printPolicy(ce)

	sec := secure.New(cfg.App.MinPasswordStr, sha1.New())
	jwt := jwt.New(cfg.JWT.Secret, cfg.JWT.SigningAlgorithm, cfg.JWT.Duration)
	log := zlog.New()

	e := server.New()
	// swagger ui
	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	at.NewHTTP(al.New(auth.Initialize(db, ce, jwt, sec), log), e, jwt.MWFunc())

	v1 := e.Group("/v1")
	v1.Use(jwt.MWFunc())
	v1.Use(rbac.MWFunc())

	rt.NewHTTP(rl.New(role.Initialize(db, ce, sec), log), v1)
	ut.NewHTTP(ul.New(user.Initialize(db, ce, sec), log), v1)
	mt.NewHTTP(ml.New(menu.Initialize(db), log), v1)
	// pt.NewHTTP(pl.New(password.Initialize(db, rbac, sec), log), v1)

	// 输出路由
	// data, err := json.MarshalIndent(e.Routes(), "", "  ")
	// if err != nil {
	// 	return err
	// }
	// ioutil.WriteFile("routes.json", data, 0644)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func printPolicy(ce *casbin.Enforcer) {
	policy := ce.GetPolicy()
	fmt.Println("p:", policy)
	groupPolicy := ce.GetGroupingPolicy()
	fmt.Println("g:", groupPolicy)
	allRoles := ce.GetAllRoles()
	fmt.Println("roles name:", allRoles)
	allSubjects := ce.GetAllSubjects()
	fmt.Println("sub:", allSubjects)
	allObjects := ce.GetAllObjects()
	fmt.Println("obj:", allObjects)
	allActions := ce.GetAllActions()
	fmt.Println("act:", allActions)
}
