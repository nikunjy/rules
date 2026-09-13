#!/bin/bash

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
