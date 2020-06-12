# Skyput

Upload files to [Skynet](https://siasky.net) from your command line.

![Demo!](skyput_demo.gif)

## Install

### Go

```bash
go get -u github.com/termoose/skyput
```

### Homebrew

```bash
brew install termoose/tap/skyput
```

## Usage

On the first run a config file will be created for you in `~/.config/skyput`.
If you want to change the default Skynet portal you can run:
```bash
skyput -portal
```

You can modify the config file and add custom portals as well!

Start uploading to Skynet:

```bash
skyput cat_picture.jpg
```

## To-do
- Add support for pushing directories
- Support encryption when it arrives
- Add some timeout when requests are hanging
- Support resuming uploads
- Add support /skynet/portals [GET]?
