
```text
$ buf generate && go test ./ -v
=== RUN   TestLocalJ5
=== RUN   TestLocalJ5/Value_by_itself
=== RUN   TestLocalJ5/Wrapped_in_order_1_2
=== RUN   TestLocalJ5/Wrapped_in_order_2_1
    test_test.go:75: Fresh Validation should have failed
--- FAIL: TestLocalJ5 (0.01s)
    --- PASS: TestLocalJ5/Value_by_itself (0.01s)
    --- PASS: TestLocalJ5/Wrapped_in_order_1_2 (0.00s)
    --- FAIL: TestLocalJ5/Wrapped_in_order_2_1 (0.00s)
FAIL
FAIL    github.com/daemonl/pvbug        0.276s
FAIL

```
