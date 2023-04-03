# Go-Dart Interoperability Toolchain
This project aims to build a toolchain for generating C wrappers and Dart bindings to Go libraries, making it easier for Dart developers to use Go libraries in their projects. The toolchain will consist of a Go package and a Dart package, along with a command-line interface for generating the necessary code.

## Background
Dart is a client-optimized programming language developed by Google, while Go is a server-side programming language designed for performance and scalability. While both languages excel in their respective domains, there is no easy way to use Go libraries in Dart projects, and vice versa. This project aims to bridge that gap by providing a set of tools that make it easy for Dart developers to use Go libraries in their projects.

## Project Goals
The goal of this project is to build a toolchain that generates C wrappers and Dart bindings to Go libraries, allowing Dart developers to use Go libraries in their projects without having to write tedious and error-prone C bindings by hand. The toolchain will consist of the following components:

A Go package that generates C wrappers for Go functions and types.
A Dart package that generates Dart bindings for C functions and types.
A command-line interface that automates the generation of C wrappers and Dart bindings.

## Project Scope
The toolchain will be designed to work with Go packages that follow the Go package layout convention, where all Go code is located in a single directory. The toolchain will not support Go packages that use multiple directories, as this would require more complex parsing and code generation logic.

The toolchain will support Go functions that take basic types (int, float, string, etc.) as arguments and return basic types as results. The toolchain will also support Go functions that take pointers to structs as arguments and return pointers to structs as results.

The toolchain will generate C code that uses the C89 standard, and Dart code that uses the FFI library to interface with the C code.

## Getting Started
To use the toolchain, simply install the Go and Dart packages, and then run the command-line interface with the path to the Go package as an argument. The command-line interface will generate the necessary C and Dart code, which you can then use in your Dart project.

