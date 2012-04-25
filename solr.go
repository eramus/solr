package solr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type SolrParams struct {
	Indent	bool		`json:"indent"`
	Format	string	`json:"wt"`
	Query	string	`json:"q"`
}

type SolrResponseHeader struct {
	Status	int	`json:"status"`
	QTime	int
	Params	SolrParams	`json:"params"`
}

type SolrResponseData struct {
	NumFound	int				`json:"numFound"`
	Start		int				`json:"start"`
	Docs		[]interface{}	`json:"docs"`
}

type SolrResponse struct {
	ResponseHeader	SolrResponseHeader	`json:"responseHeader"`
	Response		SolrResponseData	`json:"response"`
}

type Solr struct {
	host string
	port int
	index string
}

func New(host string, port int, index string) *Solr {
	s := new(Solr)
	s.host = host
	s.port = port
	s.index = index
	return s
}

// TODO: add some limiting options
func (s *Solr) Query(query string) (*SolrResponse, error) {
	// run the search query
	q := fmt.Sprintf("http://%s:%d/%s/select/?wt=json&q=%s", s.host, s.port, s.index, url.QueryEscape(query))
	resp, err := http.Get(q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sr := new(SolrResponse)
	err = json.Unmarshal(results, sr)
	return sr, err
}

func (s *Solr) Update(docs []interface{}) (*SolrResponse, error) {
	ds, err := json.Marshal(docs)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(ds)
	uri := fmt.Sprintf("http://%s:%d/%s/update/json/?commit=true&wt=json", s.host, s.port, s.index)
	resp, err := http.Post(uri, "application/json", buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	results, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sr := new(SolrResponse)
	err = json.Unmarshal(results, sr)
	if err != nil {
		return nil, err
	}
	return sr, nil
}