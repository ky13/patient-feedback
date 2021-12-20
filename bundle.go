package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	//"github.com/k0kubun/pp"
	"io/ioutil"
)

type Bundle struct {
	ResourceType string `json: resourceType`
	Id           string `json: id`
	Timestamp    string `json: timestamp`
	Entry        []struct {
		Resource Resource `json: entry`
	}
}

type Resource struct {
	ResourceType string `json: resourceType`
	Id           string `json: id`

	Active bool `json: bool`
	Actor  struct {
		Reference string `json: reference`
	}
	Address []struct {
		Use  string   `json: use`
		Line []string `json: line`
	}
	Appointment struct {
		Reference string `json: reference`
	}
	BirthDate string `json: birthDate`
	Code      struct {
		Coding []struct {
			System string `json: system`
			Code   string `json: code`
			Name   string `json: name`
		}
	}
	Contact []struct {
		System string `json: system`
		Value  string `json: value`
		Use    string `json: use`
	}
	Gender string `json: gender`
	Meta   struct {
		LastUpdated string `json: lastUpdated`
	}
	Name []struct {
		Text   string   `json: text`
		Family string   `json: family`
		Given  []string `json: given`
	}
	Period struct {
		Start string `json: start`
		End   string `json: end`
	}
	Response map[string]string
	Status   string `json: status`
	Subject  struct {
		Reference string `json: reference`
	}
	Type []struct {
		Text string `json: text`
	}
}

type Patient struct {
	ResourceType string
	Id           string
	Active       bool
	Name         []struct {
		Text   string
		Family string
		Given  []string
	}
	Contact []struct {
		System string
		Value  string
		Use    string
	}
	Gender    string
	BirthDate string
	Address   []struct {
		Use  string
		Line []string
	}
}

type Feedback struct {
	ResourceType string
	Id           string
	Response     map[string]string
}

func NewBundle(path string) Bundle {
	var bundle Bundle

	fmt.Println("Reading bundle...")
	source, ferr := ioutil.ReadFile(path)
	if ferr != nil {
		panic(ferr)
	}
	jerr := json.Unmarshal(source, &bundle)
	if jerr != nil {
		panic(jerr)
	}

	return bundle
}

func GetPatient(bundle Bundle) Resource {
	var resource Resource
	for i, _ := range bundle.Entry {
		if bundle.Entry[i].Resource.ResourceType == "Patient" {
			resource = bundle.Entry[i].Resource
		}
	}
	return resource
}

func GetDoctor(bundle Bundle) Resource {
	var resource Resource
	for i, _ := range bundle.Entry {
		if bundle.Entry[i].Resource.ResourceType == "Doctor" {
			resource = bundle.Entry[i].Resource
		}
	}
	return resource
}

func GetDiagnosis(bundle Bundle) Resource {
	var resource Resource
	for i, _ := range bundle.Entry {
		if bundle.Entry[i].Resource.ResourceType == "Diagnosis" {
			resource = bundle.Entry[i].Resource
		}
	}
	return resource
}

func GetAppointment(bundle Bundle) Resource {
	var resource Resource
	for i, _ := range bundle.Entry {
		if bundle.Entry[i].Resource.ResourceType == "Appointment" {
			resource = bundle.Entry[i].Resource
		}
	}
	return resource
}

func AddFeedback(bundle Bundle, f map[string]string) {
	var feedback Resource
	feedback.ResourceType = "Feedback"
	feedback.Id = uuid.NewString()
	feedback.Response = make(map[string]string)
	for key, value := range f {
		feedback.Response[key] = value
	}
	// TODO: Figure out how to append onto the Entry slice
	//bundle.Entry = append(bundle.Entry, feedback)
	//bundle.Entry[99] = feedback
	//pp.Println(bundle)
}

func SaveFeedback(f map[string]string) {
	var feedback Feedback
	feedback.ResourceType = "Feedback"
	feedback.Id = uuid.NewString()
	feedback.Response = make(map[string]string)
	for key, value := range f {
		feedback.Response[key] = value
	}
	jsonString, _ := json.Marshal(feedback)
	path := fmt.Sprintf("feedback-%s.json", feedback.Id)
	ioutil.WriteFile(path, jsonString, 0600)
}
