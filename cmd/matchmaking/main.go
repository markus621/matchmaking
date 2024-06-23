/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package main

import (
	matchmaking "matchmaking/app"
	"matchmaking/pkg"

	"go.osspkg.com/goppy"
	"go.osspkg.com/goppy/web"
)

var Version = "v0.0.0-dev"

func main() {
	app := goppy.New()
	app.AppName("matchmaking")
	app.AppVersion(Version)
	app.Plugins(
		web.WithServer(),
	)
	app.Plugins(matchmaking.Plugins...)
	app.Plugins(pkg.Plugins...)
	app.Run()
}
