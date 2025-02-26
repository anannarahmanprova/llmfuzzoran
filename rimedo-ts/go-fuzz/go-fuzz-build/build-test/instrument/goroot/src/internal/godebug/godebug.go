// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /usr/local/go/src/internal/godebug/godebug.go:23
package godebug

//line /usr/local/go/src/internal/godebug/godebug.go:23
import (
//line /usr/local/go/src/internal/godebug/godebug.go:23
	_go_fuzz_dep_ "go-fuzz-dep"
//line /usr/local/go/src/internal/godebug/godebug.go:23
)
//line /usr/local/go/src/internal/godebug/godebug.go:23
import (
//line /usr/local/go/src/internal/godebug/godebug.go:23
	_atomic_ "sync/atomic"
//line /usr/local/go/src/internal/godebug/godebug.go:23
)

import (
	"sync"
	"sync/atomic"
	_ "unsafe"
)

//line /usr/local/go/src/internal/godebug/godebug.go:32
type Setting struct {
	name	string
	once	sync.Once
	value	*atomic.Pointer[string]
}

//line /usr/local/go/src/internal/godebug/godebug.go:39
func New(name string) *Setting {
//line /usr/local/go/src/internal/godebug/godebug.go:39
	_go_fuzz_dep_.CoverTab[3459]++
								return &Setting{name: name}
//line /usr/local/go/src/internal/godebug/godebug.go:40
	// _ = "end of CoverTab[3459]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:44
func (s *Setting) Name() string {
//line /usr/local/go/src/internal/godebug/godebug.go:44
	_go_fuzz_dep_.CoverTab[3460]++
								return s.name
//line /usr/local/go/src/internal/godebug/godebug.go:45
	// _ = "end of CoverTab[3460]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:49
func (s *Setting) String() string {
//line /usr/local/go/src/internal/godebug/godebug.go:49
	_go_fuzz_dep_.CoverTab[3461]++
								return s.name + "=" + s.Value()
//line /usr/local/go/src/internal/godebug/godebug.go:50
	// _ = "end of CoverTab[3461]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:66
var cache sync.Map

var empty string

//line /usr/local/go/src/internal/godebug/godebug.go:77
func (s *Setting) Value() string {
//line /usr/local/go/src/internal/godebug/godebug.go:77
	_go_fuzz_dep_.CoverTab[3462]++
								s.once.Do(func() {
//line /usr/local/go/src/internal/godebug/godebug.go:78
		_go_fuzz_dep_.CoverTab[3464]++
									v, ok := cache.Load(s.name)
									if !ok {
//line /usr/local/go/src/internal/godebug/godebug.go:80
			_go_fuzz_dep_.CoverTab[3466]++
										p := new(atomic.Pointer[string])
										p.Store(&empty)
										v, _ = cache.LoadOrStore(s.name, p)
//line /usr/local/go/src/internal/godebug/godebug.go:83
			// _ = "end of CoverTab[3466]"
		} else {
//line /usr/local/go/src/internal/godebug/godebug.go:84
			_go_fuzz_dep_.CoverTab[3467]++
//line /usr/local/go/src/internal/godebug/godebug.go:84
			// _ = "end of CoverTab[3467]"
//line /usr/local/go/src/internal/godebug/godebug.go:84
		}
//line /usr/local/go/src/internal/godebug/godebug.go:84
		// _ = "end of CoverTab[3464]"
//line /usr/local/go/src/internal/godebug/godebug.go:84
		_go_fuzz_dep_.CoverTab[3465]++
									s.value = v.(*atomic.Pointer[string])
//line /usr/local/go/src/internal/godebug/godebug.go:85
		// _ = "end of CoverTab[3465]"
	})
//line /usr/local/go/src/internal/godebug/godebug.go:86
	// _ = "end of CoverTab[3462]"
//line /usr/local/go/src/internal/godebug/godebug.go:86
	_go_fuzz_dep_.CoverTab[3463]++
								return *s.value.Load()
//line /usr/local/go/src/internal/godebug/godebug.go:87
	// _ = "end of CoverTab[3463]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:97
//go:linkname setUpdate
func setUpdate(update func(string, string))

func init() {
	setUpdate(update)
}

var updateMu sync.Mutex

//line /usr/local/go/src/internal/godebug/godebug.go:109
func update(def, env string) {
//line /usr/local/go/src/internal/godebug/godebug.go:109
	_go_fuzz_dep_.CoverTab[3468]++
								updateMu.Lock()
								defer updateMu.Unlock()

//line /usr/local/go/src/internal/godebug/godebug.go:117
	did := make(map[string]bool)
								parse(did, env)
								parse(did, def)

//line /usr/local/go/src/internal/godebug/godebug.go:122
	cache.Range(func(name, v any) bool {
//line /usr/local/go/src/internal/godebug/godebug.go:122
		_go_fuzz_dep_.CoverTab[3469]++
									if !did[name.(string)] {
//line /usr/local/go/src/internal/godebug/godebug.go:123
			_go_fuzz_dep_.CoverTab[3471]++
										v.(*atomic.Pointer[string]).Store(&empty)
//line /usr/local/go/src/internal/godebug/godebug.go:124
			// _ = "end of CoverTab[3471]"
		} else {
//line /usr/local/go/src/internal/godebug/godebug.go:125
			_go_fuzz_dep_.CoverTab[3472]++
//line /usr/local/go/src/internal/godebug/godebug.go:125
			// _ = "end of CoverTab[3472]"
//line /usr/local/go/src/internal/godebug/godebug.go:125
		}
//line /usr/local/go/src/internal/godebug/godebug.go:125
		// _ = "end of CoverTab[3469]"
//line /usr/local/go/src/internal/godebug/godebug.go:125
		_go_fuzz_dep_.CoverTab[3470]++
									return true
//line /usr/local/go/src/internal/godebug/godebug.go:126
		// _ = "end of CoverTab[3470]"
	})
//line /usr/local/go/src/internal/godebug/godebug.go:127
	// _ = "end of CoverTab[3468]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:135
func parse(did map[string]bool, s string) {
//line /usr/local/go/src/internal/godebug/godebug.go:135
	_go_fuzz_dep_.CoverTab[3473]++

//line /usr/local/go/src/internal/godebug/godebug.go:141
	end := len(s)
	eq := -1
	for i := end - 1; i >= -1; i-- {
//line /usr/local/go/src/internal/godebug/godebug.go:143
		_go_fuzz_dep_.CoverTab[3474]++
									if i == -1 || func() bool {
//line /usr/local/go/src/internal/godebug/godebug.go:144
			_go_fuzz_dep_.CoverTab[3475]++
//line /usr/local/go/src/internal/godebug/godebug.go:144
			return s[i] == ','
//line /usr/local/go/src/internal/godebug/godebug.go:144
			// _ = "end of CoverTab[3475]"
//line /usr/local/go/src/internal/godebug/godebug.go:144
		}() {
//line /usr/local/go/src/internal/godebug/godebug.go:144
			_go_fuzz_dep_.CoverTab[3476]++
										if eq >= 0 {
//line /usr/local/go/src/internal/godebug/godebug.go:145
				_go_fuzz_dep_.CoverTab[3478]++
											name, value := s[i+1:eq], s[eq+1:end]
											if !did[name] {
//line /usr/local/go/src/internal/godebug/godebug.go:147
					_go_fuzz_dep_.CoverTab[3479]++
												did[name] = true
												v, ok := cache.Load(name)
												if !ok {
//line /usr/local/go/src/internal/godebug/godebug.go:150
						_go_fuzz_dep_.CoverTab[3481]++
													p := new(atomic.Pointer[string])
													p.Store(&empty)
													v, _ = cache.LoadOrStore(name, p)
//line /usr/local/go/src/internal/godebug/godebug.go:153
						// _ = "end of CoverTab[3481]"
					} else {
//line /usr/local/go/src/internal/godebug/godebug.go:154
						_go_fuzz_dep_.CoverTab[3482]++
//line /usr/local/go/src/internal/godebug/godebug.go:154
						// _ = "end of CoverTab[3482]"
//line /usr/local/go/src/internal/godebug/godebug.go:154
					}
//line /usr/local/go/src/internal/godebug/godebug.go:154
					// _ = "end of CoverTab[3479]"
//line /usr/local/go/src/internal/godebug/godebug.go:154
					_go_fuzz_dep_.CoverTab[3480]++
												v.(*atomic.Pointer[string]).Store(&value)
//line /usr/local/go/src/internal/godebug/godebug.go:155
					// _ = "end of CoverTab[3480]"
				} else {
//line /usr/local/go/src/internal/godebug/godebug.go:156
					_go_fuzz_dep_.CoverTab[3483]++
//line /usr/local/go/src/internal/godebug/godebug.go:156
					// _ = "end of CoverTab[3483]"
//line /usr/local/go/src/internal/godebug/godebug.go:156
				}
//line /usr/local/go/src/internal/godebug/godebug.go:156
				// _ = "end of CoverTab[3478]"
			} else {
//line /usr/local/go/src/internal/godebug/godebug.go:157
				_go_fuzz_dep_.CoverTab[3484]++
//line /usr/local/go/src/internal/godebug/godebug.go:157
				// _ = "end of CoverTab[3484]"
//line /usr/local/go/src/internal/godebug/godebug.go:157
			}
//line /usr/local/go/src/internal/godebug/godebug.go:157
			// _ = "end of CoverTab[3476]"
//line /usr/local/go/src/internal/godebug/godebug.go:157
			_go_fuzz_dep_.CoverTab[3477]++
										eq = -1
										end = i
//line /usr/local/go/src/internal/godebug/godebug.go:159
			// _ = "end of CoverTab[3477]"
		} else {
//line /usr/local/go/src/internal/godebug/godebug.go:160
			_go_fuzz_dep_.CoverTab[3485]++
//line /usr/local/go/src/internal/godebug/godebug.go:160
			if s[i] == '=' {
//line /usr/local/go/src/internal/godebug/godebug.go:160
				_go_fuzz_dep_.CoverTab[3486]++
											eq = i
//line /usr/local/go/src/internal/godebug/godebug.go:161
				// _ = "end of CoverTab[3486]"
			} else {
//line /usr/local/go/src/internal/godebug/godebug.go:162
				_go_fuzz_dep_.CoverTab[3487]++
//line /usr/local/go/src/internal/godebug/godebug.go:162
				// _ = "end of CoverTab[3487]"
//line /usr/local/go/src/internal/godebug/godebug.go:162
			}
//line /usr/local/go/src/internal/godebug/godebug.go:162
			// _ = "end of CoverTab[3485]"
//line /usr/local/go/src/internal/godebug/godebug.go:162
		}
//line /usr/local/go/src/internal/godebug/godebug.go:162
		// _ = "end of CoverTab[3474]"
	}
//line /usr/local/go/src/internal/godebug/godebug.go:163
	// _ = "end of CoverTab[3473]"
}

//line /usr/local/go/src/internal/godebug/godebug.go:164
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /usr/local/go/src/internal/godebug/godebug.go:164
var _ = _go_fuzz_dep_.CoverTab
