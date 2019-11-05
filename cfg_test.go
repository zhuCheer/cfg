package cfg

import (
	"fmt"
	"testing"
	"time"
)

var path = "./example/config.toml"

func TestNew(t *testing.T) {
	fmt.Println("start testing")

	New(path)

	cfg := New(path)

	if cfg.confInfo == nil {
		t.Error("New func have an error #3")
	}
}

func TestGetString(t *testing.T) {
	cfg := New(path)
	key := cfg.GetString("owner.organization")

	if key != "GitHub" {
		t.Error("GetString func have an error #1")
	}

	key = cfg.GetString("test")
	if key != "" {
		t.Error("GetString func have an error #2")
	}
}

func TestGetInt(t *testing.T) {
	cfg := New(path)
	key := cfg.GetInt("owner.organization")

	if key != 0 {
		t.Error("GetInt func have an error #1")
	}

	key = cfg.GetInt("database.connection_max")

	if key != 5000 {
		t.Error("GetInt func have an error #2")
	}
}

func TestGetInt64(t *testing.T) {
	cfg := New(path)
	key := cfg.GetInt64("owner.organization")

	if key != 0 {
		t.Error("GetInt func have an error #1")
	}

	key = cfg.GetInt64("database.connection_max")

	if key != 5000 {
		t.Error("GetInt func have an error #2")
	}
}

func TestGetBool(t *testing.T) {
	cfg := New(path)
	key := cfg.GetBool("owner.organization")

	if key != false {
		t.Error("GetBool func have an error #1")
	}

	key = cfg.GetBool("database.enabled")
	if key != true {
		t.Error("GetBool func have an error #1")
	}
}

func TestGetDuration(t *testing.T) {
	cfg := New(path)
	key := cfg.GetDuration("owner.organization")

	if key != time.Duration(0) {
		t.Error("GetDuration func have an error #1")
	}

	key = cfg.GetDuration("database.timeout")
	if key != time.Duration(30) {
		t.Error("GetDuration func have an error #2")
	}
}

func TestGetSliceInt(t *testing.T) {
	cfg := New(path)
	value := cfg.GetSliceInt("database.ports")
	if len(value) != 3 {
		t.Error("GetSliceInt func have an error #1")
	}

	value = cfg.GetSliceInt("hosts")
	if len(value) != 0 {
		t.Error("GetSliceInt func have an error #2")
	}

	value = cfg.GetSliceInt("test")
	if len(value) != 0 {
		t.Error("GetSliceInt func have an error #3")
	}

}

func TestGetSliceInt64(t *testing.T) {
	cfg := New(path)
	value := cfg.GetSliceInt64("database.ports")

	if len(value) != 3 {
		t.Error("GetSliceInt64 func have an error #1")
	}

	value = cfg.GetSliceInt64("hosts")
	if len(value) != 0 {
		t.Error("GetSliceInt64 func have an error #2")
	}

	value = cfg.GetSliceInt64("test")
	if len(value) != 0 {
		t.Error("GetSliceInt64 func have an error #3")
	}

}

func TestGetSliceString(t *testing.T) {
	cfg := New(path)
	value := cfg.GetSliceString("database.ports")

	if len(value) != 0 {
		t.Error("GetSliceString func have an error #1")
	}

	value = cfg.GetSliceString("servers.beta.dc")
	if len(value) != 3 {
		t.Error("GetSliceString func have an error #2")
	}

	value = cfg.GetSliceString("database.connection_max")

	if len(value) != 0 {
		t.Error("GetSliceString func have an error #3")
	}
}

func TestExists(t *testing.T) {
	cfg := New(path)
	exists := cfg.Exists("database.servers")

	if exists != false {
		t.Error("Exists func have an error #1")
	}

	exists = cfg.Exists("database.server")
	if exists != true {
		t.Error("Exists func have an error #2")
	}

	exists = cfg.Exists("database.ports")
	if exists != true {
		t.Error("Exists func have an error #3")
	}

	exists = cfg.Exists("database.debug")
	if exists != true {
		t.Error("Exists func have an error #4")
	}

	exists = cfg.Exists("title")
	if exists != true {
		t.Error("Exists func have an error #5")
	}

	exists = cfg.Exists("titles")
	if exists != false {
		t.Error("Exists func have an error #6")
	}

	exists = cfg.Exists("clients.data")
	if exists != true {
		t.Error("Exists func have an error #7")
	}
}
