# Go Language Mastery: Basics to Advanced

This repository documents my comprehensive learning journey through the Go programming language, covering everything from foundational concepts to advanced techniques, with extensive coverage on Protocol Buffers and gRPC.

## Learning Path & Progress

- **Basics:** ✅ **COMPLETED**

  - [x] Introduction to Go (`module-1-intro-to-go/hello.go`): Basic Go program structure, printing, and logger usage.
  - [x] Data Types, Variables, and Constants (`module-2-datatypes/datatypes-variables-constants.go`): Declaring variables, constants, type inference, and scope.
  - [x] Arithmetic Operators (`module-3-arthimetic-operators/operators.go`): Basic arithmetic operations, integer division, modulo, and overflow/underflow notes.
  - [x] Loops (`module-4-loops/loops.go`): For loops, range, break/continue, and while-style loops.
  - [x] Arrays (`module-5-arrays/arrays.go`): Array declaration, initialization, iteration, copying, multidimensional arrays, and pointers to arrays.
  - [x] Slices (`module-6-slices/slices.go`): Slice declaration, initialization, slicing, appending, copying, 2D slices, and length vs. capacity.
  - [x] Maps (`module-7-maps/maps.go`): Map declaration, initialization, adding, deleting, checking existence, iterating, and multi-dimensional maps.
  - [x] Range (`module-8-range/range.go`): Using range with arrays, slices, maps, and strings.
  - [x] Functions (`module-9-functions/`): Function declaration, parameters, return values, multiple returns, variadic functions, and function types.
  - [x] Defer (`module-10-defer/defer.go`): Defer statements, execution order, and common use cases.
  - [x] Panic (`module-11-panic/panic.go`): Panic mechanism, built-in panic function, and panic scenarios.
  - [x] Recover (`module-12-recover/recover.go`): Recover function, panic recovery, and error handling patterns.
  - [x] Exit (`module-13-exit/exit.go`): Program termination with os.Exit and exit codes.
  - [x] Init (`module-14-init/init.go`): Init function, package initialization, and execution order.

- **Intermediate:** **IN PROGRESS**

  - [x] Closures (`01-clousures/closure.go`): Function closures, lexical scoping, and closure patterns.
  - [x] Pointers (`02-pointers/pointers.go`): Pointer declaration, dereferencing, pointer arithmetic, and pointer usage patterns.
  - [x] Strings and Runes (`03-strings-runes/strings_runes.go`): String manipulation, rune types, string iteration, and Unicode handling.
  - [x] Fmt Package (`04-fmt/fmt.go`): Formatting verbs, printing functions, and string formatting.
  - [x] Structs (`05-struct/struct.go`):
    - Defining structs, nested structs, and anonymous (embedded) fields.
    - Initializing structs using composite literals and dot notation.
    - Anonymous structs for ad-hoc data.
    - Struct comparison using the equality operator.
    - Example: See `intermediate/05-struct/struct.go` for a comprehensive demonstration.
  - [x] Methods (`06-methods/methods.go`): Method declaration, value and pointer receivers, method sets, and method chaining.
  - [x] Interfaces (`07-interface/interface.go`): Interface definition, interface implementation, empty interfaces, and type assertions.
  - [x] Struct Embedding (`08-struct-embedding/structEmbedding.go`): Embedded structs, promotion, and composition patterns.
  - [x] Generics (`09-generics/generics.go`): Generic functions, generic types, type constraints, and generic patterns.
  - [x] Error Handling (`10-errors/`): Basic errors, custom errors, error wrapping, and error handling patterns.
  - [x] Time and Epoch (`11-time-epoch/main.go`): Time manipulation, epoch conversion, time formatting, and timezone handling.
  - [x] JSON (`12-json/main.go`): JSON marshalling/unmarshalling, JSON tags, and JSON handling.
  - [x] Environment Variables (`13-env/main.go`): Reading environment variables, configuration management, and env file handling.
  - [x] Logging (`14-logging/main.go`): Basic logging, log levels, and log formatting.
  - [x] Type Conversions (`15-type-conversions/main.go`): Type casting, type assertions, and type conversions.
  - [x] Logrus (`16-logrus/main.go`): Structured logging with Logrus, log levels, hooks, and formatters.
  - [x] Zap (`17-zap/main.go`): High-performance structured logging with Zap, sugar and logger modes.
  - [x] URL Parsing (`21-url-parsing/parsing.go`): URL parsing, query parameters, path manipulation, and URL building.

- **Advanced:** **IN PROGRESS**

  - [x] Goroutines (`01-goroutines/main.go`): Creating goroutines, concurrent execution, and goroutine basics.
  - [x] Channels (`02-channels/main.go`): Channel creation, sending/receiving, blocking behavior, and channel basics.
  - [x] Buffered Channels (`03-buffered-channels/main.go`): Buffered channel creation, capacity, and non-blocking behavior.
  - [x] Channel Synchronization (`04-channel-synchronization/main.go`): Using channels for synchronization, wait patterns, and coordination.
  - [x] Channel Directions (`05-channel-directions/main.go`): Bidirectional, send-only, and receive-only channels, channel direction parameters.
  - [x] Multiplexing with Select (`06-multiplexing-select/main.go`): Select statement, handling multiple channels, default cases, and timeout patterns.
  - [x] Non-blocking Channels (`07-non-blocking-channel/main.go`): Non-blocking sends and receives, select with default, and async patterns.
  - [x] Closing Channels (`08-closing-channels/main.go`): Channel closure, detecting closed channels, range over channels, and cleanup patterns.
  - [x] Context (`09-context/main.go`): Context package, context cancellation, timeouts, deadlines, and context propagation.
  - [x] Timers and Tickers (`10-times-tickers/`): Time.After, time.Tick, timers, tickers, and periodic tasks.
  - [x] Worker Pool (`11-worker-pool/main.go`): Worker pool pattern, job distribution, concurrent processing, and pool management.
  - [ ] Protocol Buffers
  - [ ] gRPC
  - [ ] Building REST and gRPC APIs
  - [ ] Data Structures
  - [ ] Database Integration (SQL, NoSQL - MongoDB, MariaDB)
  - [ ] Version Control with Git and GitHub
  - [ ] Benchmarking Techniques and Tools (wrk, h2load, ghz)
  - [ ] HTTP/2 and HTTPS Servers
  - [ ] Code Obfuscation
  - [ ] Reflection Package
  - [ ] Go Runtime Internals
  - [ ] ...

## Repository Structure

- **`basics/`**: Code examples and exercises for fundamental Go concepts.
- **`intermediate/`**: Code for intermediate Go topics covering pointers, structs, methods, interfaces, error handling, logging, and more.
- **`advanced/`**: Extensive coverage of advanced topics including goroutines, channels, concurrency patterns, context, and worker pools.
- **`projects/`**: Larger, real-world projects built during the learning process (e.g., `slack-bot/`).
- **`gophercises/`**: Solutions and exercises from the Gophercises course (e.g., `exercise-01-quiz-game/`).
- **`notes/`**: Personal notes, summaries, and key takeaways including formatting verbs and naming conventions.
- **`resources/`**: Any external resources, useful links, or recommended readings.

## How to Use This Repository

Each folder corresponds to a major section or topic in the learning path. Within each, you'll find subfolders for specific concepts, containing the Go code (`.go` files).

Feel free to explore the code!

---

**Happy Coding!**
