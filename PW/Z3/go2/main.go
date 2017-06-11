package main

import (
	"fmt"
	"time"
	"sync"
	"strconv"
	"math/rand"
	"os"
	"reflect"
	"bufio"
	"regexp"
	"strings"
	"math"
)

func main() {
	Simulation_start()
}

//History record
type Track_History struct {
	train_id  int64
	arrival   time.Time
	departure time.Time
}

//Types of tracks
type Track_Type int64

const (
	Track_Type_Unknown  Track_Type = iota
	Track_Type_Track    Track_Type = iota
	Track_Type_Platform Track_Type = iota
	Track_Type_Service  Track_Type = iota
)

//Keys for additional fields.
const (
	T_distance   string = "dist"
	T_max_speed  string = "speed"
	T_min_delay  string = "delay"
	T_station_id string = "stat"
)

//Track record with all necessary data. Additional fields based on type are stored in data field
type Track struct {
	id       int64
	t_type   Track_Type
	st_start int64
	st_end   int64
	used_by  int64
	data     map[string]int64

	out_of_order bool
	reliability  float64

	//accepts given train thus blocking track for others
	acceptTrain chan int64
	//clears out block for other trains after currently blocking train left the track.
	clearAfterTrain chan int64

	//service entries
	allowServiceTrain    chan int64
	acceptServiceTrain   chan int64
	repair               chan int64
	freeFromServiceTrain chan int64

	history []Track_History
}

//Initialises common fields for all tracks
func initTrack(id int64, st_start int64, st_end int64) *Track {
	t := new(Track)
	t.id = id
	t.out_of_order = false
	t.reliability = 0.99995
	t.st_start = st_start
	t.st_end = st_end
	t.used_by = 0
	t.data = make(map[string]int64)
	t.history = make([]Track_History, 0)
	t.acceptTrain = make(chan int64)
	t.clearAfterTrain = make(chan int64)
	t.allowServiceTrain = make(chan int64)
	t.acceptServiceTrain = make(chan int64)
	t.repair = make(chan int64)
	t.freeFromServiceTrain = make(chan int64)
	return t
}

//Creates new Track of Track_Type_Track type.
func NewTrack(id int64, st_start int64, st_end int64, distance int64, speed int64) *Track {
	t := initTrack(id, st_start, st_end)
	t.t_type = Track_Type_Track
	t.data[T_distance] = distance
	t.data[T_max_speed] = speed
	return t
}

//Creates new Track of Track_Type_Platform.
func NewPlatform(id int64, st_start int64, st_end int64, min_delay int64, station_id int64) *Track {
	t := initTrack(id, st_start, st_end)
	t.t_type = Track_Type_Platform
	t.data[T_min_delay] = min_delay
	t.data[T_station_id] = station_id
	return t
}

//Creates new Track of Track_Type_Service.
func NewServiceTrack(id int64, st_start int64, st_end int64) *Track {
	t := initTrack(id, st_start, st_end)
	t.t_type = Track_Type_Service
	return t
}

//Track task. Allows accepted trains to move/wait on them
func TrackTask(track_ptr *Track, model_ptr *Simulation_Model) {
	var type_str string
	var train_ptr *Train
	var hist *Track_History = nil

	var r *rand.Rand = rand.New(rand.NewSource(int64(time.Now().Nanosecond() + int(2*track_ptr.id) + 16384)))
	var work bool = false

	if track_ptr != nil && model_ptr != nil {
		//track naming
		if track_ptr.t_type == Track_Type_Platform {
			type_str = "Platform"
		} else if track_ptr.t_type == Track_Type_Track {
			type_str = "Track"
		} else if track_ptr.t_type == Track_Type_Service {
			type_str = "Service Track"
		} else {
			type_str = "Unknown"
		}

		var help_service_train_ptr *Train = nil
		var pass_service_train_ptr *Train = nil
		var help bool = false
		for model_ptr.work {

			if track_ptr.out_of_order == true && help == false {
				help_service_train_ptr = GetServiceTrain(model_ptr)
				if help_service_train_ptr != nil {

					select {
					case help_service_train_ptr.trackOutOfOrder <- track_ptr.id:
						help = true
					default:
						//  delay Standard.Duration(1);
					}

				} else {
					//Ada.Text_IO.Put_Line(ustr.To_String(type_str)&"["&Positive'Image(track_ptr.id)&"] received null pointer for service train" );
					fmt.Println("#2# " + type_str + "[" + strconv.FormatInt(track_ptr.id, 10) + "] received null pointer for service train")
				}
			}

			select {

			case train_id := <-when(track_ptr.used_by == 0, track_ptr.allowServiceTrain):
				pass_service_train_ptr = GetTrain(train_id, model_ptr)
				if pass_service_train_ptr != nil {
					track_ptr.used_by = train_id
					PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] received accept request from service train ["+strconv.FormatInt(pass_service_train_ptr.id, 10)+
						"]. Blocking track for other trains.", model_ptr)
				} else {
					fmt.Println("#2# " + type_str + "[" + strconv.FormatInt(track_ptr.id, 10) +
						"] received null pointer for service train ID[" + strconv.FormatInt(train_id, 10) + "]")
				}

			case train_id := <-track_ptr.acceptServiceTrain:
				if pass_service_train_ptr != nil && pass_service_train_ptr.id == train_id {
					train_ptr = pass_service_train_ptr
					work = true

					hist = new(Track_History) //_Record
					hist.arrival = time.Now()
					hist.train_id = train_id

					PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] blocked by passing service train: ["+strconv.FormatInt(pass_service_train_ptr.id, 10)+
						"]", model_ptr)
				} else {
					if track_ptr.used_by == 0 || track_ptr.used_by == train_id {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
							"] received accept signal from invalid serivce train ID["+strconv.FormatInt(train_id, 10)+
							"] no service train expected. Blocking the track anyway.", model_ptr)
						pass_service_train_ptr = GetTrain(train_id, model_ptr)
						track_ptr.used_by = train_id
						train_ptr = pass_service_train_ptr
						work = true

						hist = new(Track_History) //_Record
						hist.arrival = time.Now()
						hist.train_id = train_id

					} else {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
							"] received accept signal from invalid serivce train ID["+strconv.FormatInt(train_id, 10)+
							"] no service train expected. Currently used by other train.", model_ptr)
					}
				}

			case train_id := <-track_ptr.freeFromServiceTrain:
				if pass_service_train_ptr != nil && pass_service_train_ptr.id == train_id {
					PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] unblocked from service train["+strconv.FormatInt(pass_service_train_ptr.id, 10)+"].",
						model_ptr)
					track_ptr.used_by = 0
					train_ptr = nil
					pass_service_train_ptr = nil
				} else {
					if pass_service_train_ptr == nil {
						fmt.Println("#2# " + type_str + "[" + strconv.FormatInt(track_ptr.id, 10) +
							"] receive free signal from invalid serivce train ID[" + strconv.FormatInt(train_id, 10) +
							"] no service train accepted.")
					} else {
						fmt.Println("#2# " + type_str + "[" + strconv.FormatInt(track_ptr.id, 10) +
							"] receive free signal from invalid serivce train ID[" + strconv.FormatInt(train_id, 10) +
							"] accepted: [" + strconv.FormatInt(pass_service_train_ptr.id, 10) + "]")
					}
				}
			case train_id := <-track_ptr.repair:
				if help_service_train_ptr != nil && help_service_train_ptr.id == train_id {
					PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] was just repaired. Ready to accept incoming trains anew.",
						model_ptr)
					track_ptr.out_of_order = false
				} else {
					if help_service_train_ptr != nil {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
							"] has no information about service train but received repair signal from service train["+strconv.FormatInt(train_id, 10)+
							"]. Accepting the repair and moving along with schedule.",
							model_ptr)
						track_ptr.out_of_order = false
					} else {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
							"] received repair signal from illegal service train ["+strconv.FormatInt(train_id, 10)+
							"]. Accepting the repair and moving along with schedule.",
							model_ptr)
						track_ptr.out_of_order = false
					}
				}

			//accepts given train thus blocking track for others
			case train_id := <-when(track_ptr.out_of_order == false && train_ptr == nil, track_ptr.acceptTrain):
				train_ptr = GetTrain(train_id, model_ptr)
				if train_ptr != nil {
					hist = new(Track_History) //_Record
					hist.arrival = time.Now()
					hist.train_id = train_id
					work = true
					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] is now blocked by train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)

					track_ptr.used_by = train_ptr.id
				} else {
					fmt.Println(type_str + "[" + strconv.FormatInt(track_ptr.id, 10) +
						"] received null pointer for train ID[" + strconv.FormatInt(train_id, 10) + "]")
				}

			//otherwise waits for train to clear out the steering.
			//clears out block for other trains after currently blocking train left the track.
			case train_id := <-when(track_ptr.out_of_order == false && train_ptr != nil, track_ptr.clearAfterTrain):
				if train_ptr.id == train_id {
					hist.departure = time.Now()
					track_ptr.history = append(track_ptr.history, *hist)

					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] is now unblocked from train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)

					train_ptr = nil
					track_ptr.used_by = 0
				} else {
					fmt.Println(type_str + "[" + strconv.FormatInt(track_ptr.id, 10) +
						"] received clear out signal from invalid train:[" + strconv.FormatInt(train_id, 10) +
						"], currently used by:[" + strconv.FormatInt(track_ptr.used_by, 10) + "]")
				}
			case <-time.After(time.Second * time.Duration(GetTimeSimToRealFromModel(1.0, Time_Interval_Hour, model_ptr))):

			}

			//for given train waits specified duration and then signals the train that it's ready to depart from this track.

			if track_ptr.out_of_order == false && train_ptr != nil && work {
				work = false
				var delay_dur, real_delay_dur float64
				//track delay based on track type
				if track_ptr.t_type == Track_Type_Track {
					//for normal tracks checks speed the train can move on this track.
					if track_ptr.data[T_max_speed] < train_ptr.max_speed {
						train_ptr.current_speed = track_ptr.data[T_max_speed]
						delay_dur = float64(track_ptr.data[T_distance]) / float64(track_ptr.data[T_max_speed])
						real_delay_dur = GetTimeSimToRealFromModel(delay_dur, Time_Interval_Hour, model_ptr)

						PutLine("Train["+strconv.FormatInt(train_ptr.id, 10)+
							"] moves on track["+strconv.FormatInt(track_ptr.id, 10)+
							"] with track max speed for next"+strconv.FormatFloat(delay_dur, 'f', 3, 64)+
							" hours ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

					} else {
						train_ptr.current_speed = train_ptr.max_speed
						delay_dur = float64(track_ptr.data[T_distance]) / float64(train_ptr.max_speed)
						real_delay_dur = GetTimeSimToRealFromModel(delay_dur, Time_Interval_Hour, model_ptr)

						PutLine("Train["+strconv.FormatInt(train_ptr.id, 10)+
							"] moves with its top speed on track["+strconv.FormatInt(track_ptr.id, 10)+
							"] for next "+strconv.FormatFloat(delay_dur, 'f', 3, 64)+" hours ("+
							strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

					}
					//delay for tracks
					time.Sleep(time.Duration(real_delay_dur) * time.Second)

					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] signals the train:["+strconv.FormatInt(train_ptr.id, 10)+
						"] that it's ready to depart onto next steering", model_ptr)
					//notify train
					train_ptr.trainArrivedToTheEndOfTrack <- track_ptr.id /* chan_ready <- TrainMessageFromTrack(track_ptr.id)*/
					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] signaled the train:["+strconv.FormatInt(train_ptr.id, 10)+
						"] that it's ready to depart onto next steering", model_ptr)
					//notify train
				} else if track_ptr.t_type == Track_Type_Platform && pass_service_train_ptr == nil {
					train_ptr.current_speed = 0
					delay_dur := float64(track_ptr.data[T_min_delay])
					real_delay_dur := GetTimeSimToRealFromModel(delay_dur, Time_Interval_Minute, model_ptr)

					//fmt.Println("["&sim_delay_str&"]");

					PutLine("Train["+strconv.FormatInt(train_ptr.id, 10)+
						"] waits on platform["+strconv.FormatInt(track_ptr.id, 10)+
						"] for next "+strconv.FormatFloat(delay_dur, 'f', 3, 64)+
						" minutes ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

					//delay for platforms
					time.Sleep(time.Duration(real_delay_dur) * time.Second)

					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] signals the train train["+strconv.FormatInt(train_ptr.id, 10)+
						"] that it's ready to depart onto next steering", model_ptr)
					//notify train
					//train_ptr.chan_ready <- TrainMessageFromPlatform(track_ptr.id)
					train_ptr.trainReadyToDepartFromPlatform <- track_ptr.id /* chan_ready <- TrainMessageFromTrack(track_ptr.id)*/
					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] signaled the train train["+strconv.FormatInt(train_ptr.id, 10)+
						"] that it's ready to depart onto next steering", model_ptr)
				} else { //for service tracks and platforms on which service trains moves on
					train_ptr.current_speed = 0
					delay_dur := 1.0
					real_delay_dur := GetTimeSimToRealFromModel(delay_dur, Time_Interval_Minute, model_ptr)

					PutLine("Train["+strconv.FormatInt(train_ptr.id, 10)+
						"] waits on "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] for 1 minute ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

					//delay for unknown tracks
					time.Sleep(time.Duration(real_delay_dur) * time.Second)

					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] signals the train["+strconv.FormatInt(train_ptr.id, 10)+
						"] that it's ready to depart onto next steering", model_ptr)
					//notify train
					//train_ptr.t_task.trainArrivedToTheEndOfTrack(track_ptr.id)
					train_ptr.trainArrivedToTheEndOfTrack <- track_ptr.id /* chan_ready <- TrainMessageFromTrack(track_ptr.id)*/
					PutLine(type_str+"["+strconv.FormatInt(track_ptr.id, 10)+
						"] signaled the train["+strconv.FormatInt(train_ptr.id, 10)+
						"] that it's ready to depart onto next steering", model_ptr)
				}
			}

			if track_ptr.t_type != Track_Type_Service && track_ptr.used_by == 0 && track_ptr.out_of_order == false {
				ran := r.Float64()
				//fmt.Println(type_str + "[" + strconv.FormatInt(track_ptr.id, 10) + "] rolled " + strconv.FormatFloat(ran, 'f', 3, 64) + " at time " + TimeToString(GetRelativeTime(time.Now(), model_ptr)))

				if track_ptr.reliability < ran {
					PutLine("#2# "+type_str+"["+strconv.FormatInt(track_ptr.id, 10)+"] broke at time "+TimeToString(GetRelativeTime(time.Now(), model_ptr)), model_ptr)
					track_ptr.out_of_order = true
					help = false
				}
			}

		}
		fmt.Println(type_str + "[" + strconv.FormatInt(track_ptr.id, 10) + "] terminates its execution")
	} else {
		fmt.Println("TrackTask received null pointer! Task will not work")
	}
}

type Station struct {
	id int64

	mutex *sync.RWMutex

	trains         map[int64]bool
	passengers     map[*Worker]bool
	ready_workers  map[*Worker]bool
	chosen_workers map[*Worker]bool

	notifyAboutWorkerArrival   chan int64
	notifyAboutWorkerDeparture chan int64

	notifyAboutTrainArrival   chan int64
	notifyAboutTrainDeparture chan int64

	notifyAboutReadinessToWork  chan int64
	notifyAboutFinishingTheWork chan int64
}

func newStation(stat_id int64) *Station {
	t := new(Station)
	t.id = stat_id
	t.trains = make(map[int64]bool)
	t.passengers = make(map[*Worker]bool)
	t.ready_workers = make(map[*Worker]bool)
	t.chosen_workers = make(map[*Worker]bool)

	t.mutex = new(sync.RWMutex)

	t.notifyAboutWorkerArrival = make(chan int64)
	t.notifyAboutWorkerDeparture = make(chan int64)

	t.notifyAboutTrainArrival = make(chan int64)
	t.notifyAboutTrainDeparture = make(chan int64)

	t.notifyAboutReadinessToWork = make(chan int64)
	t.notifyAboutFinishingTheWork = make(chan int64)

	return t
}

