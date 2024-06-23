/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package app

import (
	"fmt"

	"matchmaking/pkg/models"
	"matchmaking/pkg/queue"

	"go.osspkg.com/goppy/web"
	"go.osspkg.com/goppy/xc"
)

type App struct {
	route web.Router
	queue queue.IQueue
}

func New(r web.RouterPool, q queue.IQueue) *App {
	return &App{
		route: r.Main(),
		queue: q,
	}
}

func (v *App) Up(xc.Context) error {
	v.route.Post("/users", v.AddUserToQueue)
	return nil
}

func (v *App) Down() error {
	return nil
}

func (v *App) AddUserToQueue(ctx web.Context) {
	var m models.User
	if err := ctx.BindJSON(&m); err != nil {
		ctx.Error(400, fmt.Errorf("decode request: %w", err))
		return
	}
	v.queue.Add(m)
	ctx.String(200, "ok")
}
