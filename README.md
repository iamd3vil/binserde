# Binserde

A code generator for serializing & deserializing binary data in Go.

## Motivation

Currently for serializing & deserializing binary data in Go, there is `encoding/binary`, but you either have to manually write code for structs without fixed size or use `binary.Write()` / `binary.Read()` for structs with fixed size. 

So `binserde` generates efficient code for serializing & deserializing data for structs. The goal is that this should be as fast as handwritten code.

### TODO

- [ ] Marshal, unmarshal code.
    - [x] Basic Types(int, float, string, []byte).
    - [x] Embedded structs.
    - [X] Custom types.