func StationTask(stat_ptr *Station, model_ptr *Simulation_Model) {

	var ran float64
	var work_ptr *Worker
	var created_task bool = false
	var active_task bool = false

	var notify_train int64 = 0

	if model_ptr != nil && stat_ptr != nil {
		var r *rand.Rand = rand.New(rand.NewSource(int64(time.Now().Nanosecond() + int(10*stat_ptr.id+65536))))
		for model_ptr.work {
			// fmt.Println("#3#%%%%%%%%%%%%%%%%%%%%%%%%% { ### before select" );
			select {
			case train_id := <-stat_ptr.notifyAboutTrainArrival:
				stat_ptr.mutex.Lock()
				if _, ok := stat_ptr.trains[train_id]; !ok {
					PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] welcomes train["+strconv.FormatInt(train_id, 10)+"]", model_ptr)
					stat_ptr.trains[train_id] = true
					notify_train = train_id
				} else {
					fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal arrival notification from train[" + strconv.FormatInt(train_id, 10) + "]")
				}
				stat_ptr.mutex.Unlock()

			case train_id := <-stat_ptr.notifyAboutTrainDeparture:
				stat_ptr.mutex.Lock()
				if _, ok := stat_ptr.trains[train_id]; ok {
					PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] bids farewell to train["+strconv.FormatInt(train_id, 10)+"]", model_ptr)
					delete(stat_ptr.trains, train_id)
				} else {
					fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal departure notification from train[" + strconv.FormatInt(train_id, 10) + "]")
				}
				stat_ptr.mutex.Unlock()

			case work_id := <-stat_ptr.notifyAboutWorkerArrival:
				work_ptr = GetWorker(work_id, model_ptr)
				if work_ptr != nil {
					if _, ok := stat_ptr.passengers[work_ptr]; !ok {
						stat_ptr.passengers[work_ptr] = true
						PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] welcomes passenger["+strconv.FormatInt(work_id, 10)+"]", model_ptr)
					} else {
						fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal arrival notification from worker[" + strconv.FormatInt(work_id, 10) + "]")
					}
				} else {
					fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received nil pointer for worker[" + strconv.FormatInt(work_id, 10) + "]")
				}

			case work_id := <-stat_ptr.notifyAboutWorkerDeparture:
				work_ptr = GetWorker(work_id, model_ptr)
				if work_ptr != nil {
					if _, ok := stat_ptr.passengers[work_ptr]; ok {
						PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] bids farewell to passenger["+strconv.FormatInt(work_id, 10)+"]", model_ptr)
						delete(stat_ptr.passengers, work_ptr)

						if stat_ptr.ready_workers[work_ptr] {
							delete(stat_ptr.ready_workers, work_ptr)
							fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal departure notification from worker[" + strconv.FormatInt(work_id, 10) + "] before he finished task.")
						}
						if stat_ptr.chosen_workers[work_ptr] {
							delete(stat_ptr.chosen_workers, work_ptr)
							fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal departure notification from worker[" + strconv.FormatInt(work_id, 10) + "] before he started task.")
						}

					} else {
						// log.printStations(model_ptr);
						// log.printWorkers(model_ptr);
						fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal departure notification from worker[" + strconv.FormatInt(work_id, 10) + "]")

					}
				} else {
					fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received nil pointer for worker[" + strconv.FormatInt(work_id, 10) + "]")
				}

			case work_id := <-when(created_task,
				stat_ptr.notifyAboutReadinessToWork):
				work_ptr = GetWorker(work_id, model_ptr)
				if work_ptr != nil {

					_, ok1 := stat_ptr.passengers[work_ptr]
					_, ok2 := stat_ptr.chosen_workers[work_ptr]
					_, ok3 := stat_ptr.ready_workers[work_ptr]

					if ok1 && ok2 && !ok3 {
						PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] received notification that worker["+strconv.FormatInt(work_id, 10)+"] is ready to work", model_ptr)
						stat_ptr.ready_workers[work_ptr] = true
					} else {
						fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal ready notification from worker[" + strconv.FormatInt(work_id, 10) + "]")
					}
				} else {
					fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received nil pointer for worker[" + strconv.FormatInt(work_id, 10) + "]")
				}

			case work_id := <-when(active_task,
				stat_ptr.notifyAboutFinishingTheWork):
				work_ptr = GetWorker(work_id, model_ptr)
				if work_ptr != nil {
					if stat_ptr.ready_workers[work_ptr] && stat_ptr.chosen_workers[work_ptr] {
						PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] received notification that worker["+strconv.FormatInt(work_id, 10)+"] finished his task", model_ptr)
						delete(stat_ptr.ready_workers, work_ptr)
						delete(stat_ptr.chosen_workers, work_ptr)
					} else {
						if stat_ptr.ready_workers[work_ptr] {
							delete(stat_ptr.ready_workers, work_ptr)
							fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] had worker:" + strconv.FormatInt(work_id, 10) + " only within ready worker pool.")
						} else if stat_ptr.chosen_workers[work_ptr] {
							fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal finish notification from worker[" + strconv.FormatInt(work_id, 10) + "] before he started task.")
						} else {
							fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received illegal finish notification from worker[" + strconv.FormatInt(work_id, 10) + "] .")
						}

					}
				} else {
					fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] received nil pointer for worker[" + strconv.FormatInt(work_id, 10) + "]")
				}

			case <-time.After(time.Second * time.Duration(GetTimeSimToRealFromModel(1, Time_Interval_Hour, model_ptr))):
			}

			if notify_train != 0 {
				for cur := range stat_ptr.passengers {
					work_ptr = cur
					select {
					case work_ptr.notifyAboutTrainArrival <- TSPMessage(notify_train, stat_ptr.id):
					default:
					}
				}
				notify_train = 0
			}

			if created_task {
				if reflect.DeepEqual(stat_ptr.ready_workers, stat_ptr.chosen_workers) {

					ran = 5.0 + 5.0*r.Float64()

					PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] got notifications from all chosen workers. Notifying workers thath they can start working for next "+strconv.FormatFloat(ran, 'f', 3, 64)+" hours.", model_ptr)
					created_task = false
					active_task = true
					for cur := range stat_ptr.ready_workers {
						work_ptr = cur
						work_ptr.startTask <- STSMessage(stat_ptr.id, ran)
					}
				} else {
					if len(stat_ptr.ready_workers) == len(stat_ptr.chosen_workers) {
						fmt.Println("#3# station[" + strconv.FormatInt(stat_ptr.id, 10) + "] has same amount of ready and chosen workers but sets are different")
					}
				}
			} else if (!active_task) && (!created_task) {
				ran = r.Float64()
				if ran < 0.05 {
					PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] needs to have new task performed. Notifying workers", model_ptr)
					var worker_count int64 = 0
					for it := 0; it < len(model_ptr.worker); it++ {
						work_ptr = model_ptr.worker[it]
						if work_ptr.state == AtHome && r.Float64() < 0.25 {
							select {
							case work_ptr.acceptTask <- stat_ptr.id:
								//PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] choose worker["+strconv.FormatInt(work_ptr.id, 10)+"] for task.", model_ptr)
								stat_ptr.chosen_workers[work_ptr] = true
								worker_count = worker_count + 1
							case <-time.After(time.Second * time.Duration(1)):
								//PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] couldnt choose worker["+strconv.FormatInt(work_ptr.id, 10)+"] for task.", model_ptr)
							}

						}
					}
					if worker_count > 0 {
						PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] choose "+strconv.FormatInt(worker_count, 10)+" workers for task.", model_ptr)
						created_task = true
					} else {
						PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] could not choose any workers for this task. Abandoning task.", model_ptr)

					}
				}
			} else if active_task {
				if len(stat_ptr.chosen_workers) < 0 {
					PutLine("#3# station["+strconv.FormatInt(stat_ptr.id, 10)+"] and all workers finished performing a task.", model_ptr)
					active_task = false
				}
			}

		}
	} else {
		fmt.Println("#3# StationTask received nil pointer.")
	}

}

//history record
type Steering_History struct {
	train_id  int64
	arrival   time.Time
	departure time.Time
}

//Steering record with all necessary data.
type Steering struct {
	id        int64
	min_delay int64
	used_by   int64

	out_of_order bool
	reliability  float64

	history []Steering_History

	//accepts given train thus blocking steering for others
	acceptTrain chan int64
	//clears out block for other trains after currently blocking train left the steering.
	clearAfterTrain chan int64

	allowServiceTrain    chan int64
	acceptServiceTrain   chan int64
	repair               chan int64
	freeFromServiceTrain chan int64
}

//creates new steering
func NewSteering(id int64, min_delay int64) *Steering {
	t := new(Steering)
	t.out_of_order = false
	t.reliability = 0.99995
	t.id = id
	t.min_delay = min_delay
	t.used_by = 0
	t.acceptTrain = make(chan int64)
	t.clearAfterTrain = make(chan int64)
	t.allowServiceTrain = make(chan int64)
	t.acceptServiceTrain = make(chan int64)
	t.repair = make(chan int64)
	t.freeFromServiceTrain = make(chan int64)
	return t
}

// Steering task. Allows accepted trains to switch onto their next track.
func SteeringTask(steering_ptr *Steering, model_ptr *Simulation_Model) {
	var train_ptr *Train
	var hist *Steering_History = nil

	var r *rand.Rand = rand.New(rand.NewSource(int64(time.Now().Nanosecond() + int(3*steering_ptr.id) + 256)))

	if steering_ptr != nil && model_ptr != nil {
		var help_service_train_ptr *Train = nil
		var pass_service_train_ptr *Train = nil
		var help bool = false
		var work bool = false

		for model_ptr.work {
			PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
				"] starts new loop", model_ptr)

			if steering_ptr.out_of_order == true && help == false {
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] calls for help", model_ptr)
				help_service_train_ptr = GetServiceTrain(model_ptr)
				if help_service_train_ptr != nil {

					select {
					case help_service_train_ptr.steeringOutOfOrder <- steering_ptr.id:
						help = true
					default:
						//  delay Standard.Duration(1);
					}

				} else {
					//Ada.Text_IO.Put_Line(ustr.To_String(type_str)&"["+strconv.FormatInt(steering_ptr.id, 10)+"] received null pointer for service train" );
					fmt.Println("#2# Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
						"] received null pointer for service train")
				}
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] called for help", model_ptr)
			}
			PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
				"] enters select", model_ptr)

			select {

			case train_id := <-when(steering_ptr.used_by == 0, steering_ptr.allowServiceTrain):
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] enters allowServiceTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
				pass_service_train_ptr = GetTrain(train_id, model_ptr)
				if pass_service_train_ptr != nil {
					steering_ptr.used_by = train_id
					PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] received accept request from service train ["+strconv.FormatInt(pass_service_train_ptr.id, 10)+
						"]. Blocking track for other trains.", model_ptr)
				} else {
					fmt.Println("#2# Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
						"] received null pointer for service train ID[" + strconv.FormatInt(train_id, 10) + "]")
				}
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] leaves allowServiceTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)

			case train_id := <-steering_ptr.acceptServiceTrain:
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] enters acceptServiceTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
				if pass_service_train_ptr != nil && pass_service_train_ptr.id == train_id {
					PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] blocked by passing service train: ["+strconv.FormatInt(pass_service_train_ptr.id, 10)+
						"]", model_ptr)
					train_ptr = pass_service_train_ptr

					hist = new(Steering_History) //_Record
					hist.arrival = time.Now()
					hist.train_id = train_id
					work = true
				} else {
					if pass_service_train_ptr != nil {
						if steering_ptr.used_by == 0 || steering_ptr.used_by == train_id {
							PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
								"] received accept signal from invalid serivce train ID["+strconv.FormatInt(train_id, 10)+
								"] no service train expected. Blocking the track anyway.", model_ptr)
							pass_service_train_ptr = GetTrain(train_id, model_ptr)
							steering_ptr.used_by = train_id
							train_ptr = pass_service_train_ptr

							hist = new(Steering_History) //_Record
							hist.arrival = time.Now()
							hist.train_id = train_id
							work = true

						} else {
							PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
								"] received accept signal from invalid serivce train ID["+strconv.FormatInt(train_id, 10)+
								"] no service train expected. Currently used by other train.", model_ptr)
						}
					} else {
						PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
							"] received accept signal from invalid serivce train["+strconv.FormatInt(train_id, 10)+
							"] expected: ["+strconv.FormatInt(pass_service_train_ptr.id, 10)+"]", model_ptr)
					}
				}
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] leaves acceptServiceTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)

			case train_id := <-steering_ptr.freeFromServiceTrain:
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] enters freeFromServiceTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
				if pass_service_train_ptr != nil && pass_service_train_ptr.id == train_id {
					PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] unblocked from service train["+strconv.FormatInt(pass_service_train_ptr.id, 10)+"].",
						model_ptr)
					steering_ptr.used_by = 0
					train_ptr = nil
					pass_service_train_ptr = nil
				} else {
					if pass_service_train_ptr == nil {
						fmt.Println("#2# Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
							"] receive free signal from invalid serivce train ID[" + strconv.FormatInt(train_id, 10) +
							"] no service train accepted.")
					} else {
						fmt.Println("#2# Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
							"] receive free signal from invalid serivce train ID[" + strconv.FormatInt(train_id, 10) +
							"] accepted: [" + strconv.FormatInt(pass_service_train_ptr.id, 10) + "]")
					}
				}
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] leaves freeFromServiceTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
			case train_id := <-steering_ptr.repair:
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] enters repair("+strconv.FormatInt(train_id, 10)+")", model_ptr)
				if help_service_train_ptr != nil && help_service_train_ptr.id == train_id {
					PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] was just repaired. Ready to accept incoming trains anew.",
						model_ptr)
					steering_ptr.out_of_order = false
				} else {
					if help_service_train_ptr != nil {
						PutLine("#2# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
							"] has no information about service train but received repair signal from service train["+strconv.FormatInt(train_id, 10)+
							"]. Accepting the repair and moving along with schedule.",
							model_ptr)
						steering_ptr.out_of_order = false
					} else {
						PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
							"] received repair signal from illegal service train ["+strconv.FormatInt(train_id, 10)+
							"]. Accepting the repair and moving along with schedule.",
							model_ptr)
						steering_ptr.out_of_order = false
					}
				}

				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] leaves repair("+strconv.FormatInt(train_id, 10)+")", model_ptr)
			//accepts given train thus blocking track for others
			case train_id := <-when(steering_ptr.out_of_order == false && train_ptr == nil, steering_ptr.acceptTrain):
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] enters acceptTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
				train_ptr = GetTrain(train_id, model_ptr)
				if train_ptr != nil {
					hist = new(Steering_History) //_Record
					hist.arrival = time.Now()
					hist.train_id = train_id
					work = true
					PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] is now blocked by train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)

					steering_ptr.used_by = train_ptr.id
				} else {
					fmt.Println("Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
						"] received null pointer for train ID[" + strconv.FormatInt(train_id, 10) + "]")
				}
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] leaves acceptTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)

			//otherwise waits for train to clear out the steering.
			//clears out block for other trains after currently blocking train left the track.
			case train_id := <-when(steering_ptr.out_of_order == false && train_ptr != nil, steering_ptr.clearAfterTrain):
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] enters clearAfterTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
				if train_ptr.id == train_id {
					hist.departure = time.Now()
					steering_ptr.history = append(steering_ptr.history, *hist)

					PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] is now unblocked from train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)

					train_ptr = nil
					steering_ptr.used_by = 0
				} else {
					fmt.Println("Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
						"] received clear out signal from invalid train:[" + strconv.FormatInt(train_id, 10) +
						"], currently used by:[" + strconv.FormatInt(steering_ptr.used_by, 10) + "]")
				}
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] leaves clearAfterTrain("+strconv.FormatInt(train_id, 10)+")", model_ptr)
			case <-time.After(time.Second * time.Duration(GetTimeSimToRealFromModel(1.0, Time_Interval_Hour, model_ptr))):

				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] had timeout in select", model_ptr)
			}
			PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
				"] leaves select", model_ptr)

			//fmt.Println("Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
			//	"] is after first if and has train " + strconv.FormatInt(train_ptr.id, 10))

			//for given train waits specified duration and { signals the train that it's ready to depart from this steering.
			if steering_ptr.out_of_order == false && train_ptr != nil && work {
				work = false
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] starts processing train routine ", model_ptr)
				delay_dur := float64(steering_ptr.min_delay)
				real_delay_dur := GetTimeSimToRealFromModel(delay_dur, Time_Interval_Minute, model_ptr)

				PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] with train["+strconv.FormatInt(train_ptr.id, 10)+
					"] switches tracks for "+strconv.FormatFloat(delay_dur, 'f', 3, 64)+
					" minutes ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

				//steering delay

				time.Sleep(time.Duration(real_delay_dur) * time.Second)

				PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] signals the train["+strconv.FormatInt(train_ptr.id, 10)+
					"] that it's ready to depart onto next track", model_ptr)

				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] starts train["+strconv.FormatInt(train_ptr.id, 10)+
					"].trainReadyToDepartFromSteering("+strconv.FormatInt(steering_ptr.id, 10)+")", model_ptr)
				train_ptr.trainReadyToDepartFromSteering <- steering_ptr.id
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] finished train["+strconv.FormatInt(train_ptr.id, 10)+
					"].trainReadyToDepartFromSteering("+strconv.FormatInt(steering_ptr.id, 10)+")", model_ptr)
				PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] signaled the train["+strconv.FormatInt(train_ptr.id, 10)+
					"] that it's ready to depart onto next track", model_ptr)
				PutLine("#debug# Steering["+strconv.FormatInt(steering_ptr.id, 10)+
					"] finishes processing train routine ", model_ptr)

			}
			//fmt.Println("Steering[" + strconv.FormatInt(steering_ptr.id, 10) +
			//	"] is after second if and has train " + strconv.FormatInt(train_ptr.id, 10))

			if steering_ptr.used_by == 0 && steering_ptr.out_of_order == false {
				ran := r.Float64()
				//fmt.Println("Steering[" + strconv.FormatInt(steering_ptr.id, 10) + "] rolled " + strconv.FormatFloat(ran, 'f', 3, 64) + " at time " + TimeToString(GetRelativeTime(time.Now(), model_ptr)))

				if steering_ptr.reliability < ran {
					PutLine("Steering["+strconv.FormatInt(steering_ptr.id, 10)+
						"] broke at time "+TimeToString(GetRelativeTime(time.Now(), model_ptr)), model_ptr)
					steering_ptr.out_of_order = true
					help = false
				}
			}
		}
		fmt.Println("Steering[" + strconv.FormatInt(steering_ptr.id, 10) + "] terminates its execution")
	} else {
		fmt.Println("SteeringTask received null pointer! Task will not work")
	}

}

type Railway_Object_Type int64

//Object types for history
const (
	Type_Steering Railway_Object_Type = iota
	Type_Track    Railway_Object_Type = iota
	Type_Platform Railway_Object_Type = iota
	Type_Train    Railway_Object_Type = iota
	Type_Service  Railway_Object_Type = iota
	Type_Unknown  Railway_Object_Type = iota
)

