package main

import (
	"key-value-storage/internal/app/kvs"
	"reflect"
	"testing"
)

func Test_ShouldRollbackNestedTransactions(t *testing.T) {
	store := kvs.New()
	_, _ = handle(store, "SET", "foo", "123")
	_, _ = handle(store, "BEGIN", "", "")
	_, _ = handle(store, "SET", "foo", "456")
	_, _ = handle(store, "BEGIN", "", "")
	_, _ = handle(store, "SET", "foo", "789")

	result, err := handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "789")

	_, err = handle(store, "ROLLBACK", "", "")
	assertError(t, err, false)

	result, err = handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "456")

	_, err = handle(store, "ROLLBACK", "", "")
	assertError(t, err, false)

	result, err = handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "123")

	_, err = handle(store, "ROLLBACK", "", "")
	assertError(t, err, true)
}

func Test_ShouldRollbackATransaction(t *testing.T) {
	store := kvs.New()

	_, _ = handle(store, "SET", "foo", "123")
	_, _ = handle(store, "SET", "bar", "abc")
	_, _ = handle(store, "BEGIN", "", "")
	_, _ = handle(store, "SET", "foo", "456")

	result, err := handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "456")

	_, err = handle(store, "ROLLBACK", "", "")
	assertError(t, err, false)

	result, err = handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "123")

	result, err = handle(store, "GET", "bar", "")
	assertError(t, err, false)
	assertEqual(t, result, "abc")

	_, err = handle(store, "COMMIT", "", "")
	assertError(t, err, true)
}

func Test_ShouldDeleteAKey(t *testing.T) {
	store := kvs.New()
	_, _ = handle(store, "SET", "foo", "123")
	_, err := handle(store, "DELETE", "foo", "")
	assertError(t, err, false)
	result, err := handle(store, "GET", "foo", "")
	assertError(t, err, true)
	assertEqual(t, result, "")
}

func Test_ShouldCreateAKey(t *testing.T) {
	store := kvs.New()
	_, _ = handle(store, "SET", "foo", "123")
	result, err := handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "123")
}

func Test_ShouldCommitATransaction(t *testing.T) {
	store := kvs.New()
	_, _ = handle(store, "BEGIN", "", "")
	_, _ = handle(store, "SET", "foo", "456")
	_, err := handle(store, "COMMIT", "", "")
	assertError(t, err, false)

	result, err := handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "456")

	_, err = handle(store, "ROLLBACK", "", "")
	assertError(t, err, true)

	result, err = handle(store, "GET", "foo", "")
	assertError(t, err, false)
	assertEqual(t, result, "456")
}

func assertError(t *testing.T, err error, wantErr bool) {
	if (err != nil) != wantErr {
		t.Errorf("handle() error = %v, wantErr %v", err, wantErr)
	}
}

func assertEqual(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("handle() gotResult = %v, want %v", got, want)
	}
}
