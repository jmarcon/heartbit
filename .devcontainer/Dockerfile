FROM mcr.microsoft.com/devcontainers/base:noble 

ARG BUILDPLATFORM
ARG TARGETPLATFORM
ARG TARGETARCH
ARG TARGETVARIANT
ARG RANDOM=$RANDOM
ENV container docker

RUN echo "I am running on $BUILDPLATFORM, building for $TARGETPLATFORM"

# Create a folder to put binary apps
ENV PATH="/home/vscode/.apps:${PATH}"

RUN wget -qO- https://raw.githubusercontent.com/eza-community/eza/main/deb.asc  | sudo gpg --dearmor -o /etc/apt/keyrings/gierens.gpg
RUN echo "deb [signed-by=/etc/apt/keyrings/gierens.gpg] http://deb.gierens.de stable main" | sudo tee /etc/apt/sources.list.d/gierens.list
RUN sudo chmod 644 /etc/apt/keyrings/gierens.gpg /etc/apt/sources.list.d/gierens.list

RUN if [ "$TARGETARCH" = "amd64" ]; then \
  sudo gpg -k; \
  sudo gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80  --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69 ; \
  echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list  ; \
  apt-get update; \
  apt-get install -y k6; \
  fi

RUN apt-get update && apt-get upgrade -y && apt-get dist-upgrade -y

## Installing the tools from the base image
RUN apt-get install -y \
  eza gnupg gcc gpg procps \
  apt-transport-https ca-certificates openssl \
  lsb-release build-essential 

RUN apt-get install -y software-properties-common

RUN add-apt-repository ppa:deadsnakes/ppa -y ; \
  apt-get update;

RUN apt-get install -y \
  dotnet-sdk-8.0 \
  dotnet-runtime-8.0 \
  aspnetcore-runtime-8.0 

RUN apt-get install -y \
  git git-delta git-extras git-flow

RUN apt-get install -y \
  snapd squashfuse fuse \
  zsh zsh-syntax-highlighting \
  ack btop duf jq tldr zoxide \
  gh unzip curl wget wrk tmux \
  fd-find eza bat ripgrep \
  file lua5.4 aptitude fortunes-br cowsay 

RUN apt-get install -y \
  python3.11 python3-pip python3-venv python3-dev

## Install fzf from git source | The package manager is not up to date
RUN git clone --depth 1 https://github.com/junegunn/fzf.git ~/.apps/.fzf
RUN ~/.apps/.fzf/install --key-bindings --completion --no-update-rc  --no-bash --no-fish --bin

RUN if [ "$TARGETARCH" = "arm64" ]; then \
  apt-get install -y \
  neovim \
  ;   fi

RUN if [ "$TARGETARCH" = "amd64" ] || [ "$TARGETARCH" = "arm/v8" ]; then \
  apt-get install -y \
  neovim plantuml \
  nodejs npm httpie \
  neofetch \
  ; fi

RUN apt-get autoremove -y && \
  apt-get clean -y

# Set zsh as default
RUN chsh -s /usr/bin/zsh vscode

## SET ZSH AS DEFAULT SHELL
RUN rm -rf /home/vscode/.oh-my-zsh

USER vscode
RUN mkdir -p /home/vscode/.apps

# # Install Homebrew
# RUN bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# # # Putting the themes in the same place as the homebrew uses
RUN curl -s https://ohmyposh.dev/install.sh | bash -s -- -d /home/vscode/.apps -t /home/vscode/.oh-my-posh/themes
ENV POSH_THEMES_PATH="/home/vscode/.oh-my-posh/themes"
ENV POSH_THEME="$POSH_THEMES_PATH/craver.omp.json"
RUN oh-my-posh disable notice

## Install Powershell
# WORKDIR /tmp
# RUN sudo snap update
# RUN sudo snap install powershell --classic

# Instal kubectl
WORKDIR /home/vscode/.apps
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/$TARGETARCH/kubectl"

WORKDIR /home/vscode
RUN echo ${RANDOM}
RUN git clone https://github.com/jmarcon/dotfiles ~/.apps/dotfiles
RUN  ~/.apps/dotfiles/Install/dotfiles.zsh 
RUN  ~/.apps/dotfiles/Install/git_repos.zsh

## Neovim Config
RUN ln -s /home/vscode/.apps/dotfiles/Config/nvim/Config/ ~/.config/nvim 
RUN ln -s /home/vscode/.apps/dotfiles/Config/nvim/.vimrc ~/.vimrc

ENTRYPOINT [ "zsh" ]
