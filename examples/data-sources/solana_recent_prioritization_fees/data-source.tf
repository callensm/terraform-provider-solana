# Basic Usage
data "solana_recent_prioritization_fees" "testnet" {}

# Fees for Specific Addresses
data "solana_recent_prioritization_fees" "addrs" {
  accounts = ["..."]
}
