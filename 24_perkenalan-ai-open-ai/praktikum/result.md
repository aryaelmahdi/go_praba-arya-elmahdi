- Untuk menghandle API dari Midtrans dalam bahasa Go, Anda perlu menggunakan Midtrans Go Library. Pertama, pastikan Anda telah menginstal Midtrans Go Library dengan perintah berikut:

- go get github.com/midtrans/midtrans-go

- code 
package main

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func main() {
	// Setup konfigurasi Midtrans
	clientKey := "your-client-key"
	serverKey := "your-server-key"

	midclient := midtrans.NewClient()
	midclient.ServerKey = serverKey
	midclient.ClientKey = clientKey
	midclient.APIEnvType = midtrans.Sandbox

	// Inisialisasi charge request
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeCreditCard,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "order-123",
			GrossAmt: 200000, // Jumlah total pembayaran dalam rupiah
		},
		CreditCard: &coreapi.CreditCardDetail{
			TokenID: "your-credit-card-token",
		},
		CustomerDetails: midtrans.CustomerDetails{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "johndoe@example.com",
			Phone:     "081234567890",
		},
	}

	// Membuat charge (pembayaran)
	chargeResp, err := coreapi.ChargeTransaction(chargeReq, midclient)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Mendapatkan hasil pembayaran
	fmt.Println("Payment Status:", chargeResp.StatusMessage)
	fmt.Println("Payment ID:", chargeResp.TransactionID)
}


- Berikut adalah penjelasan rinci kode di atas:

1. Impor pustaka Midtrans dan pustaka yang diperlukan.
2. Atur konfigurasi dengan menggantikan "your-client-key" dan "your-server-key" dengan kunci-kunci Anda dari akun Midtrans Anda. Pastikan untuk mengatur midtrans.Sandbox sebagai lingkungan API jika Anda sedang menguji di lingkungan pengembangan.
3. Inisialisasi permintaan pembayaran (chargeReq) dengan rincian transaksi, jenis pembayaran, dan informasi pelanggan yang relevan. Token kartu kredit digunakan jika Anda ingin melakukan pembayaran kartu kredit.
4. Panggil fungsi coreapi.ChargeTransaction() untuk membuat pembayaran berdasarkan permintaan Anda.
5. Periksa status pembayaran dan tampilkan informasi yang relevan.

- Pastikan Anda telah menggantikan nilai-nilai seperti your-client-key, your-server-key, dan your-credit-card-token dengan nilai-nilai yang sesuai dari akun Midtrans Anda. Selain itu, pastikan Anda telah menyesuaikan detail transaksi sesuai dengan kebutuhan Anda.