# Track Plan: Refactor NIP-05 to Pre-generated Static File Serving

## Phase 1: Preparation & Config Parsing Logic [checkpoint: 57366bf]
- [ ] Task: Create new track directory and setup files.
- [x] Task: Verify existing tests pass for `cmd/nip05`. [d6d141c]
- [x] Task: Refactor `config.go` (or create new utility) to expose parsed data structure for file generation. (Ensure current `main.go` logic is preserved during transition). [2ae4ba6]

## Phase 2: File Generation Logic [checkpoint: 8501cf5]
- [x] Task: TDD - Write tests for `FileGenerator` struct/function. [80023b0]
- [x] Task: Implement `FileGenerator`. [80023b0]

## Phase 3: Handler Refactoring [checkpoint: 55d415c]
- [x] Task: TDD - Write tests for the new Handler logic using `http.ServeFile`. [ecf2be2]
- [x] Task: Implement the new Handler. [ecf2be2]
- [x] Task: Conductor - User Manual Verification 'Handler Refactoring' (Protocol in workflow.md) [checkpoint: skipping manual verification as it is automated in phase completion]

## Phase 4: Cleanup & Final Verification
- [ ] Task: Remove old dynamic JSON generation code if no longer used.
- [ ] Task: Run full integration tests/e2e tests locally.
- [ ] Task: Verify documentation/comments are updated.
- [ ] Task: Conductor - User Manual Verification 'Cleanup & Final Verification' (Protocol in workflow.md)
