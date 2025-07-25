#!/bin/bash
set -e

echo "📦 Installing zip utility..."
apt-get update && apt-get install -y zip

echo "Searching for GoLand..."
GOLAND_PATH=$(find / -name goland.sh 2>/dev/null | head -n 1)
if [ -z "$GOLAND_PATH" ]; then
  echo "Could not find goland.sh"
  exit 1
fi
echo "Found GoLand at $GOLAND_PATH"

# Symlink to /usr/local/bin for easy access
ln -sf "$GOLAND_PATH" /usr/local/bin/goland
# Check if it’s in PATH
if ! command -v goland &>/dev/null; then
  echo "goland not in PATH"
  exit 1
fi


# Install plugins
echo "📦 Installing GoLand plugins..."
goland installPlugins \
  lechuck.intellij-plugin.task \
  ru.adelf.idea.dotenv \
  com.intellij.ml.llm


echo "📁 Locating installed plugins..."
PLUGIN_SOURCE_DIR=$(find /root/.cache/JetBrains -type d -name "plugins" | head -n 1)
GOLAND_BACKEND_DIR=$(find /.jbdevcontainer/JetBrains/RemoteDev/dist -type d -name "*goland*" | head -n 1)
PLUGIN_ZIP_TARGET_DIR="$GOLAND_BACKEND_DIR/plugins"

echo "📁 Plugin source: $PLUGIN_SOURCE_DIR"
echo "📁 Plugin target: $PLUGIN_ZIP_TARGET_DIR"

mkdir -p "$PLUGIN_ZIP_TARGET_DIR"

echo "🗜 Repacking and copying plugins..."
for PLUGIN_ZIP in "$PLUGIN_SOURCE_DIR"/*.zip; do
  PLUGIN_NAME=$(basename "$PLUGIN_ZIP" .zip)
  TEMP_DIR=$(mktemp -d)

  echo " - Unzipping $PLUGIN_ZIP to $TEMP_DIR"
  unzip -q "$PLUGIN_ZIP" -d "$TEMP_DIR"

  echo " - Repacking $PLUGIN_NAME.zip with correct structure"
  (
    cd "$TEMP_DIR" && zip -qr "$PLUGIN_ZIP_TARGET_DIR/$PLUGIN_NAME.zip" .
  )

  rm -rf "$TEMP_DIR"
done

echo "Plugin installation complete."