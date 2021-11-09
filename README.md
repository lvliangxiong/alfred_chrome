## Installation

### Build from source

```bash
go get -u github.com/jason0x43/go-alfred/alfred
alfred link
alfred build
```

### Install from release

Open release page and download the latest `.alfredworkflow` file，install by double click.

## Configuration

Every chrome account information resides in a seperate folder, and this wf needs information from this folder to function. So you need to configure a `Workflow Environment Variables` named `CHROME_PROFILE` to make results right.

To determine the profile folder for a running Chrome instance:

1. Navigate your chrome to `chrome://version`
2. Look for the `Profile Path` field. This gives the path to the profile directory
3. Set the `CHROME_PROFILE` to the last part of this field. Generally, it might be `Default`, `Profiel 1`, `Profile 2`

Another `LIMIT` is the maximum number of result items.

Three interval environment variables are used to control the frequency of syncing the latest bookmark, history and favicon. Valid values are duration strings. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
