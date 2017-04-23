package main

import (
	"sync"
	"strings"
    "fmt"
	"strconv"
    "io/ioutil"
	"time"
)
var mutex = &sync.Mutex{}
var stacje ([]stacja)
var tory_przejazdowe ([]tor_przejazdowy)
var tory_postojowe ([]tor_postojowy)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type pojazd struct {
	id int
	V_max int
	il_osob int
	trasa ([]stacja)
}
// jesli jest startP to pojazd znajduje sie na torze postojowym dla danej stacji

func startP(poj *pojazd) {
	trasa := poj.trasa
	dlugosc := len(trasa)
	var peron *tor_postojowy
	for in,n := range tory_postojowe {
		if n.stc.nazwa == trasa[0].nazwa {
			if(n.czyWolny) {
				mutex.Lock()
				peron = &tory_postojowe[in]
				peron.czyWolny = false
				mutex.Unlock()
				break
			}
		}
	}
	fmt.Println(poj.id,"Jestem na posotowym ",peron)
	idStacji,_ := strconv.Atoi(peron.stc.nazwa)
	for {
		stacje[idStacji].czyWolna = false
		time.Sleep(time.Second * 1)
		stacje[idStacji] = true;
	}

		
	i := 0
	for {

		var tor *tor_przejazdowy
		var predkosc int = poj.V_max
		//trzeba bedzie jeszcze zrobic czekanie na zwrotnice i jakies te tory postojowe
		//szukamy toru przejazdowego dla naszego pojazdu
		for d, n := range tory_przejazdowe {
			if n.poczatek == trasa[i%dlugosc] && n.koniec == trasa[(i+1)%dlugosc] {
				tor = &tory_przejazdowe[d]
//				fmt.Println(poj.id, tory_przejazdowe[d])
				break
			}
		}
		//sprawdzamy ograniczenie predkosci dla danego toru
		if predkosc > tor.V_max {
			predkosc = tor.V_max
		}
//		fmt.Println(poj.id, "Sprawdzam czy zwrotnica wolna",trasa[i])
		for {
			if(!(stacje[i].czyWolna)) {
				continue
			}
			mutex.Lock()
			stacje[i].czyWolna = false
			mutex.Unlock()
			break;
		}

//		fmt.Println(poj.id, "Wjezdzam na zwrotnice: ",trasa[i])
		time.Sleep(time.Second * time.Duration(trasa[i].czasZwrot))
//		fmt.Println(poj.id, "Skonczylem na zwrotnicy",trasa[i])

//		fmt.Println("Sprawdzam czy tor wolny ")

		for {
			if tor.czyWolny {
				mutex.Lock()
				stacje[i].czyWolna = true
				tor.czyWolny = false
				mutex.Unlock()
				break
			}
		}
		czas := tor.dlugosc / predkosc
		interval := time.Duration(czas)
		
//		fmt.Println(poj.id , "Wyjezdzam z ", trasa[i].nazwa)
//		fmt.Println(poj.id, "I jade ... do",trasa[(i+1)%dlugosc])
		time.Sleep(time.Second * interval)
		
		for {
			if(stacje[(i+1)%dlugosc].czyWolna) {
				continue
			}
			mutex.Lock()
			stacje[i].czyWolna = false
			tor.czyWolny = true
			mutex.Unlock()
			break;
		}


		i = (i + 1)%dlugosc
	}
}

type stacja struct {
	nazwa string
	czasZwrot int
	czyWolna bool
}

type tor_postojowy struct {
	stc stacja
	min_czas_post int
	czyWolny bool
}

type tor_przejazdowy struct {
	poczatek stacja
	koniec stacja
	czyWolny bool
	dlugosc int
	V_max int
}

func monitorPostoj(){
	for {
		for _,n := range tory_postojowe {
			fmt.Println(n)
		}
		time.Sleep(time.Second * 1)
	}
}