//History record
type Train_History struct {
	object_id   int64
	object_type Railway_Object_Type
	arrival     time.Time
	departure   time.Time
}
type Train_Type int64

const (
	Train_Type_Normal  Train_Type = iota
	Train_Type_Service Train_Type = iota
)

//Keys for additional fields.
const (
	T_capacity       string = "cap"
	T_service_track  string = "strack"
	T_going_back     string = "g back" //0 false 1 - true
	T_uniqueStations string = "ustat"
)

//Train record with all necessary data
type Train struct {
	id        int64
	max_speed int64

	t_type Train_Type

	track_it int

	on_track      int64
	on_steer      int64
	current_speed int64

	out_of_order bool
	reliability  float64

	data map[string]int64

	passengers  *map[*Worker]bool
	stationlist *[]int64

	//t_task
	//chan_ready chan Train_Message
	tracklist *[]int64
	history   []Train_History

	trainReadyToDepartFromSteering chan int64
	trainReadyToDepartFromPlatform chan int64
	trainArrivedToTheEndOfTrack    chan int64

	leaveTrain chan int64
	enterTrain chan int64
	//notifications for service train
	trackOutOfOrder    chan int64
	trainOutOfOrder    chan int64
	steeringOutOfOrder chan int64

	repair chan int64
}

func tracklistToString(train_ptr *Train) string {
	var track_list string = ""
	if train_ptr.tracklist != nil {
		for it := 0; it < len(*train_ptr.tracklist); it++ {
			if track_list != "" {
				track_list += "," + strconv.FormatInt((*train_ptr.tracklist)[it], 10)
			} else {
				track_list = strconv.FormatInt((*train_ptr.tracklist)[it], 10)
			}
		}
	}
	return track_list
}

func initTrain(id int64, max_speed int64, tracklist *[]int64) *Train {
	t := new(Train)
	t.id = id
	t.max_speed = max_speed

	t.out_of_order = false
	t.reliability = 0.99995

	t.track_it = 1
	t.on_track = 0
	t.on_steer = 0
	t.current_speed = 0
	t.tracklist = tracklist
	t.history = make([]Train_History, 0)

	t.passengers = nil
	t.stationlist = nil

	t.data = make(map[string]int64)

	t.trainReadyToDepartFromSteering = make(chan int64)
	t.trainReadyToDepartFromPlatform = make(chan int64)
	t.trainArrivedToTheEndOfTrack = make(chan int64)
	t.trackOutOfOrder = make(chan int64)
	t.trainOutOfOrder = make(chan int64)
	t.steeringOutOfOrder = make(chan int64)

	t.leaveTrain = make(chan int64)
	t.enterTrain = make(chan int64)

	t.repair = make(chan int64)

	return t
}

//Creates new train
func newTrain(id int64, max_speed int64, capacity int64, tracklist []int64) *Train {
	t := initTrain(id, max_speed, &tracklist)
	t.t_type = Train_Type_Normal
	t.data[T_capacity] = capacity
	pass := make(map[*Worker]bool)
	t.passengers = &pass

	sl := make([]int64, 0)
	t.stationlist = &sl

	return t
}

//Creates new service train
func newServiceTrain(id int64, max_speed int64, service_track int64) *Train {
	t := initTrain(id, max_speed, nil)
	t.t_type = Train_Type_Service
	t.data[T_service_track] = service_track
	t.data[T_going_back] = 1
	return t
}

// Train task. Cycles thourgh its tracklist and moves from track to steering to track and so on.
func TrainTask(train_ptr *Train, model_ptr *Simulation_Model) {
	var track_ptr *Track
	var steer_ptr *Steering

	var work_ptr *Worker
	var stat_ptr *Station

	var on_station bool = false

	var ready bool = true
	var hist *Train_History = new(Train_History) //_Record
	var hist_prev *Train_History = nil
	var type_str string

	var help_train_ptr *Train = nil
	var help_track_ptr *Track = nil
	var help_steer_ptr *Steering = nil
	var first bool = true

	var r *rand.Rand = rand.New(rand.NewSource(int64(time.Now().Nanosecond() + int(4*train_ptr.id) + 65536)))

	if train_ptr != nil && model_ptr != nil {

		if train_ptr.t_type == Train_Type_Normal {
			type_str = "Train"
			PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
				"] prepares to start its schedule", model_ptr)
			// initialisation

			train_ptr.track_it = 1
			train_ptr.on_track = 0
			train_ptr.on_steer = 0
			train_ptr.current_speed = 0
		} else if train_ptr.t_type == Train_Type_Service {
			type_str = "Service Train"
			PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
				"] begins waiting for service call", model_ptr)
		} else {
			type_str = "Unknown"
		}

		var help_service_train_ptr *Train = nil
		var help bool = false
		for model_ptr.work {
			PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
				"] starts new loop", model_ptr)

			if train_ptr.t_type == Train_Type_Normal && train_ptr.out_of_order == true && help == false {
				PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
					"] calls for help", model_ptr)
				help_service_train_ptr = GetServiceTrain(model_ptr)
				if help_service_train_ptr != nil {

					select {
					case help_service_train_ptr.trainOutOfOrder <- train_ptr.id:
						help = true
					default:
						//  delay Standard.Duration(1);
					}

				} else {
					fmt.Println("#2# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
						"] received null pointer for service train")
				}
				PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
					"] called for help", model_ptr)
			}

			if train_ptr.out_of_order == false && ready == true { // train is ready to depart from either track or steering
				PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
					"] tries to depart from track or steering", model_ptr)
				ready = false
				hist_prev = hist
				hist = new(Train_History)                                  //_Record
				if train_ptr.tracklist != nil && train_ptr.on_track != 0 { //train is currently on track
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] tries to process on track routine", model_ptr)
					//retrieve next steering
					steer_ptr = GetSteering(track_ptr.st_end, model_ptr)
					if (train_ptr.t_type == Train_Type_Service && train_ptr.data[T_going_back] == 0) && train_ptr.track_it >= len(*train_ptr.tracklist) {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] arrived near its target. Proceeding to repair routine.", model_ptr)

						real_delay_dur := GetTimeSimToRealFromModel(1.0, Time_Interval_Hour, model_ptr)

						if help_track_ptr != nil {
							PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] will be repairing track ["+strconv.FormatInt(help_track_ptr.id, 10)+
								"] for next 1 hour ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)
						} else if help_train_ptr != nil {
							PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] will be repairing train ["+strconv.FormatInt(help_train_ptr.id, 10)+
								"] for next 1 hour ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)
						} else {
							PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] will be repairing steering ["+strconv.FormatInt(help_steer_ptr.id, 10)+
								"] for next 1 hour ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)
						}

						time.Sleep(time.Second * time.Duration(real_delay_dur))

						if help_track_ptr != nil {
							select {
							case help_track_ptr.repair <- train_ptr.id:
							case <-time.After(time.Second * time.Duration(2.0*real_delay_dur)):
								PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] could not reach out to train ["+strconv.FormatInt(help_train_ptr.id, 10)+
									"] with repair offer. Will manualy force repair.", model_ptr)
								help_track_ptr.out_of_order = false
							}

						} else if help_train_ptr != nil {
							select {
							case help_train_ptr.repair <- train_ptr.id:

							case <-time.After(time.Second * time.Duration(2.0*real_delay_dur)):
								//delay Standard.Duration(2.0*real_delay_dur);
								PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] could not reach out to track ["+strconv.FormatInt(help_track_ptr.id, 10)+
									"] with repair offer. Will manualy force repair.", model_ptr)
								help_train_ptr.out_of_order = false
							}
						} else {
							select {
							case help_steer_ptr.repair <- train_ptr.id:
							case <-time.After(time.Second * time.Duration(2.0*real_delay_dur)):
								PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] could not reach out to steering ["+strconv.FormatInt(help_steer_ptr.id, 10)+
									"] with repair offer. Will manualy force repair.", model_ptr)
								help_steer_ptr.out_of_order = false
							}
						}

						help_track_ptr = nil
						help_train_ptr = nil
						help_steer_ptr = nil
						train_ptr.tracklist = findTracklistTo(train_ptr.id, false, train_ptr.on_track, train_ptr.data[T_service_track], Type_Track, model_ptr)
						train_ptr.track_it = 1
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] finished repairing target. Going back to service track using tracklist: "+tracklistToString(train_ptr), model_ptr)

						train_ptr.data[T_going_back] = 1
					} else if (train_ptr.t_type == Train_Type_Service && train_ptr.data[T_going_back] != 0) && train_ptr.track_it == len(*train_ptr.tracklist) {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] arrived back to it's home service track ", model_ptr)
						train_ptr.tracklist = nil
						train_ptr.track_it = 1
					} else {
						PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] starts to depart from track", model_ptr)

						if train_ptr.t_type == Train_Type_Normal && track_ptr.t_type == Track_Type_Platform {
							on_station = false

							stat_ptr = GetStation(track_ptr.data[T_station_id], model_ptr)
							if stat_ptr != nil {
								PutLine("#3# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] notifies station["+strconv.FormatInt(stat_ptr.id, 10)+
									"] that its about to depart from platform and cannot accept new passengers.", model_ptr)
								PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] starts station["+strconv.FormatInt(stat_ptr.id, 10)+
									"].notifyAboutTrainDeparture("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)

								stat_ptr.notifyAboutTrainDeparture <- train_ptr.id

								PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] finished station["+strconv.FormatInt(stat_ptr.id, 10)+
									"].notifyAboutTrainDeparture("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
								PutLine("#3# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] departs from station["+strconv.FormatInt(stat_ptr.id, 10)+"]", model_ptr)
							} else {
								fmt.Println("#3# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
									"] received null pointer for station[" + strconv.FormatInt(track_ptr.data[T_station_id], 10) + "].")
							}

						}
					}

					steer_ptr = GetSteering(track_ptr.st_end, model_ptr)
					if steer_ptr != nil {
						//if steering is not null then waits for it to accept this train
						train_ptr.current_speed = 0

						PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] waits for steering ["+strconv.FormatInt(track_ptr.st_end, 10)+"]", model_ptr)

						if train_ptr.t_type == Train_Type_Service && train_ptr.data[T_going_back] == 0 {
							if steer_ptr.out_of_order == true {
								real_delay_dur := GetTimeSimToRealFromModel(1.0, Time_Interval_Hour, model_ptr)

								PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] will be repairing encountered out of order steering ["+strconv.FormatInt(help_steer_ptr.id, 10)+
									"] for next 1 hour ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

								time.Sleep(time.Second * time.Duration(real_delay_dur))

								select {
								case steer_ptr.repair <- train_ptr.id:
								case <-time.After(time.Second * time.Duration(2.0*real_delay_dur)):
									PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
										"] could not reach out to steering ["+strconv.FormatInt(steer_ptr.id, 10)+
										"] with repair offer. Will manualy force repair.", model_ptr)
									steer_ptr.out_of_order = false
								}

							}

							//waiting for steering to accept
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].acceptServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							steer_ptr.acceptServiceTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].acceptServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						} else {
							//waiting for steering to accept
							//steer_ptr.chan_accept <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].acceptTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							steer_ptr.acceptTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].acceptTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						}

						hist.arrival = time.Now()
						hist.object_type = Type_Steering

						train_ptr.on_steer = steer_ptr.id
						hist.object_id = steer_ptr.id

						//&& clears out currently blocked track

						PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] leaves the track ["+strconv.FormatInt(track_ptr.id, 10)+"]", model_ptr)

						//clearing the track
						if train_ptr.t_type == Train_Type_Service && (train_ptr.data[T_going_back] == 0 && train_ptr.track_it > 1) {
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts track["+strconv.FormatInt(track_ptr.id, 10)+
								"].freeFromServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							track_ptr.freeFromServiceTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished track["+strconv.FormatInt(track_ptr.id, 10)+
								"].freeFromServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						} else {
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts track["+strconv.FormatInt(track_ptr.id, 10)+
								"].clearAfterTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							track_ptr.clearAfterTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished track["+strconv.FormatInt(track_ptr.id, 10)+
								"]. clearAfterTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						}

						//track_ptr.chan_clear <- train_ptr.id

						//track_ptr.t_task.clearAfterTrain(train_ptr.id);
						hist_prev.departure = time.Now()
						train_ptr.history = append(train_ptr.history, *hist_prev)

						train_ptr.on_track = 0
						track_ptr = nil
					} else {
						fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
							"] received null pointer for steering ID[" + strconv.FormatInt(track_ptr.st_end, 10) + "].")
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] finished processing on track routine", model_ptr)

				} else if train_ptr.tracklist != nil && train_ptr.on_steer != 0 { //train is currently on steering
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] tries to process on steering routine", model_ptr)
					//incrementing the current track iterator and retrieving next track
					train_ptr.track_it = 1 + (train_ptr.track_it % len(*train_ptr.tracklist))
					track_ptr = GetTrack((*train_ptr.tracklist)[train_ptr.track_it-1], model_ptr)

					if track_ptr != nil {
						//if track is not null { waits for it to accept this train
						train_ptr.current_speed = 0

						PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] waits for track["+strconv.FormatInt(track_ptr.id, 10)+"]", model_ptr)

						if train_ptr.t_type == Train_Type_Service && train_ptr.data[T_going_back] == 0 {
							if track_ptr.out_of_order == true {
								real_delay_dur := GetTimeSimToRealFromModel(1.0, Time_Interval_Hour, model_ptr)

								PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
									"] will be repairing encountered out of order steering ["+strconv.FormatInt(help_steer_ptr.id, 10)+
									"] for next 1 hour ("+strconv.FormatFloat(real_delay_dur, 'f', 3, 64)+"s)", model_ptr)

								time.Sleep(time.Second * time.Duration(real_delay_dur))

								select {
								case track_ptr.repair <- train_ptr.id:
								case <-time.After(time.Second * time.Duration(2.0*real_delay_dur)):
									PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
										"] could not reach out to track ["+strconv.FormatInt(track_ptr.id, 10)+
										"] with repair offer. Will manualy force repair.", model_ptr)
									track_ptr.out_of_order = false
								}

							}

							//waiting for steering to accept
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts track["+strconv.FormatInt(track_ptr.id, 10)+
								"].acceptServiceTraine("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							track_ptr.acceptServiceTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished track["+strconv.FormatInt(track_ptr.id, 10)+
								"].acceptServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						} else {
							//waiting for steering to accept
							//steer_ptr.chan_accept <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts track["+strconv.FormatInt(track_ptr.id, 10)+
								"].acceptTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							track_ptr.acceptTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished track["+strconv.FormatInt(track_ptr.id, 10)+
								"].acceptTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						}

						//waiting for track to accept
						//track_ptr.chan_accept <- train_ptr.id
						//PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						//	"] no longer waits for track["+strconv.FormatInt(track_ptr.id, 10)+"]", model_ptr)

						hist.arrival = time.Now()
						if track_ptr.t_type == Track_Type_Track {
							hist.object_type = Type_Track
						} else if track_ptr.t_type == Track_Type_Platform {
							hist.object_type = Type_Platform

							on_station = true

							if train_ptr.t_type == Train_Type_Normal {
								stat_ptr = GetStation(track_ptr.data[T_station_id], model_ptr)
								if stat_ptr != nil {
									PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
										"] starts station["+strconv.FormatInt(stat_ptr.id, 10)+
										"].notifyAboutTrainArrival("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
									stat_ptr.notifyAboutTrainArrival <- train_ptr.id
									PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
										"] finished station["+strconv.FormatInt(stat_ptr.id, 10)+
										"].notifyAboutTrainArrival("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
									PutLine("#3# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
										"] arrived at station["+strconv.FormatInt(stat_ptr.id, 10)+"]", model_ptr)

									//if not HashSet.Is_Empty(Container => train_ptr.passengers) {
									for work_ptr = range *train_ptr.passengers {
										PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
											"] starts worker["+strconv.FormatInt(work_ptr.id, 10)+
											"].trainStop("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
										select {
										case work_ptr.trainStop <- TSPMessage(train_ptr.id, track_ptr.data[T_station_id]):

										default:
										}
										PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
											"] finished worker["+strconv.FormatInt(work_ptr.id, 10)+
											"].trainStop("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
									}
									//}
								} else {
									fmt.Println("#3# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
										"] received null pointer for station[" + strconv.FormatInt(track_ptr.data[T_station_id], 10) + "].")
								}
							}

						} else if track_ptr.t_type == Track_Type_Service {
							hist.object_type = Type_Service
						} else {
							hist.object_type = Type_Unknown
						}

						train_ptr.on_track = track_ptr.id
						hist.object_id = track_ptr.id
						//&& clears out currently blocked steering

						PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] leaves the steering ["+strconv.FormatInt(steer_ptr.id, 10)+"]", model_ptr)

						//clearing the steering
						if train_ptr.t_type == Train_Type_Service && train_ptr.data[T_going_back] == 0 {
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].freeFromServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							steer_ptr.freeFromServiceTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].freeFromServiceTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						} else {
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].clearAfterTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							steer_ptr.clearAfterTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished steering["+strconv.FormatInt(steer_ptr.id, 10)+
								"].clearAfterTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
						}

						//	steer_ptr.chan_clear <- train_ptr.id

						hist_prev.departure = time.Now()
						train_ptr.history = append(train_ptr.history, *hist_prev)

						train_ptr.on_steer = 0
						steer_ptr = nil
					} else {
						fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
							"] received null pointer for track ID[" +
							strconv.FormatInt((*train_ptr.tracklist)[train_ptr.track_it-1], 10) + "].")
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] finished processing on steering routine", model_ptr)
				} else {

					if first == true {
						PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] first time entry", model_ptr)
						first = false
						if train_ptr.t_type == Train_Type_Normal {
							track_ptr = GetTrack((*train_ptr.tracklist)[train_ptr.track_it], model_ptr)
						} else {
							track_ptr = GetTrack(train_ptr.data[T_service_track], model_ptr)
						}

						if track_ptr != nil {
							hist.arrival = time.Now()
							train_ptr.on_track = track_ptr.id
							hist.object_id = track_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] starts track["+strconv.FormatInt(track_ptr.id, 10)+
								"].acceptTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
							track_ptr.acceptTrain <- train_ptr.id
							PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] finished track["+strconv.FormatInt(track_ptr.id, 10)+
								"].acceptTrain("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)

							if track_ptr.t_type == Track_Type_Track {
								hist.object_type = Type_Track
							} else if track_ptr.t_type == Track_Type_Platform {
								hist.object_type = Type_Platform

								on_station = true

								if train_ptr.t_type == Train_Type_Normal {
									stat_ptr = GetStation(track_ptr.data[T_station_id], model_ptr)
									if stat_ptr != nil {
										PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
											"] starts station["+strconv.FormatInt(stat_ptr.id, 10)+
											"].notifyAboutTrainArrival("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
										stat_ptr.notifyAboutTrainArrival <- train_ptr.id
										PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
											"] finished station["+strconv.FormatInt(stat_ptr.id, 10)+
											"].notifyAboutTrainArrival("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
										PutLine("#3# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
											"] arrived at station["+strconv.FormatInt(stat_ptr.id, 10)+"]", model_ptr)

										//if not HashSet.Is_Empty(Container => train_ptr.passengers) {
										for work_ptr = range *train_ptr.passengers {
											PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
												"] starts worker["+strconv.FormatInt(work_ptr.id, 10)+
												"].trainStop("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
											select {
											case work_ptr.trainStop <- TSPMessage(train_ptr.id, track_ptr.data[T_station_id]):
											default:
											}
											PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
												"] finished worker["+strconv.FormatInt(work_ptr.id, 10)+
												"].trainStop("+strconv.FormatInt(train_ptr.id, 10)+")", model_ptr)
										}
										//}
									} else {
										fmt.Println("#3# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
											"] received null pointer for station[" + strconv.FormatInt(track_ptr.data[T_station_id], 10) + "].")
									}
								}

							} else {
								hist.object_type = Type_Unknown
							}
						} else {
							if train_ptr.t_type == Train_Type_Service {
								fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
									"] received null pointer for track ID[" + strconv.FormatInt(train_ptr.data[T_service_track], 10) + "].")
							} else {
								fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
									"] received null pointer for track ID[" + strconv.FormatInt((*train_ptr.tracklist)[train_ptr.track_it], 10) + "].")
							}
						}
						PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
							"] ends first time entry", model_ptr)
					} else {
						fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
							"] is not on neither track or steering at the moment.")
					}
				}
			} else {
				PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
					"] enters select", model_ptr)
				select {
				// when train_ptr.out_of_order = true =>
				case train_id := <-train_ptr.repair:
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters repair("+strconv.FormatInt(train_id, 10)+")", model_ptr)
					if help_service_train_ptr != nil && help_service_train_ptr.id == train_id {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] was just repaired. Returning to schedule.", model_ptr)
						train_ptr.out_of_order = false
					} else {
						if help_service_train_ptr != nil {
							PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] has no information about service train but received repair signal from service train["+strconv.FormatInt(train_id, 10)+"]. Accepting the repair and moving along with schedule.", model_ptr)
							train_ptr.out_of_order = false
						} else {
							PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] received repair signal from illegal service train ["+strconv.FormatInt(train_id, 10)+"]. Accepting the repair and moving along with schedule.", model_ptr)
							train_ptr.out_of_order = false
						}
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves repair("+strconv.FormatInt(train_id, 10)+")", model_ptr)

					//notification from steering that train can move further
				case steer_id := <-when(train_ptr.tracklist != nil && train_ptr.out_of_order == false,
					train_ptr.trainReadyToDepartFromSteering):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters trainReadyToDepartFromSteering("+strconv.FormatInt(steer_id, 10)+")", model_ptr)
					train_ptr.current_speed = 0
					PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] finished waiting on steering ["+strconv.FormatInt(steer_id, 10)+"] and is ready to move onto next track", model_ptr)
					ready = true
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves trainReadyToDepartFromSteering("+strconv.FormatInt(steer_id, 10)+")", model_ptr)
					//notification from platform that train can move further
				case track_id := <-when(train_ptr.tracklist != nil && train_ptr.out_of_order == false,
					train_ptr.trainReadyToDepartFromPlatform):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters trainReadyToDepartFromPlatform("+strconv.FormatInt(track_id, 10)+")", model_ptr)
					train_ptr.current_speed = 0
					PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] finished waiting on platform ["+strconv.FormatInt(track_id, 10)+"] and is ready to move onto steering", model_ptr)
					ready = true
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves trainReadyToDepartFromPlatform("+strconv.FormatInt(track_id, 10)+")", model_ptr)
					//notification from track that train can move further
				case track_id := <-when(train_ptr.tracklist != nil && train_ptr.out_of_order == false,
					train_ptr.trainArrivedToTheEndOfTrack):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters trainArrivedToTheEndOfTrack("+strconv.FormatInt(track_id, 10)+")", model_ptr)
					train_ptr.current_speed = 0
					PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] finished riding on track ["+strconv.FormatInt(track_id, 10)+"] and is ready to move onto steering", model_ptr)
					ready = true
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves trainArrivedToTheEndOfTrack("+strconv.FormatInt(track_id, 10)+")", model_ptr)

					//notification for service train from track for help
				case track_id := <-when(train_ptr.t_type == Train_Type_Service && (train_ptr.data[T_going_back] != 0 && train_ptr.on_track != 0),
					train_ptr.trackOutOfOrder):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters trackOutOfOrder("+strconv.FormatInt(track_id, 10)+")", model_ptr)
					train_ptr.data[T_going_back] = 0
					//PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] printing model:", model_ptr)
					//PrintModel(model_ptr, model_ptr.mode)
					train_ptr.tracklist = findTracklistTo(train_ptr.id, true, train_ptr.on_track, track_id, Type_Track, model_ptr)
					if train_ptr.tracklist == nil {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] could not find path to the train", model_ptr)
						train_ptr.data[T_going_back] = 1
					} else {
						train_ptr.track_it = 1

						help_track_ptr = GetTrack(track_id, model_ptr)

						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] received help request from track ["+strconv.FormatInt(track_id, 10)+"]. Using tracklist: "+tracklistToString(train_ptr), model_ptr)
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves trackOutOfOrder("+strconv.FormatInt(track_id, 10)+")", model_ptr)

					//notification for service train from platform for help
				case train_id := <-when(train_ptr.t_type == Train_Type_Service && (train_ptr.data[T_going_back] != 0 && train_ptr.on_track != 0),
					train_ptr.trainOutOfOrder):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters trainOutOfOrder("+strconv.FormatInt(train_id, 10)+")", model_ptr)
					train_ptr.data[T_going_back] = 0
					//PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] printing model:", model_ptr)
					//PrintModel(model_ptr, model_ptr.mode)
					train_ptr.tracklist = findTracklistTo(train_ptr.id, true, train_ptr.on_track, train_id, Type_Train, model_ptr)
					if train_ptr.tracklist == nil {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] could not find path to the track", model_ptr)
						train_ptr.data[T_going_back] = 1
					} else {
						train_ptr.track_it = 1

						help_train_ptr = GetTrain(train_id, model_ptr)

						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] received help request from train ["+strconv.FormatInt(train_id, 10)+"]. Using tracklist: "+tracklistToString(train_ptr), model_ptr)
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves trainOutOfOrder("+strconv.FormatInt(train_id, 10)+")", model_ptr)

					//notification for service train from train for help
				case steer_id := <-when(train_ptr.t_type == Train_Type_Service && (train_ptr.data[T_going_back] != 0 && train_ptr.on_track != 0),
					train_ptr.steeringOutOfOrder):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters steeringOutOfOrder("+strconv.FormatInt(steer_id, 10)+")", model_ptr)
					train_ptr.data[T_going_back] = 0
					//PutLine(type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] printing model:", model_ptr)
					//PrintModel(model_ptr, model_ptr.mode)
					train_ptr.tracklist = findTracklistTo(train_ptr.id, true, train_ptr.on_track, steer_id, Type_Steering, model_ptr)
					if train_ptr.tracklist == nil {
						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] could not find path to the steering", model_ptr)
						train_ptr.data[T_going_back] = 1
					} else {
						train_ptr.track_it = 1

						help_steer_ptr = GetSteering(steer_id, model_ptr)

						PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] received help request from steering ["+strconv.FormatInt(steer_id, 10)+"]. Using tracklist: "+tracklistToString(train_ptr), model_ptr)
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves steeringOutOfOrder("+strconv.FormatInt(steer_id, 10)+")", model_ptr)

				case worker_id := <-when(train_ptr.t_type == Train_Type_Normal && (train_ptr.on_track != 0 && on_station),
					train_ptr.leaveTrain):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters leaveTrain("+strconv.FormatInt(worker_id, 10)+")", model_ptr)

					work_ptr = GetWorker(worker_id, model_ptr)
					if work_ptr != nil {
						if (*train_ptr.passengers)[work_ptr] {

							delete(*train_ptr.passengers, work_ptr)
							PutLine("#3# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] bids farewell to passenger["+strconv.FormatInt(worker_id, 10)+"]", model_ptr)

						} else {
							fmt.Println("#3# " + type_str + "]" + strconv.FormatInt(train_ptr.id, 10) +
								"] received illegal leave notification from worker[" + strconv.FormatInt(worker_id, 10) + "]")
						}

					} else {
						fmt.Println("#3# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
							"] received null pointer for worker[" + strconv.FormatInt(worker_id, 10) + "]")
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves leaveTrain("+strconv.FormatInt(worker_id, 10)+")", model_ptr)

				case worker_id := <-when(train_ptr.t_type == Train_Type_Normal && (train_ptr.on_track != 0 && on_station),
					train_ptr.enterTrain):

					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] enters enterTrain("+strconv.FormatInt(worker_id, 10)+")", model_ptr)

					work_ptr = GetWorker(worker_id, model_ptr)
					if work_ptr != nil {
						if !(*train_ptr.passengers)[work_ptr] {
							(*train_ptr.passengers)[work_ptr] = true
							PutLine("#3# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
								"] welcomes passenger["+strconv.FormatInt(worker_id, 10)+"]", model_ptr)

						} else {
							fmt.Println("#3# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
								"] received illegal leave notification from worker[" + strconv.FormatInt(worker_id, 10) + "]")
						}

					} else {
						fmt.Println("#3# " + type_str + "[" + strconv.FormatInt(train_ptr.id, 10) +
							"] received null pointer for worker[" + strconv.FormatInt(worker_id, 10) + "]")
					}
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] leaves enterTrain("+strconv.FormatInt(worker_id, 10)+")", model_ptr)

				case <-time.After(time.Second * time.Duration(GetTimeSimToRealFromModel(1.0, Time_Interval_Hour, model_ptr))):
					PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
						"] had timeout in select", model_ptr)

				}
				PutLine("#debug# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+
					"] leaves select", model_ptr)

			}
			if train_ptr.t_type != Train_Type_Service && train_ptr.out_of_order == false {
				ran := r.Float64()
				//fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) + "] rolled " + strconv.FormatFloat(ran, 'f', 3, 64) + " at time " + TimeToString(GetRelativeTime(time.Now(), model_ptr)))

				if train_ptr.reliability < ran {
					PutLine("#2# "+type_str+"["+strconv.FormatInt(train_ptr.id, 10)+"] broke at time "+TimeToString(GetRelativeTime(time.Now(), model_ptr)), model_ptr)
					train_ptr.out_of_order = true
					help = false
				}
			}
		}
		fmt.Println(type_str + "[" + strconv.FormatInt(train_ptr.id, 10) + "] terminates its execution")
	} else {
		fmt.Println("TrainTask received null pointer! Task will terminate")
	}
}

