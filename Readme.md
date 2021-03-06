# A [![Build Status](https://travis-ci.org/davidrjenni/A.svg?branch=master)](https://travis-ci.org/davidrjenni/A) [![GoDoc](https://godoc.org/github.com/davidrjenni/A?status.svg)](https://godoc.org/github.com/davidrjenni/A) [![Go Report Card](https://goreportcard.com/badge/github.com/davidrjenni/A)](https://goreportcard.com/report/github.com/davidrjenni/A)

A - Go tools for Acme.

## Installation

```
% go get github.com/davidrjenni/A

% go get golang.org/x/tools/cmd/guru
% go get github.com/zmb3/gogetdoc
% go get github.com/godoctor/godoctor
% go get github.com/josharian/impl
% go get golang.org/x/tools/cmd/gorename
```

## Usage

### Callees

```
A cle <scope>
```
Shows possible targets of the function call under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Callers

```
A clr <scope>
```
Shows possible callers of the function under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Callstack

```
A cs <scope>
```
Shows the path from the callgraph root to the function under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Goto Definition

```
A def
```
Shows the declaration for the identifier under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

### Describe

```
A desc
```
Describes the declaration for the syntax under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

### Documentation

```
A doc
```
Shows the documentation for the entity under the cursor.
This command uses `github.com/zmb3/gogetdoc`.

### Errors

```
A err <scope>
```
Shows possible values of the error variable under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Extract to Function/Method

```
A ex <name>
```
Extracts the selected statements to a new function/method with name `<name>`.
This command uses `github.com/godoctor/godoctor`.

### Freevars

```
A fv
```
Shows the free variables of the selected snippet.
This command uses `golang.org/x/tools/cmd/guru`.

### Generate Method Stubs

```
A impl <recv> <iface>
A impl 'f *File' io.ReadWriteCloser
```
Generates method stubs with receiver `<recv>` for implementing the interface `<iface>` and inserts them at the location of the cursor.
This command uses `github.com/josharian/impl`.

### Implements

```
A impls <scope>
```
Shows the `implements` relation for the type or method under the cursor.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Peers

```
A peers <scope>
```
Shows send/receive corresponding to the selected channel op.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Points To

```
A pto <scope>
```
Shows variables the selected pointer may point to.
`<scope>` is a comma-separated list of packages the analysis should be limited to, this parameter is optional.
This command uses `golang.org/x/tools/cmd/guru`.

### Referrers

```
A refs
```
Shows all refs to the entity denoted by identifier under the cursor.
This command uses `golang.org/x/tools/cmd/guru`.

### Renaming

```
A rn <name>
```
Renames the entity under the cursor with `<name>`.
This commands uses `golang.org/x/tools/cmd/gorename`.

### Share

```
A share
```
Uploads the selected snippet to play.golang.org and prints the URL.

### What

```
A what
```
Shows basic information about the selected syntax node.
This command uses `golang.org/x/tools/cmd/guru`.
