# ![ratchet](https://raw.githubusercontent.com/chendrix/ratchet/master/ratchet.png) Ratchet
> Make your code better little by little via ratcheting.

![travis-ci](https://travis-ci.org/chendrix/ratchet.svg?branch=master)

For every go codebase that doesn't currently pass `golint` or `go vet`.

Make your codebase better one package at a time.

## Installing / Getting started

```shell
go install github.com/chendrix/ratchet
ratchet path/to/project
```

Here you should say what actually happens when you execute the code above.

## Features

## Configuration

## Developing

Here's a brief intro about what a developer must do in order to start developing
the project further:

```shell
git clone https://github.com/chendrix/ratchet.git
cd ratchet/
./bin/bootstrap
```

Clone the repository and run the `bootstrap` script to get up and running.

## Contributing

If you'd like to contribute, please fork the repository and use a feature
branch. Pull requests are warmly welcome.

Steps to Verify Before Submitting a PR:

* Do the tests pass? `bin/test`
* Has it been formatted? `go fmt ./...`
* Have any new dependencies been vendored via [gvt](https://github.com/FiloSottile/gvt)?
* Have you written documentation for the code you wrote?
* Have you documented any major changes in the README? (e.g. changes to Usage, follow [README best practices](https://github.com/jehna/readme-best-practices))
* Are your commits [small and well-written](http://chris.beams.io/posts/git-commit/)?

## Licensing

See [LICENSE](https://raw.githubusercontent.com/chendrix/ratchet/master/LICENSE)

* Wrench by Creative Stall from the Noun Project