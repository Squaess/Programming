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
var pojazdy ([] pojazd)

func gadatliwy(){
	fmt.Println("Gadatliwy")
	
	go startP(&pojazdy[0])
	go startP(&pojazdy[1])
	go startP(&pojazdy[2])
	go startP(&pojazdy[3])
}

func spokojny(){
	fmt.Println("Spokojny")
	go startS(&pojazdy[0])
	go startS(&pojazdy[1])
	go startS(&pojazdy[2])
	go startS(&pojazdy[3])
	for {
		fmt.Println("1. Informacje o peronach")
		fmt.Println("2. Informacje o zwrotnicach")
		fmt.Println("3. Informacje o torach")
		var input string
		fmt.Scanln(&input)
		fmt.Println(input)
		switch input {
			case "1":
				monitorPostoj()			
			case "2":
				monitorStacje()
			case "3":
				monitorPrzejazd()
		}
		if(input == "quit"){
		break}
	}
}

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
func startS(poj *pojazd) {
	
	trasa := poj.trasa
	dlugosc := len(trasa)
	//umieszczamy pojazd na peronie
	var peron *tor_postojowy
	for in,n := range tory_postojowe {
		if n.stc.nazwa == trasa[0].nazwa {
			mutex.Lock()
			if(n.czyWolny) {
				
				peron = &tory_postojowe[in]
				peron.czyWolny = false
				mutex.Unlock()
				break
			}
			mutex.Unlock()
		}
	}
	
	// szukamy stacji aby moc skozystac ze zwrotnicy


		
	i := 0
	for {
//		fmt.Println(poj.id,"Jestem na postojowym ",peron)
		// torem bedzie tor ktorego chcemy uzyc aby dostac sie na nastepna stacje ktora jest na naszej trasie
		var tor *tor_przejazdowy
		var predkosc int = poj.V_max

		//szukamy toru przejazdowego dla naszego pojazdu
		for d, n := range tory_przejazdowe {
			if n.poczatek == trasa[i%dlugosc] && n.koniec == trasa[(i+1)%dlugosc] {
				tor = &tory_przejazdowe[d]
				break
			}
		}
		//sprawdzamy ograniczenie predkosci dla danego toru
		if predkosc > tor.V_max {
			predkosc = tor.V_max
		}
		//najpierw musimu sprawdzic czy tor jest wolny aby nie blokowac zwrotnicy 
		for {
			mutex.Lock()
			if tor.czyWolny {
				idS,_ := strconv.Atoi(tor.poczatek.nazwa)
				//odejmujemy 1 poniewaz stacja o nazwie 1 znajduje sie pod stacje[0]
				idS = idS -1
//				fmt.Println(poj.id,"tor wolny sprawdzam zwrotnice",tor.czyWolny, stacje[idS].czyWolna)
				//rezerwujemy sobie tor
				tor.czyWolny = false
				mutex.Unlock()

				//czekamy na zwrotnice
				for {
					mutex.Lock()	
					if stacje[idS].czyWolna==true {
//						fmt.Println(poj.id,"Tor wolny i zwrotnica tez... jade", tor.czyWolny, stacje[idS].czyWolna)	
						stacje[idS].czyWolna = false
						peron.czyWolny = true;
						mutex.Unlock()
						break
					}
					mutex.Unlock()
				}
				break
			}
			mutex.Unlock()
		}
		//czekamy az obroci sie zwrotnica
		time.Sleep(time.Second * time.Duration(trasa[i].czasZwrot))
		idS,_ := strconv.Atoi(trasa[i].nazwa)
		idS -= 1
//		fmt.Println(poj.id, "Skonczylem na zwrotnicy zwalniam ja",stacje[idS])
		mutex.Lock()
		stacje[idS].czyWolna = true
		mutex.Unlock()
//		fmt.Println(poj.id, "Wjezdzam na tor",tor)

//		fmt.Println("Sprawdzam czy tor wolny ")

//		for {
//			if tor.czyWolny {
//				mutex.Lock()
//				stacje[i].czyWolna = true
//				tor.czyWolny = false
//				mutex.Unlock()
//				break
//			}
//		}
		czas := tor.dlugosc / predkosc
		interval := time.Duration(czas)
		
//		fmt.Println(poj.id , "Wyjezdzam z ", trasa[i].nazwa)
//		fmt.Println(poj.id, "I jade ... do",trasa[(i+1)%dlugosc])
		time.Sleep(time.Second * interval)
		idS,_  = strconv.Atoi(trasa[(i+1)%dlugosc].nazwa)
		idS = idS-1
//		fmt.Println(poj.id, "Jestem pod stacja czekam na zwrotnice", stacje[idS])

		for {

			if(stacje[idS].czyWolna) {
				mutex.Lock()
//				fmt.Println(poj.id,"Zwrotnica wolna",stacje[idS].czyWolna)
				stacje[idS].czyWolna = false
				tor.czyWolny = true
				mutex.Unlock()
				break;
			}
		}

		zwrot := stacje[idS].czasZwrot
		time.Sleep(time.Second * time.Duration(zwrot))

		//sprawdzamy czy jest wolny peron
		for {
			for ind, en := range tory_postojowe {
				if en.stc.nazwa == stacje[idS].nazwa {
					if(en.czyWolny) {
						mutex.Lock()
						peron = &tory_postojowe[ind]
//						fmt.Println(poj.id, "Zjezdzam ze zwrotnicy na peron", stacje[idS], peron)
						peron.czyWolny = false
						stacje[idS].czyWolna = true
						mutex.Unlock()
						break
					}
				}
			}
			if(peron.stc.nazwa == stacje[idS].nazwa) {
				break
			}
		}

		time.Sleep(time.Second * time.Duration(peron.min_czas_post))

		i = (i + 1)%dlugosc
	}
}

