package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/verily-src/fsh-lint/internal/cli/diagnostic"
	"github.com/verily-src/fsh-lint/lint"
)

func main() {
	log.SetFlags(0) // ignore timestamp formatting

	installFlags(pflag.CommandLine)
	pflag.Parse()

	files, err := filesFromFlags(pflag.CommandLine)
	if err != nil {
		log.Fatal(err)
	}

	if Linter.Reporter == nil {
		reporter, err := diagnostic.ReporterFromFlags(pflag.CommandLine)
		if err != nil {
			log.Fatal(err)
		}
		Linter.Reporter = reporter
	}

	if Linter.Formatter == nil {
		Linter.Formatter = &lint.DefaultFormatter{}
	}

	Linter.Fix = pflag.CommandLine.Changed("fix")

	RunLinter(Linter, files)
}

// RunLinter runs the Linter against all files in the given paths. Exits with
// a non-zero exit code if any error-level problems are found.
func RunLinter(linter *lint.Linter, paths []string) {
	for _, path := range paths {
		linter.Lint(path)
	}

	// when the linter has error-level lint problems, exit with an error to indicate blocking
	if linter.HasErrors {
		os.Exit(1)
	}
}

// installFlags installs the flags --env, --paths, --fix, --output-format, and
// --debug to the given flag set.
func installFlags(fs *pflag.FlagSet) {
	// input flags
	fs.String("env", "", "Read new line delimited list of files or directories from the given environment variable.")
	fs.String("paths", "", "Read comma delimited list of files or directories. For file names with spaces, use quotes.")

	// configuration flags
	fs.Bool("fix", false, "Modify files and fix linting errors if possible.")

	// diagnostic flags
	diagnostic.MustInstallFlags(fs)
}

// filesFromFlags returns a list of files from the flags. If both --env and
// --paths are given, an error is returned.
func filesFromFlags(fs *pflag.FlagSet) ([]string, error) {
	envGiven := fs.Changed("env")
	pathsGiven := fs.Changed("paths")

	if envGiven && pathsGiven {
		return nil, fmt.Errorf("only one of --env or --paths must be used")
	}
	if !envGiven && !pathsGiven {
		return nil, fmt.Errorf("one of --env or --paths must be used")
	}

	var paths []string
	if envGiven {
		envName := fs.Lookup("env").Value.String()
		env := os.Getenv(envName)
		if env == "" {
			return nil, fmt.Errorf("environment variable %s is not set or empty", envName)
		}
		paths = pathsFromString(env, "\n")
	} else if pathsGiven {
		paths = pathsFromString(fs.Lookup("paths").Value.String(), ",")
	}

	return fshFilesFromPaths(paths), nil
}

// pathsFromString returns a list of paths from the given string. The string is
// split by the given delimiter. The paths are trimmed of whitespace and empty paths are ignored.
func pathsFromString(s string, delim string) []string {
	parts := strings.Split(s, delim)

	var paths []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}

		paths = append(paths, p)
	}

	return paths
}

// fshFilesFromPaths returns a list of files from the given paths. If a path is a directory,
// it will walk the directory and return all files with the .fsh extension. If a path is a file,
// it will return the file if it has the .fsh extension.
func fshFilesFromPaths(paths []string) []string {
	fileSet := make(map[string]struct{})
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			log.Fatalf("Error accessing path %s: %v", path, err)
			continue
		}

		if info.IsDir() {
			err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
				if err != nil {
					log.Fatalf("Error walking path %s: %v", p, err)
					return nil
				}
				if !d.IsDir() && filepath.Ext(d.Name()) == ".fsh" {
					fileSet[p] = struct{}{}
				}
				return nil
			})
			if err != nil {
				log.Fatalf("Error walking directory %s: %v", path, err)
			}
		} else if filepath.Ext(info.Name()) == ".fsh" {
			fileSet[path] = struct{}{}
		}
	}

	var files []string
	for file := range fileSet {
		files = append(files, file)
	}
	return files
}
