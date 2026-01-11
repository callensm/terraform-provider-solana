provider "solana" {
  endpoint = "https://api.testnet.solana.com"
}

data "solana_genesis_hash" "gen" {}