func startP(poj *pojazd) {
	
	trasa := poj.trasa
	dlugosc := len(trasa)
	//umieszczamy pojazd na peronie
	var peron *tor_postojowy
	for in,n := range tory_postojowe {
		if n.stc.nazwa == trasa[0].nazwa {
			mutex.Lock()
			if(n.czyWolny) {
				
				peron = &tory_postojowe[in]
				peron.czyWolny = false
				mutex.Unlock()
				break
			}
			mutex.Unlock()
		}
	}
	
	// szukamy stacji aby moc skozystac ze zwrotnicy


		
	i := 0
	for {
		fmt.Println(poj.id,"Jestem na postojowym ",peron)
		// torem bedzie tor ktorego chcemy uzyc aby dostac sie na nastepna stacje ktora jest na naszej trasie
		var tor *tor_przejazdowy
		var predkosc int = poj.V_max

		//szukamy toru przejazdowego dla naszego pojazdu
		for d, n := range tory_przejazdowe {
			if n.poczatek == trasa[i%dlugosc] && n.koniec == trasa[(i+1)%dlugosc] {
				tor = &tory_przejazdowe[d]
				break
			}
		}
		//sprawdzamy ograniczenie predkosci dla danego toru
		if predkosc > tor.V_max {
			predkosc = tor.V_max
		}
		//najpierw musimu sprawdzic czy tor jest wolny aby nie blokowac zwrotnicy 
		for {
			mutex.Lock()
			if tor.czyWolny {
				idS,_ := strconv.Atoi(tor.poczatek.nazwa)
				//odejmujemy 1 poniewaz stacja o nazwie 1 znajduje sie pod stacje[0]
				idS = idS -1
				fmt.Println(poj.id,"tor wolny sprawdzam zwrotnice",tor.czyWolny, stacje[idS].czyWolna)
				//rezerwujemy sobie tor
				tor.czyWolny = false
				mutex.Unlock()

				//czekamy na zwrotnice
				for {
					mutex.Lock()	
					if stacje[idS].czyWolna==true {
						fmt.Println(poj.id,"Tor wolny i zwrotnica tez... jade", tor.czyWolny, stacje[idS].czyWolna)	
						stacje[idS].czyWolna = false
						peron.czyWolny = true;
						mutex.Unlock()
						break
					}
					mutex.Unlock()
				}
				break
			}
			mutex.Unlock()
		}
		//czekamy az obroci sie zwrotnica
		time.Sleep(time.Second * time.Duration(trasa[i].czasZwrot))
		idS,_ := strconv.Atoi(trasa[i].nazwa)
		idS -= 1
		fmt.Println(poj.id, "Skonczylem na zwrotnicy zwalniam ja",stacje[idS])
		mutex.Lock()
		stacje[idS].czyWolna = true
		mutex.Unlock()
		fmt.Println(poj.id, "Wjezdzam na tor",tor)

//		fmt.Println("Sprawdzam czy tor wolny ")

//		for {
//			if tor.czyWolny {
//				mutex.Lock()
//				stacje[i].czyWolna = true
//				tor.czyWolny = false
//				mutex.Unlock()
//				break
//			}
//		}
		czas := tor.dlugosc / predkosc
		interval := time.Duration(czas)
		
//		fmt.Println(poj.id , "Wyjezdzam z ", trasa[i].nazwa)
//		fmt.Println(poj.id, "I jade ... do",trasa[(i+1)%dlugosc])
		time.Sleep(time.Second * interval)
		idS,_  = strconv.Atoi(trasa[(i+1)%dlugosc].nazwa)
		idS = idS-1
		fmt.Println(poj.id, "Jestem pod stacja czekam na zwrotnice", stacje[idS])

		for {

			if(stacje[idS].czyWolna) {
				mutex.Lock()
				fmt.Println(poj.id,"Zwrotnica wolna",stacje[idS].czyWolna)
				stacje[idS].czyWolna = false
				tor.czyWolny = true
				mutex.Unlock()
				break;
			}
		}

		zwrot := stacje[idS].czasZwrot
		time.Sleep(time.Second * time.Duration(zwrot))

		//sprawdzamy czy jest wolny peron
		for {
			for ind, en := range tory_postojowe {
				if en.stc.nazwa == stacje[idS].nazwa {
					if(en.czyWolny) {
						mutex.Lock()
						peron = &tory_postojowe[ind]
						fmt.Println(poj.id, "Zjezdzam ze zwrotnicy na peron", stacje[idS], peron)
						peron.czyWolny = false
						stacje[idS].czyWolna = true
						mutex.Unlock()
						break
					}
				}
			}
			if(peron.stc.nazwa == stacje[idS].nazwa) {
				break
			}
		}

		time.Sleep(time.Second * time.Duration(peron.min_czas_post))

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

func monitorPrzejazd() {
	for _, n := range tory_przejazdowe {
		fmt.Println(n)
	}
}
func monitorPostoj(){
	for _,n := range tory_postojowe {
		fmt.Println(n)
	}
	time.Sleep(time.Second * 1)
}

func monitorStacje(){
	for _,n := range stacje {
		fmt.Println(n)
	}
	time.Sleep(time.Second * 1)
}

func main() {
	// Wczytujemy dane z pliku
	dat, err := ioutil.ReadFile("./../data.txt")
	check(err)
	data := string(dat)
	arr := strings.Split(data, "\n")
	more_data := strings.Split(arr[0]," ")
	fmt.Println(arr[1])

	ilosc_stacji, err := strconv.Atoi(more_data[0])
	ilosc_przejazdowych, err := strconv.Atoi(arr[1])
	ilosc_postojowych, err := strconv.Atoi(arr[2])
	ilosc_pojazdow, err := strconv.Atoi(arr[3])

	fmt.Println("Ilosc stacji", ilosc_stacji)
	fmt.Println("Ilosc torow przejazdowych", ilosc_przejazdowych)
	fmt.Println("Ilosc torow postojowych", ilosc_postojowych)
	fmt.Println("Ilosc pojazdow",ilosc_pojazdow)

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

	pojazdy = make([]pojazd, ilosc_pojazdow)
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

	
	fmt.Println("1. Tryb gadatliwy")
	fmt.Println("2. Tryb spokojny")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	switch input {
		case "1":
			go gadatliwy()
		case "2":
			spokojny()
		default:
			fmt.Println("Error error")
	}




//	go monitorPostoj()
//	go monitorStacje()
	for {
		fmt.Scanln(&input)
		if(input == "quit") {
			break
		}
	}
//	fmt.Scanln(&input)
	fmt.Println("done")
}