func monitorStacje(){
	for {
		for _,n := range stacje {
			fmt.Println(n)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Wczytujemy dane z pliku
	dat, err := ioutil.ReadFile("./../data.txt")
	check(err)
	data := string(dat)
	arr := strings.Split(data, "\n")
	more_data := strings.Split(arr[0], " ")

	ilosc_stacji, err := strconv.Atoi(more_data[0])
	ilosc_przejazdowych, err := strconv.Atoi(arr[1])
	ilosc_postojowych, err := strconv.Atoi(arr[2])
	ilosc_pojazdow, err := strconv.Atoi(arr[3])

//	fmt.Println("Ilosc torow przejazdowych", ilosc_przejazdowych)
//	fmt.Println("Ilosc torow postojowych", ilosc_postojowych)
//	fmt.Println("Ilosc pojazdow",ilosc_pojazdow)

	stacje = make([]stacja, ilosc_stacji)
	for n := 0; n < ilosc_stacji; n++ {
		minTime,_ := strconv.Atoi(more_data[n+1])
		stacje[n] = stacja{strconv.Itoa(n+1), minTime, true}
	}

//	fmt.Println("Stacje: ", stacje)

	tory_przejazdowe = make([]tor_przejazdowy, ilosc_przejazdowych)
	i := 0
	for n := 4; n < (len(arr)-ilosc_postojowych-1-ilosc_pojazdow); n++ {
		special_data := strings.Split(arr[n]," ")
		pocz,_ := strconv.Atoi(special_data[0])
		kon,_ := strconv.Atoi(special_data[1])
		dlug,_ := strconv.Atoi(special_data[2])
		Vmax,_ := strconv.Atoi(special_data[3])
		tory_przejazdowe[i] = tor_przejazdowy{stacje[pocz-1], stacje[kon-1], true,dlug,Vmax}
		i++
	}
	
//	fmt.Println("Tory przejazdowe",tory_przejazdowe)

	i = 0

	tory_postojowe = make([]tor_postojowy, ilosc_postojowych)
	for n := (len(arr)-1-ilosc_postojowych-ilosc_pojazdow); n < (len(arr)-1-ilosc_pojazdow); n++ {
		special_data := strings.Split(arr[n], " ")
		stc,_ := strconv.Atoi(special_data[0])
		minT,_ := strconv.Atoi(special_data[1])
		tory_postojowe[i] = tor_postojowy{stacje[stc-1],minT,true}
		i++;
	}

//	fmt.Println("Tory postojowe",tory_postojowe)

	pojazdy := make([]pojazd, ilosc_pojazdow)
	it := 0
	for n := (len(arr)-1-ilosc_pojazdow); n < (len(arr)-1); n++ {
		special_data := strings.Split(arr[n]," ")
		v,_ := strconv.Atoi(special_data[0])
		il_osob,_ := strconv.Atoi(special_data[1])
		dl_trasy,_ := strconv.Atoi(special_data[2])
		trasa := make([]stacja, dl_trasy)
		i = 0
		for k := 3; k < len(special_data); k++ {
			numer_stacji,_ := strconv.Atoi(special_data[k])
			trasa[i] = stacje[numer_stacji - 1]
			i++
		}
		pojazdy[it] = pojazd{it, v, il_osob, trasa}
		it++
	}

//	fmt.Println("Pojazdy:",pojazdy)
	
//	go func(p1 *pojazd) {
//		for {
//			p1.il_osob++
//			fmt.Println("zwiekszam ilosc osob dla ", p1.id,p1.il_osob)
//			time.Sleep(time.Second*10)
//		}
//	}(&pojazdy[0])
//
//	go func(poj []pojazd) {
//		for {
//			for _,n := range poj {
//				fmt.Println(n)
//				time.Sleep(time.Second*2)
//			}
//		}
//	}(pojazdy)

	go startP(&pojazdy[0])
	go startP(&pojazdy[1])
//	go startP(&pojazdy[2])
//	go startP(&pojazdy[3])
//	go monitorPostoj()
	go monitorStacje()
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}