type WorkerState int64

const (
	AtHome           WorkerState = iota
	TravellingToWork WorkerState = iota
	WaitingForWork   WorkerState = iota
	TravellingToHome WorkerState = iota
	Working          WorkerState = iota
)

type Worker struct {
	id                      int64
	home_stat_id            int64
	on_train                int64
	on_Station              int64
	dest_Station            int64
	connectionlist          *[]*Connection
	connectionlist_iterator int64
	state                   WorkerState

	acceptTask              chan int64
	startTask               chan StartTaskStruct
	trainStop               chan TrainStatPair
	notifyAboutTrainArrival chan TrainStatPair
}

func TSPMessage(train_id int64, stat_id int64) TrainStatPair {
	t := new(TrainStatPair)
	t.stat_id = stat_id
	t.train_id = train_id
	return *t
}
func STSMessage(stat_id int64, work_time float64) StartTaskStruct {
	t := new(StartTaskStruct)
	t.stat_id = stat_id
	t.work_time_hours = work_time
	return *t
}

type TrainStatPair struct {
	train_id int64
	stat_id  int64
}

type StartTaskStruct struct {
	stat_id         int64
	work_time_hours float64
}

func newWorker(id int64, home_stat int64) *Worker {
	t := new(Worker)
	t.id = id
	t.home_stat_id = home_stat
	t.on_train = 0
	t.on_Station = home_stat
	t.dest_Station = 0
	t.connectionlist = nil
	t.connectionlist_iterator = -1
	t.state = AtHome

	t.acceptTask = make(chan int64)
	t.startTask = make(chan StartTaskStruct)
	t.trainStop = make(chan TrainStatPair)
	t.notifyAboutTrainArrival = make(chan TrainStatPair)

	return t
}

