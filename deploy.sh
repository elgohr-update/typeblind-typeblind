#!/bin/bash
set -xe

if [ $TRAVIS_BRANCH == 'master' ] ; then
  eval "$(ssh-agent -s)"
  ssh-add ~/.ssh/id_rsa

  git init

  git remote add deploy "travis@kletskovg.tech:/var/www/kletskovg.tech"
  git config user.name "Travis CI"
  git config user.email "gfgfddglebrtravis@gmail.com"

  rm .gitignore
  git add client/dist/
  git add server/dist/

  git status

  git commit -m "Deploy"
  git push --force deploy master

else
  echo "Not deploying, since this branch isn't master."
fi
