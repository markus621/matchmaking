/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package computation

import (
	"context"
	"math"
	"time"

	"matchmaking/pkg/config"
	"matchmaking/pkg/models"
	"matchmaking/pkg/queue"
	"matchmaking/pkg/stats"

	"github.com/prateek/knn-go"
	"go.osspkg.com/goppy/xc"
)

type Computation struct {
	queue   queue.IQueue
	stats   stats.IStats
	params  config.Params
	groupId uint64
}

func New(q queue.IQueue, s stats.IStats, c *config.Config) *Computation {
	return &Computation{
		queue:  q,
		stats:  s,
		params: c.Params,
	}
}

func (c *Computation) Up(ctx xc.Context) error {
	go c.handler(ctx.Context())
	return nil
}

func (c *Computation) Down() error {
	return nil
}

func (c *Computation) handler(ctx context.Context) {
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			for {
				vectors, ok := c.createVectors()
				if !ok {
					break
				}
				result, err := c.search(vectors, c.params.GroupSize)
				if err != nil {
					break
				}
				users, ok := c.getUsersByVectors(result)
				if !ok {
					break
				}

				c.groupId++

				timeNow := float64(time.Now().Unix())
				groupInfo := stats.GroupInfo{
					ID:        c.groupId,
					Names:     make([]string, 0, len(users)),
					Skill:     [3]float64{},
					Latency:   [3]float64{},
					TimeSpent: [3]float64{},
				}

				for i, user := range users {
					c.queue.Delete(user.Name)

					groupInfo.Names = append(groupInfo.Names, user.Name)
					user.TimeSpent = timeNow - user.TimeSpent

					if i == 0 {
						c.changeStat(&groupInfo, user, stats.AttrMin, sum)
						c.changeStat(&groupInfo, user, stats.AttrMax, sum)
						c.changeStat(&groupInfo, user, stats.AttrAvg, sum)
						continue
					}

					c.changeStat(&groupInfo, user, stats.AttrMin, math.Min)
					c.changeStat(&groupInfo, user, stats.AttrMax, math.Max)
					c.changeStat(&groupInfo, user, stats.AttrAvg, sum)
				}

				count := float64(c.params.GroupSize)
				groupInfo.Skill[stats.AttrAvg] /= count
				groupInfo.Latency[stats.AttrAvg] /= count
				groupInfo.TimeSpent[stats.AttrAvg] /= count

				c.stats.Receive(groupInfo)
			}
		}
	}
}

func (c *Computation) search(vectors []knn.Vector, count int) ([]knn.Vector, error) {
	obj, err := knn.NewNaiveKNN(vectors, knn.CosineDistance)
	if err != nil {
		return nil, err
	}
	inx := 0
	for i := 1; i < len(vectors); i++ {
		if vectors[i].Point[0] < vectors[inx].Point[0] {
			inx = i
		}
	}
	target := knn.Vector{ID: "target", Point: []float64{0, 0}}
	copy(target.Point, vectors[inx].Point)
	result := obj.Search(count, target)
	return result, nil
}

func (c *Computation) createVectors() ([]knn.Vector, bool) {
	if c.queue.Count() < c.params.GroupSize {
		return nil, false
	}

	vectors := make([]knn.Vector, 0, c.queue.Count())
	c.queue.Each(func(m models.User) {
		vectors = append(vectors, knn.Vector{
			ID:    knn.ID(m.Name),
			Point: []float64{m.Skill, m.Latency},
		})
	})
	return vectors, true
}

func (c *Computation) getUsersByVectors(v []knn.Vector) ([]models.User, bool) {
	users := make([]models.User, 0, len(v))
	for _, vector := range v {
		user, ok := c.queue.Get(string(vector.ID))
		if !ok {
			return nil, false
		}
		users = append(users, user)
	}
	return users, true
}

func (c *Computation) changeStat(group *stats.GroupInfo, user models.User, i int, call func(x, y float64) float64) {
	group.Skill[i] = call(user.Skill, group.Skill[i])
	group.Latency[i] = call(user.Latency, group.Latency[i])
	group.TimeSpent[i] = call(user.TimeSpent, group.TimeSpent[i])
}

func sum(x, y float64) float64 {
	return x + y
}
