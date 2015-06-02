package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Order struct {
	Username         string `json:"username"`
	PaintPrimary     string `json:"paintPrimary"`
	PaintSecondary   string `json:"paintSecondary"`
	WindowTint       string `json:"windowTint"`
	BulletproofTires string `json:"bulletproofTires"`
	EMSTune          string `json:"emsTune"`
	Brakes           string `json:"brakes"`
	TireSmokeR       string `json:"tireSmokeR"`
	TireSmokeB       string `json:"tireSmokeB"`
	TireSmokeG       string `json:"tireSmokeG"`
	Horn             string `json:"horn"`
	Armor            string `json:"armor"`
	Turbo            string `json:"turbo"`
	Suspension       string `json:"suspension"`
	Lights           string `json:"lights"`
	PlateText        string `json:"plateText"`
}

var OrderMap map[string]Order = map[string]Order{}
var Lock sync.RWMutex = sync.RWMutex{}

func startOrderHandler(port string) {

	router := http.NewServeMux()

	router.HandleFunc("/order", handleNewOrder)

	log.Println("Order handler listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func handleNewOrder(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))

	if request.Method != "POST" {
		response.WriteHeader(http.StatusNotAcceptable)
		io.WriteString(response, "Only POST requests are accepted. Please see github.com/jcrumb/turbocharger for more info")
		return
	}

	var order Order
	body, _ := ioutil.ReadAll(request.Body)
	log.Println(string(body))

	err := json.Unmarshal(body, &order)
	if err != nil {
		log.Println(err)

		response.WriteHeader(http.StatusInternalServerError)
		io.WriteString(response, "Error creating order: "+err.Error())
	} else {
		OrderMap[order.Username] = order
		log.Println("New order from " + order.Username)
		log.Println(order)
		response.WriteHeader(http.StatusCreated)
		io.WriteString(response, "Order created successfully")
	}
}

func applyOrder(originalOrder rockstarOrder) rockstarOrder {

	Lock.RLock()
	order, customOrder := OrderMap[originalOrder.Owner]
	Lock.RUnlock()

	if customOrder {
		Lock.Lock()
		delete(OrderMap, originalOrder.Owner)
		Lock.Unlock()
		return applyCustomOrder(order, originalOrder)
	} else {
		return applyDefaultOrder(originalOrder)
	}
}

func applyDefaultOrder(order rockstarOrder) rockstarOrder {
	order.Parts.Armor = Armor["Max"]
	order.Parts.BPTires = Tires["Bulletproof"]
	order.Parts.EMSTune = EMSTune["Max"]
	order.Parts.Brakes = Brakes["Max"]
	order.Parts.Turbo = Turbo["Applied"]

	return order
}

func applyCustomOrder(order Order, originalOrder rockstarOrder) rockstarOrder {
	originalOrder.Parts.PrimaryPaint = newIfNotBlank(order.PaintPrimary, originalOrder.Parts.PrimaryPaint, Paints)
	originalOrder.Parts.SecondaryPaint = newIfNotBlank(order.PaintSecondary, originalOrder.Parts.SecondaryPaint, Paints)
	originalOrder.Parts.Turbo = newIfNotBlank(order.Turbo, originalOrder.Parts.Turbo, Turbo)
	originalOrder.Parts.Lights = newIfNotBlank(order.Lights, originalOrder.Parts.Lights, Lights)
	originalOrder.Parts.BPTires = newIfNotBlank(order.BulletproofTires, originalOrder.Parts.BPTires, Tires)
	originalOrder.Parts.Brakes = newIfNotBlank(order.Brakes, originalOrder.Parts.Brakes, Brakes)
	originalOrder.Parts.EMSTune = newIfNotBlank(order.EMSTune, originalOrder.Parts.EMSTune, EMSTune)
	originalOrder.Parts.WindowTint = newIfNotBlank(order.WindowTint, originalOrder.Parts.WindowTint, WindowTint)
	originalOrder.Parts.Armor = newIfNotBlank(order.Armor, originalOrder.Parts.Armor, Armor)
	originalOrder.Parts.Suspension = newIfNotBlank(order.Suspension, originalOrder.Parts.Suspension, Suspension)

	if order.PlateText != "" {
		// Trim the string to 8 characters, the maximum plate length
		originalOrder.Plate.PlateText = string([]byte(order.PlateText)[:7])
	}

	return originalOrder
}

// If newVal is not blank, performs a lookup and returns the new value, otherwise returns the original value
func newIfNotBlank(newVal string, oldVal string, lookup map[string]string) string {
	if newVal != "" {
		return lookup[newVal]
	} else {
		return oldVal
	}
}
