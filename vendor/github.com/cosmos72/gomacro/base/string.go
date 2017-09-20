/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Lesser General Public License as published
 *     by the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Lesser General Public License for more details.
 *
 *     You should have received a copy of the GNU Lesser General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/lgpl>.
 *
 *
 * string.go
 *
 *  Created on: Feb 13, 2017
 *      Author: Massimiliano Ghilardi
 */

package base

import (
	"strconv"
)

func UnescapeChar(str string) (rune, error) {
	// Debugf("unescapeChar(): parsing CHAR %#v", str)
	n := len(str)
	if n >= 2 && str[0] == '\'' && str[n-1] == '\'' {
		str = str[1 : n-1]
	}
	ret, _, _, err := strconv.UnquoteChar(str, '\'')
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func UnescapeString(str string) string {
	ret, err := strconv.Unquote(str)
	if err != nil {
		Error(err)
	}
	return ret
}

func FindFirstToken(src []byte) int {
	n := len(src)
	const (
		Normal = iota
		Slash
		LineComment
		MultiLineComment
		MultiLineCommentStar
	)
	mode := Normal
	for i := 0; i < n; i++ {
		ch := src[i]
		switch mode {
		case Normal:
			if ch == '/' {
				mode = Slash
			} else if ch > ' ' {
				return i
			}
		case Slash:
			if ch == '/' {
				mode = LineComment
			} else if ch == '*' {
				mode = MultiLineComment
			} else {
				return i - 1
			}
		case LineComment:
			if ch == '\n' {
				mode = Normal
			}
		case MultiLineComment:
			if ch == '*' {
				mode = MultiLineCommentStar
			}
		case MultiLineCommentStar:
			if ch == '/' {
				mode = Normal
			} else {
				mode = MultiLineComment
			}
		}
	}
	return n
}

/*
func extractFirstIdentifier(src []byte) []byte {
	n := len(src)
	for i := 0; i < n; i++ {
		ch := src[i]
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') ||
			ch == '_' || ch >= 128 ||
			(i != 0 && (ch >= '0' && ch <= '9')) {
		} else {
			return src[:i]
		}
	}
	return src
}
*/
