package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maximelamure/restli-go/common"
)

func loadSchema(file string, data interface{}) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return common.Wrap(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return json.Unmarshal([]byte(byteValue), data)
}

func generate(data interface{}, tpl *template.Template) ([]byte, error) {
	var raw bytes.Buffer
	err := tpl.Execute(&raw, data)
	if err != nil {
		return nil, common.Wrap(err)
	}

	formated, err := format.Source(raw.Bytes())
	if err != nil {
		return nil, common.Wrap(err)
	}

	return formated, nil
}

func save(data string, output string) error {
	out, err := os.Create(output)
	if err != nil {
		return common.Wrap(err)
	}

	defer out.Close()
	_, err = fmt.Fprintln(out, data)
	return err
}

func capitalInitialism(in string) string {
	if u := strings.ToUpper(in); commonInitialisms[u] {
		return u
	}

	return strings.Title(in)
}

func toGoType(t string) string {
	if val, ok := pdscToGOTypes[strings.ToUpper(t)]; ok {
		return val
	}

	log.Println("Unknown type", t)
	return t
}

// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

var pdscToGOTypes = map[string]string{
	"LONG":   "int64",
	"STRING": "string",
}

func paramToString(paramName string, paramType string) string {
	switch paramType {
	case "int64":
		return "strconv.FormatInt(" + paramName + ", 10)"
	case "int":
		return "strconv.Itoa(" + paramName + ")"
	default:
		return paramName
	}
}
