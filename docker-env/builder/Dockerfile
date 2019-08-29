# Image for building prod.

FROM build-base:latest

# Note 'HOME' is not present by default.
ENV HOME /root

# Used by pyenv and rbenv
ENV PYENV_ROOT="${HOME}/.pyenv"
ENV RBENV_ROOT="${HOME}/.rbenv"
ENV PATH="${PYENV_ROOT}/bin:${RBENV_ROOT}/bin:${RBENV_ROOT}/shims:${PATH}"

# Install pyenv
RUN curl -L https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash
# Install rbenv
RUN curl -fsSL https://github.com/rbenv/rbenv-installer/raw/master/bin/rbenv-installer | bash

CMD ["/bin/bash"]
