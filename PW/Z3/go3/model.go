package main
import "time"

type Simulation_Mode int64

type Log_Mode int64

type Time_Interval int64

type Simulation_Model struct {
	speed int
	start_time time.Time
	mode Simulation_Mode

	log_mode Log_Mode

//	workers []*Worker
	// Przejazdowe
//	tracks []*Track
	//Postojowe
//	platforms []*Track
//	trains []*Train
	stations []*Station

	work bool

}