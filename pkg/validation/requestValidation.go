package validation

import (
	"sync"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type (
	EchoRequest interface {
		Bind(obj any) error
	}

	customEchoRequest struct {
		ctx       echo.Context
		validator *validator.Validate
	}
)

// Singleton pattern for validator instance
var (
	once              sync.Once
	validatorInstance *validator.Validate
)

// Constructor for customEchoRequest

func NewCustomEchoRequest(echoRequest echo.Context) EchoRequest {
	once.Do(func() {
		validatorInstance = validator.New()
	})

	return &customEchoRequest{
		ctx:       echoRequest,
		validator: validatorInstance,
	}
}

// ถ้าเราต้องประกาศ validate ทุกตัว struct ที่เราใช้ จะทำให้เราเสียเวลาในการเขียนโค้ดเพิ่มเติม
// จึงสร้างฟังก์ชัน Bind ที่รับ obj ใดๆ และทำการ bind ข้อมูลจาก request และ validate ข้อมูลด้วย

func (r *customEchoRequest) Bind(obj any) error {
	if err := r.ctx.Bind(obj); err != nil {
		return err
	}

	if err := r.validator.Struct(obj); err != nil {
		return err
	}

	return nil
}
