# Strikebot
Grab, transform, and deliver data.

## Overview

Strikebot follows the Strategy Pattern to execute a series of steps. These steps
contain information about which adaptors to invoke. The adapotors in turn do
some of job for you.

For example, you might want to:
1. Download some JSON from FTP
2. Transform said JSON
3. Convert said JSON to CSV
4. Upload CSV to Google Sheets

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

### Todo

- begin to do something useful
- validate map.json against schema

### Done

- run: `strikebot ./testdata/map.json` which executes `print` adaptor
- display usage message if used incorrectly
- handle unknown json

