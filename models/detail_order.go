package models

import(
	"time"
)

type DetailOrder struct {
	Id             int64 `gorm:"primary_key" json:"id"`
	TglPeminjaman  string `gorm:"type:date; not null" json:"tgl_peminjaman" binding:"required"`
	TglPengembalian string `gorm:"type:date; not null" json:"tgl_pengembalian" binding:"required"`
	DurasiSewa     int `gorm:"_" json:"durasi_sewa"`
	Total float64 `gorm:"double; not null" json:"total"`
	Jaminan string `gorm:"varchar(255); not null" json:"jaminan" binding:"required"`

}

func (d *DetailOrder) HitungDurasiSewa() {
	tglPeminjaman, _ := time.Parse("2006-01-02", d.TglPeminjaman)
	tglPengembalian, _ := time.Parse("2006-01-02", d.TglPengembalian)
	durasi := tglPengembalian.Sub(tglPeminjaman).Hours() / 24 // Hitung selisih hari
	d.DurasiSewa = int(durasi)
}
