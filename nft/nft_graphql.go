package nft

import (
	"fmt"

	"github.com/coming-chat/go-aptos/graphql"
)

const queryTokensFormat = `
query CurrentTokens {
	current_token_ownerships(
	  where: {owner_address: {_eq: "%v"}, table_type: {_eq: "0x3::token::TokenStore"}, amount: {_gt: "0"}}
	  order_by: {last_transaction_version: asc}
	) {
	  name
	  collection_name
	  property_version
	  amount
	  creator_address
	  last_transaction_timestamp
	  last_transaction_version
	  owner_address
	  current_token_data {
		metadata_uri
		description
	  }
	}
  }`

type GraphQLToken struct {
	Name                     string `json:"name"`
	CollectionName           string `json:"collection_name"`
	PropertyVersion          uint64 `json:"property_version"`
	Amount                   uint64 `json:"amount"`
	CreatorAddress           string `json:"creator_address"`
	LastTransactionTimestamp string `json:"last_transaction_timestamp"`
	LastTransactionVersion   uint64 `json:"last_transaction_version"`
	OwnerAddress             string `json:"owner_address"`
	CurrentTokenData         struct {
		MetadataUri string `json:"metadata_uri"`
		Description string `json:"description"`
	} `json:"current_token_data"`
}

// @param graphUrl Default is mainnet url if unspecified
func FetchGraphqlTokensOfOwner(owner string, graphUrl string) ([]GraphQLToken, error) {
	query := fmt.Sprintf(queryTokensFormat, owner)
	res := struct {
		Ownerships []GraphQLToken `json:"current_token_ownerships"`
	}{}
	err := graphql.FetchGraphQLSample(query, graphUrl, &res)
	if err != nil {
		return nil, err
	}
	return res.Ownerships, nil
}
