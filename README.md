# protoc-gen-lint
[![Build Status](http://img.shields.io/travis/ckaznocha/protoc-gen-lint.svg?style=flat)](https://travis-ci.org/ckaznocha/protoc-gen-lint)
[![Release](http://img.shields.io/github/release/ckaznocha/protoc-gen-lint.svg?style=flat)](https://github.com/ckaznocha/protoc-gen-lint/releases/latest)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://ckaznocha.mit-license.org)
<!-- [![Coverage Status](https://img.shields.io/coveralls/ckaznocha/protoc-gen-lint.svg?style=flat)](https://coveralls.io/r/ckaznocha/protoc-gen-lint?branch=master) -->

A plug-in for Google's [Protocol Buffers](https://github.com/google/protobuf)
compiler to check `.proto` files for style violations.

## About
This plug-in will check a `.proto` file for violations of Google's Protocol
Buffer [Style Guide](https://developers.google.com/protocol-buffers/docs/style).
The protobuf compiler already reports on compilation errors; by using
this plug-in you are also able to retrieve those compilation error without
writing any file which is be helpful for use cases like IDE integrations.

## Installation
Download `protoc-gen-lint` and make sure it's available in your PATH. Once it's
in your PATH, `protoc` will be able to make use of the plug-in.

### Dependencies
You must have a working version of Google's Protocol Buffers compiler `protoc`
in your PATH. You can download it
[here](https://developers.google.com/protocol-buffers/docs/downloads)

### Using go get
If you have a go environment already set up you can use `go get` to install.
```
go get github.com/ckaznocha/protoc-gen-lint
```

### Downloading binary
Download the latest release for your operating system
[here](https://github.com/ckaznocha/protoc-gen-lint/releases/latest) and ensure
the executable is available in your PATH.

## Usage
```
protoc --lint_out=. *.proto
```
### Optional Checks
Optional checks are enabled but passing a comma separated list to the protoc
"parameter". Example:
```
protoc --lint_out=sort_imports:. *.proto
```

Available optional checks:
*   `sort_imports` - checks that imports are sorted in descending alphabetical
order.

## TODO
*   Write more tests
*   Find out about any common protocol buffer smells to check for from the community

## Contributing
See the `CONTRIBUTING` file.

## License
See `LICENSE` file
