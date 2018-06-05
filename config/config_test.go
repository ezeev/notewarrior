package config

import "testing"

func TestConfigLoad(t *testing.T) {
	c, err := NewConfig("../config.yml")
	if err != nil {
		t.Error(err)
	}
	if c.MarkDownPath == "" {
		t.Error("mdpath is empty!")
	}
	t.Logf("md path: %s", c.MarkDownPath)
}
