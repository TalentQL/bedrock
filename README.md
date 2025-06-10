# TalentQL

## Bedrock

### Introduction

A collection of reusable utilities for building APIs

### Usage

```bash
# Published version
go get github.com/TalentQL/bedrock@v1.0.6

# Local development
git clone github.com/TalentQL/bedrock
go mod edit -replace=github.com/TalentQL/bedrock@v0.0.0-unpublished=../bedrock
go get github.com/TalentQL/bedrock@v0.0.0-unpublished
```

### Structure

1. `pkg` - reusables for config, dbs, repository controller etc

### TODO

1. Add unit tests
