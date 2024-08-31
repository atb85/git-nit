package nits

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
)

func whiteSpace(l string) string {
	re := regexp.MustCompile(`^\s+`)
	return re.FindString(l)
}

// NewNit returns a comment for the specific file given with a hash of details
func NewNit(f string, args ...string) string {

	dta := strings.Join(append(args, f), "")
	hshr := md5.New()
	hshr.Write([]byte(dta))
	enc := hex.EncodeToString(hshr.Sum(nil))
	return GenerateComment(f, enc)
}

// AddNit adds a nit to the content at the correct indentation
func AddNit(old []byte, nit string, chance float32) ([]byte, error) {
	br := bytes.NewReader(old)
	scn := bufio.NewScanner(br)
	var buf bytes.Buffer

	found := false
	for scn.Scan() {
		ln := scn.Bytes()
		if !found && rand.Float32() < chance {
			buf.Write([]byte(whiteSpace(string(ln)) + nit + "\n"))
			found = true
		}
		buf.Write([]byte(string(ln) + "\n"))
	}
	if err := scn.Err(); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}
