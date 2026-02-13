#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
BIN="$ROOT_DIR/bin/gor"

# Build Gor binary if needed
if [ ! -f "$BIN" ]; then
    echo "Building Gor..."
    go build -o "$BIN" "$ROOT_DIR/cmd/gor/"
fi

echo "======================================"
echo "  Gor Language Benchmark Suite"
echo "======================================"
echo ""

run_benchmark() {
    local name="$1"
    local gor_file="$SCRIPT_DIR/$name.gor"
    local py_file="$SCRIPT_DIR/$name.py"
    local lua_file="$SCRIPT_DIR/$name.lua"
    local js_file="$SCRIPT_DIR/$name.js"

    echo "--- $name ---"
    echo ""

    # Gor
    printf "  Gor:         "
    gor_start=$(date +%s%N)
    gor_out=$("$BIN" "$gor_file")
    gor_end=$(date +%s%N)
    gor_ms=$(( (gor_end - gor_start) / 1000000 ))
    printf "%6d ms  (result: %s)\n" "$gor_ms" "$gor_out"

    # Python
    printf "  Python:      "
    py_start=$(date +%s%N)
    py_out=$(python3 "$py_file")
    py_end=$(date +%s%N)
    py_ms=$(( (py_end - py_start) / 1000000 ))
    printf "%6d ms  (result: %s)\n" "$py_ms" "$py_out"

    # Lua
    if command -v lua &> /dev/null; then
        printf "  Lua:         "
        lua_start=$(date +%s%N)
        lua_out=$(lua "$lua_file")
        lua_end=$(date +%s%N)
        lua_ms=$(( (lua_end - lua_start) / 1000000 ))
        printf "%6d ms  (result: %s)\n" "$lua_ms" "$lua_out"
    else
        echo "  Lua:         (skipped - lua not installed)"
    fi

    # JavaScript (Node.js / Bun)
    local js_runtime=""
    if command -v bun &> /dev/null; then
        js_runtime="bun"
    elif command -v node &> /dev/null; then
        js_runtime="node"
    fi

    if [ -n "$js_runtime" ]; then
        printf "  JavaScript:  "
        js_start=$(date +%s%N)
        js_out=$(NO_COLOR=1 $js_runtime "$js_file" 2>/dev/null)
        js_end=$(date +%s%N)
        js_ms=$(( (js_end - js_start) / 1000000 ))
        printf "%6d ms  (result: %s)\n" "$js_ms" "$js_out"
    else
        echo "  JavaScript:  (skipped - node/bun not installed)"
    fi

    echo ""
}

run_benchmark "fibonacci"
run_benchmark "prime_sieve"

echo "======================================"
