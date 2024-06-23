/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package stats

const (
	AttrMin = 0
	AttrMax = 1
	AttrAvg = 2
)

type GroupInfo struct {
	ID        uint64
	Names     []string
	Skill     [3]float64
	Latency   [3]float64
	TimeSpent [3]float64
}
