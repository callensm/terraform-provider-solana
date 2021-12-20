# Basic Usage
resource "solana_associated_token_account" "acc" {
  owner      = "<base-58 pubkey>"
  token_mint = "<base-58 pubkey>"
}

# Custom Payer Wallet
resource "solana_associated_token_account" "acc" {
  owner      = "<base-58 pubkey>"
  payer      = "<base-58 pubkey>"
  token_mint = "<base-58 pubkey>"
}
