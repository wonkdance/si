package si_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/wonksing/si/v2"
	"github.com/wonksing/si/v2/sicore"
	"github.com/wonksing/si/v2/siutils"
	"github.com/wonksing/si/v2/tests/testmodels"
)

func makeStuReq(num int) testmodels.StudentList {
	res := make(testmodels.StudentList, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, testmodels.Student{ID: i})
	}
	return res
}
func makeStuRes(num int) testmodels.StudentList {
	res := make(testmodels.StudentList, 0, num)
	for i := 0; i < num; i++ {
		res = append(res, testmodels.Student{ID: i})
	}
	return res
}

var (
	stuReqTiny = makeStuReq(1)
	stuResTiny = makeStuRes(1)
	stuReqSml  = makeStuReq(8)
	stuResSml  = makeStuRes(8)
	stuReqMed  = makeStuReq(128)
	stuResMed  = makeStuRes(128)
	stuReqLrg  = makeStuReq(512)
	stuResLrg  = makeStuRes(512)
)

func handleTestBasicTiny(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(&stuResTiny); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
func handleTestReaderWriterTiny(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	if err := si.DecodeJson(&req, r.Body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := si.EncodeJson(w, &stuResTiny); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func handleTestReaderWriterCopiedTiny(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	copiedReqBody, err := si.DecodeJsonCopied(&req, r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedReqBody)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	copiedRespBody, err := si.EncodeJsonCopied(w, &stuResTiny)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedRespBody)
}

func handleTestBasicSml(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(&stuResSml); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
func handleTestReaderWriterSml(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	if err := si.DecodeJson(&req, r.Body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := si.EncodeJson(w, &stuResSml); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func handleTestReaderWriterCopiedSml(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	copiedReqBody, err := si.DecodeJsonCopied(&req, r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedReqBody)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	copiedRespBody, err := si.EncodeJsonCopied(w, &stuResSml)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedRespBody)
}

func handleTestBasicMed(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(&stuResMed); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
func handleTestReaderWriterMed(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	if err := si.DecodeJson(&req, r.Body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := si.EncodeJson(w, &stuResMed); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func handleTestReaderWriterCopiedMed(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	copiedReqBody, err := si.DecodeJsonCopied(&req, r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedReqBody)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	copiedRespBody, err := si.EncodeJsonCopied(w, &stuResMed)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedRespBody)
}

func handleTestBasicLrg(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(b, &req); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(&stuResLrg); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
func handleTestReaderWriterLrg(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	if err := si.DecodeJson(&req, r.Body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := si.EncodeJson(w, &stuResLrg); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}

func handleTestReaderWriterCopiedLrg(w http.ResponseWriter, r *http.Request) {
	var req testmodels.StudentList
	copiedReqBody, err := si.DecodeJsonCopied(&req, r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedReqBody)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	copiedRespBody, err := si.EncodeJsonCopied(w, &stuResLrg)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer sicore.PutBytesBuffer(copiedRespBody)
}

func TestHttpHandlerSml(t *testing.T) {
	router := http.NewServeMux()
	router.HandleFunc("/test", handleTestReaderWriterSml)

	buf := bytes.NewBuffer([]byte(`[{"id":1,"email_address":"wonk@wonk.org","name":"wonk","borrowed":false,"book_id":23}]`))
	req, err := http.NewRequest("POST", "/test", buf)
	siutils.AssertNilFail(t, err)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	fmt.Println(rec)
}

func TestDecodeJsonCopied_FailToDecode(t *testing.T) {
	str := `[{"id":1,"email_address":"wonk@wonk.org","name":"wonk","borrowed":false,"book_id":23}`
	r := bytes.NewReader([]byte(str))

	res := testmodels.StudentList{}
	c, err := si.DecodeJsonCopied(&res, r)
	siutils.AssertNotNilFail(t, err)
	siutils.AssertNotNilFail(t, c)
	fmt.Println(c.String())
}
