package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Wappsto/wedge-api/go/slx"
	"github.com/Wappsto/wedge-api/go/wedge"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var (
	nodeServer = &wedge.NodeIdentity{
		Id: uuid.New().String(),
	}
	number = &slx.Number{}
	state  = []*slx.State{}
	device = []*slx.Device{}
	// value  = []slx.Value{}
)

func main() {

	// ================================================
	// Defining model struct -> with inner embedded structs

	state = []*slx.State{
		{
			Data: "25",     //string    `json:"data"` -> we can generate pseudo.random number with rand.Intn()
			Type: slx.Type_Report, //string    `json:"type"`
			Id:   2,
		},
	}
	number = &slx.Number{
		Min:  0.01,         //int    `json:"min"`
		Max:  100,       //int    `json:"max"`
		Step: 1,         //int    `json:"step"`
		Unit: "Celsius", // string `json:"unit"`
	}

	value := []*slx.Value{
			{
				Name:       "temperature", //string `json:"name"`
				Type:       "",            //string `json:"type"`
				Status:     "OK",          //string `json:"status"`
				Permission: "RW",          //string `json:"permission"
				Number:     number,
				State:      state,
				Id:         1,
			},
		}
	device = []*slx.Device{
		{
			Name:          "Temperature-Sensor",            //string `json:"name"`
			Manufacturer:  "Mitsumi",                       //string `json:"manufacturer"`
			Product:       "MM3286",                        //string `json:"product"`
			Serial:        "bla-bla-bla",                   //string `json:"serial"`
			Description:   "Reliable Japanese temp sensor", //string `json:"description"`
			Protocol:      "PROFINET/ETHERNET-TCP/IP",      //string `json:"protocol"`
			Communication: "always",                        //string `json:"communication"`
			Id:            1,
			Value: value,
		},
	}

	
	
	
	// ================================================================================
	// Making node client -> calling methods from wedge server (setDevice, setModel etc...)
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to establish connection ")
	}
	defer connection.Close()

	Model := &wedge.Model{
		Node: nodeServer,
		Device: device,
	}
	wedgeClient := wedge.NewWedgeClient(connection)

	// modelRequest is request for SetModel method
	var modelRequest = wedge.SetModelRequest{}
	modelRequest = wedge.SetModelRequest{
		Model: Model,
	}

	// SetModel returns slx.Reply{} as response, from wedgeServer
	response, err := wedgeClient.SetModel(context.Background(), &modelRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response about model received from wedge server: %v\n", response)
	fmt.Println("==============================================")

	// var updatedData int
	// var wg *sync.WaitGroup
	for {
		rand.Seed(time.Now().Unix())
		updatedData := rand.Intn(101)
		fmt.Printf("New temperature is %d Celsius\n", updatedData)
		setStateRequest := wedge.SetStateRequest{
			Node:   nodeServer,
			DeviceId: 1,
			ValueId:  1,
			State: &slx.State{
				Data: strconv.Itoa(updatedData),
				Id:   2,
			},
		}
		response, err := wedgeClient.SetState(context.Background(), &setStateRequest)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Response about state received from wedge server: %v\n", response)
		fmt.Println("==============================================")
		time.Sleep(2 * time.Second)
		// wg.Add(1)
		// go func(sr *wedge.SetStateRequest) {
		// 	response, err := wedgeClient.SetState(context.Background(), sr)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	log.Printf("Response about state received from wedge server: %v\n", response)
		// 	fmt.Println("==============================================")
		// 	wg.Done()
		// }(&setStateRequest)
		// wg.Wait()
	}
}
