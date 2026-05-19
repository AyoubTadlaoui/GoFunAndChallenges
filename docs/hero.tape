# vhs tape for the hero still (docs/hero.png).
#
# Composition rule: leave generous empty bands at top and bottom of the
# frame. The atlas-ragnarok vignette paints into those bands (thunder-
# blue at top, crimson at bottom) so an over-stuffed frame defeats it.
#
# Regenerate via the shader pipeline:
#   sh docs/_theme/gen.sh

Output docs/_hero_raw.gif

Set Shell zsh
Set FontSize 18
Set Width 1280
Set Height 760
Set Padding 56
Set TypingSpeed 0ms
Set WindowBar Colorful
Set WindowBarSize 40

Set Theme { "name": "atlas-ragnarok", "background": "#000000", "foreground": "#ffffff", "cursor": "#3b82f6", "selection": "#1e3a5f", "black": "#000000", "red": "#ff8080", "green": "#99ffe4", "yellow": "#2563eb", "blue": "#a0a0a0", "magenta": "#3b82f6", "cyan": "#99ffe4", "white": "#ffffff", "brightBlack": "#505050", "brightRed": "#ff9999", "brightGreen": "#b3ffe4", "brightYellow": "#60a5fa", "brightBlue": "#b0b0b0", "brightMagenta": "#93c5fd", "brightCyan": "#99ffe4", "brightWhite": "#ffffff" }

Hide
Type "PROMPT='%F{8}$%f '" Enter
Type "go build ./..." Enter
Sleep 35s
Type "clear" Enter
Show

# A few empty prompt lines so the command lands in the middle band,
# leaving room for the top vignette. Looks like an idle terminal.
Enter
Enter
Enter

Type "go run ./lessons/13-concurrency" Sleep 300ms Enter
Sleep 3000ms

# Hold the final frame so gen.sh's ffmpeg can grab a steady last frame.
Sleep 1500ms
