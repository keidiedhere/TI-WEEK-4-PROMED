package main
import "fmt"

func main () {
	//deklarasi variabel
	var nama_kustomer, nama_barang string
	var harga float32 = 88000
	var quantity int32
	var hasil_discount, total_harga float32

	//input nama kustomer
	fmt.Print("Nama kustomer: ")
	fmt.Scanf("%s\n", &nama_kustomer)

	//input nama barang
	fmt.Print("Nama barang: ")
	fmt.Scanf("%s\n", &nama_barang)

	//input quantity barang
	fmt.Print("Jumlah barang: ")
	fmt.Scanf("%d\n", &quantity)

	// kondisi diskon, kalau lebih dari 5 dapat 10% selain itu 2%
	if quantity > 5 {
		hasil_discount = 0.1 // 10%
	} else {
		hasil_discount = 0.02 // 2%
	}

	//proses
	sub_total := float32(quantity)*harga
	total_harga = sub_total - (hasil_discount*sub_total)

	//tampilkan hasil harga
	fmt.Println("hasil discount:", hasil_discount)
	fmt.Println("quantity: ", quantity)
	fmt.Println("harga: ", harga)
	fmt.Println("Total Harga adalah: ", total_harga)

}
