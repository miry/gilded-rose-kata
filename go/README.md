# GO Starter

- Run :

```shell
$ make run
```

- Run tests :

```shell
$ make test
```

- Run tests and coverage :

```shell
$ make coverage
```

# Packages

Use standard libraries for testing, for easy setup, and skip create a packages for incapsulate code.
Ideally services and modules should be extracted to `pkg` folder.

# Linter

Check that code is designed by Team practicies and save team on code review.
Use https://github.com/golangci/golangci-lint as tool to run multiple checks.

```shell
$ go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
# Mac OS could be installed
$ brew install golangci/tap/golangci-lint
```

To check code, before be sure that `golangci-lint` in `$PATH`

```shell
$ make lint
```

## Git

### Hooks

To run git hooks from repo. Be sure that Linter installed.

```shell
$ git config core.hooksPath ./go/bin/hooks
```
