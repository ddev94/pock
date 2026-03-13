# Architecture Overview

## System Architecture

```
┌─────────────────────────────────────────────────────────────────────┐
│                          CLI Entry Point                            │
│                         cmd/pock/main.go                            │
│                                                                      │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │                    Cobra Command Router                       │  │
│  │  - Root command setup                                         │  │
│  │  - Version & help management                                  │  │
│  │  - Subcommand registration                                    │  │
│  └──────────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                    ┌───────────────┴────────────────┐
                    │                                │
                    ▼                                ▼
┌─────────────────────────────────┐   ┌──────────────────────────────┐
│      Command Layer              │   │      Utility Layer           │
│   internal/commands/            │   │   internal/utils/            │
│                                 │   │                              │
│  Core Commands:                 │   │  • colors.go                 │
│  • add.go                       │   │    - Color formatting        │
│  • list.go                      │   │                              │
│  • run.go                       │   │  • exec.go                   │
│  • remove.go                    │   │    - Command execution       │
│                                 │   │    - Interactive/batch       │
│  Management:                    │   │                              │
│  • history.go                   │   │  • table.go                  │
│  • config.go                    │   │    - Table rendering         │
│                                 │   └──────────────────────────────┘
│  Import/Export:                 │
│  • export.go                    │
│  • import.go                    │
│                                 │
│  Marketplace (placeholder):     │
│  • install.go                   │
│  • browse.go                    │
│  • publish.go                   │
└─────────────────────────────────┘
                    │
                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                       Storage Layer                                  │
│                   internal/storage/                                  │
│                                                                      │
│  ┌────────────────────┐  ┌────────────────────┐  ┌──────────────┐  │
│  │   types.go         │  │   database.go      │  │  commands.go │  │
│  │                    │  │                    │  │              │  │
│  │  Data Structures:  │  │  Database Manager: │  │  CRUD Ops:   │  │
│  │  • SavedCommand    │  │  • Initialization  │  │  • Create    │  │
│  │  • CommandHistory  │  │  • Load/Save       │  │  • Read      │  │
│  │  • Settings        │  │  • Mutex locking   │  │  • Update    │  │
│  │  • StorageData     │  │  • Thread safety   │  │  • Delete    │  │
│  └────────────────────┘  └────────────────────┘  └──────────────┘  │
│                                                                      │
│  ┌────────────────────┐  ┌────────────────────┐                      │
│  │   history.go       │  │   settings.go      │                      │
│  │                    │  │                    │                      │
│  │  • Track execution │  │  • Get/Set config  │                      │
│  │  • Get history     │  │  • Layout prefs    │                      │
│  │  • Get stats       │  │  • Date format     │                      │
│  │  • Clear history   │  │  • Reset defaults  │                      │
│  └────────────────────┘  └────────────────────┘                      │
└─────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────┐
│                       Persistent Storage                             │
│                                                                      │
│                   ~/.local/share/pock/db.json                        │
│                                                                      │
│  {                                                                   │
│    "savedCommands": [...],                                           │
│    "commandHistories": [...],                                        │
│    "settings": {...}                                                 │
│  }                                                                   │
└─────────────────────────────────────────────────────────────────────┘
```

## Data Flow

### Adding a Command

```
User Input: pock add hello "echo 'Hello'"
        │
        ▼
┌─────────────────┐
│  Cobra Parser   │  Parse command line arguments
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  add.go         │  Validate and prepare data
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  commands.go    │  Check for duplicates
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  database.go    │  Thread-safe update
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  db.json        │  Persist to disk
└─────────────────┘
```

### Running a Command

```
User Input: pock run hello
        │
        ▼
┌─────────────────┐
│  Cobra Parser   │  Parse command name
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  run.go         │  Find command by name
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  commands.go    │  Retrieve command details
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  exec.go        │  Execute command
│                 │  • Shell selection
│                 │  • Interactive I/O
│                 │  • Time tracking
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  history.go     │  Record execution
│                 │  • Status
│                 │  • Output
│                 │  • Duration
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  db.json        │  Update history
└─────────────────┘
```

### Listing Commands

```
User Input: pock list --stats
        │
        ▼
┌─────────────────┐
│  Cobra Parser   │  Parse flags
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  list.go        │  Get display preferences
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  settings.go    │  Load layout settings
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  commands.go    │  Get all commands
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  history.go     │  Get stats (if --stats)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  table.go or    │  Render output
│  colors.go      │  (table or simple)
└─────────────────┘
```

## Package Dependencies

