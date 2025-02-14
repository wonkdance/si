package sicore

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_EncodeJson(t *testing.T) {
	dst := bytes.Buffer{}
	src := _testStruct{
		Msg: "hello world",
	}
	err := EncodeJson(&dst, &src)
	require.Nil(t, err)
	b, _ := json.Marshal(&src)
	b = append(b, '\n')
	require.EqualValues(t, b, dst.Bytes())
}

type _testStruct struct {
	Msg string `json:"msg"`
}
