/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package stats

import "fmt"

type (
	Stats struct {
	}

	IStats interface {
		Receive(info GroupInfo)
	}
)

func New() IStats {
	return &Stats{}
}

func (s *Stats) Receive(v GroupInfo) {
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("GroupID:", v.ID, "Users:", v.Names)
	fmt.Println("STATS", "[min/max/avg]")
	fmt.Printf("- skill     : %.3f / %.3f / %.3f\n",
		v.Skill[AttrMin], v.Skill[AttrMax], v.Skill[AttrAvg])
	fmt.Printf("- latency   : %.3f / %.3f / %.3f\n",
		v.Latency[AttrMin], v.Latency[AttrMax], v.Latency[AttrAvg])
	fmt.Printf("- time spent: %.3f / %.3f / %.3f\n",
		v.TimeSpent[AttrMin], v.TimeSpent[AttrMax], v.TimeSpent[AttrAvg])
}
