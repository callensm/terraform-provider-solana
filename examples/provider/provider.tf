# Derive Endpoint from Cluster
provider "solana" {
  cluster = "testnet"
}

# Set Custom RPC endpoint
provider "solana" {
  endpoint = "http://10.11.12.13:8899"
}

