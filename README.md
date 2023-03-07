# Go universal time format parser

[![GitHub Workflow Status](https://github.com/aohorodnyk/gotimeparser/actions/workflows/go.yml/badge.svg)](https://github.com/aohorodnyk/gotimeparser/actions/workflows/go.yml)
[![License](https://img.shields.io/github/license/aohorodnyk/gotimeparser)](https://github.com/aohorodnyk/gotimeparser/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/aohorodnyk/gotimeparser)](https://github.com/aohorodnyk/gotimeparser/issues)
[![GitHub issues](https://img.shields.io/github/issues-pr/aohorodnyk/gotimeparser)](https://github.com/aohorodnyk/gotimeparser/pulls)
[![The latest release version](https://img.shields.io/github/v/release/aohorodnyk/gotimeparser)](https://github.com/aohorodnyk/gotimeparser/releases)
[![GoDoc](https://godoc.org/github.com/aohorodnyk/gotimeparser?status.svg)](https://pkg.go.dev/github.com/aohorodnyk/gotimeparser)

## Background

Many companies use service to service integration through different API porotocols.
I had the same situation in my production environment. At some point, clients of my service started to send me dates in different formats.

Some of them sent timestamps in int, some of them used ISO 8601, etc. But what was clear, these dates were valid and we could not request to change it in one day.

As a result of this issue, I implemented some solution that solved the issue and wrote a couple of articles that shared the issue and the approach:

1. [Parse timestamp formats](https://aohorodnyk.com/post/2022-05-06-parse-timestamp/)
1. [Parse time from different non timestamp formats](https://aohorodnyk.com/post/2022-05-07-parse-time-strings/)
1. [Universal time UnmarshalJSON implementation](https://aohorodnyk.com/post/2022-05-08-universal-time-unmarshaljson/)

But I did not want to create an Open Source package, because of I believe that Open Source cannot be just few lines of code. It's a right formatting of the code, some code coverage by tests, examples and the most important - is commitment to support it.

## Documentation
To make documentation more standartized and always up-to-date, we use [GoDoc (pkg.go.dev/github.com/aohorodnyk/gotimeparser)](https://pkg.go.dev/github.com/aohorodnyk/gotimeparser). It is the best documentation tool for Go.

Every feature or feature change should introduce some changes in comments, doc.go, README.md and tests include examples.

## Contributing
The library is open source and you can contribute to it.

Before contrbution, make sure that githook is configured for you and all your commits contain the correct issue tag.

### Branch Name
Before you start the contribution, make sure that you are on the correct branch. Branch name should start from the issue number dash and short explanation with spaces replaced by underscores. Example:
* `1-my_feature`
* `2-fix_bug`
* `234-my_important_pr`

### Git Hook
To configure the git hook, you need to simply run the command: `git config core.hooksPath .githooks`

It will configure the git hook to run the `pre-commit` script. Source code of the hook is in `.githooks/prepare-commit-msg`.

This git-hook will parse branch name and add the issue tag to the commit message.
