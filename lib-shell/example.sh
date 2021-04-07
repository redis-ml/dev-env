function my_examples() {
  echo "
# Tmux cmd
tmux -CC new -A -s main

# Tmux + ssh
ssh -t "host" "tmux -CC new -A -s main"
  "
}
