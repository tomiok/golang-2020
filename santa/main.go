package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	numOfElf       = 12 //numero de elfos totales
	totalReindeers = 9  //renos totales
)

var (
	counter      uint32            //contador de elfos que tienen problemas
	santaChUp    = make(chan bool) //canal para Santa despierto
	santaChSleep = make(chan bool) //canal para santa dormido
	s            = santa{
		mutex: sync.Mutex{},
	}
)

type santa struct {
	mutex sync.Mutex
}

func elfsAreWorking(extWG *sync.WaitGroup) { //la funcion lanza cada una de las rutinas de los elfos

	var wg sync.WaitGroup //WaitGroup para esperar a cada una de las gorutinas de los elfos

	wg.Add(numOfElf) //añado el numero de elfos, que seran el numero de rutinas que voy a esperar
	for i := 1; i <= numOfElf; i++ {
		log.Println("elf", i, "crafting a toy")
		random := rand.Intn(1) //saco un numero aleatorio 0,1 o 2 porque la probabilidad de que se rompa un juguete es 1/3

		if random == 0 {
			atomic.AddUint32(&counter, 1) //incremento el contador de elfos con problemas
			go santaHelp(i)               //en el caso de que se rompe un juguete llamamos a la rutina santaHelp que controla el numero de juguetes totales que se han roto
		} else {
			log.Println("toy of elf", i, "is done")

		}
		wg.Done()                                          //decremento en 1 el contador de gorutinas que espero
		time.Sleep(time.Duration(2000) * time.Millisecond) //espero un periodo de un segundo para lanzar cada una de las gorutinas
	}

	wg.Wait()    //una vez que he esperado a todas las gorutinas hago el Wait
	extWG.Done() //decremento en 1 el WaitGroup principal de las 3 gorutinas lanzadas en el programa principal, ya que la gorutina elfsAreWorking ha terminado
}

func santaHelp(i int) {

	log.Println("The toy of elf", i, "is broken, he needs help")

	if atomic.LoadUint32(&counter) == 3 { //necesitamos 3 elfos con problemas para llamar a santa
		santaChUp <- true                                  //despertamos a santa claus porque necesitan los elfos su ayuda
		time.Sleep(time.Duration(1000) * time.Millisecond) //esperamos un segundo para que aparezca antes el mensaje "Santa is up now..." que "santa fixing the toys"
		log.Println("santa fixing the toys")
		helptime := time.Duration(2000+rand.Intn(3000)) * time.Millisecond //santa tiene que solucionar el problema de los elfos entre 2 y 5 segundos
		time.Sleep(helptime)
		log.Println("Problem fixed in", helptime, "seconds")
		santaChSleep <- true //una vez que santa soluciona el problema de los elfos vuelve a dormir hasta volver a ser despertado por otro grupo de elfos o por los renos
		s.mutex.Unlock()

		atomic.StoreUint32(&counter, 0) //reinicio el contador de elfos con problemas
	}
}

func reindeerArrival(wg *sync.WaitGroup) { //esta funcion recibe como parametro un puntero a un WaitGroup
	for i := 1; i <= totalReindeers; i++ { //bucle para cada uno de los renos
		time.Sleep(time.Duration(5000+rand.Intn(2000)) * time.Millisecond) //entre cada reno que llega tiene que haber al menos 5 segundos. En mi caso habrá entre 5 y 7 seg
		log.Println("reindeer", i, "arrived!")
	}
	log.Println("all the reindeers are here") //cuando llegan todos los renos
	wg.Done()
}

func santaStateZero() { //funcion que inicializa a santa al estado de dormido, pasando al canal santaChSleep un true
	go func() {
		santaChSleep <- true
	}()
}

func santaRoutine() {
	for {
		select {
		case <-santaChUp:
			s.mutex.Lock()
			log.Println("Santa is up now...")

		case <-santaChSleep:
			log.Println("Santa is sleeping...")
		}
	}
}

func wakeUpSanta() {
	go func() { santaChUp <- true }()
}

func main() {
	rand.Seed(time.Now().UnixNano()) //lanzo una semilla aleatoria para que no realice lo mismo en cada iteración
	var wg sync.WaitGroup
	santaStateZero() //en primer lugar lanzamos la funcion santaStateZero que inicializa a santa al estado de dormido

	wg.Add(3)                                          //agregamos las gorutinas que vamos a esperar, en este caso 3, santaRoutine, reinderArrival y elfsAreWorking
	go santaRoutine()                                  //lanzamos la gorutina de santa
	time.Sleep(time.Duration(1000) * time.Millisecond) //este sleep es para que aparezca en primer lugar el mensaje "Santa is sleeping..."
	go elfsAreWorking(&wg)                             //lanzamos la gorutina para lo elfos pasandole como referencia el WaitGroup que esperara a otras gorutinas
	go reindeerArrival(&wg)                            //lanzamos la gorutina reindeerArrival pasandole como referencia el WaitGroup que esperara a otras gorutinas

	wg.Done() //reducimos en 1 el contador de las rutinas a las que esperamos por la gorutina santaRoutine()
	wg.Wait()
	log.Println("Santa is ready to go!")
}
