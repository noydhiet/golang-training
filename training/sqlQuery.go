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
	idpeserta   int
	namapeserta string
	addrpeserta string
	agepeserta  int
}

type pesertaMataKuliah struct {
	namapeserta string
	mataKuliah  string
	agepeserta  int
}

// type Post struct {

// }

func main() {
	//sqlQuery()
	// db, err := connect()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	defer db.Close() //menjalankan semua logika baru lakukan proses defer
	// 	// fmt.Println("sukses")
	// }
	//queryGetPeserta()
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("nama peserta : ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(strings.Trim(str, "\n"))
	queryGetnamaPeserta(str)

}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang1")
	if err != nil {
		// if err := nil {
		return nil, err
	}
	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	rows, err := db.Query("select id_peserta, nama_peserta, addr_peserta, age_peserta from t_mtr_peserta")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []peserta

	for rows.Next() {
		var each = peserta{}
		var err = rows.Scan(&each.addrpeserta, &each.namapeserta, &each.addrpeserta, &each.agepeserta)

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
		fmt.Println(each.idpeserta, each.namapeserta, each.addrpeserta, each.agepeserta)
	}
}

func queryGetPeserta() {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var query string = `select nama_peserta, nama_matkul, age_peserta from 
						t_trx_matakuliah_peserta a,
						t_mtr_matakuliah b,
						t_mtr_peserta c where a.fk_peserta = c.id_peserta and 
						a.fk_matakuliah = b.id_matkul`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pesertaMataKuliah

	for rows.Next() {
		var each = pesertaMataKuliah{}
		var err = rows.Scan(&each.namapeserta, &each.mataKuliah, &each.agepeserta)

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
		fmt.Println(each.namapeserta, each.mataKuliah, each.agepeserta)
	}
}

func queryGetnamaPeserta(namepesertaInput string) {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}
	// namepesertaInput = strings.TrimSpace(strings.Trim(namepesertaInput, "\n"))
	// fmt.Println(namepesertaInput)

	defer db.Close()

	var query string = `select nama_peserta, nama_matkul, age_peserta from 
						t_trx_matakuliah_peserta a,
						t_mtr_matakuliah b,
						t_mtr_peserta c 
						where a.fk_peserta = c.id_peserta and 
						a.fk_matakuliah = b.id_matkul 
						and c.nama_peserta = ?`

	rows, err := db.Query(query, namepesertaInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []pesertaMataKuliah

	for rows.Next() {
		var each = pesertaMataKuliah{}
		var err = rows.Scan(&each.namapeserta, &each.mataKuliah, &each.agepeserta)

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
		fmt.Println(each.namapeserta, each.mataKuliah, each.agepeserta)
	}
}
