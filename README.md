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
`papertrailer <group_id>`

`<group_id>` is the id of the group to watch in Papertrail.

You can find this group id in the url on Papertrail.

![Group ID in URL](/groupid.png)