func WorkerTask(work_ptr *Worker, model_ptr *Simulation_Model) {
	var work_duration float64
	var work bool = false
	var train_ptr *Train
	var stat_ptr *Station

	var check_train int64 = 0
	var check_station int64 = 0

	if model_ptr != nil && work_ptr != nil {
		for model_ptr.work {
			//fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] enters select. is at home?" + strconv.FormatBool(work_ptr.state == AtHome))
			select {
			case stat_id := <-when(work_ptr.state == AtHome,
				work_ptr.acceptTask):
				work_ptr.state = TravellingToWork
				work_ptr.dest_Station = stat_id

			case ST := <-whenSTS(work_ptr.state == WaitingForWork,
				work_ptr.startTask):
				if ST.stat_id == work_ptr.dest_Station {
					work_duration = ST.work_time_hours
					work = true
				} else {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received illegal start task notification from station[" + strconv.FormatInt(ST.stat_id, 10) + "]")
				}

			case ST := <-whenTSP(work_ptr.on_train != 0,
				work_ptr.trainStop):
				if ST.train_id == work_ptr.on_train {
					check_station = ST.stat_id
				} else {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received illegal train stop notification from train[" + strconv.FormatInt(ST.train_id, 10) + "]")
				}

			case ST := <-whenTSP(work_ptr.on_Station != 0 && (work_ptr.state == TravellingToWork || work_ptr.state == TravellingToHome),
				work_ptr.notifyAboutTrainArrival):
				if ST.stat_id == work_ptr.on_Station {
					check_train = ST.train_id
				} else {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received illegal train arrival notification from station[" + strconv.FormatInt(ST.stat_id, 10) + "]")
				}

			case <-time.After(time.Second * time.Duration(GetTimeSimToRealFromModel(100.0, Time_Interval_Minute, model_ptr))):
			}
			//fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] leaves select")

			if check_station != 0 && (*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].arrive_station_id == check_station {
				train_ptr = GetTrain(work_ptr.on_train, model_ptr)
				stat_ptr = GetStation(check_station, model_ptr)
				if train_ptr != nil && stat_ptr != nil {
					if (*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].arrive_station_id == check_station {

						PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] tries to enter station["+strconv.FormatInt(check_station, 10)+"]", model_ptr)
						select {
						case stat_ptr.notifyAboutWorkerArrival <- work_ptr.id:
							train_ptr.leaveTrain <- work_ptr.id
							work_ptr.on_train = 0
							work_ptr.on_Station = stat_ptr.id
							work_ptr.connectionlist_iterator = work_ptr.connectionlist_iterator + 1
							PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] left the train["+strconv.FormatInt(train_ptr.id, 10)+"] and entered station["+strconv.FormatInt(stat_ptr.id, 10)+"]", model_ptr)
						case <-time.After(time.Second * time.Duration(10.0)):
							PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] failed to leave train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)
						}

					}
				} else {
					if train_ptr == nil {
						fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for train[" + strconv.FormatInt(work_ptr.on_train, 10) + "]")
					} else {
						fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for station[" + strconv.FormatInt(check_station, 10) + "]")
					}
				}
				check_station = 0
			}

			if check_train != 0 && work_ptr.connectionlist != nil && ((*work_ptr.connectionlist)[work_ptr.connectionlist_iterator]) != nil {
				stat_ptr = GetStation(work_ptr.on_Station, model_ptr)
				if stat_ptr != nil {
					if (*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].train_id == check_train {
						train_ptr = GetTrain(check_train, model_ptr)
						if train_ptr != nil {
							if (*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].train_id == check_train {
								PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] tries to enter train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)
								select {
								case train_ptr.enterTrain <- work_ptr.id:
									stat_ptr.notifyAboutWorkerDeparture <- work_ptr.id
									PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] aboards train["+strconv.FormatInt(train_ptr.id, 10)+"] and leaves station["+strconv.FormatInt(stat_ptr.id, 10)+"]", model_ptr)
									work_ptr.on_train = train_ptr.id
									work_ptr.on_Station = 0
								case <-time.After(time.Second * time.Duration(10.0)):
									fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] failed to hop on train[" + strconv.FormatInt(train_ptr.id, 10) + "]")
								}
							} else {
								fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for train[" + strconv.FormatInt(check_train, 10) + "]")
							}
						} else {
							fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for station[" + strconv.FormatInt(work_ptr.on_Station, 10) + "]")
						}
					}
				}
				check_train = 0
			}

			if work_ptr.state == TravellingToWork && work_ptr.connectionlist == nil {
				work_ptr.connectionlist = GetConnection(work_ptr.home_stat_id, work_ptr.dest_Station, model_ptr)
				work_ptr.connectionlist_iterator = 0
				if work_ptr.connectionlist == nil {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] did not found connection from " + strconv.FormatInt(work_ptr.home_stat_id, 10) + " to " + strconv.FormatInt(work_ptr.dest_Station, 10) + " stations.")
					work_ptr.state = AtHome
					work_ptr.dest_Station = 0
				} else if len(*work_ptr.connectionlist) == 0 {
					work_ptr.connectionlist = nil
					work_ptr.connectionlist_iterator = -1
					work_ptr.state = WaitingForWork
					stat_ptr = GetStation(work_ptr.dest_Station, model_ptr)
					if stat_ptr != nil {
						PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] is already at target station and is ready to start working.", model_ptr)
						stat_ptr.notifyAboutReadinessToWork <- work_ptr.id
					} else {
						fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for station[" + strconv.FormatInt(work_ptr.dest_Station, 10) + "]")
					}
				} else {
					PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] has accepted request for work. Moving to station["+strconv.FormatInt(work_ptr.dest_Station, 10)+"]", model_ptr)
				}
			}

			if work {
				stat_ptr = GetStation(work_ptr.dest_Station, model_ptr)
				if stat_ptr != nil {
					work = false
					PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] starts working for next "+strconv.FormatFloat(work_duration, 'f', 3, 64)+" hours.", model_ptr)

					time.Sleep(time.Duration(GetTimeSimToRealFromModel(work_duration, Time_Interval_Hour, model_ptr)) * time.Second)
					PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] finished working on task.", model_ptr)
					stat_ptr.notifyAboutFinishingTheWork <- work_ptr.id

				} else {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for station[" + strconv.FormatInt(work_ptr.dest_Station, 10) + "]")
				}

				work_ptr.connectionlist = GetConnection(work_ptr.dest_Station, work_ptr.home_stat_id, model_ptr)
				work_ptr.connectionlist_iterator = 0
				work_ptr.dest_Station = work_ptr.home_stat_id
				work_ptr.state = TravellingToHome
				if work_ptr.connectionlist == nil {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] does not found connection from " + strconv.FormatInt(work_ptr.dest_Station, 10) + " to " + strconv.FormatInt(work_ptr.home_stat_id, 10) + " stations.")
				} else if len(*work_ptr.connectionlist) == 0 {
					work_ptr.connectionlist = nil
					work_ptr.connectionlist_iterator = -1
					work_ptr.state = AtHome
					PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] left work and went directly back home.", model_ptr)
				} else {
					PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] left work and is going back to home station["+strconv.FormatInt(work_ptr.home_stat_id, 10)+"]", model_ptr)
				}
			} else if work_ptr.state == TravellingToWork && work_ptr.on_Station == work_ptr.dest_Station {
				work_ptr.connectionlist = nil
				work_ptr.connectionlist_iterator = -1
				work_ptr.state = WaitingForWork
				stat_ptr = GetStation(work_ptr.dest_Station, model_ptr)
				if stat_ptr != nil {
					PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] arrived to target station and is ready to start working.", model_ptr)
					stat_ptr.notifyAboutReadinessToWork <- work_ptr.id
					// PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] notified that it's ready to work.",model_ptr);
				} else {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for station[" + strconv.FormatInt(work_ptr.dest_Station, 10) + "]")
				}
			} else if work_ptr.state == TravellingToHome && work_ptr.on_Station == work_ptr.home_stat_id {
				work_ptr.connectionlist = nil
				work_ptr.connectionlist_iterator = -1
				work_ptr.state = AtHome
				PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] arrived back at home.", model_ptr)
			} else if work_ptr.connectionlist != nil && work_ptr.on_Station != 0 {
				stat_ptr = GetStation(work_ptr.on_Station, model_ptr)
				if stat_ptr != nil {

					stat_ptr.mutex.RLock()
					for cur := range stat_ptr.trains {
						if cur == (*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].train_id {

							train_ptr = GetTrain((*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].train_id, model_ptr)
							if train_ptr != nil {

								PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] tries to enter train["+strconv.FormatInt(train_ptr.id, 10)+"]", model_ptr)
								select {
								case train_ptr.enterTrain <- work_ptr.id:
									stat_ptr.notifyAboutWorkerDeparture <- work_ptr.id
									PutLine("#3# worker["+strconv.FormatInt(work_ptr.id, 10)+"] aboards train["+strconv.FormatInt(train_ptr.id, 10)+"] and leaves station["+strconv.FormatInt(stat_ptr.id, 10)+"]", model_ptr)
									work_ptr.on_train = train_ptr.id
									work_ptr.on_Station = 0
								case <-time.After(time.Second * time.Duration(10.0)):
									fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] failed to hop on train[" + strconv.FormatInt(train_ptr.id, 10) + "]")
								}
							} else {
								fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for train[" + strconv.FormatInt((*work_ptr.connectionlist)[work_ptr.connectionlist_iterator].train_id, 10) + "]")
							}

						}
					}
					stat_ptr.mutex.RUnlock()
				} else {
					fmt.Println("#3# worker[" + strconv.FormatInt(work_ptr.id, 10) + "] received nil pointer for station[" + strconv.FormatInt(work_ptr.dest_Station, 10) + "]")
				}
			}
		}
	} else {
		fmt.Println("#3# WorkTask received nil pointer.")
	}
}

func getInput(filename string) []string {
	file, err := os.Open(filename)
	var rtrn []string = make([]string, 0)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 0 {
				if line[0] != '#' {
					rtrn = append(rtrn, line)
				}
			}
		}
	} else {
		fmt.Println("Could not open file " + filename)
	}
	return rtrn
}

/*
 * gets input from file and creates Simulation_Model object.
 */
func getModel(filename string) *Simulation_Model {
	//var inp String
	var str []string = getInput(filename)

	var line_state int64 = 0
	var model_ptr *Simulation_Model = new(Simulation_Model)
	model_ptr.log_mode = all_output
	model_ptr.mode = Mixed_Mode
	model_ptr.debug = false
	steer_regex, err1 := regexp.Compile("(\\d+)\\s(\\d+)")
	platf_regex, err3 := regexp.Compile("(\\d+)\\s(\\d+)\\s(\\d+)\\sstop\\s(\\d+)\\s(\\d+)")
	track_regex, err4 := regexp.Compile("(\\d+)\\s(\\d+)\\s(\\d+)\\spass\\s(\\d+)\\s(\\d+)")
	service_track_regex, err5 := regexp.Compile("(\\d+)\\s(\\d+)\\s(\\d+)\\sservice")
	station_regex, err11 := regexp.Compile("(\\d+)")
	worker_regex, err12 := regexp.Compile("(\\d+)\\s(\\d+)")

	train_regex, err2 := regexp.Compile("(\\d+)\\s(\\d+)\\snormal\\s(\\d+)\\s([\\d+,?]+)")
	service_train_regex, err6 := regexp.Compile("(\\d+)\\s(\\d+)\\sservice\\s(\\d+)")

	train_tracklist_regex, err7 := regexp.Compile("(\\d+)")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err11 != nil || err12 != nil {
		fmt.Println("Could not compile one of the input regexes")
		return nil
	}
	for i := 0; i < len(str); i++ {
		//fmt.Println("[", i, "]"+str[i])
		var curr_line string = str[i]

		if curr_line == "@simulation_speed:" {
			line_state = 1
		} else if curr_line == "@steering:" {
			line_state = 2
		} else if curr_line == "@tracks:" {
			line_state = 3
		} else if curr_line == "@trains:" {
			line_state = 4
		} else if curr_line == "@stations:" {
			line_state = 5
		} else if curr_line == "@workers:" {
			line_state = 6
		} else {
			switch line_state {
			case 1:
				speed, err := strconv.ParseInt(curr_line, 10, 64)
				if err != nil {
					fmt.Println("Could not parse simulation speed value from this line: [" + curr_line + "]")
					return nil
				} else {
					model_ptr.speed = speed
				}
			case 2:
				if steer_regex.MatchString(curr_line) {
					res := steer_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 3 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						delay, err2 := strconv.ParseInt(res[0][2], 10, 64)
						if err1 == nil && err2 == nil {
							//inserting steering pointer into array
							model_ptr.steer = append(model_ptr.steer, NewSteering(id, delay))
							//fmt.Println(SteeringToString(model_ptr.steer[len(model_ptr.steer)-1]))
						} else {
							fmt.Println("Could not parse data for steering record from line[", i, "]\""+curr_line+"\"")
						}

					} else {
						fmt.Println("Problem with steer result for line [", i, "]\""+curr_line+"\"")
					}
				} else {
					fmt.Println("line[", i, "]\""+curr_line+"\" does not matches the steering record format.")
				}
			case 3: // tracks
				if track_regex.MatchString(curr_line) {
					res := track_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 6 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						v1, err2 := strconv.ParseInt(res[0][2], 10, 64)
						v2, err3 := strconv.ParseInt(res[0][3], 10, 64)
						dist, err4 := strconv.ParseInt(res[0][4], 10, 64)
						max_speed, err5 := strconv.ParseInt(res[0][5], 10, 64)
						if err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
							//inserting track pointer into array
							model_ptr.track = append(model_ptr.track, NewTrack(id, v1, v2, dist, max_speed))
							//fmt.Println(TrackToString(model_ptr.track[len(model_ptr.track)-1]))
						} else {
							fmt.Println("Could not parse data for track record from line[", i, "]\""+curr_line+"\"")
						}
					} else {
						fmt.Println("Problem with track regex result for line [", i, "]\""+curr_line+"\"")
					}
				} else if platf_regex.MatchString(curr_line) {
					res := platf_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 6 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						v1, err2 := strconv.ParseInt(res[0][2], 10, 64)
						v2, err3 := strconv.ParseInt(res[0][3], 10, 64)
						delay, err4 := strconv.ParseInt(res[0][4], 10, 64)
						stat, err5 := strconv.ParseInt(res[0][5], 10, 64)
						if err1 == nil && err2 == nil && err3 == nil && err4 == nil && err5 == nil {
							//inserting track pointer into array
							model_ptr.track = append(model_ptr.track, NewPlatform(id, v1, v2, delay, stat))

							model_ptr.platf = append(model_ptr.platf, model_ptr.track[len(model_ptr.track)-1])

							//fmt.Println(TrackToString(model_ptr.track[len(model_ptr.track)-1]))
						} else {
							fmt.Println("Could not parse data for platform record from line[", i, "]\""+curr_line+"\"")
						}
					} else {
						fmt.Println("Problem with platf regex result for line [", i, "]\""+curr_line+"\"")
					}
				} else if service_track_regex.MatchString(curr_line) {
					res := service_track_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 4 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						v1, err2 := strconv.ParseInt(res[0][2], 10, 64)
						v2, err3 := strconv.ParseInt(res[0][3], 10, 64)
						if err1 == nil && err2 == nil && err3 == nil {
							//inserting track pointer into array
							model_ptr.track = append(model_ptr.track, NewServiceTrack(id, v1, v2))
							//fmt.Println("service track" + TrackToString(model_ptr.track[len(model_ptr.track)-1]))
						} else {
							fmt.Println("Could not parse data for service track record from line[", i, "]\""+curr_line+"\"")
						}
					} else {
						fmt.Println("Problem with service track result for line [", i, "]\""+curr_line+"\"")
					}
				} else {
					fmt.Println("line[", i, "]\""+curr_line+"\" does not matches the track or platform record format.")
				}
			case 4: //trains
				if train_regex.MatchString(curr_line) {
					res := train_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 5 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						speed, err2 := strconv.ParseInt(res[0][2], 10, 64)
						capacity, err3 := strconv.ParseInt(res[0][3], 10, 64)
						res_track := train_tracklist_regex.FindAllStringSubmatch(res[0][4], -1)
						var tracklist []int64 = make([]int64, 0)
						for i := 0; i < len(res_track); i++ {
							v, err_t := strconv.ParseInt(res_track[i][0], 10, 64)
							if err_t == nil {
								tracklist = append(tracklist, v)
							} else {
								fmt.Println("Could not parse one of the parameters \""+res[i][0]+"\" for tracklist from line[", i, "]\""+curr_line+"\"")
							}
						}
						if err1 == nil && err2 == nil && err3 == nil /*&& err4 == nil*/ {
							//inserting train pointer into array
							model_ptr.train = append(model_ptr.train, newTrain(id, speed, capacity, tracklist))
							//fmt.Println(TrainToString(model_ptr.train[len(model_ptr.train)-1]))
						} else {
							fmt.Println("Could not parse data for train record from line[", i, "]\""+curr_line+"\"")
						}
					} else {
						fmt.Println("Problem with train result for line [", i, "]\""+curr_line+"\"")
					}

				} else if service_train_regex.MatchString(curr_line) {
					res := service_train_regex.FindAllStringSubmatch(curr_line, -1)
					//fmt.Println("service train " + strconv.FormatInt(int64(len(res[0])), 10))
					if len(res) > 0 && len(res[0]) >= 4 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						speed, err2 := strconv.ParseInt(res[0][2], 10, 64)
						track, err3 := strconv.ParseInt(res[0][3], 10, 64)
						if err1 == nil && err2 == nil && err3 == nil {
							//inserting train pointer into array
							model_ptr.train = append(model_ptr.train, newServiceTrain(id, speed, track))
							//fmt.Println("service train" + TrainToString(model_ptr.train[len(model_ptr.train)-1]))
						} else {
							fmt.Println("Could not parse data for service train record from line[", i, "]\""+curr_line+"\"")
						}
					} else {
						fmt.Println("Problem with service train result for line [", i, "]\""+curr_line+"\"")
					}

				} else {
					fmt.Println("line[", i, "]\""+curr_line+"\" does not matches the train record format.")
				}
			case 5: //stations
				if station_regex.MatchString(curr_line) {
					res := station_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 2 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						if err1 == nil {
							//inserting station pointer into array
							model_ptr.station = append(model_ptr.station, newStation(id))
						} else {
							fmt.Println("Could not parse data for station record from line[", i, "]\""+curr_line+"\"")
						}

					} else {
						fmt.Println("Problem with steer result for line [", i, "]\""+curr_line+"\"")
					}
				} else {
					fmt.Println("line[", i, "]\""+curr_line+"\" does not matches the station record format.")
				}

			case 6: //workers
				if worker_regex.MatchString(curr_line) {
					res := worker_regex.FindAllStringSubmatch(curr_line, -1)
					if len(res) > 0 && len(res[0]) >= 3 {
						id, err1 := strconv.ParseInt(res[0][1], 10, 64)
						stat_id, err2 := strconv.ParseInt(res[0][2], 10, 64)
						if err1 == nil && err2 == nil {
							//inserting worker pointer into array
							model_ptr.worker = append(model_ptr.worker, newWorker(id, stat_id))
						} else {
							fmt.Println("Could not parse data for worker record from line[", i, "]\""+curr_line+"\"")
						}

					}
				} else {
					fmt.Println("line[", i, "]\""+curr_line+"\" does not matches the worker record format.")
				}
			}
		}
	}

	var unique bool = false
	for it := 0; it < len(model_ptr.train); it++ {
		train_ptr := model_ptr.train[it]
		if train_ptr.t_type == Train_Type_Normal {
			var uniq int64 = 0
			for itt := 0; itt < len(*train_ptr.stationlist); itt++ {

				unique = true
				for ittt := itt + 1; ittt < len(*train_ptr.stationlist); ittt++ {
					if (*train_ptr.stationlist)[itt] == (*train_ptr.stationlist)[ittt] {
						unique = false
						break
					}
				}
				if unique {
					uniq++
				}
			}
			train_ptr.data[T_uniqueStations] = uniq
		}
	}
	for it := 0; it < len(model_ptr.train); it++ {
		train_ptr := model_ptr.train[it]
		if train_ptr.t_type == Train_Type_Normal {
			for it := 0; it < len(*train_ptr.tracklist); it++ {
				track_id := (*train_ptr.tracklist)[it]
				for itt := 0; itt < len(model_ptr.platf); itt++ {
					if track_id == model_ptr.platf[itt].id {
						ap := append(*train_ptr.stationlist, model_ptr.platf[itt].data[T_station_id])
						train_ptr.stationlist = &ap
						break
					}
				}
			}
		}
	}

	for it := 0; it < len(model_ptr.station); it++ {
		stat_ptr := model_ptr.station[it]
		if stat_ptr == nil {
			fmt.Println("null statioin")
		}
		for itt := 0; itt < len(model_ptr.worker); itt++ {
			if model_ptr.worker[itt] == nil {
				fmt.Println("null worker")
			}
			if model_ptr.worker[itt].home_stat_id == stat_ptr.id {
				stat_ptr.passengers[model_ptr.worker[itt]] = true
			}
		}
	}

	return model_ptr

}

