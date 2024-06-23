/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package models

//go:generate easyjson

//easyjson:json
type User struct {
	Name      string  `json:"name"`
	Skill     float64 `json:"skill"`
	Latency   float64 `json:"latency"`
	TimeSpent float64 `json:"-"`
}
