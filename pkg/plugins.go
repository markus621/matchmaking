/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package pkg

import (
	"go.osspkg.com/goppy/plugins"

	"matchmaking/pkg/computation"
	"matchmaking/pkg/config"
	"matchmaking/pkg/queue"
	"matchmaking/pkg/stats"
)

var Plugins = plugins.Inject(
	queue.New,
	stats.New,
	computation.New,
	plugins.Plugin{
		Config: &config.Config{},
	},
)
