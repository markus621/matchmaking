/*
 *  Copyright (c) 2024 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a GPL-3.0 license that can be found in the LICENSE file.
 */

package queue

import (
	"time"

	"matchmaking/pkg/models"

	"go.osspkg.com/goppy/iosync"
)

type (
	Queue struct {
		list map[string]models.User
		mux  iosync.Lock
	}
	IQueue interface {
		Add(m models.User)
		Get(name string) (m models.User, ok bool)
		Delete(name string)
		Count() (c int)
		Each(call func(m models.User))
	}
)

func New() IQueue {
	return &Queue{
		list: make(map[string]models.User, 100),
		mux:  iosync.NewLock(),
	}
}

func (q *Queue) Add(m models.User) {
	m.TimeSpent = float64(time.Now().Unix())
	q.mux.Lock(func() {
		q.list[m.Name] = m
	})
}

func (q *Queue) Get(name string) (m models.User, ok bool) {
	q.mux.Lock(func() {
		m, ok = q.list[name]
	})
	return
}

func (q *Queue) Delete(name string) {
	q.mux.Lock(func() {
		delete(q.list, name)
	})
	return
}

func (q *Queue) Count() (c int) {
	q.mux.RLock(func() {
		c = len(q.list)
	})
	return
}

func (q *Queue) Each(call func(m models.User)) {
	q.mux.RLock(func() {
		for _, m := range q.list {
			call(m)
		}
	})
}
