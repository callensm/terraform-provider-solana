package solana

import (
	"context"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	associatedtokenaccount "github.com/gagliardetto/solana-go/programs/associated-token-account"
	token "github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAssociatedTokenAccount() *schema.Resource {
	return &schema.Resource{
		Description: "Create a new associated token account for a specified wallet address.",

		Create: resourceAssociatedTokenAccountCreate,
		Read:   resourceAssociatedTokenAccountRead,
		Delete: resourceAssociatedTokenAccountDelete,

		Importer: &schema.ResourceImporter{
			StateContext: resourceAssociatedTokenAccountImporter,
		},

		Schema: map[string]*schema.Schema{
			"owner": {
				Type:        schema.TypeString,
				Description: "Public key of the wallet to create the token account for as a base-58 encoded string.",
				Required:    true,
				ForceNew:    true,
			},
			"token_mint": {
				Type:        schema.TypeString,
				Description: "The token mint for the associated token account as a base-58 encoded string.",
				Required:    true,
				ForceNew:    true,
			},
			"payer": {
				Type:        schema.TypeString,
				Description: "Public key of the transaction payer as a base-58 encoded string.",
				Optional:    true,
				ForceNew:    true,
			},
			"public_key": {
				Type:        schema.TypeString,
				Description: "The base-58 encoded string for the new token account's public key.",
				Computed:    true,
			},
		},
	}
}

func resourceAssociatedTokenAccountCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	owner, err := solana.PublicKeyFromBase58(d.Get("owner").(string))
	if err != nil {
		return err
	}

	tokenMint, err := solana.PublicKeyFromBase58(d.Get("token_mint").(string))
	if err != nil {
		return err
	}

	payer := owner
	if p, ok := d.GetOk("payer"); ok {
		pubkey, err := solana.PublicKeyFromBase58(p.(string))
		if err != nil {
			return err
		}

		payer = pubkey
	}

	inst := associatedtokenaccount.NewCreateInstruction(payer, owner, tokenMint)
	recentHash, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentConfirmed)
	if err != nil {
		return err
	}

	tx, err := solana.NewTransaction([]solana.Instruction{inst.Build()}, recentHash.Value.Blockhash)
	if err != nil {
		return err
	}

	_, err = client.SendTransactionWithOpts(context.Background(), tx, rpc.TransactionOpts{})
	if err != nil {
		return err
	}

	res, err := client.GetTokenAccountsByOwner(
		context.Background(),
		owner,
		&rpc.GetTokenAccountsConfig{
			Mint:      &tokenMint,
			ProgramId: &associatedtokenaccount.ProgramID,
		},
		&rpc.GetTokenAccountsOpts{
			Commitment: rpc.CommitmentConfirmed,
		},
	)
	if err != nil {
		return err
	} else if len(res.Value) == 0 {
		return fmt.Errorf("the new associated token account was not found in the confirmed accounts")
	}

	tokenAccountPubkey := res.Value[0].Pubkey.String()

	d.SetId(fmt.Sprintf("associated_token_account:%s", tokenAccountPubkey))
	d.Set("payer", payer.String())
	d.Set("public_key", tokenAccountPubkey)
	return nil
}

func resourceAssociatedTokenAccountRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	owner, err := solana.PublicKeyFromBase58(d.Get("owner").(string))
	if err != nil {
		return err
	}

	tokenMint, err := solana.PublicKeyFromBase58(d.Get("token_mint").(string))
	if err != nil {
		return err
	}

	res, err := client.GetTokenAccountsByOwner(
		context.Background(),
		owner,
		&rpc.GetTokenAccountsConfig{
			Mint:      &tokenMint,
			ProgramId: &associatedtokenaccount.ProgramID,
		},
		&rpc.GetTokenAccountsOpts{
			Commitment: rpc.CommitmentConfirmed,
		},
	)
	if err != nil {
		return err
	} else if len(res.Value) == 0 {
		return fmt.Errorf("the associated token account was not found in the confirmed accounts for owner")
	}

	d.Set("public_key", res.Value[0].Pubkey.String())

	return nil
}

func resourceAssociatedTokenAccountDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*providerConfig).rpcClient

	tokenAccount, err := solana.PublicKeyFromBase58(d.Get("public_key").(string))
	if err != nil {
		return err
	}

	owner, err := solana.PublicKeyFromBase58(d.Get("owner").(string))
	if err != nil {
		return err
	}

	payer, err := solana.PublicKeyFromBase58(d.Get("payer").(string))
	if err != nil {
		return err
	}

	inst := token.NewCloseAccountInstruction(tokenAccount, payer, owner, nil)

	recentHash, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentConfirmed)
	if err != nil {
		return err
	}

	tx, err := solana.NewTransaction([]solana.Instruction{inst.Build()}, recentHash.Value.Blockhash)
	if err != nil {
		return err
	}

	_, err = client.SendTransactionWithOpts(context.Background(), tx, rpc.TransactionOpts{})
	if err != nil {
		return err
	}

	return nil
}

func resourceAssociatedTokenAccountImporter(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*providerConfig).rpcClient
	idParts := strings.Split(d.Id(), "/")

	owner, err := solana.PublicKeyFromBase58(idParts[0])
	if err != nil {
		return nil, err
	}

	tokenMint, err := solana.PublicKeyFromBase58(idParts[1])
	if err != nil {
		return nil, err
	}

	res, err := client.GetTokenAccountsByOwner(
		ctx,
		owner,
		&rpc.GetTokenAccountsConfig{
			Mint:      &tokenMint,
			ProgramId: &associatedtokenaccount.ProgramID,
		},
		&rpc.GetTokenAccountsOpts{
			Commitment: rpc.CommitmentConfirmed,
		},
	)
	if err != nil {
		return nil, err
	} else if len(res.Value) == 0 {
		return nil, fmt.Errorf("the associated token account was not found in the confirmed accounts for owner")
	}

	d.SetId(fmt.Sprintf("associated_token_account:%s", res.Value[0].Pubkey.String()))
	d.Set("owner", owner.String())
	d.Set("payer", owner.String())
	d.Set("public_key", res.Value[0].Pubkey.String())
	d.Set("token_mint", tokenMint.String())

	return []*schema.ResourceData{d}, nil
}
