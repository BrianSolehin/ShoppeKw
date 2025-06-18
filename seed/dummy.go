package seed

// import (
// 	"ecommerce-api/config"
// 	"ecommerce-api/model"
// 	"fmt"
// 	"time"
// )

// func InsertDummyData() {
// 	// ✅ Dummy user
// 	user := model.User{
// 		Name:     "Fufufafa",
// 		Email:    "fufufafa@example.com",
// 		Password: "123456",
// 		Role:     "customer",
// 	}
// 	if err := config.DB.Create(&user).Error; err != nil {
// 		fmt.Println("❌ Gagal insert user:", err)
// 		return
// 	}
// 	fmt.Println("✅ User ID:", user.ID)

// 	// ✅ Dummy category
// 	category := model.Category{
// 		Name: "Elektronik",
// 	}
// 	if err := config.DB.Create(&category).Error; err != nil {
// 		fmt.Println("❌ Gagal insert category:", err)
// 		return
// 	}
// 	fmt.Println("✅ Category ID:", category.ID)

// 	// ✅ Dummy product
// 	product := model.Product{
// 		Name:        "Laptop ASUS",
// 		Description: "Laptop gaming 16GB RAM",
// 		Price:       15000000,
// 		Stock:       10,
// 		CategoryID:  category.ID,
// 		UserID:      user.ID,
// 	}
// 	if err := config.DB.Create(&product).Error; err != nil {
// 		fmt.Println("❌ Gagal insert product:", err)
// 		return
// 	}
// 	fmt.Println("✅ Product ID:", product.ID)

// 	// ✅ Dummy payment method
// 	payment := model.PaymentMethod{
// 		Name: "QRIS",
// 	}
// 	if err := config.DB.Create(&payment).Error; err != nil {
// 		fmt.Println("❌ Gagal insert payment method:", err)
// 		return
// 	}
// 	fmt.Println("✅ PaymentMethod ID:", payment.ID)

// 	// ✅ Dummy order
// 	order := model.Order{
// 		UserID:      user.ID,
// 		OrderDate:   time.Now(),
// 		Status:      "completed",
// 		TotalAmount: product.Price * 2,
// 	}
// 	if err := config.DB.Create(&order).Error; err != nil {
// 		fmt.Println("❌ Gagal insert order:", err)
// 		return
// 	}
// 	fmt.Println("✅ Order ID:", order.ID)

// 	// ✅ Dummy order detail
// 	orderDetail := model.OrderDetail{
// 		OrderID:   order.ID,
// 		ProductID: product.ID,
// 		Quantity:  2,
// 		Price:     product.Price,
// 	}
// 	if err := config.DB.Create(&orderDetail).Error; err != nil {
// 		fmt.Println("❌ Gagal insert order detail:", err)
// 		return
// 	}
// 	fmt.Println("✅ OrderDetail masuk dengan OrderID:", orderDetail.OrderID)

// 	// ✅ Dummy transaction
// 	transaction := model.Transaction{
// 		OrderID:         order.ID,
// 		PaymentMethodID: payment.ID,
// 		Status:          "success",
// 		TransactionDate: time.Now(),
// 	}
// 	if err := config.DB.Create(&transaction).Error; err != nil {
// 		fmt.Println("❌ Gagal insert transaction:", err)
// 		return
// 	}
// 	fmt.Println("✅ Transaction ID:", transaction.ID)

// 	fmt.Println("🎉 Semua dummy data berhasil disimpan ke database!")
// }
