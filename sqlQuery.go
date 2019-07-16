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
	namaPeserta    string
	namaMataKuliah string
	agePeserta     int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("masukan nama")
	inputString, _ := reader.ReadString('\n')
	inputString = strings.Trim(inputString, "\n")
	inputString = strings.TrimSpace(inputString)
	queryGetPesertaNama(inputString)
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1)/cel_training_db")
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

	defer db.Close()
	var query string = `select name_peserta, nama_matakuliah, age_peserta from
						t_trx_matakuliah_peserta a,
						t_mtr_matakuliah b,
						t_mtr_peserta c
						where a.fk_peserta = c.id_peserta and
						a.fk_matakuliah = b.id_matakuliah`

	rows, err := db.Query(query)
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

func queryGetPesertaMataKuliah() {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	var query string = `select name_peserta, nama_matakuliah, age_peserta from
						t_trx_matakuliah_peserta a,
						t_mtr_matakuliah b,
						t_mtr_peserta c
						where a.fk_peserta = c.id_peserta and
						a.fk_matakuliah = b.id_matakuliah`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []pesertaMataKuliah

	for rows.Next() {
		var each = pesertaMataKuliah{}
		var err = rows.Scan(&each.namaPeserta, &each.namaMataKuliah, &each.agePeserta)

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
		fmt.Println(each.namaPeserta, each.namaMataKuliah, each.agePeserta)
	}

}

func scanner() {

}

func queryGetPesertaNama(str string) {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	var query string = `select name_peserta, nama_matakuliah, age_peserta from
						t_trx_matakuliah_peserta a,
						t_mtr_matakuliah b,
						t_mtr_peserta c
						where a.fk_peserta = c.id_peserta and
						a.fk_matakuliah = b.id_matakuliah and nama_peserta=?`

	rows, err := db.Query(query, str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var result []pesertaMataKuliah

	for rows.Next() {
		var each = pesertaMataKuliah{}
		var err = rows.Scan(&each.namaPeserta, &each.namaMataKuliah, &each.agePeserta)

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
		fmt.Println(each.namaPeserta, each.namaMataKuliah, each.agePeserta)
	}

}
