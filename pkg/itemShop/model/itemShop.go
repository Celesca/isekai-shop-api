package model

// Model คือสิ่งที่ผู้ใช้เห็น ไม่จำเป็นต้องมี UpdatedAt ตาม entities item.go

type (
	Item struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Picture     string `json:"picture"`
		Price       uint   `json:"price"`
	}

	ItemFilter struct {
		Name        string `query:"name" validate:"omitempty,max=64"`
		Description string `query:"description" validate="omitempty,max=128"`
	}
)
