package repository

import "github.com/Celesca/isekai-shop-api/entities"

// ที่เป็น Interface เพราะ ถ้าสมมติเราจะสร้างเป็น Mock
type ItemShopRepository interface {
	Listing() ([]*entities.Item, error)
}
