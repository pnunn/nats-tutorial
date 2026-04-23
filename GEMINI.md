# NATS Cluster Configuration Project

## Project Overview
This project contains configuration files for setting up a NATS cluster named `MDC_Cluster`. The cluster consists of three nodes (`mdc_nats1`, `mdc_nats2`, and `mdc_nats3`) with JetStream enabled.

### Key Technologies
- **NATS Server**: A high-performance messaging system.
- **JetStream**: NATS's persistence layer for streaming and message queuing.
- **Clustering**: Configured for high availability and scalability.

## Directory Overview
The directory contains three primary configuration files, each corresponding to a node in the NATS cluster.

### Key Files
- `nats_server1.conf`: Configuration for server `mdc_nats1`.
- `nats_server2.conf`: Configuration for server `mdc_nats2`.
- `nats_server3.conf`: Configuration for server `mdc_nats3`.

## Configuration Details
Each server configuration includes:
- **Server Name**: Unique identifier within the cluster.
- **Listen Address**: Default client port `4222`.
- **Accounts**: A `$SYS` account is configured with an `admin` user.
- **JetStream**: Enabled with storage at `/tmp/nats/`.
- **Cluster**:
  - Name: `MDC_Cluster`
  - Listen: Port `6222` for cluster communication.
  - Routes: Explicit routes defined to other cluster members via IP addresses (Note: Some configurations appear to have typos in the cluster ports, e.g., `:622` instead of `:6222`).

## Usage

### Running the Servers
To start a NATS server with one of these configurations, use the `nats-server` binary:

```bash
# To start node 1
nats-server -c nats_server1.conf

# To start node 2
nats-server -c nats_server2.conf

# To start node 3
nats-server -c nats_server3.conf
```

### Testing the Cluster
Once the servers are running, you can verify cluster connectivity using the NATS CLI or by checking the monitoring endpoints (if enabled in the config, though not explicitly present in these files).

## Go Application
A Go project is included to demonstrate JetStream publishing and consuming.

### Structure
- `common/`: Shared configuration and connection logic.
- `publisher/`: A simple JetStream publisher that sends messages to `foo.test`.
- `subscriber/`: A JetStream pull subscriber that consumes and acknowledges messages from `my-pull-consumer` in `test-stream`.

### Running the Go Apps
Make sure you have Go installed and the NATS cluster is running.

1.  **Configure Environment**: Create a `.env` file in the root directory (based on the provided template) to store your NATS URLs and credentials.
    ```env
    NATS_SERVERS=nats://10.144.144.2:4222,nats://10.144.144.3:4222,nats://10.144.144.4:4222
    NATS_USER=clientuser
    NATS_PASSWORD=clientpassword
    NATS_LEAF_SERVER=nats://localhost:4223
    ```

2.  **Install dependencies**:
    ```bash
    go mod tidy
    ```

3.  **Run applications**:
    ```bash
    # To run the hub
    go run hub/main.go

    # To run the leaf app (on the leaf node VM)
    go run leaf_app/main.go
    ```

## TODO / Observations
- **Port Inconsistencies**: Several cluster routes use port `622` instead of the expected `6222`. These should be verified and corrected if they prevent cluster formation. (Fixed)
- **IP Addresses**: The configurations use static IP addresses in the `10.144.144.x` range. Ensure these match the environment where the servers are deployed.
- **Security**: The `admin` user has a bcrypt hashed password. Ensure appropriate secret management in production.
- **Environment Configuration**: The Go client now uses a `.env` file for configuration. Ensure the `.env` file is not committed to source control in a production environment (add it to `.gitignore`).
