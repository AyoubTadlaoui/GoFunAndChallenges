# vhs tape for the animated demo (docs/demo.gif / .webp / .mp4).
# Regenerate via the shader pipeline:
#   sh docs/_theme/gen.sh

Output docs/demo.gif

Set Shell zsh
Set FontSize 17
Set Width 1280
Set Height 680
Set Padding 32
Set TypingSpeed 45ms
Set WindowBar Colorful
Set WindowBarSize 40

# atlas-ragnarok palette — pure black background, thunder-blue cursor,
# peppermint accents, crimson for errors.
# https://github.com/AyoubTadlaoui/atlas-ragnarok
Set Theme { "name": "atlas-ragnarok", "background": "#000000", "foreground": "#ffffff", "cursor": "#3b82f6", "selection": "#1e3a5f", "black": "#000000", "red": "#ff8080", "green": "#99ffe4", "yellow": "#2563eb", "blue": "#a0a0a0", "magenta": "#3b82f6", "cyan": "#99ffe4", "white": "#ffffff", "brightBlack": "#505050", "brightRed": "#ff9999", "brightGreen": "#b3ffe4", "brightYellow": "#60a5fa", "brightBlue": "#b0b0b0", "brightMagenta": "#93c5fd", "brightCyan": "#99ffe4", "brightWhite": "#ffffff" }

# Off-camera: minimal prompt + pre-warm Go build cache so the visible
# `go test` lands in ~3s instead of ~40s on a cold cache. The test packages
# are the same ones rendered visibly below.
Hide
Type "PROMPT='%F{8}$%f '" Enter
Sleep 200ms
Type "go build ./..." Enter
Sleep 35s
Type "go test ./lessons/01-hello ./lessons/13-concurrency ./challenges/02-linked-list >/dev/null 2>&1" Enter
Sleep 20s
Type "clear" Enter
Sleep 200ms
Show

# 1. Quickest sanity check — one runnable lesson.
Type "go run ./lessons/01-hello" Sleep 300ms Enter
Sleep 2500ms

# 2. A heftier lesson: goroutines, channels, mutex, atomic.
Type "go run ./lessons/13-concurrency" Sleep 300ms Enter
Sleep 3000ms

# 3. The "every package has tests" promise — three real packages, all pass.
Type "go test ./lessons/01-hello ./lessons/13-concurrency ./challenges/02-linked-list" Sleep 300ms Enter
Sleep 4500ms

# Hold the final frame so viewers can read the test summary.
Sleep 2500ms
