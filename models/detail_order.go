package models

import (
	"strconv"
	"time"
)

type DetailOrder struct {
	Id             int64 `gorm:"primary_key" json:"id"`
	BorrowDate  string `gorm:"type:date; not null" json:"borrow_date" binding:"required"`
	ReturnDate string `gorm:"type:date; not null" json:"return_date" binding:"required"`
	BorrowDuration     int `gorm:"_" json:"borrow_duration"`
	Total float64 `gorm:"double; not null" json:"total"`
	Guarantee string `gorm:"varchar(255); not null" json:"guarantee" binding:"required"`
}

func (d *DetailOrder) HitungDurasiSewa() {
    // Parse tanggal peminjaman dan tanggal pengembalian
    borrowDate, _ := time.Parse("2006-01-02", d.BorrowDate)
    returnDate, _ := time.Parse("2006-01-02", d.ReturnDate)
    
    // Hitung durasi sewa dalam hari
    durasi := returnDate.Sub(borrowDate).Hours() / 24
    
    // Simpan durasi sewa ke dalam properti guarantee
    d.Guarantee = strconv.Itoa(int(durasi))
}
