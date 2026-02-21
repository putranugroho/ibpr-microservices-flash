package service

import (
	"bytes"
	"io"
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"
	"time"
)

type BalanceRequest struct {
	NoRek    string `json:"no_rek"`
	BprID    string `json:"bpr_id"`
	TrxCode  string `json:"trx_code"`
	TrxType  string `json:"trx_type"`
	TglTrans string `json:"tgl_trans"`
	RRN      string `json:"rrn"`
}

func CheckBalance(db *sql.DB, req BalanceRequest) (map[string]interface{}, error) {


	fmt.Println("REQUEST:", req)

	var gateway string

	err := db.QueryRow(
		`SELECT gateway FROM kd_bpr WHERE bpr_id=$1 AND status='1'`,
		req.BprID,
	).Scan(&gateway)


	fmt.Println("GATEWAY:", gateway)

	if err == sql.ErrNoRows {
		return map[string]interface{}{
			"code":    "002",
			"status":  "Failed",
			"message": "Gagal, Inquiry BPR Tidak Ditemukan",
			"data":    []interface{}{},
		}, nil
	}

	payload := map[string]interface{}{
		"bpr_id":       req.BprID,
		"no_rek":       req.NoRek,
		"trx_code":     req.TrxCode,
		"trx_type":     req.TrxType,
		"tgl_trans":    req.TglTrans,
		"tgl_transmis": time.Now().Format("060102150405"),
		"rrn":          req.RRN,
	}

	body, _ := json.Marshal(payload)

	client := &http.Client{
	Timeout: 10 * time.Second,
}

reqHttp, _ := http.NewRequest(
	"POST",
	gateway+"gateway_bpr/inquiry_account",
	bytes.NewBuffer(body),
)

reqHttp.Header.Set("Content-Type", "application/json")

resp, err := client.Do(reqHttp)
if err != nil {
	fmt.Println("HTTP ERROR:", err)
	return map[string]interface{}{
		"code": "E98",
		"message": err.Error(),
	}, nil
}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawBody, _ := io.ReadAll(resp.Body)

	fmt.Println("HTTP STATUS:", resp.Status)
	fmt.Println("RAW RESPONSE:", string(rawBody))

	var result map[string]interface{}
	json.Unmarshal(rawBody, &result)
	// json.NewDecoder(resp.Body).Decode(&result)

	// if result["code"] != "000" {
	// 	return result, nil
	// }

	// fmt.Println("RESULT:", result)

	// return map[string]interface{}{
	// 	"code":    "000",
	// 	"status":  "ok",
	// 	"message": "Success",
	// 	"data":    result,
	// }, nil
	return result, nil
}
