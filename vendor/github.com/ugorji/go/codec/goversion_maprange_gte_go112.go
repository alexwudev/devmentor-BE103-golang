// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

//go:build go1.12 && (safe || codec.safe || appengine)
// +build go1.12
// +build safe codec.safe appengine

package codec

import "reflect"

type mapIter struct {
	t      *reflect.MapIter
	m      reflect.Value
	values bool
}

func (t *mapIter) Next() (r bool) {
	return t.t.Next()
}

func (t *mapIter) Key() reflect.Value {
	return t.t.Key()
}

func (t *mapIter) Value() (r reflect.Value) {
	if t.values {
		return t.t.Value()
	}
	return
}

func (t *mapIter) Done() {}

func mapRange(t *mapIter, m, k, v reflect.Value, values bool) {
	*t = mapIter{
		m:      m,
		t:      m.MapRange(),
		values: values,
	}
}
