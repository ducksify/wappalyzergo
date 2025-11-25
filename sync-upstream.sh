#!/bin/bash
# Script to sync your fork with upstream projectdiscovery/wappalyzergo

set -e

echo "🔄 Syncing with upstream repository..."

# Fetch latest changes from upstream
echo "📥 Fetching latest changes from upstream..."
git fetch upstream

# Get current branch name
CURRENT_BRANCH=$(git branch --show-current)
echo "📍 Current branch: $CURRENT_BRANCH"

# Check if there are uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "⚠️  Warning: You have uncommitted changes."
    echo "   Please commit or stash them before syncing."
    exit 1
fi

# Merge upstream/main into current branch
echo "🔀 Merging upstream/main into $CURRENT_BRANCH..."
git merge upstream/main --no-edit

echo "✅ Successfully synced with upstream!"
echo ""
echo "📤 To push changes to your fork, run:"
echo "   git push origin $CURRENT_BRANCH"

