# Cek Resi
Hanya kode yang saya gunakan untuk keperluan cek resi (tambah layanan kalau memang diperlukan) silakan buat issues jika ada yang ingin d tambahkan (lebih baik lagi diberikan juga contoh resi yang akan digunakan untuk testing)

Semoga bermanfaat

# Shopee

cek resi shopee xpress

```go
resi := "ID012706574768"
shopee, err := Shopee(resi)
if err != nil {
    panic(err)
}
fmt.Println(shopee)
```