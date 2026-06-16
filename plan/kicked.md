## Feature 6: Interactive mode

- [ ] 6.1 `/` key enters filter input; typing narrows the visible entries
- [ ] 6.2 `default_interactive=true` config option drops into interactive mode by default
- [ ] 6.3 Esc / Enter exits interactive mode
  - notes: use Bubble Tea inline mode (no full-screen takeover). See existing `internal/ui/list.go` for reference pattern.

## Feature 3: Filtering & navigation
- [ ] 3.2 Directory hop: if query exactly matches a directory name, `cd` into it and list
  - [ ] 3.2.1 If query matches exactly one directory (fuzzy), hop into it
  - [ ] 3.2.2 Two positional args: first is the directory to hop, second is the filter query inside that directory
  - [ ] 3.2.3 Warn when second arg is ignored (ambiguous first arg, no hop occurred)
- [ ] 3.3 show arrow indicating sort direction on field being sorted in the direction of the sort   


