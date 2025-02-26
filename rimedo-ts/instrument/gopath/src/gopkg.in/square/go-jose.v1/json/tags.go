// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
package json

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
import (
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:5
)

import (
	"strings"
)

// tagOptions is the string following a comma in a struct field's "json"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:11
// tag, or the empty string. It does not include the leading comma.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:13
type tagOptions string

// parseTag splits a struct field's json tag into its name and
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:15
// comma-separated options.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:17
func parseTag(tag string) (string, tagOptions) {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:17
	_go_fuzz_dep_.CoverTab[185849]++
											if idx := strings.Index(tag, ","); idx != -1 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:18
		_go_fuzz_dep_.CoverTab[185851]++
												return tag[:idx], tagOptions(tag[idx+1:])
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:19
		// _ = "end of CoverTab[185851]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:20
		_go_fuzz_dep_.CoverTab[185852]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:20
		// _ = "end of CoverTab[185852]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:20
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:20
	// _ = "end of CoverTab[185849]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:20
	_go_fuzz_dep_.CoverTab[185850]++
											return tag, tagOptions("")
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:21
	// _ = "end of CoverTab[185850]"
}

// Contains reports whether a comma-separated list of options
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:24
// contains a particular substr flag. substr must be surrounded by a
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:24
// string boundary or commas.
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:27
func (o tagOptions) Contains(optionName string) bool {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:27
	_go_fuzz_dep_.CoverTab[185853]++
											if len(o) == 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:28
		_go_fuzz_dep_.CoverTab[185856]++
												return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:29
		// _ = "end of CoverTab[185856]"
	} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:30
		_go_fuzz_dep_.CoverTab[185857]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:30
		// _ = "end of CoverTab[185857]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:30
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:30
	// _ = "end of CoverTab[185853]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:30
	_go_fuzz_dep_.CoverTab[185854]++
											s := string(o)
											for s != "" {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:32
		_go_fuzz_dep_.CoverTab[185858]++
												var next string
												i := strings.Index(s, ",")
												if i >= 0 {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:35
			_go_fuzz_dep_.CoverTab[185861]++
													s, next = s[:i], s[i+1:]
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:36
			// _ = "end of CoverTab[185861]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:37
			_go_fuzz_dep_.CoverTab[185862]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:37
			// _ = "end of CoverTab[185862]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:37
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:37
		// _ = "end of CoverTab[185858]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:37
		_go_fuzz_dep_.CoverTab[185859]++
												if s == optionName {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:38
			_go_fuzz_dep_.CoverTab[185863]++
													return true
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:39
			// _ = "end of CoverTab[185863]"
		} else {
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:40
			_go_fuzz_dep_.CoverTab[185864]++
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:40
			// _ = "end of CoverTab[185864]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:40
		}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:40
		// _ = "end of CoverTab[185859]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:40
		_go_fuzz_dep_.CoverTab[185860]++
												s = next
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:41
		// _ = "end of CoverTab[185860]"
	}
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:42
	// _ = "end of CoverTab[185854]"
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:42
	_go_fuzz_dep_.CoverTab[185855]++
											return false
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:43
	// _ = "end of CoverTab[185855]"
}

//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:44
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/gopkg.in/square/go-jose.v1@v1.1.2/json/tags.go:44
var _ = _go_fuzz_dep_.CoverTab
