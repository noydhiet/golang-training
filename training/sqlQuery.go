package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type peserta struct {
	idPeserta   int
	namePeserta string
	addrPeserta string
	agePeserta  int
}

type pesertaMataKuliah struct {
	namePeserta string
	mataKuliah  string
	agePeserta  int
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ID Peserta : ")
	str, _ := reader.ReadString('\n')

	str = strings.Trim(str, "\n")
	str = strings.TrimSpace(str)

	queryGetPesertaMataKuliah(str)
	// sqlQuery()

	// db, err := connect()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	/*keyword 'defer' digunakan untuk mengeksekusi semua baris code
	// 	  terlebih dahulu baru mengeksekusi dirinya sendiri*/
	// 	defer db.Close()
	// 	fmt.Println("Db connection success!")
	// }
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/cel_training_db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select id_peserta, name_peserta, addr_peserta, age_peserta from t_mtr_peserta")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []peserta

	for rows.Next() {
		var each = peserta{}
		var err = rows.Scan(&each.idPeserta, &each.namePeserta, &each.addrPeserta, &each.agePeserta)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.idPeserta, each.namePeserta, each.addrPeserta, each.agePeserta)
	}
}

func queryGetPesertaMataKuliah(namePesertaInput string) {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}

	// namePesertaInput = strings.Trim(namePesertaInput, "\n")

	defer db.Close()

	var query string = `select name_peserta, nama_matakuliah, age_peserta from
						t_trx_matakuliah_peserta a,
						t_mtr_mata_kuliah b,
						t_mtr_peserta c
						where a.fk_peserta = c.id_peserta and
						a.fk_matakuliah = b.id_matakuliah
						and c.name_peserta = ?`

	rows, err := db.Query(query, namePesertaInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pesertaMataKuliah

	for rows.Next() {
		var each = pesertaMataKuliah{}
		var err = rows.Scan(&each.namePeserta, &each.mataKuliah, &each.agePeserta)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.namePeserta, each.mataKuliah, each.agePeserta)
	}
}
