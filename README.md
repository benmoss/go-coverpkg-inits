# Weird behavior with go test

The `foo` package imports bar and `baz`, while the `qux` package only imports `bar`.

```
$ go test ./... -v -cover
bar 1
baz 2
foo 3
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
PASS
coverage: 100.0% of statements
ok      foo     0.001s  coverage: 100.0% of statements
?       foo/bar [no test files]
?       foo/baz [no test files]
bar 1
qux 2
=== RUN   TestQux
--- PASS: TestQux (0.00s)
PASS
coverage: 100.0% of statements
ok      foo/qux 0.001s  coverage: 100.0% of statements
```

If we change our command to use `-coverpkg=all` or even just
`-coverpkg=foo/qux`, we see that now the `qux` `init()` function gets run from
within the `foo_test` process.

```
$ go test ./... -v -count=1 -coverpkg all
bar 1
baz 2
foo 3
qux 4
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
PASS
coverage: 14.5% of statements in all
ok      foo     0.002s  coverage: 14.5% of statements in all
?       foo/bar [no test files]
?       foo/baz [no test files]
bar 1
qux 2
baz 3
foo 4
=== RUN   TestQux
--- PASS: TestQux (0.00s)
PASS
coverage: 14.5% of statements in all
ok      foo/qux 0.002s  coverage: 14.5% of statements in all
```

```
$ go test ./... -v -count=1 -coverpkg foo/qux
bar 1
baz 2
foo 3
qux 4
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
PASS
coverage: 66.7% of statements in foo/qux
ok      foo     0.001s  coverage: 66.7% of statements in foo/qux
?       foo/bar [no test files]
?       foo/baz [no test files]
bar 1
qux 2
=== RUN   TestQux
--- PASS: TestQux (0.00s)
PASS
coverage: 100.0% of statements in foo/qux
ok      foo/qux 0.001s  coverage: 100.0% of statements in foo/qux
```
