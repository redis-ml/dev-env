#Installation

1. Install Docker

```bash
set -ex
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
```

1. Install OpenVPN server.

```bash
sudo bash install-openvpn-docker.sh
```

1. Start OpenVPN server.

```bash
# Usage: bash start-openvpn-docker.sh [server_domain]
sudo bash start-openvpn-docker.sh
```

1. Generate client config file.

```bash
# Usage bash gen-openvpn-client-cert.sh [client_name]
sudo bash gen-openvpn-client-cert.sh
```
Copy [clientname].ovpn to local and use in OpenVPN client.

