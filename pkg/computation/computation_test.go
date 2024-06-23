/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package computation

import (
	"testing"
	"time"

	"matchmaking/pkg/config"
	"matchmaking/pkg/models"
	"matchmaking/pkg/queue"
	"matchmaking/pkg/stats"

	"go.osspkg.com/goppy/xc"
	"go.osspkg.com/goppy/xtest"
)

type statMock struct {
	G []stats.GroupInfo
}

func (s *statMock) Receive(info stats.GroupInfo) {
	s.G = append(s.G, info)
}

func TestUnitComputation_handler(t *testing.T) {
	q := queue.New()
	c := &config.Config{Params: config.Params{GroupSize: 3}}
	s := &statMock{G: make([]stats.GroupInfo, 0, 3)}

	ctx := xc.New()

	q.Add(models.User{Name: "A1", Skill: 5, Latency: 1.0})
	q.Add(models.User{Name: "A2", Skill: 6, Latency: 1.1})
	q.Add(models.User{Name: "A3", Skill: 7, Latency: 1.2})
	q.Add(models.User{Name: "B1", Skill: 11, Latency: 1.3})
	q.Add(models.User{Name: "B2", Skill: 13, Latency: 1.4})
	q.Add(models.User{Name: "B3", Skill: 16, Latency: 1.5})
	q.Add(models.User{Name: "C1", Skill: 25, Latency: 1.6})
	q.Add(models.User{Name: "C2", Skill: 30, Latency: 1.7})

	app := New(q, s, c)
	xtest.NoError(t, app.Up(ctx), "run handler")

	<-time.After(3 * time.Second)
	ctx.Close()

	for _, info := range s.G {
		switch info.ID {
		case 1:
			xtest.Equal(t, []string{"A1", "A2", "A3"}, info.Names)
		case 2:
			xtest.Equal(t, []string{"B1", "B2", "B3"}, info.Names)
		default:
			t.FailNow()
		}
	}

}
