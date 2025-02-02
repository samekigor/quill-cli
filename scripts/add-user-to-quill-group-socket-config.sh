#!/bin/bash

# sudo groupadd quill #add

sudo touch /var/run/quill.sock

sudo chmod 660 /var/run/quill.sock

sudo usermod -aG quill $USER

newgrp quill  # Aktywacja zmian w sesji

echo "logout and login"