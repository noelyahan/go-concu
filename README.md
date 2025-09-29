# go-concurrency-examples

A **hands-on collection of Go concurrency examples**.  
This repository contains small, self-contained programs that demonstrate Go’s concurrency primitives — goroutines, channels, synchronization, timeouts, and more.

---

## 🚀 Features

- Simple, focused examples of Go concurrency
- Covers goroutines, channels, select statements, timeouts, and sync primitives
- Easy to run and experiment with
- Useful as a reference or learning resource

---

## 📂 Repository Structure

```
├── 01-basic-routine.go # simple goroutine execution
├── 02-basic-channel.go # basic channel send/receive
├── 03-multiple-routine.go # launching multiple goroutines
├── 04-channel-direction.go # directional channels (send-only / receive-only)
├── 05-channel-selector.go # using select on multiple channels
├── 06-channel-timeout.go # channel operations with timeouts
├── 07-sync-routine.go # synchronization with WaitGroup / sync
├── 08-just-signalling.go # signaling patterns with channels
├── go.mod
└── go.sum
```


---

## 📦 Getting Started

### Prerequisites
- Go 1.20+ (or latest stable)

### Run an Example
Pick any file and run it directly:
```bash
go run 05-channel-selector.go
```

