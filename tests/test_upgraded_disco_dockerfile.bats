#!/usr/bin/env bats
#
# Framework: Bats (Bash Automated Testing System)
# Purpose: Validate Dockerfile changes introduced in this PR.
# Focus: Base image, ARG/ENV wiring, WORKDIR, COPY, Air install, CMD, and directive ordering.
# Notes:
# - Tests avoid external dependencies and do not build images; they statically validate the Dockerfile contents.
# - The Dockerfile path is auto-discovered by searching for unique markers from the diff.

setup() {
  # Try to locate the Dockerfile by a unique marker from the diff
  DOCKERFILE="${DOCKERFILE_PATH:-$(grep -R -l --exclude-dir='.git' --exclude-dir='node_modules' 'github.com/cosmtrek/air@latest' . | head -n1)}"
  if [ -z "$DOCKERFILE" ]; then
    # Fallback by base image string
    DOCKERFILE="$(grep -R -l --exclude-dir='.git' --exclude-dir='node_modules' '^FROM golang:1\.25\.0-alpine' . | head -n1)"
  fi
}

teardown() {
  true
}

@test "Dockerfile is found and readable" {
  [ -n "$DOCKERFILE" ]
  [ -f "$DOCKERFILE" ]
  [ -r "$DOCKERFILE" ]
}

@test "uses Go 1.25.0-alpine builder stage named 'builder'" {
  run grep -E '^FROM[[:space:]]+golang:1\.25\.0-alpine[[:space:]]+as[[:space:]]+builder$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "defines build args: CONFIG_SOURCE and CONFIG_PATH" {
  run grep -E '^ARG[[:space:]]+CONFIG_SOURCE$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
  run grep -E '^ARG[[:space:]]+CONFIG_PATH$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "sets WORKDIR to /app" {
  run grep -E '^WORKDIR[[:space:]]+/app$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "copies base_dir contents to /app" {
  # Expect exactly: COPY ${base_dir}/ /app
  run grep -E '^COPY[[:space:]]+\$\{base_dir\}\/[[:space:]]+\/app$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "propagates build args to environment variables" {
  run grep -E '^ENV[[:space:]]+CONFIG_SOURCE=\$\{CONFIG_SOURCE\}$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
  run grep -E '^ENV[[:space:]]+CONFIG_PATH=\$\{CONFIG_PATH\}$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "installs Air using go install at @latest" {
  run grep -E '^RUN[[:space:]]+go[[:space:]]+install[[:space:]]+github\.com/cosmtrek/air@latest$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "default command starts Air with .air.toml" {
  run grep -E '^CMD[[:space:]]+air[[:space:]]+-c[[:space:]]+\.air\.toml$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "ordering: FROM -> ARGs -> WORKDIR -> COPY -> ENVs -> RUN (air install) -> CMD" {
  from_ln=$(grep -nE '^FROM[[:space:]]+golang:1\.25\.0-alpine[[:space:]]+as[[:space:]]+builder$' "$DOCKERFILE" | cut -d: -f1)
  arg_src_ln=$(grep -nE '^ARG[[:space:]]+CONFIG_SOURCE$' "$DOCKERFILE" | head -n1 | cut -d: -f1)
  arg_path_ln=$(grep -nE '^ARG[[:space:]]+CONFIG_PATH$' "$DOCKERFILE" | head -n1 | cut -d: -f1)
  workdir_ln=$(grep -nE '^WORKDIR[[:space:]]+/app$' "$DOCKERFILE" | cut -d: -f1)
  copy_ln=$(grep -nE '^COPY[[:space:]]+\$\{base_dir\}\/[[:space:]]+\/app$' "$DOCKERFILE" | cut -d: -f1)
  env_src_ln=$(grep -nE '^ENV[[:space:]]+CONFIG_SOURCE=\$\{CONFIG_SOURCE\}$' "$DOCKERFILE" | cut -d: -f1)
  env_path_ln=$(grep -nE '^ENV[[:space:]]+CONFIG_PATH=\$\{CONFIG_PATH\}$' "$DOCKERFILE" | cut -d: -f1)
  run_air_ln=$(grep -nE '^RUN[[:space:]]+go[[:space:]]+install[[:space:]]+github\.com/cosmtrek/air@latest$' "$DOCKERFILE" | cut -d: -f1)
  cmd_ln=$(grep -nE '^CMD[[:space:]]+air[[:space:]]+-c[[:space:]]+\.air\.toml$' "$DOCKERFILE" | cut -d: -f1)

  # Ensure all markers are present
  for n in "$from_ln" "$arg_src_ln" "$arg_path_ln" "$workdir_ln" "$copy_ln" "$env_src_ln" "$env_path_ln" "$run_air_ln" "$cmd_ln"; do
    [ -n "$n" ]
  done

  # Verify strict increasing order
  [ "$from_ln" -lt "$arg_src_ln" ]
  [ "$arg_src_ln" -lt "$arg_path_ln" ]
  [ "$arg_path_ln" -lt "$workdir_ln" ]
  [ "$workdir_ln" -lt "$copy_ln" ]
  [ "$copy_ln" -lt "$env_src_ln" ]
  [ "$env_src_ln" -lt "$env_path_ln" ]
  [ "$env_path_ln" -lt "$run_air_ln" ]
  [ "$run_air_ln" -lt "$cmd_ln" ]
}

@test "defensive: exactly one CMD and it matches expected form" {
  cmd_count=$(grep -cE '^CMD[[:space:]]+' "$DOCKERFILE")
  [ "$cmd_count" -eq 1 ]
  run grep -E '^CMD[[:space:]]+air[[:space:]]+-c[[:space:]]+\.air\.toml$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}

@test "defensive: Air installation is pinned to @latest (no other tag/ref)" {
  run grep -E 'github\.com/cosmtrek/air@' "$DOCKERFILE"
  [ "$status" -eq 0 ]
  run grep -E '^RUN[[:space:]]+go[[:space:]]+install[[:space:]]+github\.com/cosmtrek/air@latest$' "$DOCKERFILE"
  [ "$status" -eq 0 ]
}