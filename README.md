# stcompilerlib

## About
This is intended as a standalone library for adding "compilation" support to a subset of the Structured Text language to any Go project.
It does not support variable declarations or the like. Instead, it works over simple constructs such as loops, if/then, expressions, etc.

## Usage

* *Parsing:* Simply call `stcompilerlib.ParseString()` to convert a structured text String into a slice of instructions and expressions.
* *Compiling:* Simply call `stcompilerlib.XXXCompileSequence()` with your set of instructions, where `XXX` is either `C`, `Vhdl`, or `Verilog`. You may also wish to help the compiler understand your code further by using the `stcompilerlib.SetKnownVarNames()` function to clarify known variable names (it does a best guess however).