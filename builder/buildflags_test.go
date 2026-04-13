package builder

import (
	"testing"
)

func TestBuildFlagsValidate(t *testing.T) {
	tests := []struct {
		name    string
		flags   BuildFlags
		wantErr bool
	}{
		{
			name:    "empty flags are valid",
			flags:   BuildFlags{},
			wantErr: false,
		},
		{
			name:    "valid opt level z",
			flags:   BuildFlags{Opt: "z"},
			wantErr: false,
		},
		{
			name:    "valid opt level 2",
			flags:   BuildFlags{Opt: "2"},
			wantErr: false,
		},
		{
			name:    "invalid opt level",
			flags:   BuildFlags{Opt: "3"},
			wantErr: true,
		},
		{
			name:    "valid GC conservative",
			flags:   BuildFlags{GC: "conservative"},
			wantErr: false,
		},
		{
			name:    "invalid GC",
			flags:   BuildFlags{GC: "generational"},
			wantErr: true,
		},
		{
			name:    "valid scheduler tasks",
			flags:   BuildFlags{Scheduler: "tasks"},
			wantErr: false,
		},
		{
			name:    "invalid scheduler",
			flags:   BuildFlags{Scheduler: "goroutines"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.flags.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildFlagsTagsString(t *testing.T) {
	f := &BuildFlags{
		Tags: []string{"baremetal", "cortexm", "tinygo"},
	}
	got := f.TagsString()
	want := "baremetal cortexm tinygo"
	if got != want {
		t.Errorf("TagsString() = %q, want %q", got, want)
	}
}

func TestDefaultBuildFlags(t *testing.T) {
	f := DefaultBuildFlags()
	if f.Opt != "z" {
		t.Errorf("expected default Opt=z, got %q", f.Opt)
	}
	if f.GC != "conservative" {
		t.Errorf("expected default GC=conservative, got %q", f.GC)
	}
	if f.Scheduler != "tasks" {
		t.Errorf("expected default Scheduler=tasks, got %q", f.Scheduler)
	}
	if !f.Debug {
		t.Error("expected default Debug=true")
	}
	if err := f.Validate(); err != nil {
		t.Errorf("DefaultBuildFlags() produced invalid flags: %v", err)
	}
}
