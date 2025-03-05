cmd:git() {
  # Interactive rebase
  git rebase -i "$(arg:commit)"

  # Interactive fixup
  git commit --fixup="$(arg:commit)"

  # Interactive revert
  git revert "$(arg:commit)"

  # Switch branch
  git switch "$(arg:branch)"

  # See logs - including deleted
  git log --all --full-history
}

# Select commit from log
arg:commit() {
  git flog --all |
    fzf --ansi --bind 'ctrl-p:toggle-preview' --preview-window=':hidden' \
      --preview "echo {} | grep -oEi '[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]?' | xargs -r git show" |
    grep -oEi '[0-9a-f]{7,8}' |
    head -1
}

# Select branch
arg:branch() { git branch | fzf; }