//Author: Piotr Olejarz 220398
//Dijkstra alghoritm

const MAX = 100000000

type vertice struct {
	steer *Steering
	dist  float64
	prev  int64
}

func newVert(steer *Steering) *vertice {
	v := new(vertice)
	v.steer = steer
	v.dist = MAX
	v.prev = 0
	return v
}

func findTracklistTo(train_id int64, block bool, start_track_id int64, target_id int64, target_type Railway_Object_Type, model_ptr *Simulation_Model) *[]int64 {
	var tl []int64 = nil

	var vert []*vertice = make([]*vertice, 0)

	var train_ptr *Train = GetTrain(train_id, model_ptr)

	var start_track *Track = GetTrack(start_track_id, model_ptr)

	var target_steer_id int64 = 0
	var target_steer_id_2 int64 = 0

	var target_track_ptr *Track
	var target_Train_ptr *Train

	if target_type == Type_Steering {
		target_steer_id = target_id
	} else if target_type == Type_Track {
		target_track_ptr = GetTrack(target_id, model_ptr)
		target_steer_id = target_track_ptr.st_end
		target_steer_id_2 = target_track_ptr.st_start
	} else if target_type == Type_Train {
		target_Train_ptr = GetTrain(target_id, model_ptr)
		if target_Train_ptr.on_steer != 0 {
			target_steer_id = target_Train_ptr.on_steer
		} else {
			target_track_ptr = GetTrack(target_Train_ptr.on_track, model_ptr)
			target_steer_id = target_track_ptr.st_end
			target_steer_id_2 = target_track_ptr.st_start
		}
	}

	//Ada.Text_IO.Put_Line("Looking for path from "+ int64'Image(start_track.st_start)+" or " + int64'Image(start_track.st_end)
	//                      + " to "+ int64'Image(target_steer_id)+" or " + int64'Image(target_steer_id_2)
	//)
	if (start_track.st_start == target_steer_id || start_track.st_end == target_steer_id) ||
		(target_steer_id_2 != 0 && (start_track.st_start == target_steer_id_2 || start_track.st_end == target_steer_id_2)) {
		tl = append(tl, start_track_id)
		return &tl
	}

	for it := 0; it < len(model_ptr.steer); it++ {
		// Ada.Text_IO.Put_Line("%%%%steering "+ int64'Image(model_ptr.steer[it].id ) + " used by" + int64'Image(model_ptr.steer[it].used_by) )
		if model_ptr.steer[it].used_by == 0 {
			if block == true {
				select {
				case model_ptr.steer[it].allowServiceTrain <- train_id:
					//  Ada.Text_IO.Put_Line("Steering "+ int64'Image(model_ptr.steer[it].id )+" accepted Train " + int64'Image(d_length))
					vert = append(vert, newVert(model_ptr.steer[it]))
				//case <- time.After(time.Microsecond * 5):
				default:
					//   Ada.Text_IO.Put_Line("Steering "+ int64'Image(model_ptr.steer[it].id )+" did not respond " + int64'Image(d_length))
				}
			} else {
				vert = append(vert, newVert(model_ptr.steer[it]))
			}
		} else if model_ptr.steer[it].used_by == train_id || (model_ptr.steer[it].id == target_steer_id || (target_steer_id_2 != 0 && model_ptr.steer[it].id == target_steer_id_2)) {
			//   Ada.Text_IO.Put_Line("Steering "+ int64'Image(model_ptr.steer[it].id )+" is already used by this Train " + int64'Image(d_length))
			vert = append(vert, newVert(model_ptr.steer[it]))

			if model_ptr.steer[it].id == start_track.st_end || model_ptr.steer[it].id == start_track.st_start {
				vert[len(vert)-1].dist = 0.0
			}

			// } else {
			//  Ada.Text_IO.Put_Line("Steering "+ int64'Image(model_ptr.steer[it].id )+" is used by other Train " + int64'Image(d_length))
		}
	}

	if block == true {
		for it := 0; it < len(model_ptr.track); it++ {
			select {
			case model_ptr.track[it].allowServiceTrain <- train_id:
				//nothing
				//case <- time.After(time.Microsecond * 5):
			default:
				//Ada.Text_IO.Put_Line("Track "+ int64'Image(model_ptr.track[it].id )+" failed to respond")
			}
		}
	}
	//Ada.Text_IO.Put_Line("Finished length: " + int64'Image(d_length))

	//dist = new ArrF(1..d_length)
	//prev = new ArrN(1..d_length)
	//steers = new ArrS(1..d_length)

	// arrays initialisation

	/* if block == true {
	            for it := 0 ; it<len(model_ptr.steer) it++ {
	               //Ada.Text_IO.Put_Line("$$$$$Steering "+ int64'Image(model_ptr.steer[it].id ) + " used by" + int64'Image(model_ptr.steer[it].used_by) )
	               if model_ptr.steer[it].used_by == train_id ||
				    model_ptr.steer[it].id == target_steer_id ||
				 ( target_steer_id_2 != 0 && model_ptr.steer[it].id == target_steer_id_2 )  {
	                  d_it = d_it + 1
	                  if d_it <= len(steers) {
	                     //Ada.Text_IO.Put_Line("Adding "+ int64'Image(model_ptr.steer[it].id )+" to list " + int64'Image(d_it))
	                     //Ada.Text_IO.Put_Line("#$%#%#@%#@%#%#$% d_it" + int64'Image(d_it) + " steers length:" + (int64'Image(steers'Length)) )
	                     steers(d_it) = model_ptr.steer[it]
	                     if model_ptr.steer[it].id == start_track.st_end || model_ptr.steer[it].id == start_track.st_start {
	                        dist(d_it) = 0.0
	                     } else {
	                        dist(d_it) = 100000000 //to lazy to makem max float
	                     }
	                     prev(d_it) = 0
	                  }
	              // } else {
	                  //Ada.Text_IO.Put_Line("Adding "+ int64'Image(model_ptr.steer[it].id )+" not added " + int64'Image(d_it))
	               }
	            }
	         } else {
	            for it := 0 ; it<len(model_ptr.steer) ; it++ {
	               if model_ptr.steer[it].used_by == train_id {
	                  d_it = d_it + 1
	                  steers[it] = model_ptr.steer[it]
	                  if model_ptr.steer[it].id == start_track.st_end || model_ptr.steer[it].id == start_track.st_start {
	                     dist[it] = 0.0
	                  } else {
	                     dist[it] = 100000000 //to lazy to makem max float
	                  }
	                  prev[it] = 0
	               }
	            }
	         }*/

	var d_it int = len(vert) - 1

	for d_it >= 1 {
		//  Ada.Text_IO.Put_Line("#$%#%  d_it:"+ int64'Image(d_it))
		//for it in steers) it++ {
		//   if steers[it] != nil {
		//      Ada.Text_IO.Put_Line("steer id:"+ int64'Image(steers[it].id) +" dist:" + Float'Image(dist[it]) + " prev " + int64'Image(prev[it]) )
		//   } else {
		//      Ada.Text_IO.Put_Line("nil" )
		//   }
		//
		//}

		//min vertex u from q
		var min_d = vert[0].dist
		var min_i = 0
		for it := 1; it <= d_it; it++ {
			// Ada.Text_IO.Put_Line("it:"+ int64'Image[it] +" d_it:" + int64'Image(d_it) + " dist: " + int64'Image(dist'Length) )
			if vert[it].dist <= min_d {
				min_d = vert[it].dist
				min_i = it
			}
		}

		//min u
		var u = vert[min_i]

		//remove u from q

		// Ada.Text_IO.Put_Line("min_i:" + int64'Image(min_i) + " steer: " + int64'Image(steer.id))
		vert[min_i] = vert[d_it]
		vert[d_it] = u
		d_it = d_it - 1

		for it := 0; it < len(model_ptr.track); it++ {

			//var copy_prev *vertice
			var v int64
			var is_v bool
			var it_v int
			var del float64
			var alt float64

			//  Ada.Text_IO.Put_Line("track " + int64'Image(model_ptr.track[it].id )
			//                       + " used by: " + int64'Image(model_ptr.track[it].used_by )
			//                        + " start " + int64'Image(model_ptr.track[it].st_start )
			//                        + " end " + int64'Image(model_ptr.track[it].st_end )
			//                        + " steer: " + int64'Image(steer.id)  )

			if (block == false || model_ptr.track[it].used_by == train_id) && (model_ptr.track[it].st_start == u.steer.id || model_ptr.track[it].st_end == u.steer.id) {
				//    Ada.Text_IO.Put_Line("neighbouring track: " +  int64'Image(model_ptr.track[it].id) )
				if model_ptr.track[it].st_start == u.steer.id {
					v = model_ptr.track[it].st_end
				} else {
					v = model_ptr.track[it].st_start
				}
				is_v = false
				it_v = 0
				for it2 := 0; it2 <= d_it; it2++ {
					if vert[it2].steer.id == v {
						is_v = true
						it_v = it2
						break
					}
				}
				//  Ada.Text_IO.Put_Line("v: " +  int64'Image(v) + " is ok? " +  bool'Image(is_v))

				if is_v == true {

					if model_ptr.track[it].t_type == Track_Type_Track {
						if model_ptr.track[it].data[T_max_speed] < train_ptr.max_speed {
							del = float64(model_ptr.track[it].data[T_distance]) / float64(model_ptr.track[it].data[T_max_speed]) * 60.0
						} else {
							del = float64(model_ptr.track[it].data[T_distance]) / float64(train_ptr.max_speed) * 60.0
						}
					} else {
						del = 1.0 //determining that platforms && service tracks will use only 1 minute for service track to Get through
					}

					alt = float64(u.steer.min_delay) + min_d + del

					//   Ada.Text_IO.Put_Line("alt: " +  Float'Image(alt) + " dist" +  Float'Image(dist(it_v)))

					if alt <= vert[it_v].dist {
						vert[it_v].dist = alt
						vert[it_v].prev = u.steer.id
					}

				}
			}
		}
		//  Ada.Text_IO.Put_Line("")

	}

	var min_1 float64
	var min_i_1 int = -1
	var min_i_2 int = -1
	var min_2 float64 = MAX //to lazy to makem max float

	//var leng int64 = 1

	for it := 0; it < len(vert); it++ {
		if vert[it].steer != nil { //==
			//   nil
			//  Ada.Text_IO.Put_Line("$#%#@$%#@%@#%#@% nil pointer" )
			//} else {
			//Ada.Text_IO.Put_Line( int64'Image(vert[it].steers.id) +" ==? "+ int64'Image(target_steer_id) +" || "+int64'Image(target_steer_id_2) )
			if vert[it].steer.id == target_steer_id {
				min_i_1 = it
				min_1 = vert[it].dist
			}
			if target_steer_id_2 != 0 && vert[it].steer.id == target_steer_id_2 {
				min_i_2 = it
				min_2 = vert[it].dist
			}
			if min_i_1 != -1 && (target_steer_id_2 == 0 || min_i_2 != -1) {
				break
			}
		}

	}
	//if min_i_1 == 0 {
	//Ada.Text_IO.Put_Line("#$%#%#@%#@% min_1:"+ Float'Image(min_1) +" min_i_1:"+ int64'Image(min_i_1) +" min_2:"+ Float'Image(min_2) +" min_i_2:"+ int64'Image(min_i_2) +"   target1"+ int64'Image(target_steer_id) +" target2:" + int64'Image(target_steer_id_2) )
	//}

	if min_i_2 != -1 && min_1 > min_2 {
		min_1 = min_2
		min_i_1 = min_i_2
	}
	min_i_2 = min_i_1
	if min_i_1 == -1 {

		if block == true {
			for it := 0; it < len(model_ptr.track); it++ {
				if model_ptr.track[it].used_by == train_id && model_ptr.track[it].id != train_ptr.on_track {
					model_ptr.track[it].freeFromServiceTrain <- train_id
				}
			}
			for it := 0; it < len(model_ptr.steer); it++ {
				if model_ptr.steer[it].used_by == train_id && model_ptr.steer[it].id != train_ptr.on_steer {
					model_ptr.steer[it].freeFromServiceTrain <- train_id
				}
			}
		}

		return nil
	}

	/*     for vert[min_i_1].prev != 0 {
	       for it:=0 ; it< len(vert) ; it++ {
	          if vert[it].steers.id == vert[min_i_1].prev {
	             min_i_1 = it
	             break
	          }
	       }
	    }*/

	/*    if block == false && target_type == Type_Track {
	         tl = new Track_ARRAY(1 .. len+1)
	         tl(len+1) = target_id
	      } else {
	         tl = new Track_ARRAY(1 .. len)
	      }*/

	//tl must be reversed after!
	if block == false && target_type == Type_Track {
		tl = append(tl, target_id)
	}

	//  Ada.Text_IO.Put_Line("#$%#%#@%#@%#%#$%#$%#$%#@%#%#$%#$%#$%#$ tracklist length:"+ int64'Image(tl'Length) )

	for vert[min_i_1].prev != 0 {
		//    Ada.Text_IO.Put_Line("#$%#%#@%#@%#%#$%#$%#$%#@%#%#$%#$%#$%#$ itt"+ int64'Image[itt] + " iterator:" +  int64'Image(min_i_1) + " " )

		for it := 0; it < len(model_ptr.track); it++ {
			//     Ada.Text_IO.Put_Line("#$%#%#@%#@%#%#$"
			//                             +" start: " +   int64'Image(model_ptr.track[it].st_start)
			//                             +" end: " + int64'Image(model_ptr.track[it].st_end)
			//                             +" curr: " + int64'Image(vert[min_i_1].steers.id)
			//                             +" prev: " + int64'Image(vert[min_i_1].prev)
			//     )
			if (model_ptr.track[it].st_start == vert[min_i_1].steer.id && model_ptr.track[it].st_end == vert[min_i_1].prev) ||
				(model_ptr.track[it].st_end == vert[min_i_1].steer.id && model_ptr.track[it].st_start == vert[min_i_1].prev) {

				tl = append(tl, model_ptr.track[it].id)
				break
			}
		}

		for it := 0; it < len(vert); it++ {
			if vert[it].steer.id == vert[min_i_1].prev {
				min_i_1 = it
				break
			}
		}

	}
	tl = append(tl, start_track_id)

	//reverse tl
	for it := 0; it < len(tl)/2; it++ {
		tl_cpy := tl[it]
		tl[it] = tl[len(tl)-it-1]
		tl[len(tl)-it-1] = tl_cpy

	}

	//for it in tl) it++ {
	//    Ada.Text_IO.Put_Line("it"+ int64'Image[it] + " tl:" +  int64'Image(tl[it]))
	//}
	if block == true {
		var found bool
		var track_ptr *Track

		for it := 0; it < len(model_ptr.track); it++ {
			if model_ptr.track[it].used_by == train_id && model_ptr.track[it].id != train_ptr.on_track {
				found = false
				for itt := 0; itt < len(tl); itt++ {
					if tl[itt] == model_ptr.track[it].id {
						found = true
						break
					}
				}
				if found == false {
					model_ptr.track[it].freeFromServiceTrain <- train_id
				}
			}
		}
		for it := 0; it < len(model_ptr.steer); it++ {
			if model_ptr.steer[it].used_by == train_id && model_ptr.steer[it].id != train_ptr.on_steer {
				found = false
				for itt := 0; itt < len(tl); itt++ {
					track_ptr = GetTrack(tl[itt], model_ptr)
					if track_ptr.st_end == model_ptr.steer[it].id || track_ptr.st_start == model_ptr.steer[it].id {
						found = true
						break
					}
				}
				if found == false {
					model_ptr.steer[it].freeFromServiceTrain <- train_id
				}
			}
		}
	}
	//    }
	//}

	return &tl
}

//@Author: Piotr Olejarz 220398
//Sim file starts whole simulation. And requires arguments below to work properly:
//<input file path> <'talking'/'responding'>");
//<input file path> - path where simulation can find file with model configuration
//<'talking'/'waiting'> - mode in which the simulation fill run:
//+ talking - information will be printed all the time
//+ waiting - information will be printed at user request

//does not work
func clear_scr() {
	fmt.Print("\033[2J")
}

//Available options
const (
	OPTION_CLEAR               string = "clear"
	OPTION_EXIT                string = "exit"
	OPTION_HELP                string = "help"
	OPTION_TRAINS              string = "trains"
	OPTION_TRACKS              string = "tracks"
	OPTION_STEERINGS           string = "steerings"
	OPTION_WORKERS             string = "workers"
	OPTION_STATIONS            string = "stations"
	OPTION_MODEL               string = "model"
	OPTION_TIMETABLE_TRAINS    string = "timetable-train"
	OPTION_TIMETABLE_PLATFORMS string = "timetable-platform"
)

/*
 *Task used for silent mode where simulation runs in background and additional task waits for user input
 */
