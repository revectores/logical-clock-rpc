# README

### 1. Brief Introduction

This project implemented three types of logical clocks introduced by [Why Logical Clocks are Easy](https://github.com/revectores/logical-clock-rpc/blob/main/docs/Why%20Logical%20Clocks%20are%20Easy.pdf):

- **Casual Histories**
- **Vector Clock**
- **Version Vector**

A concise notes summarizing the three methods is provided [here](https://github.com/revectores/logical-clock-rpc/blob/main/docs/Why%20Logical%20Clocks%20are%20Easy%20Notes.pdf).

Refer to [Vector Clock Implementation Based on RPC](https://github.com/revectores/logical-clock-rpc/blob/main/docs/Vector%20Clock%20Implementation%20Based%20on%20RPC.md) for more details about the design and structure of this project.







### 2. Build Tutorial

1. Install Go and make sure that `go` is in your `PATH`.

2. In `src/vector_clock/node` folder, `go build node.go` to create a node executable.

3. In `src/vector_clock/cli` folder, `go build cil.go` to create a cli executable.

4. Run `cli`. The commands and their interpretations can be found in [Vector Clock Implementation Based on RPC](https://github.com/revectores/logical-clock-rpc/blob/main/docs/Vector%20Clock%20Implementation%20Based%20on%20RPC.md).



Here an example is given for the  interaction in vector clock:

```
>> create
5
../node/node 30003
../node/node 30004
../node/node 30000
../node/node 30001
../node/node 30002
>> init 0
0
>> init 1
1
>> get 0
[0 0 0 0 0]
>> update 0
0
>> update 1
0
>> send 0 1
0
>> get 1
[1 2 0 0 0]
>> exit
```

