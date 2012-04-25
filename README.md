# solr
not much here. just needed something quick to wrap the http calls

# example
	type Doc struct {
		Id		int	`json:"id"`
		Keywords	string	`json:"keywords"`
	}

	s := solr.New("localhost", 8080, "posts")
	// put some docs in
	docs := make([]interface{}, 1)
	docs[0] = Doc{12345, "cars car vehicle"}
	res, err := s.Update(docs)
	log.Println("ERR:", err, "\nRESP:", res)

	// get some docs out
	res, err = s.Query("post:car")
	log.Println("ERR:", err, "\nRESP:", res)

# todo
add some sorting and limiting abilities to Query()
do something about the return from Query() -- good enough for now