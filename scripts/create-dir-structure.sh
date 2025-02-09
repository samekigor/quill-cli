#!/bin/bash

#
# This script data structure for quill
#


# Variables
GROUP="quill"
USER=$(whoami)

# Create quill group
if ! getent group $GROUP > /dev/null; then
    sudo groupadd $GROUP
    echo "Grupa $GROUP została utworzona."
fi

# Add user and root to quill group

sudo usermod -aG $GROUP $USER
sudo usermod -aG $GROUP root


# Credentials file - it stores registry credetials

REGISTRY_CREDITS_FILE_PATH="/etc/quill/credentials.yml"

sudo touch $REGISTRY_CREDITS_FILE_PATH
sudo chown $USER:$GROUP $REGISTRY_CREDITS_FILE_PATH
sudo chmod 664 $REGISTRY_CREDITS_FILE_PATH
sudo chmod -x $REGISTRY_CREDITS_FILE_PATH

#