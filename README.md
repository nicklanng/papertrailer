# PaperTrailer
Faster grokking of event logs from PaperTrail

[![Build Status](https://travis-ci.org/nicklanng/papertrailer.svg?branch=master)](https://travis-ci.org/nicklanng/papertrailer)

## Installation
If you have golang installed, you can install papertrailer by running
```bash
go get github.com/nicklanng/papertrailer
```

Otherwise, you can click go to https://github.com/nicklanng/papertrailer/releases to get the latest builds.

## Usage


Run `papertrailer <group_id>`

This will create a config file on the first run.
Fill in the `token` with your Papertrail API token: https://help.papertrailapp.com/kb/how-it-works/http-api/#authentication

Fill in the `knownissuespath` with a path to a plain text file of strings to ignore

`<group_id>` is the id of the group to watch in Papertrail.

You can find this group id in the url on Papertrail.

![Group ID in URL](/groupid.png)
