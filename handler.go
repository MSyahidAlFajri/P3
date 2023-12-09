package P3

import (
	"encoding/json"
	"net/http"
)

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

func InsertDataProduk(Mongoenv, dbname string, r *http.Request) string {
	resp := new(Credential)
	produkdata := new(Produk)
	resp.Status = false
	conn := SetConnection(Mongoenv, dbname)
	err := json.NewDecoder(r.Body).Decode(&produkdata)
	if err != nil {
		resp.Message = "error parsing application/json: " + err.Error()
	} else if produkdata.Nama == "" || produkdata.Harga == "" || produkdata.Deskripsi == "" || produkdata.Stok == "" {
		resp.Message = "Data Tidak Boleh Kosong"
	} else {
		resp.Status = true
		insertedID, err := InsertProduk(conn, "produk", *produkdata)
		if err != nil {
			resp.Message = "Gagal memasukkan data ke database: " + err.Error()
		} else {
			resp.Message = "Berhasil Input data dengan ID: " + insertedID.Hex()
		}
	}
	return GCFReturnStruct(resp)
}

func GetAllData(MONGOCONNSTRINGENV, dbname, collectionname string) string {
	mconn := SetConnection(MONGOCONNSTRINGENV, dbname)
	data := GetAllDataProduk(mconn, collectionname)
	return GCFReturnStruct(data)
}
