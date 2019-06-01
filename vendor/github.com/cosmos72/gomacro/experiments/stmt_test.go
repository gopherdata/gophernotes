/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * stmt_0-3_test.go
 *
 *  Created on Apr 04, 2017
 *      Author Massimiliano Ghilardi
 */

package experiments

const (
	n int = 2000
)

func init() {
	println("n =", n)
}

/*
	benchmark results on Intel Core i7 4770 @3.2GHz, Debian 9, Linux 4.15.13 amd64, Go 1.10.1 linux/amd64

    -------- n = 10 --------
	BenchmarkStmt4-8           	100000000	        23.0 ns/op
	BenchmarkStmt4Unroll-8     	100000000	        23.7 ns/op
	BenchmarkStmt4Spin-8       	50000000	        27.4 ns/op
	BenchmarkStmt4Adaptive13-8	50000000	        26.0 ns/op
	BenchmarkStmt5-8           	50000000	        26.2 ns/op
	BenchmarkStmt6-8           	50000000	        27.6 ns/op
	BenchmarkStmt6Unroll-8     	50000000	        27.7 ns/op
	BenchmarkStmt6Spin-8       	50000000	        35.4 ns/op
	BenchmarkStmt6Adaptive13-8	50000000	        28.1 ns/op
	BenchmarkStmt7-8           	100000000	        23.1 ns/op
	BenchmarkStmt7Unroll-8     	50000000	        24.6 ns/op
	BenchmarkStmt7Spin-8       	50000000	        26.9 ns/op
	BenchmarkStmt7Adaptive13-8	100000000	        24.0 ns/op

    -------- n = 20 --------
	BenchmarkStmt4-8           	30000000	        48.2 ns/op
	BenchmarkStmt4Unroll-8     	30000000	        49.7 ns/op
	BenchmarkStmt4Spin-8       	30000000	        56.7 ns/op
	BenchmarkStmt4Adaptive13-8	30000000	        55.1 ns/op
	BenchmarkStmt5-8           	30000000	        51.4 ns/op
	BenchmarkStmt6-8           	30000000	        55.4 ns/op
	BenchmarkStmt6Unroll-8     	30000000	        55.9 ns/op
	BenchmarkStmt6Spin-8       	20000000	        72.0 ns/op
	BenchmarkStmt6Adaptive13-8	30000000	        55.9 ns/op
	BenchmarkStmt7-8           	30000000	        45.7 ns/op
	BenchmarkStmt7Unroll-8     	30000000	        45.8 ns/op
	BenchmarkStmt7Spin-8       	30000000	        57.1 ns/op
	BenchmarkStmt7Adaptive13-8	30000000	        46.6 ns/op

    -------- n = 50 --------
	BenchmarkStmt4-8           	10000000	       127 ns/op
	BenchmarkStmt4Unroll-8     	20000000	       115 ns/op
	BenchmarkStmt4Spin-8       	10000000	       135 ns/op
	BenchmarkStmt4Adaptive13-8	10000000	       125 ns/op
	BenchmarkStmt5-8           	10000000	       136 ns/op
	BenchmarkStmt6-8           	10000000	       142 ns/op
	BenchmarkStmt6Unroll-8     	10000000	       136 ns/op
	BenchmarkStmt6Spin-8       	10000000	       151 ns/op
	BenchmarkStmt6Adaptive13-8	10000000	       137 ns/op
	BenchmarkStmt7-8           	10000000	       133 ns/op
	BenchmarkStmt7Unroll-8     	20000000	       120 ns/op
	BenchmarkStmt7Spin-8       	10000000	       137 ns/op
	BenchmarkStmt7Adaptive13-8	20000000	       114 ns/op

    -------- n = 100 --------
	BenchmarkStmt4-8           	 5000000	       238 ns/op
	BenchmarkStmt4Unroll-8     	10000000	       236 ns/op
	BenchmarkStmt4Spin-8       	10000000	       233 ns/op
	BenchmarkStmt4Adaptive13-8	 5000000	       260 ns/op
	BenchmarkStmt5-8           	 5000000	       263 ns/op
	BenchmarkStmt6-8           	 5000000	       273 ns/op
	BenchmarkStmt6Unroll-8     	 5000000	       281 ns/op
	BenchmarkStmt6Spin-8       	 5000000	       289 ns/op
	BenchmarkStmt6Adaptive13-8	 5000000	       311 ns/op
	BenchmarkStmt7-8           	10000000	       237 ns/op
	BenchmarkStmt7Unroll-8     	10000000	       233 ns/op
	BenchmarkStmt7Spin-8       	10000000	       235 ns/op
	BenchmarkStmt7Adaptive13-8	 5000000	       269 ns/op

    -------- n = 200 --------
	BenchmarkStmt4-8           	 3000000	       479 ns/op
	BenchmarkStmt4Unroll-8     	 3000000	       483 ns/op
	BenchmarkStmt4Spin-8       	 3000000	       496 ns/op
	BenchmarkStmt4Adaptive13-8	 3000000	       480 ns/op
	BenchmarkStmt5-8           	 3000000	       520 ns/op
	BenchmarkStmt6-8           	 3000000	       558 ns/op
	BenchmarkStmt6Unroll-8     	 3000000	       563 ns/op
	BenchmarkStmt6Spin-8       	 3000000	       588 ns/op
	BenchmarkStmt6Adaptive13-8	 3000000	       585 ns/op
	BenchmarkStmt7-8           	 3000000	       449 ns/op
	BenchmarkStmt7Unroll-8     	 3000000	       455 ns/op
	BenchmarkStmt7Spin-8       	 3000000	       456 ns/op
	BenchmarkStmt7Adaptive13-8	 3000000	       471 ns/op

*/
