# Contributing

## Setup

```bash
gh repo fork imdevan/l --clone
cd l
```

## Development

```bash
just build        # build the binary
just build-run    # build and run
just watch        # rebuild on file changes
just test         # run tests
```

## Submitting Changes

1. Fork the repo and create a branch from `main`
2. Make your changes and ensure tests pass (`just test`)
3. Open a pull request with a clear description of what changed and why
