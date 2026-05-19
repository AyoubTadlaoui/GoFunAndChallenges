#!/usr/bin/env sh
#
# gen.sh — regenerate the README demo (animated GIF/WebP + hero PNG) with
# the atlas-ragnarok storm-fire vignette baked in.
#
# Requirements (install once):
#   brew install vhs ffmpeg webp
#   pip3 install --user --break-system-packages pillow numpy
#
# Run from the repo root:
#   sh docs/_theme/gen.sh
#
# Produces:
#   docs/demo.webp — animated WebP (README embed), shader applied per pixel
#                    with the luminance mask so text stays crisp.
#   docs/demo.gif  — animated GIF (download / sub-caption link), same shader.
#   docs/demo.mp4  — supplementary MP4 download.
#   docs/hero.png  — static hero, shader applied.
#
# The shader is a CPU reimplementation of atlas-ragnarok.glsl. See
# docs/_theme/_shader.py for the math.

set -eu

cd "$(dirname "$0")/../.."   # repo root
HERE="docs"
THEME="docs/_theme"

command -v vhs      >/dev/null 2>&1 || { echo "vhs not installed (brew install vhs)" >&2; exit 1; }
command -v ffmpeg   >/dev/null 2>&1 || { echo "ffmpeg not installed (brew install ffmpeg)" >&2; exit 1; }
command -v gif2webp >/dev/null 2>&1 || { echo "gif2webp not installed (brew install webp)" >&2; exit 1; }
command -v python3  >/dev/null 2>&1 || { echo "python3 not installed" >&2; exit 1; }
python3 -c "import numpy, PIL" 2>/dev/null \
  || { echo "missing python deps (pip3 install --user --break-system-packages pillow numpy)" >&2; exit 1; }
command -v go >/dev/null 2>&1 \
  || { echo "go not installed (https://go.dev/dl/)" >&2; exit 1; }

echo "→ raw GIF (vhs, atlas-ragnarok base palette)"
vhs "${HERE}/demo.tape"
mv "${HERE}/demo.gif" "${HERE}/_demo_raw.gif"

echo "→ apply GLSL shader per pixel (luma mask keeps text crisp)"
rm -rf "${HERE}/_frames"
python3 "${THEME}/_shader.py" "${HERE}/_demo_raw.gif" "${HERE}/_frames"

# PIL reports per-frame duration in milliseconds. Average → fps for ffmpeg.
AVG_MS=$(awk '{s+=$1; n++} END {printf "%d", (s/n) + 0.5}' "${HERE}/_frames/duration.txt")
FPS=$(awk -v ms="$AVG_MS" 'BEGIN {if (ms <= 0) ms = 100; printf "%g", 1000.0 / ms}')
echo "    avg frame delay ${AVG_MS}ms (${FPS} fps)"

echo "→ reassemble shaded GIF (ffmpeg palettegen for cross-frame delta)"
ffmpeg -hide_banner -loglevel error -y \
  -framerate "$FPS" -i "${HERE}/_frames/frame_%05d.png" \
  -filter_complex "split[s0][s1];[s0]palettegen=stats_mode=full[p];[s1][p]paletteuse=dither=bayer:bayer_scale=5" \
  -loop 0 \
  "${HERE}/demo.gif"

echo "→ demo.webp (animated WebP, embedded in README)"
gif2webp -quiet -q 80 -m 6 -mt -mixed \
  "${HERE}/demo.gif" \
  -o "${HERE}/demo.webp"

echo "→ demo.mp4 (supplementary download)"
ffmpeg -hide_banner -loglevel error -y \
  -framerate "$FPS" -i "${HERE}/_frames/frame_%05d.png" \
  -movflags +faststart \
  -pix_fmt yuv420p \
  -vf "pad=ceil(iw/2)*2:ceil(ih/2)*2" \
  -c:v libx264 -preset slow -crf 22 \
  "${HERE}/demo.mp4"

echo "→ hero.png (separate vhs pose with empty top/bottom bands)"
vhs "${HERE}/hero.tape"
ffmpeg -hide_banner -loglevel error -y -sseof -1 -i "${HERE}/_hero_raw.gif" \
  -vsync vfr -frames:v 1 -update 1 "${HERE}/_hero_raw.png"
python3 "${THEME}/_shader.py" "${HERE}/_hero_raw.png" "${HERE}/hero.png"

# Drop intermediates.
rm -rf "${HERE}/_frames"
rm -f "${HERE}/_demo_raw.gif" "${HERE}/_hero_raw.gif" "${HERE}/_hero_raw.png"

echo "done."
echo "  hero.png   $(du -h ${HERE}/hero.png   | cut -f1)"
echo "  demo.webp  $(du -h ${HERE}/demo.webp  | cut -f1)  ← embedded in README"
echo "  demo.gif   $(du -h ${HERE}/demo.gif   | cut -f1)"
echo "  demo.mp4   $(du -h ${HERE}/demo.mp4   | cut -f1)"
