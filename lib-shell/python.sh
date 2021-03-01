export PATH="$PATH:$HOME/.local/bin"

if [ -f $HOME/.poetry/env ]; then
  . $HOME/.poetry/env
fi

format_code() {
  black -S -l 79 .
}
