# Weird behavior with go test

The foo package imports bar and baz, so it makes sense we see the log lines from
the init functions in those packages:

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
?       foo/qux [no test files]
```

If we want coverage of those packages, we can run using the `-coverpkg` flag and
specifying packages we want to also get coverage data for, despite them having
no tests of their own:

```
go test ./... -v -coverpkg foo,foo/bar,foo/baz
bar 1
baz 2
foo 3
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
PASS
coverage: 100.0% of statements in foo, foo/bar, foo/baz
ok      foo     0.001s  coverage: 100.0% of statements in foo, foo/bar, foo/baz
?       foo/bar [no test files]
?       foo/baz [no test files]
?       foo/qux [no test files]
```

When we add `foo/qux` to the list of files we generate coverage on via
the`-coverpkg` argument, we suddenly see that our foo_test.go is now running qux's
init, despite there being no imports to it!
```
$ go test ./... -v -coverpkg foo,foo/bar,foo/baz,foo/qux
bar 1
baz 2
foo 3
qux 4 <------- why do we see this??
=== RUN   TestFoo
--- PASS: TestFoo (0.00s)
PASS
coverage: 91.7% of statements in foo, foo/bar, foo/baz, foo/qux
ok      foo     0.001s  coverage: 91.7% of statements in foo, foo/bar, foo/baz, foo/qux
?       foo/bar [no test files]
?       foo/baz [no test files]
?       foo/qux [no test files]
```
