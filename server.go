package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	SERVER_ADDRESS = "http://localhost"
	SERVER_PORT    = "32224"
	SERVER_TYPE    = "tcp"
	SERVER         = "http://localhost:32224"
)


type User struct {
	ID   int
	Name string
}

func server() {

	// GET
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		firstLine := r.Method + " " + r.Proto + " " + "200 OK" + " " + " Requested path: " + r.RequestURI
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(firstLine))
		return
	})

	// POST
	http.HandleFunc("/create-user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
			return
		}
		user := &User{}
		data := json.NewDecoder(r.Body).Decode(user)

		if data == nil {

			firstLine := r.Method + " " + r.Proto + " " + "200 OK" + " " + " Requested path: " + r.RequestURI

			w.WriteHeader(http.StatusCreated)
			_, err := w.Write([]byte(firstLine))

			if err != nil {
				fmt.Println("Error writing response")
			}
			return
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	})

	// DELETE
	http.HandleFunc("/delete-user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method now allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Body == nil {
			http.Error(w, "Cannot delete user without ID", http.StatusBadRequest)
			return
		}
		firstLine := r.Method + " " + r.Proto + " " + "200 OK" + " " + " Requested path: " + r.RequestURI

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(firstLine))

		if err != nil {
			fmt.Println("Error writing response")
			return
		}
	})

	// PUT
	http.HandleFunc("/update-user", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if r.Body == nil {
			http.Error(w, "Cannot update user without ID", http.StatusBadRequest)
			return
		}

		firstLine := r.Method + " " + r.Proto + " " + "200 OK" + " " + " Requested path: " + r.RequestURI
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(firstLine))

		if err != nil {
			fmt.Println("Error writing response")
			return
		}
	})

	if err := http.ListenAndServe(":32224", nil); err != http.ErrServerClosed {
		panic(err)
	}

}

func client() error {

	// data for requests
	user := &User{
		ID:   2,
		Name: "asd",
	}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(user)
	if err != nil {
		return err
	}
	resp, err := http.Post(SERVER+"/create-user", "application/json", b)

	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	resp, err = http.Get(SERVER + "/")
	if err != nil {
		return err
	}

	body, err = io.ReadAll(resp.Body)
	fmt.Println(string(body))

	defer resp.Body.Close()
	return nil

}

func main() {

	go server()

	time.Sleep(time.Second)

	if err := client(); err != nil {
		fmt.Println(err)
	}

}
