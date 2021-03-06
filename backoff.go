package restful

//Copyright 2016 MediaMath <http://www.mediamath.com>.  All rights reserved.
//Use of this source code is governed by a BSD-style
//license that can be found in the LICENSE file.

import "time"

//CreateBackOff is a function that returns a restful.BackOff
type CreateBackOff func() BackOff

//Notify corresponds directly with github.com/cenk/backoff/Notify
type Notify func(error, time.Duration)

//BackOff is modelled off of "github.com/cenk/backoff/BackOff"
type BackOff interface {
	Reset()
	Stop() (bool, time.Duration)
}

/*
	Wrapper around any of the backoff policies in github.com/cenk/backoff/BackOff
	func WrapBackoff(b backoff.BackOff) restful.BackOff {
		return &wrapper{b}

	}

	type wrapper struct {
		*backoff.BackOff
	}

	func (w *wrapper) Stop() (bool, time.Duration) {
		next := w.NextBackOff()

		return next == backoff.Stop, next
	}

*/
