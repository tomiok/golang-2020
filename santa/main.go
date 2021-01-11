package main

import (
	"sync"
	"time"
	"math/rand"
	"log"
	"os"
	"os/signal"
)

const (
	numOfElf = 12				//numero de elfos totales
	brokenToys = 3				//juguetes necesarios rotos para despertar a santa
	totalReindeers = 9					//renos totales
)

var (
	counter int		//contador de elfos que tienen problemas
	ticker  = time.NewTicker(20 * time.Second)
	santaChUp    = make(chan bool)								//canal para Santa despierto
	santaChSleep = make(chan bool)								//canal para santa dormido
	s            = santa{
		mutex: sync.Mutex{},
	}
	//	leave = make(chan bool)											//canal para saber si el programa ha terminado en el caso de llegar todos los renos
)

type santa struct {
	mutex sync.Mutex
}

func elfsAreWorking(extWG *sync.WaitGroup, chCallingSanta chan bool) {	//la funcion lanza cada una de las rutinas de los elfos

	var wg sync.WaitGroup		//WaitGroup para esperar a cada una de las gorutinas de los elfos

	chBrokenToy := make(chan bool)			//canal que le pasaremos a la rutina makeToy. True en el caso de que se rompa el juguete
	wg.Add(numOfElf)		//añado el numero de elfos, que seran el numero de rutinas que voy a esperar
	for i:=1; i<= numOfElf; i++ {
		go makeToy(&wg, chBrokenToy, chCallingSanta,i)
		time.Sleep(time.Duration(1000) * time.Millisecond)
	}

	wg.Wait()	//una vez que he esperado a todas las gorutinas hago el Wait
	extWG.Done()	//decremento en 1 el WaitGroup principal
}

func makeToy(wg *sync.WaitGroup, chBrokenToy, chCallingSanta chan bool, i int) {		//rutina que controla a los elfos haciendo juguetes
	log.Println("elf", i ,"crafting a toy")
	rand := rand.Intn(3)		//saco un numero aleatorio 0,1 o 2 porque la probabilidad de que se rompa un juguete es 1/3

	if rand == 0 {
		counter++		//incremento el contador de elfos con problemas
		go santaHelp(chBrokenToy, chCallingSanta, i)		//en el caso de que se rompe un juguete llamamos a la rutina santaHelp que controla el numero de juguetes totales que se han roto
		select {
		case <-ticker.C:
			chBrokenToy <- true
		}
		<-chBrokenToy
	} else {
		log.Println("toy of elf", i , "is done")

	}
	wg.Done()	//decremento en 1 el contador de gorutinas que espero
}

func santaHelp(chBrokenToy, chCallingSanta chan bool, i int) {
	log.Println("The toy of elf", i ,"is broken, he needs help")
	if counter == 3 {		//necesitamos 3 elfos con problemas para llamar a santa
		log.Println("santa fixing the toys")
		helptime := time.Duration(2000+rand.Intn(3000)) * time.Millisecond	//santa tiene que solucionar el problema de los elfos entre 2 y 5 segundos
		time.Sleep(helptime)
		log.Println("Problem fixed in", helptime, "seconds")
		santaChSleep <- true		//una vez que santa soluciona el problema de los elfos vuelve a dormir hasta volver a ser despertado por otro grupo de elfos o por los renos
		chBrokenToy <- true		//mandamos un true al canal que controla si los juguetes estan rotos
		chCallingSanta <- true
		counter = 0		//reinicio el contador de elfos con problemas
	}
}

func reindeerArrival(wg *sync.WaitGroup) bool {	//esta funcion recibe como parametro un puntero a un WaitGroup
	for i := 1; i <= totalReindeers; i++ {		//bucle para cada uno de los renos
		time.Sleep(time.Duration(5000+rand.Intn(2000)) * time.Millisecond)	//entre cada reno que llega tiene que haber al menos 5 segundos. En mi caso habrá entre 5 y 7 seg
		log.Println("reindeer", i ,"arrived!")
	}
	log.Println("all the reindeers are here")		//cuando llegan todos los renos
	wg.Done()
	//leave <- true
	return true
}

func santaStateZero() {		//funcion que inicializa a santa al estado de dormido, pasando al canal santaChSleep un true
	go func() {
		santaChSleep <- true
	}()
}

//santaRoutine Santa when all the reindeer are at home and 3 elf have troubles with the toy's production
func santaRoutine() {
	for {
		select {
		case <-santaChUp:
			log.Println("Santa is up now...")
			s.mutex.Lock()
		case <-santaChSleep:
			log.Println("Santa is sleeping...")
		}
	}
}

func wakeUpSanta() {
	go func() { santaChUp <- true }()
}

func main() {
	leave := make(chan os.Signal, 1)
	callingSanta := make(chan bool, 1)
	var wg sync.WaitGroup
	santaStateZero()		//en primer lugar lanzamos la funcion santaStateZero que inicializa a santa al estado de dormido

	wg.Add(3)		//agregamos las gorutinas que vamos a esperar, en este caso 3, santaRoutine, reinderArrival y elfsAreWorking
	go santaRoutine()
	time.Sleep(time.Duration(1000) * time.Millisecond)		//este sleep es para que aparezca en primer lugar el mensaje "Santa is sleeping..."
	go elfsAreWorking(&wg, callingSanta)
	go reindeerArrival(&wg)		//lanzamos la gorutina reindeerArrival pasandole como referencia el WaitGroup que esperara a otras gorutinas

	wg.Done()		//reducimos en 1 el contador de las rutinas que estamos esperando cuando estas acaban
	time.Sleep(time.Duration(5000) * time.Millisecond)
	//wg.Wait()

	go func() {
		signal.Notify(leave, os.Interrupt)
	}()

	select {
	case <-callingSanta:
		log.Println("santa is called")
		wakeUpSanta()
	case <-leave:
		os.Exit(100)
	}
	wg.Wait()
	log.Println("Santa is ready to go!")
}