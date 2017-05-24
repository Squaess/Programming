package main

import (
	"fmt"
	"math/rand"
	//"sync"
	"time"
	"strings"
	"strconv"
	"io/ioutil"
)

var stacje ([]*stacja)
var tory_przejazdowe ([]*tor_przejazdowy)
var tory_postojowe ([]*tor_postojowy)
var pojazdy ([]*pojazd)
var ratunek *pojazd

type stacja struct {
	id int
	czasZwrot int
	resp chan bool
	czyWolna bool
}

type tor_przejazdowy struct {
	id int
	id_poczatek int
	id_koniec int
	dlugosc int
	V_max int
	czyWolny bool
	awaria chan bool
}

type tor_postojowy struct {
	id int
	id_stacji int
	min_czas_post int
	czyWolny bool
	awaria chan bool
}

type pojazd struct {
	id int
	V_max int
	il_osob int
	trasa []int
	nast_stac int
	resp chan bool
	post bool
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func find_path(id_poczatek int, id_koniec int) []int{
	trasa := make([]int,0)
	trasa = append(trasa, id_poczatek)
	
	for _,n := range tory_przejazdowe {
		if(n.id_poczatek == id_poczatek && n.id_koniec == id_koniec){
			trasa = append(trasa, id_koniec)
			break
		}
	}	
	for{
		if trasa[0] == id_poczatek && trasa[len(trasa)-1] == id_koniec {
			break
		}
		for i := len(tory_przejazdowe)-1; i>=0; i--{
			if tory_przejazdowe[i].id_koniec > id_koniec {
				continue
			}
			if(trasa[len(trasa)-1]==tory_przejazdowe[i].id_poczatek) {
				trasa = append(trasa, tory_przejazdowe[i].id_koniec)
				break
			}
		}
	}
	return trasa
}

func find_back(id_poczatek int, id_koniec int) []int{
	trasa := find_path(id_poczatek,8)
	trasa = append(trasa,1)
	return trasa
}

func main() {
	
	dat, err := ioutil.ReadFile("./../data.txt")
	check(err)
	data := string(dat)
	// arr[0] zawiera zwrotnice z czasem obrotu
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

 	stacje = make([]*stacja, ilosc_stacji)
    for n := 0; n < ilosc_stacji; n++ {
        minTime,_ := strconv.Atoi(more_data[n+1])
        stacje[n] = &stacja{ n+1, minTime, make(chan bool),true}
    }

	tory_przejazdowe = make([]*tor_przejazdowy, ilosc_przejazdowych)
	i := 0
	for n:=4; n < len(arr)- ilosc_postojowych-1-ilosc_pojazdow; n++ {
		special_data := strings.Split(arr[n]," ")
        pocz,_ := strconv.Atoi(special_data[0])
        kon,_ := strconv.Atoi(special_data[1])
        dlug,_ := strconv.Atoi(special_data[2])
        Vmax,_ := strconv.Atoi(special_data[3])
		tory_przejazdowe[i] = &tor_przejazdowy{i+1,pocz, kon,dlug,
								Vmax, true, make(chan bool)}
		i++
	}
	i = 0
	tory_postojowe = make([]*tor_postojowy, ilosc_postojowych)
    for n := (len(arr)-1-ilosc_postojowych-ilosc_pojazdow); n < (len(arr)-1-ilosc_pojazdow); n++ {
        special_data := strings.Split(arr[n], " ")
        stc,_ := strconv.Atoi(special_data[0])
        minT,_ := strconv.Atoi(special_data[1])
        tory_postojowe[i] = &tor_postojowy{i+1, stc, minT, true, make(chan bool)}
        i++;
    }

    pojazdy = make([]*pojazd, ilosc_pojazdow)
    it := 0

	for n := (len(arr)-1-ilosc_pojazdow); n < (len(arr)-1); n++ {
        special_data := strings.Split(arr[n]," ")
        v,_ := strconv.Atoi(special_data[0])
        il_osob,_ := strconv.Atoi(special_data[1])
        dl_trasy,_ := strconv.Atoi(special_data[2])
		trasa := make([]int, dl_trasy)
		poj := &pojazd{it, v, il_osob, trasa, trasa[1],make(chan bool), true}
		i = 0
		for k := 3; k < len(special_data); k++ {
            numer_stacji,_ := strconv.Atoi(special_data[k])
            poj.trasa[i] = numer_stacji
            i++
        }

        pojazdy[it] = poj

        it++
    }
	//-----------------------------------------Koniec tworzenia obiektow i formatowania danych---------------//
	stacje_chan := make([]chan *pojazd,ilosc_stacji+1)
	for i := range stacje_chan {
		stacje_chan[i] = make(chan *pojazd)
	}

	tory_przejazdowe_chan := make ([]chan *pojazd, ilosc_przejazdowych+1)
	for i := range tory_przejazdowe_chan {
		tory_przejazdowe_chan[i] = make(chan *pojazd)
	}

	tory_postojowe_chan := make ([]chan *pojazd, ilosc_postojowych+1)
	for i := range tory_postojowe_chan {
		tory_postojowe_chan[i] = make(chan *pojazd)
	}

	kar_chan_poj := make(chan *pojazd)
	kar_chan_post := make(chan *tor_postojowy)
	kar_chan_przej := make(chan *tor_przejazdowy)
	kar_chan_stac := make(chan *stacja)

	// gourityna zarzadzajaca zwrotnica
	for i := range stacje {
		go func(stac *stacja){
			for {	
				select {	
				case poj := <-stacje_chan[stac.id]:
					stac.czyWolna = false
					
					fmt.Println( poj.id, "jest na zwrotnicy",
						stac.id," zwrotnica sie obraca ", stac.czasZwrot)
					time.Sleep(time.Duration(stac.czasZwrot) * time.Second)
					//sprawdzamy czy pojazd chce wjechac na postojowy czy na przejazdowy
					if poj.post {
					// wiemy ze chce wjechac na postojowy
					//Zwrotnica sprawdza dostepne tory postojowe
					var tor *tor_postojowy
					for {
						for _,n := range tory_postojowe {
							if n.id_stacji == stac.id {
								if n.czyWolny {
									tor = n
									tory_postojowe_chan[n.id]<-poj
									break
								}		
							}
						}
						if tor.czyWolny {
							break
						}
					}
					}else {
						fmt.Println(poj.id, "chce wiechac na przejazdowy do",
									poj.nast_stac)
						//wiemy ze chce wjechac na przejazdowy
						var id_toru int
						state := false
						for{
							for i := range tory_przejazdowe {
       	     					if tory_przejazdowe[i].id_poczatek == stac.id && tory_przejazdowe[i].id_koniec == poj.nast_stac {
									if tory_przejazdowe[i].czyWolny{
										fmt.Println(poj.id, "tor wolny wjezdzam")
										state = true
										id_toru = tory_przejazdowe[i].id
										tory_przejazdowe_chan[id_toru]<-poj		
                						break
									}
								}
							}
							if state {
								break
							}
						}
					}
					fmt.Println("Zwrotnica", stac.id, "wolna")
					stac.czyWolna = true
				}
			}
		}(stacje[i])
	}

	//gorutyna zarzadzajaca postojowymi, peronami
	for i := range tory_postojowe {
		go func(tor *tor_postojowy) {	
			state := false
			counter := 0
			for {
				select {
				case <-tor.awaria:
					counter++
					if(counter > 1){
						fmt.Println(tor.id,"Zwalniam blokade")
						state = false
					} else {
						fmt.Println(tor.id,"blokuje postojowy")
						state = true
					}
				case poj := <-tory_postojowe_chan[tor.id]:
					counter =0
					tor.czyWolny = false
					fmt.Println(poj.id, "jest na postojowym ", tor.id_stacji,"(",tor.id,")")
					if(state) {
						// czekamya z karetka nas poinformuje ze mozemy dalej jechac
						<-tor.awaria
						fmt.Println(tor.id,"zwalniam blokade")
					}
					time.Sleep(time.Duration(tor.min_czas_post)*time.Second)
					// tearz trzeba podac pojazd na stacje 
					for{
						if stacje[tor.id_stacji-1].czyWolna {
							poj.post = false
							stacje_chan[tor.id_stacji]<-poj
							break
						}
					}
					tor.czyWolny = true
				}
			}
		}(tory_postojowe[i])
	}

	// gorutyna zarzadzajaca torami przejadowymi
	for i := range tory_przejazdowe {
		go func(tor *tor_przejazdowy) {
			for {
				// sprawdzamy czy sie pospul
				t := rand.Float64()
				t2 := rand.Float64()
				t = (t+t2)/2
				if t>0.78{
					tor.czyWolny = false
					kar_chan_przej<-tor
					// tor powinien czekac az go naprawia
					<-tor.awaria
					fmt.Println("tor zostal naprawiony")
				}
				select {
				case poj := <- tory_przejazdowe_chan[tor.id]:
					tor.czyWolny = false
					fmt.Println(poj.id, "jedzie torem",tor.id_poczatek,
													"-->",tor.id_koniec)
					dl := tor.dlugosc
					pr := poj.V_max
					if pr > tor.V_max {
						pr = tor.V_max
					}
					t := dl / pr
					time.Sleep(time.Duration(t) * time.Second)
					if(poj.id != 999){
						poj.post = true
					} else {
						poj.post = false
					}
					poj.resp<-true
					for{
						if stacje[tor.id_koniec-1].czyWolna {
							stacje_chan[tor.id_koniec]<-poj
							break
						}
					}
					tor.czyWolny = true
				}		
			}
		}(tory_przejazdowe[i])
	}

	// gorutyna zarzadzajaca jednym pojazdami
	for i := range pojazdy {
		go func(poj *pojazd){
			poj.nast_stac = poj.trasa[1]
			i := 0
			fmt.Println(poj.id,"czekam na zwrotnice",poj.trasa[0])
			for{
				if(stacje[poj.trasa[0]-1].czyWolna){
					stacje_chan[poj.trasa[i]] <- poj
					break
				}
			}
			for {
				//chcemy wjechac na tor postojowy
				// dodajemy sie do kolejki czekamy az stacja nas wposci
				// dostalismy informacje zeby zwolnic tor przejazdowy
				// stacja nas wposcila na zwrotnice


				// chcialbym jechac do stacji o id = trasa[i+1]
				// czyli nastepnej na mojej trasie
				// jeszcze musze skozystac ze zwrotnicy
				// ponizszy kawalek kody powinien byc w sekcji zarzadzajcej stacja
				// dostalismy informacje zeby zwolnic postojowy

				// zamienic nastepna stacje

				select{
				case	<-poj.resp:
					fmt.Println(poj.id, "Przejechalem tor")
					i = (i+1)%len(poj.trasa)
					poj.nast_stac = poj.trasa[(i+1)%len(poj.trasa)]
				}
			}
		}(pojazdy[i])
	}

	karetka_post_chan := make(chan *pojazd)
	karetka_start_chan := make(chan bool)

	ratunek = &pojazd{999, 400, 10, make([] int,0),0, make(chan bool),true}
	go func(karetka *pojazd){
		karetka_post_chan<-karetka
		docelowy := 0
		i:=0
		var tor_prze *tor_przejazdowy
		for{
			select {
			case tor:=<-kar_chan_post:
				fmt.Println(karetka.id,"Awaria toru postojowego", tor.id)
			case tor:=<-kar_chan_przej:
				fmt.Println(karetka.id,"Awaria toru przejazdowego", tor.id_poczatek, tor.id_koniec)
				// musimy jeszcze ustalic trase
				karetka.trasa = find_path(1, tor.id_poczatek)
				docelowy = tor.id_poczatek
				fmt.Println(karetka.trasa)
				karetka.nast_stac = karetka.trasa[1]
				//informujemy zeby zarezerwowac dla nas trase dla postojowych 
				for _,n := range tory_postojowe {
					for _,k := range karetka.trasa {
						if(n.id_stacji == k){
							n.awaria<-true
						}
					}
				}

				// wysylamy do peronu zeby nas poscil
				karetka_start_chan<-true
				time.Sleep(time.Second * 10)
				tor_prze = tor
			case poj:=<-kar_chan_poj:
				fmt.Println(karetka.id, "Awaria pojazdu", poj.id)
			case stac:=<-kar_chan_stac:
				fmt.Println(karetka.id, "Awaria stacji", stac.id)
			case	<-karetka.resp:
				fmt.Println(karetka.id, "Przejechalem tor")
				i = (i+1)%len(karetka.trasa)
				if(karetka.trasa[i]== docelowy) {
					fmt.Println(karetka.id, "na miejscu zaczyna naprawe")
					tor_prze.awaria<-true
					for _,n := range tory_postojowe {
						for _,k := range karetka.trasa {
							if(n.id_stacji == k){
								n.awaria<-true
							}
						}
					}
					i = 0
					karetka.trasa = find_back(tor_prze.id_poczatek,1)
					karetka.nast_stac = karetka.trasa[1]
				}
				karetka.nast_stac = karetka.trasa[(i+1)%len(karetka.trasa)]
			}
		}
	}(ratunek)


	karetka_post := &tor_postojowy{999, 1,300,false,make(chan bool)}
	go func(tor *tor_postojowy) {	
		for {
			select {
			case poj := <-karetka_post_chan:
				tor.czyWolny = false
				fmt.Println(poj.id, "jest na postojowym ", tor.id_stacji,"(",tor.id,")")
				<-karetka_start_chan
				fmt.Println("karetka rusza")
				// tearz trzeba podac pojazd na stacje 
				for{
					if stacje[tor.id_stacji-1].czyWolna {
						poj.post = false
						stacje_chan[tor.id_stacji]<-poj
						break
					}
				}
				tor.czyWolny = true
			}
		}
	}(karetka_post)

	var input string
    fmt.Scanln(&input)
    fmt.Println("done")

}
