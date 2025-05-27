#!/usr/bin/env bash
set -euo pipefail

RELEASE_DIR="$(pwd)/release/"
PACKAGE_NAME="alfred-gcp-workflow.alfredworkflow"
BUNDLE_ID="com.dineshchikkanna.alfred.gcp"

bold() { echo -e "\033[1m$1\033[0m"; }
green() { echo -e "\033[32m$1\033[0m"; }
red() { echo -e "\033[31m$1\033[0m"; }
yellow() { echo -e "\033[33m$1\033[0m"; }
spinner() { 
  local pid=$1
  local delay=0.1
  local spinstr='|/-\'
  while [ "$(ps a | awk '{print $1}' | grep "$pid")" ]; do
    local temp=${spinstr#?}
    printf " [%c]  " "$spinstr"
    local spinstr=$temp${spinstr%"$temp"}
    sleep $delay
    printf "\b\b\b\b\b\b"
  done
  printf "    \b\b\b\b"
}


prompt_version() {
  bold "Input new version (e.g., 1.2.3):"
  read -r VERSION_INPUT
  VERSION="v${VERSION_INPUT//v/}"
}

update_info_plist() {
  bold "Updating version in info.plist to ${VERSION_INPUT}..."
  /usr/libexec/PlistBuddy -c "Set :version ${VERSION_INPUT}" info.plist
}

commit_version_update() {
  bold "Do you want to commit and push the updated info.plist? (y/n)"
  read -r SHOULD_COMMIT
  if [[ "$SHOULD_COMMIT" == "y" || "$SHOULD_COMMIT" == "Y" ]]; then
    bold "Committing and pushing changes..."
    git add info.plist
    git commit -m "Release ${VERSION}" && git push origin main &
    spinner $!
    green "✔️ Committed and pushed to main."
  else
    yellow "⚡ Skipped commit and push."
  fi
}

create_and_push_tag() {
  bold "Do you want to create and push a Git tag ${VERSION}? (y/n)"
  read -r SHOULD_TAG
  if [[ "$SHOULD_TAG" == "y" || "$SHOULD_TAG" == "Y" ]]; then
    git tag "${VERSION}"
    git push origin "${VERSION}" &
    spinner $!
    green "✔️ Tag ${VERSION} created and pushed."
  else
    yellow "⚡ Skipped tag creation."
  fi
}

clean_release_dir() {
  bold "Cleaning up release directory..."
  rm -rf "$RELEASE_DIR"
  mkdir -p "$RELEASE_DIR"
}

build_binaries() {
  bold "Building binaries..."
  GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION" -o "$RELEASE_DIR/alfred-gcp-workflow-amd64" &
  spinner $!
  
  GOARCH=arm64 go build -ldflags "-X main.Version=$VERSION" -o "$RELEASE_DIR/alfred-gcp-workflow-arm64" &
  spinner $!
  
  green "✔️ Binaries built."
}

merge_binaries() {
  bold "Merging binaries into a universal binary..."
  lipo -create -output "$RELEASE_DIR/alfred-gcp-workflow" \
    "$RELEASE_DIR/alfred-gcp-workflow-arm64" \
    "$RELEASE_DIR/alfred-gcp-workflow-amd64"
  green "✔️ Merged binaries."
}

clean_intermediate_binaries() {
  bold "Cleaning up intermediate binaries..."
  rm -f "$RELEASE_DIR/alfred-gcp-workflow-arm64" \
        "$RELEASE_DIR/alfred-gcp-workflow-amd64"
}

copy_assets() {
  bold "Copying workflow assets..."
  cp -R assets services.yml regions.yml icon.png info.plist LICENSE README.md "$RELEASE_DIR"
}

code_sign() {
  bold "Code signing the workflow..."
  cd "$RELEASE_DIR" || exit 1

  if [[ $SHOULD_TAG != "y" && $SHOULD_TAG != "Y" ]]; then
    yellow "⚡ Skipping code signing as no tag was created."
    return
  fi

  if [[ -z "${APPLE_DEVELOPER_ID_CERT_ID:-}" ]]; then
    red "❌ Missing Apple Developer ID Certificate ID. Please set the APPLE_DEVELOPER_ID_CERT_ID environment variable."
    exit 1
  fi
  
  codesign -s "$APPLE_DEVELOPER_ID_CERT_ID" -f -v --timestamp --options runtime ./alfred-gcp-workflow
  if [[ $? -ne 0 ]]; then
    red "❌ Code signing failed. Please ensure you have a valid code signing identity."
    exit 1
  fi
  
  green "✔️ Code signed successfully."
  cd - || exit 1
}

package_workflow() {
  bold "Packaging .alfredworkflow file..."
  ditto -ck "$RELEASE_DIR" "$PACKAGE_NAME"
  green "✔️ Packaged workflow."
}

zip_workflow() {
  bold "Zipping .alfredworkflow for GitHub upload..."
  zip -q "${PACKAGE_NAME}.zip" "$PACKAGE_NAME"
  green "✔️ Zipped workflow."
}

notarize_app() {
  bold "Notarizing the build..."

  if [[ $SHOULD_TAG != "y" && $SHOULD_TAG != "Y" ]]; then
    yellow "⚡ Skipping notarization as no tag was created."
    return
  fi

  if [[ -z "${APPLE_ID:-}" || -z "${APPLE_TEAM_ID:-}" || -z "${APPLE_DEVELOPER_APP_PASSWORD:-}" ]]; then
    red "❌ Missing Apple ID, Team ID, or Developer App Password. Please set these environment variables."
    exit 1
  fi
 
  xcrun notarytool submit "${PACKAGE_NAME}.zip" \
    --wait \
    --apple-id "$APPLE_ID" \
    --team-id "$APPLE_TEAM_ID" \
    --password "$APPLE_DEVELOPER_APP_PASSWORD" 
  
  if [[ $? -ne 0 ]]; then
    red "❌ Notarization failed. Please check your credentials and try again."
    exit 1
  fi
  
  green "✔️ Notarization completed successfully."
}

open_github_release_page() {
  bold "Opening GitHub release page for tag ${VERSION}..."
  open "https://github.com/dineshgowda24/alfred-gcp-workflow/releases/new?tag=${VERSION}&title=${VERSION}&body=%23%23%20Changes%0A%0AUser-facing%0A-%20TODO"
}

open_finder() {
  bold "Opening Finder to release folder..."
  open .
}

show_files() {
  bold "✅ Release build complete. Files generated:"
  ls -lh "${PACKAGE_NAME}" "${PACKAGE_NAME}.zip"
}

main() {
  prompt_version
  update_info_plist
  commit_version_update
  create_and_push_tag
  clean_release_dir
  build_binaries
  merge_binaries
  clean_intermediate_binaries
  copy_assets
  code_sign
  package_workflow
  zip_workflow
  notarize_app
  show_files
  open_github_release_page
  open_finder
}

main "$@"
