#!/bin/bash

# variables
REPO_URL="https://github.com/alexrondon89/furryfam.git"
DEST_DIR="./tmp"
DEPLOYMENT_PATH="./furryfam/infrastructure/deployments"
SCRIPT_LOCATION="$DEST_DIR/deployments/scripts/create_jenkins_container.sh"

# download the repository
echo "cloning repository $REPO_URL ..."
git clone "$REPO_URL"

#to check if cloning is ok
if [ $? -ne 0 ]; then
  echo "error downloading repository $REPO_URL"
  exit 1
fi

# create the destination directory if it doesn't exist
echo "creating directory $DEST_DIR ..."
mkdir -p "$DEST_DIR"

# moving deployment folder to new location
echo "moving folder $DEPLOYMENT_PATH to $DEST_DIR ..."
mv "$DEPLOYMENT_PATH" "$DEST_DIR"

# verify if the movement of documents was ok
if [ $? -ne 0 ]; then
  echo "error moving files from $DEPLOYMENT_PATH to $DEST_DIR"
  exit 1
fi

# to give rights to execute the sh script
chmod +x "$SCRIPT_LOCATION"

# executing script file
echo "executing script file $SCRIPT_LOCATION ..."
"$SCRIPT_LOCATION"

# to check if script execution was ok
if [ $? -ne 0 ]; then
  echo "error execution script $SCRIPT_LOCATION"
  exit 1
fi

echo "initial script execute successfully...."
