# go-git-staged

> Go! Git' staged!

Run commands on files staged in [git](https://git-scm.com/). Filter files using globs and pass them to their respective commands as arguments.

This cli application is a spiritual successor to [lint-staged](https://github.com/okonet/lint-staged) focused on being very simple and fast. It is also something for me to write in the[Go programming language](https://golang.org/).

```
❯ go-git-staged                                                                              
Run commands on files staged in git

Usage:
  go-git-staged [flags]

Examples:
go-git-staged --glob '*.js' --command 'eslint' --command 'prettier'

Flags:
  -a, --all                   Glob all files known to git instead of just staged
  -w, --working-dir string    Working directory for commands (default "/Users/iiro/git/go-git-staged")
      --relative              Use file paths relative to --working-dir instead of absolute
  -v, --verbose               Print command stdout after success instead of only fail
  -g, --glob stringArray      Glob of files passed to following --command
  -c, --command stringArray   Command to run with files matching previous --glob
  -h, --help                  help for go-git-staged
```

# Installation

See the [Latest Release](https://github.com/iiroj/go-git-staged/releases/latest) page for the binaries that can be copied directly.

## macOS

### Homebrew

This repository contains a [Homebrew Tap Formula](https://docs.brew.sh/Taps) for easier installation:

```sh
❯ brew tap iiroj/go-git-staged https://github.com/iiroj/go-git-staged
❯ brew install go-git-staged
```

# Usage

The `go-git-staged` cli uses no configuration files and is run directly using flags. Build pairs of `--glob` and their following `--command` flags to run multiple commands on multiple globs. By default files staged in git are globbed. Using the `--all` flag every file known to git will be used instead.

Because `go-git-staged` uses the `git` command, it needs to be run inside a git repository. By default the current working directory is used, and the git repo root is resolved from there. Using the `--working-dir` flag it is possible to run `go-git-staged` in another directory, but it also has to be a git directly. Commands always use the specified working directly to run in.

Commands receive the complete file paths that matched their respective globs, as their arguments. By default file paths are resolved as absolute. By using the `--relative` flag they will be relative to the `--working-dir` (_not necessarily the git repository root_).

There should be at least a single `--glob` flag followed by a `--command` flag, otherwise `go-lint-staged` will exit with an error. You can chain as many `--glob` and `--command` flags as necessary to build the desired patterns, but every new `--glob` splits into a new separate group.

All groups of commands are run concurrently. However, inside a group, all commands are run serially in the order specified in the cli flags.

# Examples

- **"For all .js files, run eslint"**

  ```
  go-git-staged --glob "**/*.js" --command "eslint"
  ```

- **"For all .js and .ts files, run eslint"**

  ```
  go-git-staged --glob "**/*.js" --glob "**/*.ts" --command "eslint"
  ```

- **"For all .js and .ts files, run eslint and prettier"**

  ```
  go-git-staged --glob "**/*.js" --glob "**/*.ts" --command "eslint" --command "prettier"
  ```

- **"For all .js files, run eslint, and for all .ts files, run prettier**

  ```
  go-git-staged --glob "**/*.js" --command "eslint" --glob "**/*.ts" --command "prettier"
  ```
