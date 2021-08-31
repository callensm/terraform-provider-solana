# Basic Usage
data "solana_account" "acc" {
  public_key = "11111111111111111111111111111111"
}

# Specify Account Data Encoding
data "solana_account" "acc" {
  public_key = "11111111111111111111111111111111"
  encoding   = "jsonParsed"
}
