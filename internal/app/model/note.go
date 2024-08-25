package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Note struct {
	ID      int    `json:"id"`
	User_ID int    `json:"user_id"`
	Note    string `json:"note"`
}

type ResponseSpellerNote struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

var worldsNote []ResponseSpellerNote

func (n *Note) ValidateBeforeCreate() error {
	s := strings.Join(strings.Fields(n.Note), "+")
	url := fmt.Sprintf("https://speller.yandex.net/services/spellservice.json/checkText?text=%s", s)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("%e %s", err, url)
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	json.Unmarshal(body, &worldsNote)

	for _, j := range worldsNote {
		i := strings.Index(s, j.Word)
		s = s[:i] + j.S[0] + s[i+len(j.Word):]
	}
	n.Note = strings.Join(strings.Split(s, "+"), " ")
	return nil
}
