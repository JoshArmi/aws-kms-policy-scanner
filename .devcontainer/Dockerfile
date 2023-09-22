FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# Install Granted
RUN curl -OL releases.commonfate.io/granted/v0.16.0/granted_0.16.0_linux_arm64.tar.gz && sudo tar -zxvf ./granted_0.16.0_linux_arm64.tar.gz -C /usr/local/bin/

RUN mkdir /home/vscode/.granted
COPY config /home/vscode/.granted/config
RUN chown -R vscode /home/vscode/.granted
RUN echo 'alias assume=". assume"' >> /home/vscode/.profile
