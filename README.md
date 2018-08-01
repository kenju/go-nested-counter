# go-nested-counter

Abstraction of Nested Counter

# Install

```
go get -u github.com/kenju/go-nested-counter
```

# Usage

```go
counter := NewCounter()

counter.Increment("foo", 201807101200, 100)
counter.Increment("foo", 201807101200, 100)
counter.Increment("foo", 201807101200, 100)

counter.Increment("bar", 201807101200, 200)
counter.Increment("bar", 201807101200, 200)

if count := counter.GetCount("foo", 201807101200); count != 300 {
  t.Errorf("GetCount expects %d, got=%d", 300, count)
}

if count := counter.GetCount("bar", 201807101200); count != 400 {
  t.Errorf("GetCount expects %d, got=%d", 400, count)
}

// when key (identifier) does not exist
if count := counter.GetCount("not exist", 201807101200); count != 0 {
  t.Errorf("GetCount expects %d, got=%d", 0, count)
}

if size := counter.Size(); size != 2 {
  t.Errorf("Size expects %d, got=%d", 2, size)
}
```
