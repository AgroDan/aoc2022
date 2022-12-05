#!/usr/bin/fish
#
# This is a fish script to pull the input file of the specific day.
#
# In order to execute, run:
#
# set -x AOC_SESSION "<session key>"
# . ./prep.fish <day>
#
# where "key" is the session key in your logged-in cookie, and "day" is
# the day you want to prep.
#
# Credit where it is due, the concept of this script was inspired heavily by 0xdf's version.

set AOC_DIR "/home/dan/Documents/aoc2022"
set AOC_URL_BASE "https://adventofcode.com/2022/day"

# AOC Session is not set
if test -z "$AOC_SESSION"
    echo "AOC Session is not set!"
    echo "Please enter 'set -x AOC_SESSION <session key>"
    exit 1
end

if test -z "$argv[1]"
    echo "Need the day we are working with, please enter the advent day"
    echo "Example: ./prep.fish 5"
    exit 2
end

set WORKDIR "$AOC_DIR/day$argv[1]/"

mkdir -p $WORKDIR
cd $WORKDIR

set AOC_FULL "$AOC_URL_BASE/$argv[1]"

curl -s --cookie "session=$AOC_SESSION" "$AOC_FULL/input" > input