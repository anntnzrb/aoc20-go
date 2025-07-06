# Architecture and Utilities

## Core Architecture

### Day Structure
Each AOC day follows this pattern:
- `dayXX/dayXX.go` - Main implementation
- `dayXX/dayXX_test.go` - Comprehensive test suite
- Input files in `input/dayXX.in` and `input/dayXX-sample.in`

## Design Patterns
- **Test-Driven Development**: Red-Green-Refactor cycle
- **Generic Programming**: Use Go generics for type safety
- **Separation of Concerns**: Clear package boundaries
- **Reusable Utilities**: Build components for future days
