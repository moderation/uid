package main

// https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

func genShortUUID(w http.ResponseWriter, req *http.Request) {
	id := shortuuid.New()
	io.WriteString(w, id)
	// fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
}

func genUUID(w http.ResponseWriter, req *http.Request) {
	id := guuid.New()
	io.WriteString(w, id.String())
	// fmt.Printf("github.com/google/uuid:         %s\n", id.String())
}

func genXid(w http.ResponseWriter, req *http.Request) {
	id := xid.New()
	io.WriteString(w, id.String())
	// fmt.Printf("github.com/rs/xid:              %s\n", id.String())
}

func genKsuid(w http.ResponseWriter, req *http.Request) {
	id := ksuid.New()
	io.WriteString(w, id.String())
	// fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())
}

func genBetterGUID(w http.ResponseWriter, req *http.Request) {
	id := betterguid.New()
	io.WriteString(w, id)
	// fmt.Printf("github.com/kjk/betterguid:      %s\n", id)
}

func genulid(w http.ResponseWriter, req *http.Request) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	io.WriteString(w, id.String())
	// fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
}

func genSonyflake(w http.ResponseWriter, req *http.Request) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	io.WriteString(w, strconv.FormatUint(id, 10))
	// fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
}

func genSid(w http.ResponseWriter, req *http.Request) {
	id := sid.Id()
	io.WriteString(w, id)
	// fmt.Printf("github.com/chilts/sid:          %s\n", id)
}

func genUUIDv4(w http.ResponseWriter, req *http.Request) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("uuid.NewV4() failed with %s\n", err)
	}
	io.WriteString(w, id.String())
	// fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
}

func genAdmin(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	
	var s = "<pre>" 
	s += "/<a href=betterguid>betterguid</a>  <a href=https://github.com/kjk/betterguid>github.com/kjk/betterguid</a><br>"
	s += "/<a href=ksuid>ksuid</a>       <a href=https://github.com/segmentio/ksuid>github.com/segmentio/ksuid</a><br>"
	s += "/<a href=shortuuid>shortuuid</a>   <a href=https://github.com/lithammer/shortuuid>github.com/lithammer/shortuuid</a><br>"
	s += "/<a href=sid>sid</a>         <a href=https://github.com/chilts/sid>github.com/chilts/sid</a><br>"
	s += "/<a href=sonyflake>sonyflake</a>   <a href=https://github.com/sony/sonyflake>github.com/sony/sonyflake</a><br>"
	s += "/<a href=ulid>ulid</a>        <a href=https://github.com/oklog/ulid>github.com/oklog/ulid</a><br>"
	s += "/<a href=uuid>uuid</a>        <a href=https://github.com/google/uuid>github.com/google/uuid</a><br>"
	s += "/<a href=uuidv4>uuidv4</a>      <a href=https://github.com/satori/go.uuid>github.com/satori/go.uuid</a><br>"
	s += "/<a href=xid>xid</a>         <a href=https://github.com/rs/xid>github.com/rs/xid</a><br>"
	s += "<br>* see <a href=https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html>generating good unique ids in go</a>"
	s += "</pre>"
	io.WriteString(w, s)
}

func main() {
	http.HandleFunc("/", genAdmin)
	http.HandleFunc("/betterguid", genBetterGUID)
	http.HandleFunc("/ksuid", genKsuid)
	http.HandleFunc("/shortuuid", genShortUUID)
	http.HandleFunc("/sid", genSid)
	http.HandleFunc("/sonyflake", genSonyflake)
	http.HandleFunc("/ulid", genulid)
	http.HandleFunc("/uuid", genUUID)
	http.HandleFunc("/uuidv4", genUUIDv4)
	http.HandleFunc("/xid", genXid)

	log.Fatal(http.ListenAndServe(":8888", nil))
}
