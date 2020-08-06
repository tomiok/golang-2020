package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

//common errors in golang part 1

//1. declarar constantes con iota y sin valor por defecto / declare with iota with no default value
//2. no usar timeouts en llamadas http / do not use timeouts in http requests
//3. ignorar errores, o manejarlos desde una misma funcion // ignore errors or wrap them in only one function
//4. muchas interfaces / to many interfaces
//5. cuando ocurre un error, recordar cerrar, cancelar, etc los componentes usados / when an error occurs, do not forget to close, cancel everything used
//6. si usamos read/write lock, separar bien lectura y escritura, ademas de liberar los locks // using RWLock, split well in write and read, and release the lock always
//7. no test unitarios // not unit testing

//ejemplos / examples
// 1

type status int

const (
	statusUnknown  status = iota
	statusActive          // 1
	_                     // skips number 2
	statusInactive        // 3
)

// ejemplo / example  2

func httpCall() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// WRONG - do not ignore errors
	req, _ := http.NewRequestWithContext(ctx, "POST", "google.com", nil)
	client := http.Client{}
	client.Do(req)

	// as a server
	serv := http.Server{
		Addr:              "",
		Handler:           nil,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    1024,
	}
	serv.ListenAndServe()
}

func ignoreErr() {
	f, err := os.Create("this.txt")

	if err != nil {
		log.Print("the file cannot be created " + err.Error())
		return //cut the function
	}

	defer f.Close()

	if true {
		return
	}
}

func usingLocks() {
	l := sync.Mutex{}

	l.Lock()
	fmt.Print("locked!")
	defer l.Unlock()

}

func usingRWLocks() {
	l := sync.RWMutex{}

	// writing
	l.Lock()
	fmt.Print("writing")
	defer l.Unlock()
	//

	// reading
	l.RUnlock()
	fmt.Print("reading")
	l.RUnlock()
	//

}

// ejemplo / example 7
type PatientStorage interface {
	SavePatient(name string)
}

type PatientStg struct {}

func (p *PatientStg) SavePatient(name string) {
	// SQL Calls
	// logs...
}

type MockPatientStg struct {
	// mock
}

func (p *MockPatientStg) SavePatient(name string) {
	// SQL Calls
	// logs...

	if name == "fail" {
		panic("")
	}


}


func main() {
	fmt.Println(statusActive)
	fmt.Println(statusInactive)
}
