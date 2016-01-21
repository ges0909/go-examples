#!/usr/bin/env bash

# A collection of .gitignore templates:
# - https://github.com/github/gitignore

# To be done before:
# - export GITHUB_USER=...
# - export GITHUB_PASS=...

#
# def
#
CMD="push"        # default git command
MSG="left empty"  # default commit message
URL="https://${GITHUB_USER}:${GITHUB_PASS}@github.com/ges0909/go-examples.git"

#
# func
#
function panic
{ # error handling
  echo "PANIC: ${1}" >&2 # stderr
  exit 1
}

function push
{ # workflow for push
  git add -A
  git commit -m "${MSG}"
  git push origin master
}

function init
{ # workflow for init
  rm -rf .git
  git init
  git remote add origin "${URL}"
  push
}

#
# parse options
#
OPTS=`getopt --options m: --longoptions message: --name "$0" -- "$@"`

# check if 'getopt' suceeded
if (( "${?}" != 0 )); then
  panic "option error"
fi

eval set -- "${OPTS}";

# store option args to vars
while true; do
  case "${1}" in
    -m|--message) MSG="${2}"; shift 2 ;;
    --) shift; break ;;
    *) panic "unknown option error" ;;
  esac
done

# store remaining args (command)
if (( "${#@}" == 1 )); then
  CMD="${1}" # overwrite default
fi

# check command
case "${CMD}" in
  init | push ) ;;
  * ) panic "unknown command '${CMD}', only 'init' or 'push' is allowed" ;;
esac

#
# run command
#
case "${CMD}" in
  init) init ;;
  push) push ;;
esac
