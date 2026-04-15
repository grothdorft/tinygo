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
			// opt level 3 is not supported by tinygo (llvm backend limitation)
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
		{
			// none scheduler is useful for single-threaded bare-metal targets
			name:    "valid scheduler none",
			flags:   BuildFlags{Scheduler: "none"},
			wantErr: false,
		},
		{
			// asyncify scheduler is used for WebAssembly targets
			name:    "valid scheduler asyncify",
			flags:   BuildFlags{Scheduler: "asyncify"},
			wantErr: false,
		},
		{
			// opt level s optimizes for size, similar to z but slightly less aggressive
			name:    "valid opt level s",
			flags:   BuildFlags{Opt: "s"},
			wantErr: false,
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

// TestBuildFlagsTagsStringEmpty verifies that an empty tag list returns an
// empty string rather than a string with a trailing space.
func TestBuildFlagsTagsStringEmpty(t *testing.T) {
	f := &BuildFlags{}
	got := f.TagsString()
	if got != "" {
		t.Errorf("TagsString() with no tags = %q, want empty string", got)
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
