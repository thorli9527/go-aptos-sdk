// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package aptostypes

import (
	"encoding/json"
)

var _ = (*AccountCoreDataMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (a AccountCoreData) MarshalJSON() ([]byte, error) {
	type AccountCoreData struct {
		SequenceNumber    jsonUint64 `json:"sequence_number"`
		AuthenticationKey string     `json:"authentication_key"`
	}
	var enc AccountCoreData
	enc.SequenceNumber = jsonUint64(a.SequenceNumber)
	enc.AuthenticationKey = a.AuthenticationKey
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *AccountCoreData) UnmarshalJSON(input []byte) error {
	type AccountCoreData struct {
		SequenceNumber    *jsonUint64 `json:"sequence_number"`
		AuthenticationKey *string     `json:"authentication_key"`
	}
	var dec AccountCoreData
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.SequenceNumber != nil {
		a.SequenceNumber = uint64(*dec.SequenceNumber)
	}
	if dec.AuthenticationKey != nil {
		a.AuthenticationKey = *dec.AuthenticationKey
	}
	return nil
}