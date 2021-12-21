# lancaster

UDP multicast LAN IP as a device finder service.

Your NUC/NAS/HDPC runs the lanc servie via systemd.
Your laptap runs the lans server to receive IP address of your device via UDP multicast.

## Install

* `go install github.com/Akagi201/lancaster/cmd/lanc@latest`
* `go install github.com/Akagi201/lancaster/cmd/lans@latest`

## Usage

* lanc: Client side, Send UDP broadcast Message.
* lans: Server side, Receive UDP broadcast Message.
