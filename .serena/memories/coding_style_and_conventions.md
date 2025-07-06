# Coding Style and Conventions

## Testing Conventions
- Test data stored in `input/` directory with naming pattern:
  - `dayXX-sample.in` for sample input
  - `dayXX.in` for real input

## Package Organization
- `day01/`, `day02/`, etc. for daily solutions
- `input/` for input data files

## Code Structure
- Build incrementally using TDD
- Create generic, reusable functions
- Use meaningful function names that describe purpose

## Type-Driven Development
- Use **domain modeling with types** (semantic typing/newtype pattern)
- Create type aliases for domain concepts instead of raw primitives:
- Think in terms of domain concepts rather than raw data types

## Tool Selection Guidelines
- **Use semantic symbol tools** for renaming types, functions, variables, methods.
- **Use `replace_regex` only as fallback** for text-based operations that symbol tools cannot handle
- **Preferred hierarchy**: Symbol tools → semantic search tools → `replace_regex`
- **Key principle**: Always leverage semantic understanding of the codebase rather than treating code as plain text

## Testing Standards
- **Use table-driven tests** for all test functions
- Group related test cases into logical test functions
- Use descriptive test names that clearly indicate what is being tested
- Use `testify/assert` and `testify/require` for assertions
- Aim for 100% statement coverage with edge cases
- Keep test cases focused and atomic
