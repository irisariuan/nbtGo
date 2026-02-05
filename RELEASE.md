# Release Process

This repository uses GitHub Actions to automatically create releases when version tags are pushed.

## How to Create a Release

1. Ensure all changes are committed and pushed to the `main` or `master` branch
2. Create and push a version tag:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

3. The GitHub Actions workflow will automatically:
   - Build the shared libraries for Linux, macOS, and Windows
   - Run tests on all platforms
   - Create a GitHub release with auto-generated release notes
   - Upload the following artifacts to the release:
     - `libnbt-linux-amd64.so` - Linux shared library
     - `libnbt-macos-amd64.dylib` - macOS shared library
     - `libnbt-windows-amd64.dll` - Windows shared library
     - `libnbt.h` - C header file
     - `nbt-linux-amd64` - Linux CLI tool
     - `nbt-macos-amd64` - macOS CLI tool
     - `nbt-windows-amd64.exe` - Windows CLI tool

## Version Naming

Use semantic versioning for tags:
- `v1.0.0` - Major release
- `v1.1.0` - Minor release
- `v1.1.1` - Patch release

All tags should start with `v` to trigger the release workflow.
