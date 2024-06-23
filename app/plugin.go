/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package app

import (
	"go.osspkg.com/goppy/plugins"
)

var Plugins = plugins.Inject(
	New,
)
