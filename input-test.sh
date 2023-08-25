#!/usr/bin/env sh

for i in test-inputs/*; do \
  x="$(echo "$i" | sed 's/.\+_\(.\+\)/\1/')"; \
  a="$(build/find-bin-width < "$i")"; \
  if [ "$x" != "$a" ]; then \
    echo "Test failed for input '$i'.  Expected '$x', got '$a'."; \
  fi; \
done

