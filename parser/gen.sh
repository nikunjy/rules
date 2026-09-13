#!/bin/bash
set -euo pipefail

cd "$(dirname "$0")"

# Try antlr4 first, fall back to antlr if not available
if command -v antlr4 >/dev/null 2>&1; then
    ANTLR_CMD="antlr4"
elif command -v antlr >/dev/null 2>&1; then
    ANTLR_CMD="antlr"
else
    echo "Error: Neither antlr4 nor antlr command found. Please install ANTLR."
    exit 1
fi

$ANTLR_CMD -Dlanguage=Go -visitor -no-listener JsonQuery.g4 -o ./

# ANTLR's Go target emits an unreachable `goto errorExit` after each
# `return localctx` so the errorExit label is always referenced. go vet
# flags that as unreachable code. Every generated function already jumps
# to errorExit on real error paths, so the dummy jump is safe to drop.
sed -i '/goto errorExit \/\/ Trick to prevent compiler error if the label is not used/d' jsonquery_parser.go