func Silent_Task(model_ptr *Simulation_Model, silentDone chan bool) {
	// task body Silent_Task {
	var input string = ""
	var id int64
	if model_ptr != nil {
		fmt.Println("Type '" + OPTION_HELP + "' to receive command list")
		for input != OPTION_EXIT {

			if model_ptr.mode == Silent_Mode {
				fmt.Print("Type here: ")
			}
			_, err := fmt.Scanln(&input)
			//fmt.Println("input:" + input + "(" + strconv.Itoa(last) + ")")
			if err == nil {
				switch input {
				case OPTION_CLEAR:
					clear_scr()
				case OPTION_TRACKS:
					PrintTracksAlways(model_ptr)
				case OPTION_TRAINS:
					PrintTrainsAlways(model_ptr)
				case OPTION_STEERINGS:
					PrintSteeringsAlways(model_ptr)
				case OPTION_WORKERS:
					PrintWorkersAlways(model_ptr)
				case OPTION_STATIONS:
					PrintSteeringsAlways(model_ptr)
				case OPTION_MODEL:
					PrintModelAlways(model_ptr)
				case OPTION_TIMETABLE_TRAINS:
					fmt.Print("type train ID for which to calculate timetable: ")
					_, err1 := fmt.Scanln(&id)
					if err1 == nil {

						//fmt.Println("train:" + strconv.FormatInt(id, 10) + "(" + strconv.Itoa(last) + ")")

						PrintTrainTimetable(id, model_ptr)
					} else {
						fmt.Scanln()
						fmt.Println("Not a number")
					}
				case OPTION_TIMETABLE_PLATFORMS:
					fmt.Print("type platform ID for which to calculate timetable: ")
					_, err1 := fmt.Scanln(&id)
					if err1 == nil {
						//fmt.Println("track:" + strconv.FormatInt(id, 10) + "(" + strconv.Itoa(last) + ")")

						PrintTrackTimetable(id, model_ptr)
					} else {
						fmt.Scanln()
						fmt.Println("Not a number")
					}
				case OPTION_EXIT:
					EndSimulation(model_ptr)
					//os.Exit(0)
					silentDone <- true
				case OPTION_HELP:
					fmt.Println("Available commands:")
					fmt.Println(OPTION_TRAINS + "\t" + OPTION_TRACKS + "\t" + OPTION_STEERINGS)
					fmt.Println(OPTION_STATIONS + "\t" + OPTION_WORKERS + "\t" + OPTION_MODEL)
					fmt.Println(OPTION_TIMETABLE_TRAINS + "\t" + OPTION_TIMETABLE_PLATFORMS)
					fmt.Println( /*OPTION_CLEAR + "\t" +*/ OPTION_HELP + "\t" + OPTION_EXIT)
					// fmt.Println(OPTION_ + "\t" + OPTION_ + "\t" + OPTION_);
				default:
					fmt.Println("Illegal Command")
				}
			} else {
				fmt.Println("Invalid input")
			}
		}
	} else {
		fmt.Println("Silent Mode Task received null pointer")
	}
}

//starts simulation
func Simulation_start() {
	var proceed bool = false
	var model_ptr *Simulation_Model
	arg_it := 2
	if len(os.Args) > 1 {
		model_ptr = getModel(os.Args[1])
		if model_ptr != nil {
			if len(os.Args) > 2 {
				if os.Args[2] == "talking" {
					model_ptr.mode = Talking_Mode
					PutLineAlways("Selected talking mode for this simulation.")
					proceed = true
					arg_it = 3
				} else if os.Args[2] == "waiting" {
					model_ptr.mode = Silent_Mode
					PutLineAlways("Selected waiting mode for this simulation.")
					proceed = true
					arg_it = 3
				}
			}
			if !proceed {
				PutLineAlways("Selected mixed mode for this simulation.")
				if len(os.Args) > 2 && os.Args[2] == "mixed" {
					arg_it = 3
				}
				proceed = true
			}
		}
	}

	if proceed {

		if model_ptr.mode != Silent_Mode {
			if len(os.Args) > arg_it {
				if os.Args[arg_it] == "all" {
					PutLineAlways("All logs will be displayed.")
					arg_it = 4
				} else if os.Args[arg_it] == "2" {
					PutLineAlways("Only logs for second task will be displayed.")
					arg_it = 4
					model_ptr.log_mode = second_task
				} else if os.Args[arg_it] == "3" {
					PutLineAlways("Only logs for third task will be displayed.")
					arg_it = 4
					model_ptr.log_mode = third_task
				} else {
					PutLineAlways("All logs will be displayed.")
				}

				if len(os.Args) > arg_it && os.Args[arg_it] == "debug" {
					PutLineAlways("Debug logs will be displayed.")
					model_ptr.debug = true
				}

			}
		}

		model_ptr.work = true
		PrintModel(model_ptr, model_ptr.mode)
		model_ptr.start_time = time.Now()

		var silentDone = make(chan bool)

		for it := 0; it < len(model_ptr.steer); it++ {
			PutLine("creating SteeringTask for steering: "+strconv.FormatInt(model_ptr.steer[it].id, 10), model_ptr)
			go SteeringTask(model_ptr.steer[it], model_ptr)
		}
		for it := 0; it < len(model_ptr.track); it++ {
			PutLine("creating TrackTask for track: "+strconv.FormatInt(model_ptr.track[it].id, 10), model_ptr)
			go TrackTask(model_ptr.track[it], model_ptr)
		}
		for it := 0; it < len(model_ptr.worker); it++ {
			PutLine("creating WorkerTask for worker: "+strconv.FormatInt(model_ptr.worker[it].id, 10), model_ptr)
			go WorkerTask(model_ptr.worker[it], model_ptr)
		}
		for it := 0; it < len(model_ptr.station); it++ {
			PutLine("creating StationTask for station: "+strconv.FormatInt(model_ptr.station[it].id, 10), model_ptr)
			go StationTask(model_ptr.station[it], model_ptr)
		}
		for it := 0; it < len(model_ptr.train); it++ {
			PutLine("creating TrainTask for train: "+strconv.FormatInt(model_ptr.train[it].id, 10), model_ptr)
			go TrainTask(model_ptr.train[it], model_ptr)
		}

		if model_ptr.mode != Talking_Mode {
			fmt.Println("creating go routine for Silent_Task")
			go Silent_Task(model_ptr, silentDone)

			<-silentDone
			//fmt.Println("go routine of Silent_Task has finished working")
		} else {
			for {
				time.Sleep(24 * time.Hour)
				fmt.Println("The ride never ends")
			}
		}
	} else {
		fmt.Println("usage: ")
		fmt.Println("<input file path> ('mixed'/'talking'/'waiting') (!waiting -> ('all'/'2'/'3')) ('debug')")
		fmt.Println("required arguments: ")
		fmt.Println("<input file path>")
		fmt.Println("<input file path> - path where simulation can find file with model configuration")
		fmt.Println("optional arguments after required:")
		fmt.Println("('mixed'/'talking'/'waiting') ('all'/'2'/'3') ('debug')")
		fmt.Println("('mixed'/'talking'/'waiting') - mode in which the simulation fill run:")
		fmt.Println("+ mixed - will both print logs and allow user to request information. It is default parameter and can be ommited.")
		fmt.Println("+ talking - logs will be printed all the time.")
		fmt.Println("+ waiting - information will be printed only at user request.")
		fmt.Println("(!waiting -> ('all'/'2'/'3')) - only for mixed/talking mode! Selects which normal logs will be displayed. Does not affect debug logs")
		fmt.Println("+ all - all normal logs will be displayed. It is default parameter and can be ommited.")
		fmt.Println("+ 2 - only normal logs directly related to second task will be displayed.")
		fmt.Println("+ 3 - only normal logs directly related to third task will be displayed.")
		fmt.Println("('debug') - allows debug logs to be displayed. Does not affect normal logs. ")

	}
}


//File model has Simulation_Model record declaration which contains all necessery data for simulation to work
//Also has methods used to get objects like tracks from it's ID and string representations of those objects

//Enum for simulation modes
type Simulation_Mode int64

type Log_Modes int64

//Enum for time intervals
type Time_Interval int64

func when(b bool, c chan int64) chan int64 {
	if !b {
		return nil
	} else {
		return c
	}
}

func whenSTS(b bool, c chan StartTaskStruct) chan StartTaskStruct {
	if !b {
		return nil
	} else {
		return c
	}
}
func whenTSP(b bool, c chan TrainStatPair) chan TrainStatPair {
	if !b {
		return nil
	} else {
		return c
	}
}

const (
	Time_Interval_Minute Time_Interval = iota
	Time_Interval_Hour   Time_Interval = iota
)

const (
	Mixed_Mode   Simulation_Mode = iota
	Silent_Mode  Simulation_Mode = iota
	Talking_Mode Simulation_Mode = iota
)
const (
	all_output  Log_Modes = iota
	second_task Log_Modes = iota
	third_task  Log_Modes = iota
)

/*
 * Simulation record used to store all necessary data
 */
type Simulation_Model struct {
	speed      int64
	start_time time.Time
	mode       Simulation_Mode

	log_mode Log_Modes
	debug    bool

	steer   []*Steering
	track   []*Track
	train   []*Train
	platf   []*Track
	station []*Station
	worker  []*Worker

	work bool
}

func EndSimulation(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		model_ptr.work = false
	}
}

//returns train with given ID
func GetServiceTrain(model_ptr *Simulation_Model) *Train {
	var train_ptr *Train
	var service_ptr *Train = nil
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.train); it++ {
			train_ptr = model_ptr.train[it]
			if train_ptr != nil && train_ptr.t_type == Train_Type_Service {
				service_ptr = train_ptr
				if service_ptr.data[T_going_back] != 0 {
					return service_ptr
				}
			}
		}
	}
	return service_ptr
}

//returns worker with given ID
func GetWorker(work_id int64, model_ptr *Simulation_Model) *Worker {
	var work_ptr *Worker
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.worker); it++ {
			work_ptr = model_ptr.worker[it]
			if work_ptr != nil && work_ptr.id == work_id {
				return work_ptr
			}
		}
	}
	return nil
}

//returns station with given ID
func GetStation(stat_id int64, model_ptr *Simulation_Model) *Station {
	var stat_ptr *Station
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.station); it++ {
			stat_ptr = model_ptr.station[it]
			if stat_ptr != nil && stat_ptr.id == stat_id {
				return stat_ptr
			}
		}
	}
	return nil
}

//returns train with given ID
func GetTrain(train_id int64, model_ptr *Simulation_Model) *Train {

	var train_ptr *Train
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.train); it++ {
			train_ptr = model_ptr.train[it]
			if train_ptr != nil && train_ptr.id == train_id {
				return train_ptr
			}
		}
	}
	return nil
}

//returns track with given ID
func GetTrack(track_id int64, model_ptr *Simulation_Model) *Track {
	var track_ptr *Track

	if model_ptr != nil {
		for it := 0; it < len(model_ptr.track); it++ {
			track_ptr = model_ptr.track[it]
			if track_ptr != nil && track_ptr.id == track_id {
				return track_ptr
			}
		}
	}
	return nil
}

//returns next track for given train
func GetNextTrack(train_ptr *Train, model_ptr *Simulation_Model) *Track {

	if train_ptr != nil {
		return GetTrack((*train_ptr.tracklist)[int(train_ptr.track_it)%int(len(*train_ptr.tracklist))], model_ptr)
	}
	return nil
}

//returns steering with given ID
func GetSteering(steering_id int64, model_ptr *Simulation_Model) *Steering {
	var steering_ptr *Steering

	if model_ptr != nil {
		for it := 0; it < len(model_ptr.steer); it++ {
			steering_ptr = model_ptr.steer[it]
			if steering_ptr != nil && steering_ptr.id == steering_id {
				return steering_ptr
			}
		}
	}
	return nil
}

//returns steering at either start of
func GetSteeringFromTrack(track_ptr *Track, end_of_track bool, model_ptr *Simulation_Model) *Steering {
	if track_ptr != nil {
		if end_of_track {
			return GetSteering(track_ptr.st_end, model_ptr)
		} else {
			return GetSteering(track_ptr.st_start, model_ptr)
		}
	}
	return nil
}

// string representation of given track
func TrackToString(track_ptr *Track) string {

	if track_ptr != nil {
		txt := "id:" + strconv.FormatInt(track_ptr.id, 10) +
			" , steerings:([" + strconv.FormatInt(track_ptr.st_start, 10) +
			"],[" + strconv.FormatInt(track_ptr.st_end, 10) + "])"
		if track_ptr.t_type == Track_Type_Track {

			if track_ptr.used_by == 0 {
				return ("track " + txt +
					" , out of order: " + strconv.FormatBool(track_ptr.out_of_order) +
					" , not used , speed:" + strconv.FormatInt(track_ptr.data[T_max_speed], 10) + "kmph , " +
					"dist:" + strconv.FormatInt(track_ptr.data[T_distance], 10) + "km")
			} else {
				return ("track " + txt +
					" , out of order: " + strconv.FormatBool(track_ptr.out_of_order) +
					" , used by train[" + strconv.FormatInt(track_ptr.used_by, 10) + "] , " +
					"speed:" + strconv.FormatInt(track_ptr.data[T_max_speed], 10) + "kmph , " +
					"dist:" + strconv.FormatInt(track_ptr.data[T_distance], 10) + "km")
			}

		} else if track_ptr.t_type == Track_Type_Platform {
			if track_ptr.used_by == 0 {
				return ("platform " + txt +
					" , station: " + strconv.FormatInt(track_ptr.data[T_station_id], 10) +
					" , out of order: " + strconv.FormatBool(track_ptr.out_of_order) +
					" , not used , delay:" + strconv.FormatInt(track_ptr.data[T_min_delay], 10) + "min")
			} else {
				return "platform " + txt +
					" , station: " + strconv.FormatInt(track_ptr.data[T_station_id], 10) +
					" , out of order: " + strconv.FormatBool(track_ptr.out_of_order) +
					" , used by train[" + strconv.FormatInt(track_ptr.used_by, 10) +
					"] , delay:" + strconv.FormatInt(track_ptr.data[T_min_delay], 10) + "min"
			}
		} else {
			if track_ptr.used_by == 0 {
				return "Service track " + txt +
					" , out of order: " + strconv.FormatBool(track_ptr.out_of_order) +
					" , not used"
			} else {
				return "Service track " + txt +
					" , out of order: " + strconv.FormatBool(track_ptr.out_of_order) +
					" , used by train[" + strconv.FormatInt(track_ptr.used_by, 10) + "]"
			}
		}
	} else {
		return "null"
	}
}

// string representation of given train
func TrainToString(train_ptr *Train) string {
	var track_list string = ""
	var station_list string = ""
	var passengers string = ""
	var pos string

	if train_ptr != nil {
		if train_ptr.t_type == Train_Type_Normal {
			for it := 0; it < len(*train_ptr.tracklist); it++ {
				if track_list != "" {
					track_list += "," + strconv.FormatInt((*train_ptr.tracklist)[it], 10)
				} else {
					track_list = strconv.FormatInt((*train_ptr.tracklist)[it], 10)
				}
			}
			for it := 0; it < len(*train_ptr.stationlist); it++ {
				if station_list != "" {
					station_list += "," + strconv.FormatInt((*train_ptr.stationlist)[it], 10)
				} else {
					station_list = strconv.FormatInt((*train_ptr.stationlist)[it], 10)
				}
			}
			for work_ptr := range *train_ptr.passengers {
				if passengers != "" {
					passengers += "," + strconv.FormatInt(work_ptr.id, 10)
				} else {
					passengers = strconv.FormatInt(work_ptr.id, 10)
				}
			}

			if train_ptr.on_track != 0 {
				pos = "on track[" + strconv.FormatInt(train_ptr.on_track, 10) + "]"
			} else if train_ptr.on_steer != 0 {
				pos = "on steering[" + strconv.FormatInt(train_ptr.on_steer, 10) + "]"
			} else {
				pos = "nowhere"
			}

			return "train id:" + strconv.FormatInt(train_ptr.id, 10) +
				" , out of order: " + strconv.FormatBool(train_ptr.out_of_order) +
				" , location: " + pos +
				" , max spd:" + strconv.FormatInt(train_ptr.max_speed, 10) +
				"kmph , curr spd:" + strconv.FormatInt(train_ptr.current_speed, 10) +
				"kmph , cap:" + strconv.FormatInt(train_ptr.data[T_capacity], 10) +
				" , tracklist(at " + strconv.FormatInt(int64(train_ptr.track_it), 10) +
				"){" + track_list + "}" +
				" , stations{" + station_list + "} uniq: " + strconv.FormatInt(int64(train_ptr.data[T_uniqueStations]), 10) +
				" , passengers{" + passengers + "}"

		} else {
			if train_ptr.tracklist != nil {
				track_list = ""
				for it := 0; it < len(*train_ptr.tracklist); it++ {
					if track_list != "" {
						track_list += "," + strconv.FormatInt((*train_ptr.tracklist)[it], 10)
					} else {
						track_list = strconv.FormatInt((*train_ptr.tracklist)[it], 10)
					}
				}
				track_list = " , tracklist(at " + strconv.FormatInt(int64(train_ptr.track_it), 10) + "): {" + track_list + "}"
			}
			if train_ptr.on_track != 0 {
				pos = "on track[" + strconv.FormatInt(train_ptr.on_track, 10) + "]"
			} else if train_ptr.on_steer != 0 {
				pos = "on steering[" + strconv.FormatInt(train_ptr.on_steer, 10) + "]"
			} else {
				pos = "nowhere"
			}

			return "service train id:" + strconv.FormatInt(train_ptr.id, 10) +
				" , out of order: " + strconv.FormatBool(train_ptr.out_of_order) +
				" , location: " + pos +
				" , max spd:" + strconv.FormatInt(train_ptr.max_speed, 10) +
				"kmph , curr spd:" + strconv.FormatInt(train_ptr.current_speed, 10) +
				" , tracklist(at " + strconv.FormatInt(int64(train_ptr.track_it), 10) + ")" +
				track_list

		}
	} else {
		return "null"
	}
}

