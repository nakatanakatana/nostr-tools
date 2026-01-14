# Track Plan: Refactor NIP-05 to Pre-generated Static File Serving

## Phase 1: Preparation & Config Parsing Logic
- [ ] Task: Create new track directory and setup files.
- [x] Task: Verify existing tests pass for `cmd/nip05`. [d6d141c]
- [ ] Task: Refactor `config.go` (or create new utility) to expose parsed data structure for file generation. (Ensure current `main.go` logic is preserved during transition).

## Phase 2: File Generation Logic
- [ ] Task: TDD - Write tests for `FileGenerator` struct/function.
    - Test generating full JSON.
    - Test generating individual user JSONs.
    - Test handling of file paths in TempDir.
- [ ] Task: Implement `FileGenerator`.
    - Logic to take the `nip05.Data` (or similar map) and write `nip05_full.json`.
    - Logic to loop through users and write `nip05_user_<name>.json`.
    - Return a map or lookup mechanism for "Name -> FilePath".

## Phase 3: Handler Refactoring
- [ ] Task: TDD - Write tests for the new Handler logic using `http.ServeFile`.
    - Test serving full file when no param.
    - Test serving user file when param exists.
    - Test serving 404/empty when user not found.
- [ ] Task: Implement the new Handler.
    - Integrate `FileGenerator` at startup in `main.go`.
    - Replace existing request handling logic with file serving logic.
- [ ] Task: Conductor - User Manual Verification 'Handler Refactoring' (Protocol in workflow.md)

## Phase 4: Cleanup & Final Verification
- [ ] Task: Remove old dynamic JSON generation code if no longer used.
- [ ] Task: Run full integration tests/e2e tests locally.
- [ ] Task: Verify documentation/comments are updated.
- [ ] Task: Conductor - User Manual Verification 'Cleanup & Final Verification' (Protocol in workflow.md)
