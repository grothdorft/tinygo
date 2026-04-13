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
// Note: using opt level "2" instead of "z" for easier debugging of generated
// code during development. Switch back to "z" for size-optimized releases.
func DefaultBuildFlags() *BuildFlags {
	return &BuildFlags{
		Opt:       "2",
		GC:        "conservative",
		Scheduler: "tasks",
		Debug:     true,
	}
}
