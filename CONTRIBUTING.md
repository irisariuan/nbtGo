# Contributing to nbtGo

Thank you for your interest in contributing to nbtGo! This document provides guidelines and instructions for contributing to the project.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/your-username/nbtGo.git
   cd nbtGo
   ```
3. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## Development Setup

### Prerequisites

- Go 1.25.6 or later
- Git

### Building

```bash
go build -o nbtGo
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run a specific test
go test -v -run TestDeserializeTagByte ./lib/nbt
```

## Code Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Write clear, descriptive variable and function names
- Add comments for exported functions and types
- Keep functions focused and reasonably sized

### Running Linters

```bash
# Format code
go fmt ./...

# Vet code for common mistakes
go vet ./...
```

## Testing

- Write tests for all new functionality
- Ensure existing tests pass
- Aim for good test coverage
- Use table-driven tests where appropriate
- Test both success and error cases

### Test File Naming

- Test files should be named `*_test.go`
- Place tests in the same package as the code they test

### Example Test Structure

```go
func TestNewFeature(t *testing.T) {
    // Test setup
    input := []byte{...}
    
    // Execute
    result, err := SomeFunction(input)
    
    // Assert
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result != expected {
        t.Errorf("Expected %v, got %v", expected, result)
    }
}
```

## Pull Request Process

1. Update documentation (README.md, doc.go files) if needed
2. Add or update tests for your changes
3. Ensure all tests pass
4. Run `go fmt ./...` and `go vet ./...`
5. Commit your changes with clear, descriptive commit messages
6. Push to your fork
7. Create a Pull Request with a clear description of the changes

### Commit Message Guidelines

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests when applicable

Example:
```
Add support for custom NBT tags

- Implement CustomTag struct
- Add parsing logic for custom tags
- Update documentation

Fixes #123
```

## What to Contribute

### Bug Reports

- Use the GitHub issue tracker
- Describe the bug clearly
- Include steps to reproduce
- Provide expected vs actual behavior
- Include relevant code snippets or test cases

### Feature Requests

- Use the GitHub issue tracker
- Clearly describe the feature and its use case
- Explain why this feature would be useful
- Consider implementation details if possible

### Code Contributions

Areas where contributions are especially welcome:

- Bug fixes
- Performance improvements
- Additional utility functions
- Documentation improvements
- Example code and tutorials
- Test coverage improvements

## Code Review Process

- All submissions require review before merging
- Reviewers may request changes
- Address feedback promptly
- Keep discussions respectful and constructive

## Documentation

- Update README.md for user-facing changes
- Update doc.go files for API changes
- Add code comments for complex logic
- Include examples for new features

## License

By contributing to nbtGo, you agree that your contributions will be licensed under the same license as the project.

## Questions?

If you have questions about contributing, feel free to:
- Open an issue with the label "question"
- Reach out to the maintainers

## Code of Conduct

- Be respectful and inclusive
- Focus on constructive feedback
- Assume good intentions
- Help create a welcoming environment

Thank you for contributing to nbtGo!
