package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Content-Type-Options", "nosniff")

		flushCopy(w, r.Body, r.ContentLength)

	})

	log.Print("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func flushCopy(dst io.Writer, src io.Reader, size int64) (written int64, err error) {
	buf := make([]byte, 1024 * 8)

	flusher, canFlush := dst.(http.Flusher)
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				if canFlush {
					flusher.Flush()
				}
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er == io.EOF {
			break
		}
		if er != nil {
			err = er
			break
		}
	}
	return written, err
}

/*
	for _, b := range buf {
		fmt.Fprintf(w, "Chunk %s \n", string(b))
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
		time.Sleep(100 * time.Millisecond)
		fmt.Println(string(b))
	}
*/