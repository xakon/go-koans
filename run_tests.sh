#!/bin/sh

find -name '*.go' | entr go test
