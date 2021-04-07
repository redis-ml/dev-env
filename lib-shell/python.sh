export PATH="$PATH:$HOME/.local/bin"

if [ -f $HOME/.poetry/env ]; then
  . $HOME/.poetry/env
fi

function format_code() {
  black -S -l 79 .
}
