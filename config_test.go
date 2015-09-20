package main

import (
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig("config.yaml")
	if err != nil {
		t.Errorf("Failed to load test configuration")
	}
	if c.TemplatePath != "templates" {
		t.Errorf("invalid template path")
	}
}

func TestInvalidConfig(t *testing.T) {
	_, err := loadConfig("invalid.yaml")
	if err == nil {
		t.Errorf("No error presented when invalid configuration is loaded")
	}
}

func TestGetPXEConfig(t *testing.T) {
	c, err := loadConfig("config.yaml")
	if err != nil {
		t.Errorf("Failed to load test configuration")
	}
	template, _ := c.getPXEConfig(Machine{Hostname: "testing"})
	if !strings.Contains(template, "hostname=testing") {
		t.Errorf("Template format failed")
	}
}
