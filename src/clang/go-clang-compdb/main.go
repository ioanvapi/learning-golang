// go-clang-compdb dumps the content of a clang compilation database
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/go-clang/bootstrap/clang"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	CompileCommandDir string `long:"cmddir" default:"" description:"a directory containing a 'compile_commands.json' file"`
}

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.InfoLevel
	f := new(logrus.TextFormatter)
	f.TimestampFormat = "2006-01-02 15:04:05"
	f.FullTimestamp = true
	log.Formatter = f
}

func main() {
	p := flags.NewParser(&opts, flags.HelpFlag|flags.PrintErrors|flags.IgnoreUnknown)
	_, err := p.Parse()
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatalf("cli parse error: %v", err)
		} else {
			return
		}
	}

	dir := os.ExpandEnv(opts.CompileCommandDir)
	log.Printf(":: inspecting [%s]...\n", dir)

	fname := filepath.Join(dir, "compile_commands.json")
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("**error: could not open file [%s]: %v\n", fname, err)
	}
	f.Close()

	err, db := clang.FromDirectory(dir)
	if err != clang.CompilationDatabase_NoError {
		log.Fatalf("**error: could not open compilation database at [%s]: %v\n", dir, err)
	}
	defer db.Dispose()

	cmds := db.AllCompileCommands()
	ncmds := cmds.Size()

	log.Printf(":: got %d compile commands\n", ncmds)

	for i := uint32(0); i < ncmds; i++ {
		cmd := cmds.Command(i)

		log.Printf("::  --- cmd=%d ---\n", i)
		log.Printf("::  dir= %q\n", cmd.Directory())

		nargs := cmd.NumArgs()
		log.Printf("::  nargs= %d\n", nargs)

		sargs := make([]string, 0, nargs)
		for iarg := uint32(0); iarg < nargs; iarg++ {
			arg := cmd.Arg(iarg)
			sfmt := "%q, "
			if iarg+1 == nargs {
				sfmt = "%q"
			}
			sargs = append(sargs, fmt.Sprintf(sfmt, arg))

		}

		log.Printf("::  args= {%s}\n", strings.Join(sargs, ""))
		if i+1 != ncmds {
			log.Println("::")
		}
	}
	log.Printf(":: inspecting [%s]... [done]\n", dir)
}
