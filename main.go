package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/l-vitaly/cryptopro"
)

var (
	sha1      = flag.String("sha1", "", "")
	operation = flag.String("o", "sign", "-o=<sign|check>")
)

func init() {
	flag.Parse()
}

func main() {
	store, err := cryptopro.SystemStore("MY")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer store.Close()

	crt, err := store.GetBySHA1(*sha1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer crt.Close()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataBuf := bytes.NewBuffer(data)

	var (
		out []byte
	)

	switch *operation {
	default:
		fmt.Println("Unknown operation")
		os.Exit(1)
	case "sign":
		out, err = signData(crt, dataBuf)
	case "check":
		out, err = checkData(crt, dataBuf)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Stdout.Write(out)
}

func signData(crt cryptopro.Cert, dataBuf *bytes.Buffer) ([]byte, error) {
	dest := new(bytes.Buffer)

	msg, err := cryptopro.OpenToEncode(dest, cryptopro.EncodeOptions{
		Signers: []cryptopro.Cert{crt},
	})
	if err != nil {
		return nil, err
	}
	_, err = dataBuf.WriteTo(msg)
	if err != nil {
		return nil, err
	}

	msg.Close()

	return dest.Bytes(), nil
}

func checkData(crt cryptopro.Cert, dataBuf *bytes.Buffer) ([]byte, error) {
	msg, err := cryptopro.OpenToDecode(dataBuf)
	if err != nil {
		return nil, err
	}

	dest := new(bytes.Buffer)

	_, err = io.Copy(dest, msg)
	if err != nil {
		return nil, err
	}
	err = msg.Verify(crt)
	if err != nil {
		return nil, err
	}
	return dest.Bytes(), nil
}
