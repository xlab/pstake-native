package types

import (
	"encoding/json"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WeightedAddressAmounts []WeightedAddressAmount

var _ sort.Interface = WeightedAddressAmounts{}

func NewWeightedAddressAmounts(ws []WeightedAddressAmount) (WeightedAddressAmounts) {
	w := WeightedAddressAmounts{}
	for _, element := range ws {
		w = append(w, element)
	}
	return w
}

func (w WeightedAddressAmounts) Len() int {
	return len(w)
}

func (w WeightedAddressAmounts) Less(i, j int) bool {
	return w[i].Coin.Amount.LT(w[j].Coin.Amount)
}

func (w WeightedAddressAmounts) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WeightedAddressAmounts) Sort() WeightedAddressAmounts {
	sort.Sort(w)
	return w
}

func (w WeightedAddressAmounts) Marshal() ([]byte, error) {
	if w == nil {
		return json.Marshal(WeightedAddressAmounts{})
	}
	return json.Marshal(w)
}

func (w WeightedAddressAmounts) Unmarshal(bz []byte) error {
	err := json.Unmarshal(bz, &w)
	if err != nil {
		return err
	}
	return nil
}

// TotalAmount returns the total amount for a given denom
func (w WeightedAddressAmounts) TotalAmount(denom string) sdk.Coin {
	total := sdk.NewCoin(denom, sdk.ZeroInt())
	
	for _, weightedAddr := range w {
		if weightedAddr.Coin.Denom == denom {
			total.Amount = total.Amount.Add(weightedAddr.Coin.Amount)
		}
	}
	return total
}