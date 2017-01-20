// go-clang-dump shows how to dump the AST of a C/C++ file via the Cursor
// visitor API.
//
// ex:
// $ go-clang-dump -fname=foo.cxx
package main

import (
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/go-clang/bootstrap/clang"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	FileName string `long:"fname" default:"" description:"the file to analyze"`
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
	args, err := p.Parse()
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Fatalf("cli parse error: %v", err)
		} else {
			return
		}
	}

	fmt.Println(":: go-clang-dump...")
	fmt.Println("opts:", opts)
	fmt.Println("args:", args)

	idx := clang.NewIndex(0, 1)
	defer idx.Dispose()

	tu := idx.ParseTranslationUnit(opts.FileName, args, nil, 0)
	defer tu.Dispose()

	fmt.Printf("tu: %s\n", tu.Spelling())

	diagnostics := tu.Diagnostics()
	for _, d := range diagnostics {
		fmt.Println("PROBLEM:", d.Spelling())
	}

	cursor := tu.TranslationUnitCursor()
	fmt.Printf("cursor-isnull: %v\n", cursor.IsNull())
	fmt.Printf("cursor: %s\n", cursor.Spelling())
	fmt.Printf("cursor-kind: %s\n", cursor.Kind().Spelling())

	fmt.Printf("tu-fname: %s\n", tu.File(opts.FileName).Name())

	cursor.Visit(func(cursor, parent clang.Cursor) clang.ChildVisitResult {
		if cursor.IsNull() {
			fmt.Println("cursor: <none>")

			return clang.ChildVisit_Continue
		}

		fmt.Printf("%s: %s (%s)\n", cursor.Kind().Spelling(), cursor.Spelling(), cursor.USR())

		switch cursor.Kind() {
		case clang.Cursor_ClassDecl, clang.Cursor_EnumDecl, clang.Cursor_StructDecl, clang.Cursor_Namespace:
			return clang.ChildVisit_Recurse
		}

		return clang.ChildVisit_Continue
	})

	if len(diagnostics) > 0 {
		fmt.Println("NOTE: There were problems while analyzing the given file")
	}

	fmt.Println(":: bye.")
}