```
cmd/pock/main.go
    │
    ├─> github.com/spf13/cobra (CLI framework)
    │
    └─> internal/commands/*
            │
            ├─> internal/storage/*
            │       │
            │       ├─> github.com/google/uuid (UUID generation)
            │       └─> encoding/json (JSON marshaling)
            │
            └─> internal/utils/*
                    │
                    ├─> github.com/fatih/color (Color output)
                    ├─> github.com/olekukonko/tablewriter (Tables)
                    └─> os/exec (Command execution)
```

## Concurrency Model

```
┌─────────────────────────────────────────────────────────────┐
│                    Database Concurrency                     │
│                                                              │
│  ┌──────────────┐                                           │
│  │ Request 1    │──┐                                        │
│  └──────────────┘  │                                        │
│                    │   ┌─────────────────────┐              │
│  ┌──────────────┐  │   │  Mutex Lock (RW)    │              │
│  │ Request 2    │──┼──>│  - Read: RLock()    │              │
│  └──────────────┘  │   │  - Write: Lock()    │              │
│                    │   └─────────────────────┘              │
│  ┌──────────────┐  │              │                         │
│  │ Request 3    │──┘              │                         │
│  └──────────────┘                 │                         │
│                                   ▼                         │
│                         ┌──────────────────┐                │
│                         │   StorageData    │                │
│                         │   (In-Memory)    │                │
│                         └──────────────────┘                │
│                                   │                         │
│                                   ▼                         │
│                         ┌──────────────────┐                │
│                         │    db.json       │                │
│                         │   (On Disk)      │                │
│                         └──────────────────┘                │
└─────────────────────────────────────────────────────────────┘
```

## Component Interactions

```
┌────────────┐     ┌────────────┐     ┌────────────┐
│   User     │────>│   CLI      │────>│  Commands  │
└────────────┘     └────────────┘     └────────────┘
                                              │
                                              │ uses
                                              │
                        ┌─────────────────────┼─────────────────────┐
                        │                     │                     │
                        ▼                     ▼                     ▼
                  ┌──────────┐        ┌──────────┐        ┌──────────┐
                  │ Storage  │        │  Utils   │        │  Types   │
                  │          │        │          │        │          │
                  │ • DB     │        │ • Colors │        │ • Structs│
                  │ • CRUD   │        │ • Exec   │        │ • Consts │
                  │ • Lock   │        │ • Table  │        │          │
                  └──────────┘        └──────────┘        └──────────┘
                        │
                        │ reads/writes
                        │
                        ▼
                  ┌──────────┐
                  │ db.json  │
                  │          │
                  │ Persist  │
                  │ Storage  │
                  └──────────┘
```

## Error Handling Flow

```
Command Execution
        │
        ▼
    ┌───────────────┐
    │ Input Valid?  │────No───> Return Error + Usage Info
    └───────┬───────┘
           Yes
            │
            ▼
    ┌───────────────┐
    │ DB Access OK? │────No───> Return DB Error + Suggestion
    └───────┬───────┘
           Yes
            │
            ▼
    ┌───────────────┐
    │ Business Logic│
    │               │
    │ Validation?   │────No───> Return Validation Error
    └───────┬───────┘
           Yes
            │
            ▼
    ┌───────────────┐
    │ Perform Action│
    │               │────Error──> Rollback + Return Error
    └───────┬───────┘
          Success
            │
            ▼
    ┌───────────────┐
    │ Return Success│
    │ With Feedback │
    └───────────────┘
```

## File Statistics

- **Total Go Files**: 22
- **Total Lines of Code**: ~1,792
- **Packages**: 4 (main, commands, storage, utils)
- **Commands Implemented**: 12
- **Documentation Files**: 5
- **Build Scripts**: 3

## Key Design Patterns

1. **Command Pattern**: Each CLI command is a separate module
2. **Repository Pattern**: Storage layer abstracts data access
3. **Singleton Pattern**: Database instance is shared
4. **Factory Pattern**: Command creation in main.go
5. **Strategy Pattern**: Different output formats (table/simple)

## Security Considerations

- File permissions: 0644 for db.json (readable by user only)
- No network access except for import from URLs
- Command execution uses user's shell (security implications noted)
- Input validation on all user inputs

## Performance Optimizations

- Lazy database loading (on first access)
- Efficient JSON marshaling/unmarshaling
- Minimal memory allocations
- Direct file I/O (no buffering overhead)
- Fast command lookup with linear search (acceptable for typical use)

---

This architecture provides a clean separation of concerns, maintainable code structure, and room for future enhancements while maintaining simplicity and performance.
