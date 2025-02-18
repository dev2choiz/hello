#!/bin/bash

if [ "" == "$TAG" ];then
  echo "Usage: Missing TAG argument (TAG=v1.0.1 make tag)"
  exit 1
fi

echo "delete tag $TAG"
git tag -d $TAG &> /dev/null || true
git push --delete origin $TAG &> /dev/null || true

echo "git tag $TAG"
git tag -a $TAG -m "Release $TAG"
git push origin $TAG
