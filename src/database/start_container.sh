#!/bin/bash

# Start the specified Docker container
sudo docker start 11bb9fd2150b

# Keep the container running
tail -f /dev/null
