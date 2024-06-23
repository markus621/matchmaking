/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package config

type Config struct {
	Params Params `yaml:"match"`
}

type Params struct {
	GroupSize int `yaml:"group_size"`
}
