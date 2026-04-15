package builder

import (
	"fmt"
	"strings"
)

// BuildFlags holds the flags used during compilation.
type BuildFlags struct {
	// Target is the LLVM target triple (e.g. "thumbv6m-unknown-unknown-eabi").
	Target string

	// Opt is the optimization level: 0, 1, 2, s, or z.
	Opt string

	// GC is the garbage collector to use (e.g. "conservative", "leaking").
	GC string

	// Scheduler is the scheduler to use (e.g. "tasks", "asyncify", "none").
	Scheduler string

	// Tags are the build tags to use during compilation.
	Tags []string

	// Debug enables DWARF debug information.
	Debug bool

	// PrintIR prints the LLVM IR before and after optimization.
	PrintIR bool

	// VerifyIR runs the LLVM IR verifier after each pass.
	VerifyIR bool
}

// Validate checks that the build flags are valid.
func (f *BuildFlags) Validate() error {
	validOpts := map[string]bool{
		"0": true, "1": true, "2": true, "s": true, "z": true,
	}
	if f.Opt != "" && !validOpts[f.Opt] {
		return fmt.Errorf("invalid optimization level %q: must be one of 0, 1, 2, s, z", f.Opt)
	}

	validGCs := map[string]bool{
		"conservative": true, "leaking": true, "none": true, "custom": true,
	}
	if f.GC != "" && !validGCs[f.GC] {
		return fmt.Errorf("invalid GC %q: must be one of conservative, leaking, none, custom", f.GC)
	}

	validSchedulers := map[string]bool{
		"tasks": true, "asyncify": true, "none": true,
	}
	if f.Scheduler != "" && !validSchedulers[f.Scheduler] {
		return fmt.Errorf("invalid scheduler %q: must be one of tasks, asyncify, none", f.Scheduler)
	}

	return nil
}

// TagsString returns the build tags as a space-separated string.
func (f *BuildFlags) TagsString() string {
	return strings.Join(f.Tags, " ")
}

// DefaultBuildFlags returns a BuildFlags with sensible defaults.
// Using opt level "s" instead of "z" — "z" can sometimes produce slower code
// due to aggressive inlining suppression. "s" is a better balance between
// size and performance for my use cases (RP2040, ESP32-C3).
// Switch to "z" if flash space is critically tight.
//
// Disabling Debug by default to reduce binary size during normal iteration;
// pass -debug explicitly when I need to attach a debugger or inspect DWARF.
func DefaultBuildFlags() *BuildFlags {
	return &BuildFlags{
		Opt:       "s",
		GC:        "conservative",
		Scheduler: "tasks",
		Debug:     false,
	}
}
