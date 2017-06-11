package main
import "io/ioutil"
import "strings"
import "fmt"
import "strconv"

func getModel(filepath string) *Simulation_Model {

	var model_ptr *Simulation_Model = new(Simulation_Model)

	dat,_ := ioutil.ReadFile(filepath);
	data := string(dat)
	it := 0

	arr := strings.Split(data, "\n")
	fmt.Println("Czas: ", arr[0])
	speed, _ := strconv.Atoi(arr[0])
	model_ptr.speed = speed
	it++
	

	more_data := strings.Split(arr[it]," ")
	il_stacji, _ := strconv.Atoi(more_data[0])
	fmt.Println("Ilosc stacji: ", il_stacji)
	counter := 1
	for i := 1; i < (il_stacji+1); i++ {
		model_ptr.stations = append(model_ptr.stations, newStation(counter, more_data[i]))
		counter++
	}

	for i,_ := range model_ptr.stations {
		fmt.Println("Stajce: ", model_ptr.stations[i])
	}
	it++
	il_przejazdowych,_ := strconv.Atoi(arr[it])
	it++
	for i := 0; i<il_przejazdowych; i++{
		more_data := strings.Split(arr[it])
		
		it++
	}
	return nil
}