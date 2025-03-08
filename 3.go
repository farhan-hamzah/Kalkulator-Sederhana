package main
import "fmt"

func menu(pilih int){
	var  a,  b, hasil float64
	for pilih != 4{
		if pilih == 1{
			fmt.Print("Masukan dua bilangan yang akan dijumlahkan :")
			fmt.Scan(&a, &b)
			hasil = a+b
			fmt.Println("Hasil penjumlahan :", hasil)
		}else if pilih == 2{
			fmt.Print("Masukan dua bilangan yang akan dikalikan :")
			fmt.Scan(&a, &b)
			hasil = (a)*(b)
			fmt.Println("Hasil perkalian :", hasil)
		}else if pilih == 3{
			fmt.Print("Masukan dua bilangan yang akan dibagikan :")
			fmt.Scan(&a, &b)
			hasil = (a)/(b)
			fmt.Println("Hasil pembagian :", hasil)
		}
		fmt.Println("----------------------")
		fmt.Println("        MENU          ")
		fmt.Println("----------------------")
		fmt.Println("1. Hitung Penjumlahan")
		fmt.Println("2. Hitung Perkalian")
		fmt.Println("3. Hitung Pembagian")
		fmt.Println("4. Exit")
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
	}

}

func main(){
	var masukan int
	fmt.Println("----------------------")
	fmt.Println("        MENU          ")
	fmt.Println("----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Print("Pilih (1/2/3/4)? ")
	fmt.Scan(&masukan)
	menu(masukan)
}