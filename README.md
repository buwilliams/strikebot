# Strikebot
Grab, transform, and deliver data.

## Overview

Strikebot is a `iPaaS` solution. It follows the Strategy Pattern to
execute a series of `Steps`. These steps determine which `Adaptors` to
invoke and what `Config` should be passed to them. The adaptors in
turn do some `Job` for you. You can think of it as the linux
command-line on steriods for the cloud.

For example, you might want to:

1. Download some JSON from FTP
2. Transform said JSON
3. Convert said JSON to CSV
4. Import CSV to a Google Sheet
5. Upload CSV to AWS S3 bucket

Strikebot allows you the ability to decide what will happen by following the
instructions you supply in a config JSON file.

## Project Goals

- Use Strategy Pattern as a development guide
- Robust library of useful adaptors
- HTTP Server invocation
- Command-line invocation
- Performant
- Run on Heroku
- Require only a JSON configuration file to run
- Allow for easy adaptor creation with a simple interface
- Log all processing

## Backlog

### Todo

- begin to do something useful
- validate map.json against schema

### Done

- run: `strikebot ./testdata/map.json` which executes `print` adaptor
- display usage message if used incorrectly
- handle unknown json

