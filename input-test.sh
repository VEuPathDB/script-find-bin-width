#!/usr/bin/env sh

EXECUTABLE=${1}

if [ -z "$EXECUTABLE" ]; then
  echo "USAGE:"
  echo "  test-input.sh path/to/find-bin-width"
  echo ""
  exit 1
fi

for i in test-inputs/*; do \
  x="$(echo "$i" | sed 's/.\+_\(.\+\)/\1/')"; \
  a="$($EXECUTABLE --rm-na "$i")"; \
  if [ "$x" != "$a" ]; then \
    echo "Test failed for input '$i'.  Expected '$x', got '$a'."; \
  fi; \
done

