package main

import (
	"encoding/json"
	"log"
	"strings"
)

type vehicleParts struct {
	OrderNumber string `json:"0"`
	VehicleID   string `json:"1"`

	// It's unknown what properties 3-6 do
	// so they're listed here as UP3-6
	// I would omit them but they're needed
	// when the struct is reserialized to be sent
	// to rockstar
	UP3 string `json:"3"`
	UP4 string `json:"4"`
	UP5 string `json:"5"`
	UP6 string `json:"6"`

	PrimaryPaint   string `json:"7"`
	SecondaryPaint string `json:"8"`

	// Entry 9 is in vehiclePlate
	WindowTint string `json:"10"`
	BPTires    string `json:"11"`
	EMSTune    string `json:"12"`
	Brakes     string `json:"13"`
	Exhaust    string `json:"14"`
	Wheels     string `json:"15"`

	// Entries 16-24 are also unknown
	UP16 string `json:"16"`
	UP17 string `json:"17"`
	UP18 string `json:"18"`
	UP19 string `json:"19"`
	UP20 string `json:"20"`
	UP21 string `json:"21"`
	UP22 string `json:"22"`
	UP23 string `json:"23"`
	UP24 string `json:"24"`

	TireSmokeChange string `json:"25"`
	TireSmokeR      string `json:"26"`
	TireSmokeG      string `json:"27"`
	TireSmokeB      string `json:"28"`

	Horn       string `json:"29"`
	Armor      string `json:"30"`
	Turbo      string `json:"31"`
	Suspension string `json:"32"`
	Lights     string `json:"33"`

	// Another string of unknown props
	UP34 string `json:"34"`
	UP35 string `json:"35"`
	UP36 string `json:"36"`
	UP37 string `json:"37"`
	UP38 string `json:"38"`
	UP39 string `json:"39"`
	UP40 string `json:"40"`
	UP41 string `json:"41"`
	UP42 string `json:"42"`
	UP43 string `json:"43"`
	UP44 string `json:"44"`
	UP45 string `json:"45"`
	UP46 string `json:"46"`
	UP47 string `json:"47"`
	UP48 string `json:"48"`
	UP49 string `json:"49"`
	UP50 string `json:"50"`
	UP51 string `json:"51"`
	UP52 string `json:"52"`
	UP53 string `json:"53"`
	UP54 string `json:"54"`
	UP67 string `json:"67"`
	UP72 string `json:"72"`

	Character string    `json:"character"`
	MpUnlocks mpUnlocks `json:"mpUnlocks"`
	Vehicle   string    `json:"vehicle"`
}

type mpUnlocks struct {
	UP55 string `json:"55"`
	UP56 string `json:"56"`
}

type vehiclePlate struct {
	PlateType string `json:"9"`
	PlateText string `json:"carPlateText"`
}

type rockstarOrder struct {
	Version int          `json:"version"`
	Parts   vehicleParts `json:"mp0_order"`
	Owner   string       `json:"ownerID"`
	Plate   vehiclePlate `json:"plate"`
}

func rockstarOrderFromString(order string) rockstarOrder {
	// The json originally sent by rockstar is not valid
	// so the json parser will throw an error if it's not corrected
	order = strings.Replace(order, "version", "\"version\"", 1)

	var orderStruct rockstarOrder
	err := json.Unmarshal([]byte(order), &orderStruct)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(orderStruct)
	}

	return orderStruct
}

func rockstarOrderToString(order rockstarOrder) string {
	orderBytes, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		return ""
	}

	orderString := string(orderBytes)
	orderString = strings.Replace(orderString, "\"version\"", "version", 1)

	return orderString
}
