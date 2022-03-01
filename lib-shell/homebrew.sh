
function switch_to_homebrew_bash() {
  export SHELL="$(brew --prefix)/bin/bash"
}

function homebrew_services_log() {
  find /usr/local/var/log/
}
