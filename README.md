
AsanaCLI
===

asana is a command-line tool facilitating development of Asana-based application.

[![Build Status](https://img.shields.io/travis/asana/asanacli.svg?branch=master&label=master)](https://travis-ci.org/goasana/asanacli)
[![Build Status](https://img.shields.io/travis/asana/asanacli.svg?branch=develop&label=develop)](https://travis-ci.org/goasana/asanacli)

## Requirements

- Go version >= 1.3.

## Installation

To install `asanacli` use the `go get` command:

```bash
go get github.com/goasana/asanacli
```

Then you can add `asanacli` binary to PATH environment variable in your `~/.bashrc` or `~/.bash_profile` file:

```bash
export PATH=$PATH:<your_main_gopath>/bin
```

> If you already have `asana` installed, updating `asana` is simple:

```bash
go get -u github.com/goasana/asanacli
```

## Basic commands

asanacli provides a variety of commands which can be helpful at various stages of development. The top level commands include:

```
    version     Prints the current Asana version
    migrate     Runs database migrations
    api         Creates a Asana API application
    bale        Transforms non-Go files to Go source files
    fix         Fixes your application by making it compatible with newer versions of Asana
    dlv         Start a debugging session using Delve
    dockerize   Generates a Dockerfile for your Asana application
    generate    Source code generator
    hprose      Creates an RPC application based on Hprose and Asana frameworks
    new         Creates a Asana application
    pack        Compresses a Asana application into a single file
    rs          Run customized scripts
    run         Run the application by starting a local development server

```

### asana version

To display the current version of `asana`, `asanacli` and `go` installed on your machine:

```bash
$ asanacli version
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0

├── Asana     : 1.0
├── GoVersion : go1.7.4
├── GOOS      : linux
├── GOARCH    : amd64
├── NumCPU    : 2
├── GOPATH    : /home/asanauser/.go
├── GOROOT    : /usr/lib/go
├── Compiler  : gc
└── Date      : Monday, 28 Apr 2019
```

You can also change the output format using `-o` flag:

```bash
$ asanacli version -o json
{
    "GoVersion": "go1.7.4",
    "GOOS": "linux",
    "GOARCH": "amd64",
    "NumCPU": 2,
    "GOPATH": "/home/asanauser/.go",
    "GOROOT": "/usr/lib/go",
    "Compiler": "gc",
    "AsanaCLIVersion": "1.0",
    "AsanaVersion": "1.0"
}
```

For more information on the usage, run `asanacli help version`.

### asanacli new

To create a new Asana web application:

```bash
$ asanacli new my-web-app
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:28:11 INFO     ▶ 0001 Creating application...
    create   /home/asanauser/.go/src/github.com/user/my-web-app/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/conf/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/controllers/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/models/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/routers/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/tests/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/static/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/static/js/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/static/css/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/static/img/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/views/
    create   /home/asanauser/.go/src/github.com/user/my-web-app/conf/app.yaml
    create   /home/asanauser/.go/src/github.com/user/my-web-app/controllers/default.go
    create   /home/asanauser/.go/src/github.com/user/my-web-app/views/index.tpl
    create   /home/asanauser/.go/src/github.com/user/my-web-app/routers/router.go
    create   /home/asanauser/.go/src/github.com/user/my-web-app/tests/default_test.go
    create   /home/asanauser/.go/src/github.com/user/my-web-app/main.go
2016/12/26 22:28:11 SUCCESS  ▶ 0002 New application successfully created!
```

For more information on the usage, run `asanacli help new`.

### asanacli run

To run the application we just created, you can navigate to the application folder and execute:

```bash
$ cd my-web-app && asana run
```

Or from anywhere in your machine:

```
$ asanacli run github.com/user/my-web-app
```

For more information on the usage, run `asanacli help run`.

### asanacli pack

To compress a Asana application into a single deployable file:

```bash
$ asanacli pack
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:29:29 INFO     ▶ 0001 Packaging application on '/home/asanauser/.go/src/github.com/user/my-web-app'...
2016/12/26 22:29:29 INFO     ▶ 0002 Building application...
2016/12/26 22:29:29 INFO     ▶ 0003 Using: GOOS=linux GOARCH=amd64
2016/12/26 22:29:31 SUCCESS  ▶ 0004 Build Successful!
2016/12/26 22:29:31 INFO     ▶ 0005 Writing to output: /home/asanauser/.go/src/github.com/user/my-web-app/my-web-app.tar.gz
2016/12/26 22:29:31 INFO     ▶ 0006 Excluding relpath prefix: .
2016/12/26 22:29:31 INFO     ▶ 0007 Excluding relpath suffix: .go:.DS_Store:.tmp
2016/12/26 22:29:32 SUCCESS  ▶ 0008 Application packed!
```

For more information on the usage, run `asana help pack`.

### asanacli rs 
Inspired by makefile / npm scripts.
  Run script allows you to run arbitrary commands using asanacli.
  Custom commands are provided from the "scripts" object inside asana.json or Asanafile.

  To run a custom command, use: $ asanacli rs mycmd ARGS

```bash
$ asanacli help rs

USAGE
  asanacli rs

DESCRIPTION
  Run script allows you to run arbitrary commands using asanacli.
  Custom commands are provided from the "scripts" object inside asana.json or Asanafile.

  To run a custom command, use: $ asanacli rs mycmd ARGS
  
AVAILABLE SCRIPTS
  gtest
      APP_ENV=test APP_CONF_PATH=$(pwd)/conf go test -v -cover
  gtestall
      APP_ENV=test APP_CONF_PATH=$(pwd)/conf go test -v -cover $(go list ./... | grep -v /vendor/)

```

*Run your scripts with:*
```$ asanacli rs gtest tests/*.go```
```$ asanacli rs gtestall```


### asanacli api

To create a Asana API application:

```bash
$ asanacli api my-api
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:30:12 INFO     ▶ 0001 Creating API...
    create   /home/asanauser/.go/src/github.com/user/my-api
    create   /home/asanauser/.go/src/github.com/user/my-api/conf
    create   /home/asanauser/.go/src/github.com/user/my-api/controllers
    create   /home/asanauser/.go/src/github.com/user/my-api/tests
    create   /home/asanauser/.go/src/github.com/user/my-api/conf/app.yaml
    create   /home/asanauser/.go/src/github.com/user/my-api/models
    create   /home/asanauser/.go/src/github.com/user/my-api/routers/
    create   /home/asanauser/.go/src/github.com/user/my-api/controllers/object.go
    create   /home/asanauser/.go/src/github.com/user/my-api/controllers/user.go
    create   /home/asanauser/.go/src/github.com/user/my-api/tests/default_test.go
    create   /home/asanauser/.go/src/github.com/user/my-api/routers/router.go
    create   /home/asanauser/.go/src/github.com/user/my-api/models/object.go
    create   /home/asanauser/.go/src/github.com/user/my-api/models/user.go
    create   /home/asanauser/.go/src/github.com/user/my-api/main.go
2016/12/26 22:30:12 SUCCESS  ▶ 0002 New API successfully created!
```

For more information on the usage, run `asana help api`.

### asanacli hprose

To create an Hprose RPC application based on Asana:

```bash
$ asanacli hprose my-rpc-app
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:30:58 INFO     ▶ 0001 Creating application...
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/conf/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/controllers/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/models/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/routers/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/tests/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/static/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/static/js/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/static/css/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/static/img/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/views/
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/conf/app.yamlº
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/controllers/default.go
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/views/index.tpl
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/routers/router.go
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/tests/default_test.go
    create   /home/asanauser/.go/src/github.com/user/my-rpc-app/main.go
2016/12/26 22:30:58 SUCCESS  ▶ 0002 New application successfully created!
```

For more information on the usage, run `asana help hprose`.

### asanacli bale

To pack all the static files into Go source files:

```bash
$ asanacli bale
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:32:41 INFO     ▶ 0001 Loading configuration from asana.json...
2016/12/26 22:32:41 SUCCESS  ▶ 0002 Baled resources successfully!
```

For more information on the usage, run `asana help bale`.

### asanacli migrate

For database migrations, use `asana migrate`.

For more information on the usage, run `asana help migrate`.

### asanacli generate

asanacli also comes with a source code generator which speeds up the development.

For example, to generate a new controller named `hello`:

```bash
$ asanacli generate controller hello
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:33:58 INFO     ▶ 0001 Using 'Hello' as controller name
2016/12/26 22:33:58 INFO     ▶ 0002 Using 'controllers' as package name
    create   /home/asanauser/.go/src/github.com/user/my-web-app/controllers/hello.go
2016/12/26 22:33:58 SUCCESS  ▶ 0003 Controller successfully generated!
```

For more information on the usage, run `asana help generate`.

### asanacli dockerize

asanacli also helps you dockerize your Asana application by generating a Dockerfile.

For example, to generate a Dockerfile with `Go version 1.6.4` and exposing port `9000`:

```bash
$ asanacli dockerize -image="library/golang:1.6.4" -expose=9000
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2016/12/26 22:34:54 INFO     ▶ 0001 Generating Dockerfile...
2016/12/26 22:34:54 SUCCESS  ▶ 0002 Dockerfile generated.
```

For more information on the usage, run `asana help dockerize`.

### asanacli dlv

asanacli can also help with debugging your application. To start a debugging session:

```bash
    ___   _____ ___    _   _____ 
   /   | / ___//   |  / | / /   |
  / /| | \__ \/ /| | /  |/ / /| |
 / ___ |___/ / ___ |/ /|  / ___ |
/_/  |_/____/_/  |_/_/ |_/_/  |_| v1.0
2017/03/22 11:17:05 INFO     ▶ 0001 Starting Delve Debugger...
Type 'help' for list of commands.
(dlv) break main.main
Breakpoint 1 set at 0x40100f for main.main() ./main.go:8

(dlv) continue
> main.main() ./main.go:8 (hits goroutine(1):1 total:1) (PC: 0x40100f)
     3:	import (
     4:		_ "github.com/user/myapp/routers"
     5:		"github.com/goasana/asana"
     6:	)
     7:	
=>   8:	func main() {
     9:		sana.Run()
    10:	}
    11:
```

For more information on the usage, run `asanacli help dlv`.

## Shortcuts

Because you'll likely type these generator commands over and over, it makes sense to create aliases:

```bash
# Generator Stuff
alias g:a="asanacli generate appcode"
alias g:m="asanacli generate model"
alias g:c="asanacli generate controller"
alias g:v="asanacli generate view"
alias g:mi="asanacli generate migration"
```

These can be stored , for example, in your `~/.bash_profile` or `~/.bashrc` files.

## Help

To print more information on the usage of a particular command, use `asana help <command>`.

For instance, to get more information about the `run` command:

```bash
$ asanacli help run
USAGE
  asanacli run [appname] [watchall] [-main=*.go] [-downdoc=true]  [-gendoc=true] [-vendor=true] [-e=folderToExclude]  [-tags=goBuildTags] [-runmode=ASANA_RUNMODE]

OPTIONS
  -downdoc
      Enable auto-download of the swagger file if it does not exist.

  -e=[]
      List of paths to exclude.

  -gendoc
      Enable auto-generate the docs.

  -main=[]
      Specify main go files.

  -runmode
      Set the Asana run mode.

  -tags
      Set the build tags. See: https://golang.org/pkg/go/build/

  -vendor=false
      Enable watch vendor folder.

DESCRIPTION
  Run command will supervise the filesystem of the application for any changes, and recompile/restart it.
```

## Contributing
Bug reports, feature requests and pull requests are always welcome.

We work on two branches: `master` for stable, released code and `develop`, a development branch.
It might be important to distinguish them when you are reading the commit history searching for a feature or a bugfix,
or when you are unsure of where to base your work from when contributing.

### Found a bug?

Please [submit an issue][new-issue] on GitHub and we will follow up.
Even better, we would appreciate a [Pull Request][new-pr] with a fix for it!

- If the bug was found in a release, it is best to base your work on `master` and submit your PR against it.
- If the bug was found on `develop` (the development branch), base your work on `develop` and submit your PR against it.

Please follow the [Pull Request Guidelines][new-pr].

### Want a feature?

Feel free to request a feature by [submitting an issue][new-issue] on GitHub and open the discussion.

If you'd like to implement a new feature, please consider opening an issue first to talk about it.
It may be that somebody is already working on it, or that there are particular issues that you should be aware of
before implementing the change. If you are about to open a Pull Request, please make sure to follow the [submissions guidelines][new-pr].

## Submission Guidelines

### Submitting an issue

Before you submit an issue, search the archive, maybe you will find that a similar one already exists.

If you are submitting an issue for a bug, please include the following:

- An overview of the issue
- Your use case (why is this a bug for you?)
- The version of `asana` you are running (include the output of `asana version`)
- Steps to reproduce the issue
- Eventually, logs from your application.
- Ideally, a suggested fix

The more information you give us, the more able to help we will be!

### Submitting a Pull Request

- First of all, make sure to base your work on the `develop` branch (the development branch):

```
  # a bugfix branch for develop would be prefixed by fix/
  # a bugfix branch for master would be prefixed by hotfix/
  $ git checkout -b feature/my-feature develop
```

- Please create commits containing **related changes**. For example, two different bugfixes should produce two separate commits.
A feature should be made of commits splitted by **logical chunks** (no half-done changes). Use your best judgement as to
how many commits your changes require.

- Write insightful and descriptive commit messages. It lets us and future contributors quickly understand your changes
without having to read your changes. Please provide a summary in the first line (50-72 characters) and eventually,
go to greater lengths in your message's body. A good example can be found in [Angular commit message format](https://github.com/angular/angular.js/blob/master/CONTRIBUTING.md#commit-message-format).

- Please **include the appropriate test cases** for your patch.

- Make sure all tests pass before submitting your changes.

- Rebase your commits. It may be that new commits have been introduced on `develop`.
Rebasing will update your branch with the most recent code and make your changes easier to review:

  ```
  $ git fetch
  $ git rebase origin/develop
  ```

- Push your changes:

  ```
  $ git push origin -u feature/my-feature
  ```

- Open a pull request against the `develop` branch.

- If we suggest changes:
  - Please make the required updates (after discussion if any)
  - Only create new commits if it makes sense. Generally, you will want to amend your latest commit or rebase your branch after the new changes:

    ```
    $ git rebase -i develop
    # choose which commits to edit and perform the updates
    ```

  - Re-run the tests
  - Force push to your branch:

    ```
    $ git push origin feature/my-feature -f
    ```

[new-issue]: #submitting-an-issue
[new-pr]: #submitting-a-pull-request

## Licence

```text
Copyright 2016 asana authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
