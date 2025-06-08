# Finance Tracker CLI - Development Plan

## Project Overview
A command-line finance tracking application built with Go and Bubbletea, featuring local data storage, configuration management, and an optional Google Sheets sync capability.

## Tech Stack
- **Language**: Go
- **TUI Framework**: Bubbletea + Lipgloss (for styling)
- **Data Storage**: JSON files (most appropriate for Go's native marshaling)
- **Configuration**: YAML (human-readable, good Go library support)

## Project Structure
```
finance-tracker/
├── cmd/
│   └── main.go # Main entry point
├── internal/
│   ├── config/ # Configuration Management
│   ├── models/ # Data Structures
│   ├── storage/ # Data persistence
│   ├── sync/ # (stretch) Syncing with Google Sheets
│   └── ui/ # TUI components
└── data/ # Data templates
```

---

## Phase 1: Foundation & Setup

### Task 1.1: Project Structure Setup
- [x] Set up Go module with proper directory structure
- [x] Initialize git repository with appropriate .gitignore
- [x] Add Bubbletea and Lipgloss dependencies
- [x] Create basic package structure following Go conventions

**Dependencies**: None

---

### Task 1.2: Data Models Definition
- [x] Define `Transaction` struct (ID, Amount, Description, Category, Date, Type)
- [x] Define `Config` struct (Currency, DateFormat, Categories, UserName)
- [x] Create enums/constants for transaction types (Income/Expense)
- [x] Add JSON tags for serialization and validation tags
- [x] Implement String() methods for display formatting

**Dependencies**: Task 1.1

**Go-specific notes**: Use `time.Time` for dates, `decimal` package for currency precision, struct tags for JSON marshaling

---

### Task 1.3: Configuration Management
- [ ] Create config package with default values for UK locale
- [ ] Implement config file creation in `~/.config/finance-tracker/`
- [ ] Add YAML parsing using `gopkg.in/yaml.v3`
- [ ] Create config validation and migration logic
- [ ] Implement config loading with fallbacks to defaults

**Dependencies**: Task 1.2

**Default values**: GBP currency, DD/MM/YYYY format, common UK expense categories

---

### Task 1.4: Local Storage Layer
- [ ] Create storage interface for data operations (CRUD)
- [ ] Implement JSON file-based storage in `~/.config/finance-tracker/data.json`
- [ ] Add file locking mechanism to prevent corruption
- [ ] Implement data backup/restore functionality
- [ ] Create data migration system for future schema changes

**Dependencies**: Task 1.2, Task 1.3

**Go-specific notes**: Use `os.UserConfigDir()` for cross-platform config directory, implement proper error handling with wrapped errors

---

## Phase 2: Core TUI Implementation

### Task 2.1: Main Application Structure
- [ ] Create main Bubbletea model with navigation state
- [ ] Implement routing between different views/screens
- [ ] Add global key bindings (quit, help, navigation)
- [ ] Create shared styling using Lipgloss
- [ ] Implement proper error handling and user feedback

**Dependencies**: Task 1.1

---

### Task 2.2: Initial Setup Flow
- [ ] Create first-run detection logic
- [ ] Build welcome screen with setup wizard
- [ ] Implement config directory creation with proper permissions
- [ ] Add user input forms for initial configuration
- [ ] Create setup completion confirmation

**Dependencies**: Task 1.3, Task 1.4, Task 2.1

---

### Task 2.3: Main Menu Implementation
- [ ] Design main menu layout with options list
- [ ] Add keyboard navigation (arrow keys, enter, escape)
- [ ] Implement menu item highlighting and selection
- [ ] Add status bar showing current user/config info
- [ ] Create help overlay with keyboard shortcuts

**Dependencies**: Task 2.1, Task 2.2

---

## Phase 3: Core Features

### Task 3.1: Transaction List View
- [ ] Create paginated transaction display
- [ ] Implement sorting options (date, amount, category)
- [ ] Add filtering capabilities (by category, date range, type)
- [ ] Show running balance and summary statistics
- [ ] Add search functionality with real-time filtering

**Dependencies**: Task 1.4, Task 2.1

**Go-specific notes**: Use slices package for sorting, implement proper pagination to handle large datasets

---

### Task 3.2: Add Transaction Form
- [ ] Create form with input fields (amount, description, category, date)
- [ ] Implement form validation with error display
- [ ] Add category selection with autocomplete/dropdown
- [ ] Support for quick entry shortcuts
- [ ] Add transaction type toggle (income/expense)

**Dependencies**: Task 1.2, Task 1.4, Task 2.1

---

### Task 3.3: Transaction Management
- [ ] Implement transaction editing functionality
- [ ] Add delete confirmation dialog
- [ ] Create bulk operations (multi-select, bulk delete)
- [ ] Add transaction duplication feature
- [ ] Implement undo/redo for recent operations

**Dependencies**: Task 3.1, Task 3.2

---

### Task 3.4: Categories Management
- [ ] Create category management screen
- [ ] Allow adding/editing/deleting categories
- [ ] Implement category usage tracking
- [ ] Add category color coding for visual distinction
- [ ] Create category-based spending insights

**Dependencies**: Task 1.3, Task 3.1

---

## Phase 4: Enhanced Features

### Task 4.1: Reporting & Analytics
- [ ] Create spending summary by category
- [ ] Implement monthly/yearly spending trends
- [ ] Add income vs expenses comparison
- [ ] Generate basic charts using ASCII art
- [ ] Export reports to CSV/text format

**Dependencies**: Task 3.1, Task 3.4

---

### Task 4.2: Data Import/Export
- [ ] Implement CSV import functionality with mapping
- [ ] Add data export in multiple formats (CSV, JSON)
- [ ] Create data backup and restore commands
- [ ] Add data validation and error reporting for imports
- [ ] Support for common bank statement formats

**Dependencies**: Task 1.4, Task 3.1

---

### Task 4.3: Enhanced UI Polish
- [ ] Add loading states and progress indicators
- [ ] Implement smooth animations and transitions
- [ ] Create responsive layout for different terminal sizes
- [ ] Add theme customization options
- [ ] Improve error messages and user guidance

**Dependencies**: Task 2.1, All Phase 3 tasks

---

## Phase 5: Google Sheets Integration (Stretch Goal)

### Task 5.1: Google Sheets Authentication
- [ ] Set up Google Sheets API credentials and OAuth2
- [ ] Implement authentication flow with token storage
- [ ] Add credential management and renewal logic
- [ ] Create secure token storage in config directory
- [ ] Handle authentication errors gracefully

**Dependencies**: Task 1.3

**Go-specific notes**: Use `golang.org/x/oauth2` and Google's official Go client libraries

---

### Task 5.2: Sheets Data Sync
- [ ] Create spreadsheet template with proper headers
- [ ] Implement bidirectional sync logic (local ↔ sheets)
- [ ] Add conflict resolution for concurrent edits
- [ ] Create sync status tracking and error reporting
- [ ] Implement incremental sync to avoid full data replacement

**Dependencies**: Task 5.1, Task 1.4

---

### Task 5.3: Sync Management UI
- [ ] Add sync configuration screen
- [ ] Create sync status display and progress tracking
- [ ] Implement manual sync trigger and scheduling
- [ ] Add sync conflict resolution interface
- [ ] Create sync history and error logs

**Dependencies**: Task 5.2, Task 2.1

---

## Testing Strategy

### Task T.1: Unit Testing
- [ ] Write tests for all data models and validation
- [ ] Test storage operations with mock data
- [ ] Create config loading/saving test cases
- [ ] Test transaction CRUD operations
- [ ] Add edge case testing for data validation

**Dependencies**: Corresponding implementation tasks

---

### Task T.2: Integration Testing
- [ ] Test complete user workflows end-to-end
- [ ] Create test fixtures for various data scenarios
- [ ] Test file system operations with temporary directories
- [ ] Verify cross-platform compatibility
- [ ] Test Google Sheets integration with mock API

**Dependencies**: Major feature completion

---

## Documentation

### Task D.1: User Documentation
- [ ] Create comprehensive README with installation instructions
- [ ] Write user guide with screenshots/examples
- [ ] Document configuration options and defaults
- [ ] Create troubleshooting guide
- [ ] Add contribution guidelines

### Task D.2: Developer Documentation
- [ ] Document code architecture and design decisions
- [ ] Create API documentation for internal packages
- [ ] Add inline code comments following Go conventions
- [ ] Document build and deployment processes

---

## Go-Specific Implementation Notes

- **Error Handling**: Use wrapped errors with `fmt.Errorf` and `errors.Is/As`
- **Concurrency**: Use goroutines for file I/O and API calls, channels for communication
- **Package Organization**: Follow Go module conventions, use internal/ for private packages
- **Testing**: Use table-driven tests, testify for assertions
- **CLI Framework**: Consider cobra for command structure if the app grows beyond TUI
- **Logging**: Use structured logging with `log/slog` for debugging
- **Configuration**: Use environment variables for overrides, following 12-factor principles
