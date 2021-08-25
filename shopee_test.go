package cekresi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShopeeXpress(t *testing.T) {
	resi := "ID012706574768"
	shopee, err := Shopee(resi)
	if err != nil {
		panic(err)
	}
	t.Run("success", func(t *testing.T) {
		assert.Equal(t, "Success", shopee.Message)
	})
	t.Run("retcode == 0", func(t *testing.T) {
		assert.Equal(t, 0, shopee.Retcode)
	})
	t.Run("resi == tracknumber", func(t *testing.T) {
		assert.Equal(t, resi, shopee.Data.SlsTrackingNumber)
	})
}
