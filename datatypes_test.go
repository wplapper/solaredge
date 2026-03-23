package solaredge

import (
	"encoding/json"
	"testing"
	"time"
)

func TestDate_JSON(t *testing.T) {
	type record struct {
		Date Date `json:"date"`
	}
	tests := []struct {
		name  string
		input string
		pass  bool
	}{
		{"valid", `{"date":"2020-02-01"}`, true},
		{"invalid", `{"date":"invalid-date"}`, false},
		{"empty", `{"date":""}`, true},
		{"null", `{"date":null}`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r record
			err := json.Unmarshal([]byte(tt.input), &r)
			if tt.pass != (err == nil) {
				t.Errorf("expected pass=%v, got err=%v", tt.pass, err)
			}
			if err != nil {
				return
			}
			if time.Time(r.Date).IsZero() {
				return
			}

			body, err := json.Marshal(r)
			if err != nil {
				t.Fatalf("marshal failed. err: %v", err)
			}
			if string(body) != tt.input {
				t.Errorf("Expected %v, got %v", tt.input, string(body))
			}
		})
	}
}

func TestTime_JSON(t *testing.T) {
	type record struct {
		Time Time `json:"time"`
	}
	tests := []struct {
		name  string
		input string
		pass  bool
	}{
		{"valid", `{"time":"2020-02-01 12:30:45"}`, true},
		{"invalid", `{"time":"invalid-date"}`, false},
		{"empty", `{"time":""}`, true},
		{"null", `{"time":null}`, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r record
			err := json.Unmarshal([]byte(tt.input), &r)
			if tt.pass != (err == nil) {
				t.Errorf("expected pass=%v, got err=%v", tt.pass, err)
			}
			if err != nil {
				return
			}
			if time.Time(r.Time).IsZero() {
				return
			}

			body, err := json.Marshal(r)
			if err != nil {
				t.Fatalf("marshal failed. err: %v", err)
			}
			if string(body) != tt.input {
				t.Errorf("Expected %v, got %v", tt.input, string(body))
			}
		})
	}
}