// string representation of given steering
func SteeringToString(steer_ptr *Steering) string {
	if steer_ptr != nil {

		txt := "Steering id:" + strconv.FormatInt(steer_ptr.id, 10) +
			" , delay:" + strconv.FormatInt(steer_ptr.min_delay, 10) + "min"

		if steer_ptr.used_by == 0 {
			return txt + " , out of order: " + strconv.FormatBool(steer_ptr.out_of_order) + " , not used"
		} else {
			return txt + " , out of order: " + strconv.FormatBool(steer_ptr.out_of_order) + " , used by train[" + strconv.FormatInt(steer_ptr.used_by, 10) + "]"
		}

	} else {
		return "null"
	}
}

// string representation of given worker
func WorkerToString(work_ptr *Worker) string {
	if work_ptr != nil {
		var connection_list string = ""

		if work_ptr.connectionlist != nil {
			for it := 0; it < len(*work_ptr.connectionlist); it++ {
				conn := (*work_ptr.connectionlist)[it]
				conn_ent := "(t:" + strconv.FormatInt(conn.train_id, 10) + " ,d:" + strconv.FormatInt(conn.depart_station_id, 10) + " ,a:" + strconv.FormatInt(conn.arrive_station_id, 10) + ")"
				if connection_list != "" {
					connection_list += "," + conn_ent
				} else {
					connection_list = conn_ent
				}
			}
		}

		txt := "Worker id:" + strconv.FormatInt(work_ptr.id, 10) +
			" , home station:" + strconv.FormatInt(work_ptr.home_stat_id, 10) +
			" , at stat:" + strconv.FormatInt(work_ptr.on_Station, 10) +
			" , on train:" + strconv.FormatInt(work_ptr.on_train, 10)
		switch work_ptr.state {
		case AtHome:
			return txt + " , at home: "
		case TravellingToWork:
			txt += " , going to work at station:" + strconv.FormatInt(work_ptr.dest_Station, 10)
			if work_ptr.on_train != 0 {
				return txt +
					" , on train:" + strconv.FormatInt(work_ptr.on_train, 10) +
					" , connection list: { " + connection_list + "}"
			} else {
				return txt +
					" , at station:" + strconv.FormatInt(work_ptr.on_Station, 10) +
					" , connection list: { " + connection_list + "}"
			}
		case WaitingForWork:
			return txt + " , waiting to start work at station: " + strconv.FormatInt(work_ptr.dest_Station, 10)
		case TravellingToHome:
			txt += " , going back to home"
			if work_ptr.on_train != 0 {
				return txt +
					" , on train:" + strconv.FormatInt(work_ptr.on_train, 10) +
					" , connection list: { " + connection_list + "}"
			} else {
				return txt +
					" , at station:" + strconv.FormatInt(work_ptr.on_Station, 10) +
					" , connection list: { " + connection_list + "}"
			}
		case Working:
			return txt + " , working at station: " + strconv.FormatInt(work_ptr.dest_Station, 10)
		default:
			return txt + " , undefined state"
		}

	} else {
		return "null"
	}
}

// string representation of given station
func StationToString(stat_ptr *Station) string {
	if stat_ptr != nil {
		var passengers string = ""
		var r_work string = ""
		var c_work string = ""
		var trains string = ""

		for work_ptr := range stat_ptr.passengers {
			if passengers != "" {
				passengers += "," + strconv.FormatInt(work_ptr.id, 10)
			} else {
				passengers = strconv.FormatInt(work_ptr.id, 10)
			}
		}
		for work_ptr := range stat_ptr.ready_workers {
			if r_work != "" {
				r_work += "," + strconv.FormatInt(work_ptr.id, 10)
			} else {
				r_work = strconv.FormatInt(work_ptr.id, 10)
			}
		}
		for work_ptr := range stat_ptr.chosen_workers {
			if c_work != "" {
				c_work += "," + strconv.FormatInt(work_ptr.id, 10)
			} else {
				c_work = strconv.FormatInt(work_ptr.id, 10)
			}
		}
		for train := range stat_ptr.trains {
			if trains != "" {
				trains += "," + strconv.FormatInt(train, 10)
			} else {
				trains = strconv.FormatInt(train, 10)
			}
		}

		return "Station id: " + strconv.FormatInt(stat_ptr.id, 10) +
			" ,trains on station: { " + trains + "}" +
			" ,passengers: { " + passengers + "}" +
			" ,chosen workers: { " + c_work + "}" +
			" ,ready workers: { " + r_work + "}"

	} else {
		return "null"
	}
}

// Translates given simulation time to real time seconds
func GetTimeSimToRealFromModel(time_in_interval float64, interval Time_Interval, model_ptr *Simulation_Model) float64 {
	//speed  int64; // real-time seconds to simulation-hour ratio

	if model_ptr != nil {
		return GetTimeSimToReal(time_in_interval, interval, model_ptr.speed)
	} else {
		fmt.Println("getTrain received nil pointer. Returning 1")
		return 1.0
	}

}

// Translates given simulation time to real time seconds
func GetTimeSimToReal(time_in_interval float64, interval Time_Interval, speed int64) float64 {
	//speed  int64; // real-time seconds to simulation-hour ratio

	if interval == Time_Interval_Hour {
		return time_in_interval * float64(speed)
	} else if interval == Time_Interval_Minute {
		return time_in_interval * float64(speed) / 60.0
	} else {
		fmt.Println("getTrain received illegal interval. Returning 1")
		return 1.0
	}
}

type Connection struct {
	train_id          int64
	depart_station_id int64
	arrive_station_id int64
}

func GetConnection(source_station int64, destination_station int64, model_ptr *Simulation_Model) *[]*Connection {
	rtrn := make([]*Connection, 0)

	if source_station != destination_station {
		//var max_d_it int64 = 0
		//var rtrn_st_it int64 = 1
		var found_connection bool = false
		var depar = make([]*Train, 0)
		var arriv = make([]*Train, 0)
		//var arriv_it int64 = 0
		//var depar_it int64 = 0
		var src int64 = source_station
		var des int64 = destination_station

		var found_arrv bool
		var found_dep bool
		//var max_a_stat int64 = 0
		//var max_a_it int64 = 0
		//var max_d_stat int64 = 0
		var train_ptr *Train

		//find all trains that start from source station
		for it := 0; it < len(model_ptr.train); it++ {
			train_ptr = model_ptr.train[it]
			if train_ptr.t_type == Train_Type_Normal {
				found_arrv = false
				found_dep = false
				for itt := 0; itt < len(*train_ptr.stationlist); itt++ {

					if (*train_ptr.stationlist)[itt] == src {
						/*if max_d_stat <= train_ptr.data[T_uniqueStations] {
							max_d_it = int64(itt)
							max_d_stat = train_ptr.data[T_uniqueStations]
						}*/
						found_dep = true
					}

					if (*train_ptr.stationlist)[itt] == des {
						/*if max_a_stat <= train_ptr.data[T_uniqueStations] {
							max_a_it = int64(itt)
							max_a_stat = train_ptr.data[T_uniqueStations]
						}*/
						found_arrv = true
					}

					if found_dep && found_arrv {
						break
					}
				}
				if found_dep || found_arrv { //found direct connection between source and destination
					if found_dep && found_arrv {
						con := new(Connection)
						con.train_id = train_ptr.id
						con.depart_station_id = src
						con.arrive_station_id = des
						rtrn = append(rtrn, con)
						found_connection = true
						break
					} else {
						if found_dep {
							depar = append(depar, train_ptr)
						} else {
							arriv = append(arriv, train_ptr)
						}
					}
				}
			}
		}
		if !found_connection {
			//trying to find connection with one train switch

			for it_a := 0; it_a < len(arriv); it_a++ {
				arriv_ptr := arriv[it_a]
				if arriv_ptr != nil {
					for it_d := 0; it_d < len(depar); it_d++ {
						depar_ptr := depar[it_d]
						if depar_ptr != nil {
							for s_a := 0; s_a < len(*arriv_ptr.stationlist); s_a++ {
								for s_d := 0; s_d < len(*depar_ptr.stationlist); s_d++ {
									if (*arriv_ptr.stationlist)[s_a] == (*depar_ptr.stationlist)[s_d] {
										found_connection = true
										con1 := new(Connection)
										con2 := new(Connection)

										con1.train_id = depar_ptr.id
										con1.depart_station_id = src
										con2.train_id = arriv_ptr.id
										con2.arrive_station_id = des

										con1.arrive_station_id = (*arriv_ptr.stationlist)[s_a]
										con2.depart_station_id = (*arriv_ptr.stationlist)[s_a]
										rtrn = append(rtrn, con1)
										rtrn = append(rtrn, con2)

										break
									}
								}
								if found_connection == true {
									break
								}
							}
						}
						if found_connection == true {
							break
						}
					}
				}
				if found_connection == true {
					break
				}
			}
			//choosing train to
			if !found_connection {
				return nil
			}
		}
	}
	return &rtrn
}

//@Author: Piotr Olejarz 220398
//File log is used for output and logging purposes. All output should be redirected here so it will be processed correspondingly with simulation options.

//prints new line
func PutLineAlways(line string) {
	//fmt.Println(line);
	fmt.Println( /*Ada.Calendar.Formatting.Image(Ada.Calendar.Clock)+":"&*/ line)
}

//prints new line depending on mode
func PutLine(line string, model_ptr *Simulation_Model) {
	//fmt.Println("test");
	if model_ptr != nil && model_ptr.mode != Silent_Mode {

		if !model_ptr.debug && strings.HasPrefix(line, "#debug#") {
			return
		}
		if model_ptr.log_mode == second_task && !strings.HasPrefix(line, "#2#") {
			return
		}
		if model_ptr.log_mode == third_task && !strings.HasPrefix(line, "#3#") {
			return
		}
		PutLineAlways(line)
	}

}

//prints model based on mode
func PrintModel(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintModelAlways(model_ptr)
	}

}

//prints model
func PrintModelAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		fmt.Println("Simulation speed: " + strconv.FormatInt(model_ptr.speed, 10) + "s -> 1h")
		PrintSteeringsAlways(model_ptr)
		PrintStationsAlways(model_ptr)
		PrintTracksAlways(model_ptr)
		PrintTrainsAlways(model_ptr)
		PrintWorkersAlways(model_ptr)
	}
}

//prints steerings based on mode
func PrintSteerings(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintSteeringsAlways(model_ptr)
	}

}

//prints steerings
func PrintSteeringsAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.steer); it++ {
			fmt.Println(SteeringToString(model_ptr.steer[it]))
		}
	}
}

//prints tracks based on mode
func PrintTracks(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTracksAlways(model_ptr)
	}

}

//prints tracks
func PrintTracksAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.track); it++ {
			fmt.Println(TrackToString(model_ptr.track[it]))
		}
	}
}

//prints trains based on mode
func PrintTrains(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTrainsAlways(model_ptr)
	}

}

//prints trains
func PrintTrainsAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.train); it++ {
			fmt.Println(TrainToString(model_ptr.train[it]))
		}
	}
}

//prints stations based on mode
func PrintStations(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintStationsAlways(model_ptr)
	}

}

//prints stations
func PrintStationsAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.station); it++ {
			fmt.Println(StationToString(model_ptr.station[it]))
		}
	}
}

//prints workers based on mode
func PrintWorkers(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintWorkersAlways(model_ptr)
	}

}

//prints workers
func PrintWorkersAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.worker); it++ {
			fmt.Println(WorkerToString(model_ptr.worker[it]))
		}
	}
}

//prints train locations based on mode
func PrintTrainLocations(model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTrainLocationsAlways(model_ptr)
	}

}

//prints train locations
func PrintTrainLocationsAlways(model_ptr *Simulation_Model) {
	if model_ptr != nil {
		for it := 0; it < len(model_ptr.train); it++ {
			c_t := model_ptr.train[it]
			if c_t.on_track != 0 {
				fmt.Println(strconv.FormatInt(c_t.id, 10) +
					" { at track: " + strconv.FormatInt((*c_t.tracklist)[c_t.track_it], 10) +
					" and moves at " + strconv.FormatInt(c_t.current_speed, 10) + "kmph")
			} else {
				fmt.Println(strconv.FormatInt(c_t.id, 10) +
					" { at steering: " + strconv.FormatInt((*c_t.tracklist)[c_t.track_it], 10))
			}
		}
	}
}

//prints given train status status based on mode
func PrintTrainStatusFromIDAlways(train_id int64, model_ptr *Simulation_Model) {
	PrintTrainStatusAlways(GetTrain(train_id, model_ptr))
}

//prints given train status status
func PrintTrainStatusFromID(train_id int64, model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTrainStatusFromIDAlways(train_id, model_ptr)
	}
}

//prints given train status status based on mode
func PrintTrainStatus(train_ptr *Train, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTrainStatusAlways(train_ptr)
	}
}

//prints given train status status
func PrintTrainStatusAlways(train_ptr *Train) {
	fmt.Println(TrainToString(train_ptr))
}

//prints given steering status based on mode
func PrintSteeringFromIDStatus(steer_id int64, model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintSteeringStatusFromIDAlways(steer_id, model_ptr)
	}
}

//prints given steering status
func PrintSteeringStatusFromIDAlways(steer_id int64, model_ptr *Simulation_Model) {
	PrintSteeringStatusAlways(GetSteering(steer_id, model_ptr))
}

//prints given steering status based on mode
func PrintSteeringStatus(steer_ptr *Steering, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintSteeringStatusAlways(steer_ptr)
	}
}

//prints given steering status
func PrintSteeringStatusAlways(steer_ptr *Steering) {
	fmt.Println(SteeringToString(steer_ptr))
}

//prints given track status based on mode
func PrintTrackStatusFromID(track_id int64, model_ptr *Simulation_Model, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTrackStatusFromIDAlways(track_id, model_ptr)
	}
}

//prints given track status
func PrintTrackStatusFromIDAlways(track_id int64, model_ptr *Simulation_Model) {
	PrintTrackStatusAlways(GetTrack(track_id, model_ptr))
}

//prints given track status based on mode
func PrintTrackStatus(track_ptr *Track, talking_mode Simulation_Mode) {
	if talking_mode != Silent_Mode {
		PrintTrackStatusAlways(track_ptr)
	}
}

//prints given track status
func PrintTrackStatusAlways(track_ptr *Track) {
	fmt.Println(TrackToString(track_ptr))
}

type Time struct {
	day    int64
	hour   int64
	minute int64
}

func GetRelativeTime(current_time time.Time, model_ptr *Simulation_Model) Time {
	var t Time

	dur := current_time.Sub(model_ptr.start_time)
	sim_time := dur.Seconds() / float64(model_ptr.speed)

	//Float'Value(Duration'Image(Ada.Real_Time.To_Duration(current_time-model_ptr.start_time))) / Float(model_ptr.speed);

	//fmt.Println("dur sec:" + strconv.FormatFloat(dur.Seconds(), 'f', 3, 64))
	//fmt.Println("Sim+time:" + strconv.FormatFloat(sim_time, 'f', 3, 64))

	t.day = int64(sim_time) / 24
	t.hour = int64(sim_time) % 24
	//ada.Text_IO.Put_Line("&&&:"&Float'Image(Float'Fraction(sim_time)));
	_, frac := math.Modf(sim_time)
	min := int64(frac * 60.0)
	if min < 0 {
		t.minute = 60 + min
		t.hour = t.hour - 1
	} else {
		t.minute = min
	}

	return t
}

func TimeToString(t Time) string {
	// if t.day != 0 {
	//  return strconv.FormatInt(t.day)+"d "+strconv.FormatInt(t.hour)+":"+strconv.FormatInt(t.minute);
	//}else{
	return "+" + strconv.FormatInt(t.day, 10) + "d " + strconv.FormatInt(t.hour, 10) + "h " + strconv.FormatInt(t.minute, 10) + "m"
	//}

}

//prints timetable for given train
func PrintTrainTimetable(id int64, model_ptr *Simulation_Model) {
	if model_ptr != nil {
		train_ptr := GetTrain(id, model_ptr)
		if train_ptr != nil {
			fmt.Println("Timetable for train:" + strconv.FormatInt(train_ptr.id, 10))
			fmt.Println("platform\tarrival\tdeparture")
			for it := 0; it < len(train_ptr.history); it++ {
				th := train_ptr.history[it]
				if th.object_type == Type_Platform {

					t_arr := GetRelativeTime(th.arrival, model_ptr)
					t_dep := GetRelativeTime(th.departure, model_ptr)
					fmt.Println(strconv.FormatInt(th.object_id, 10) + "\t" + TimeToString(t_arr) + "\t" + TimeToString(t_dep))

				}

			}
		} else {
			fmt.Println("Train not found!")
		}
	} else {
		fmt.Println("Null model!")
	}
}

//prints timetable for given track
func PrintTrackTimetable(id int64, model_ptr *Simulation_Model) {
	if model_ptr != nil {
		track_ptr := GetTrack(id, model_ptr)
		if track_ptr != nil {
			fmt.Println("Timetable for track:" + strconv.FormatInt(track_ptr.id, 10))
			fmt.Println("train\tarrival\tdeparture")
			for it := 0; it < len(track_ptr.history); it++ {
				th := track_ptr.history[it]
				t_arr := GetRelativeTime(th.arrival, model_ptr)
				t_dep := GetRelativeTime(th.departure, model_ptr)
				fmt.Println(strconv.FormatInt(th.train_id, 10) + "\t" + TimeToString(t_arr) + "\t" + TimeToString(t_dep))
			}
		} else {
			fmt.Println("Track not found!")
		}
	} else {
		fmt.Println("Null model!")
	}

}
