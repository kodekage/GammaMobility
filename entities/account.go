package entities

type Account struct {
	Id              string  `db:"id"`
	CustomerId      string  `db:"customer_id"`
	Balance         float32 `db:"balance"`
	TotalAssetValue string  `db:"total_asset_value"`
	CreatedAt       string  `db:"created_at"`
}
