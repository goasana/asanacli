// Copyright 2019 asana authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package generate

import (
	"fmt"
	"os"
	"path"

	asanaLogger "github.com/goasana/asana/logger"
	"github.com/goasana/asana/logger/colors"
	"github.com/goasana/asana/utils"
)

// recipe
// admin/recipe
func GenerateView(viewpath, currpath string) {
	w := colors.NewColorWriter(os.Stdout)

	asanaLogger.Log.Info("Generating view...")

	absViewPath := path.Join(currpath, "views", viewpath)
	err := os.MkdirAll(absViewPath, os.ModePerm)
	if err != nil {
		asanaLogger.Log.Fatalf("Could not create '%s' view: %s", viewpath, err)
	}

	cfile := path.Join(absViewPath, "index.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		_, _ = f.WriteString(cfile)
		_, _ = fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		asanaLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "show.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		_, _ = f.WriteString(cfile)
		_, _ = fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		asanaLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "create.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		_, _ = f.WriteString(cfile)
		_, _ = fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		asanaLogger.Log.Fatalf("Could not create view file: %s", err)
	}

	cfile = path.Join(absViewPath, "edit.tpl")
	if f, err := os.OpenFile(cfile, os.O_CREATE|os.O_EXCL|os.O_RDWR, 0666); err == nil {
		defer utils.CloseFile(f)
		_, _ = f.WriteString(cfile)
		_, _ = fmt.Fprintf(w, "\t%s%screate%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", cfile, "\x1b[0m")
	} else {
		asanaLogger.Log.Fatalf("Could not create view file: %s", err)
	}
}
