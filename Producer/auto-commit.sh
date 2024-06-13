#!/bin/bash
set -e

echo $1 | echo `sudo -S cat autofile.txt`