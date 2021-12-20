package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	//"github.com/k0kubun/pp"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	Status  string `json: status`
	Subject struct {
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
	BirthDate string `json: birthDate`
	Address   []struct {
		Use  string   `json: use`
		Line []string `json: line`
	}
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

func main() {

	inputPath := flag.String("i", "input.txt", "input file")
	flag.Parse()

	var patient Resource
	var doctor Resource
	var appointment Resource
	var diagnosis Resource

	bundle := NewBundle(*inputPath)
	for i, _ := range bundle.Entry {
		if bundle.Entry[i].Resource.ResourceType == "Patient" {
			patient = bundle.Entry[i].Resource
		}
		if bundle.Entry[i].Resource.ResourceType == "Doctor" {
			doctor = bundle.Entry[i].Resource
		}
		if bundle.Entry[i].Resource.ResourceType == "Appointment" {
			appointment = bundle.Entry[i].Resource
		}
		if bundle.Entry[i].Resource.ResourceType == "Diagnosis" {
			diagnosis = bundle.Entry[i].Resource
		}
	}

	reader := bufio.NewReader(os.Stdin)

	t, _ := time.Parse(time.RFC3339, appointment.Period.Start)
	fmt.Printf("\nHello %s, we are reaching out to you after your appointment with Dr %s on %s at %s. We would greatly appreciate if you could take some time to provide some feedback. We assure you that all of your answers are confidential.\n\n", patient.Name[0].Given[0], doctor.Name[0].Family, t.Format("Monday Jan 06"), t.Format("3:04PM"))

	fmt.Printf("\nHi %s, on a scale of 1-10, would you recommend Dr %s to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend\n", patient.Name[0].Given[0], doctor.Name[0].Family)

	rating := 0
	for rating == 0 {
		fmt.Print("> ")
		ratingInput, _ := reader.ReadString('\n')
		ratingInt, _ := strconv.Atoi(strings.TrimSuffix(ratingInput, "\n"))
		if 0 < ratingInt && ratingInt < 11 {
			rating = ratingInt
		} else {
			fmt.Println("Please, only respond with a number between 1 and 10.")
		}
	}

	fmt.Printf("\nThank you. You were diagnosed with %s. Did Dr %s explain how to manage this diagnosis in a way you could understand? (yes/no)\n", diagnosis.Code.Coding[0].Name, doctor.Name[0].Family)

	var understandingWording string
	understanding := ""
	for understanding == "" {
		fmt.Print("> ")
		understanding, _ = reader.ReadString('\n')
		yesRegex, _ := regexp.Compile("(y|yes)")
		noRegex, _ := regexp.Compile("(n|no)")

		if yesRegex.MatchString(strings.ToLower((understanding))) == true {
			understandingWording = "do"
		} else if noRegex.MatchString(strings.ToLower((understanding))) == true {
			understandingWording = "do not"
		} else {
			understanding = ""
			fmt.Println("Please, only respond with yes or no.")
		}
	}

	fmt.Printf("\nWe appreciate the feedback, one last question: how do you feel about being diagnosed with diagnosis %s?\n", diagnosis.Code.Coding[0].Name)

	fmt.Print("> ")
	feeling, _ := reader.ReadString('\n')

	fmt.Println("\n\nThanks again! Here’s what we heard:\n")
	fmt.Printf("    You said you were %d out of 10 likely to recommend Dr %s\n", rating, doctor.Name[0].Family)
	fmt.Printf("    You felt that you %s understand Dr %s's explanation for how to manage this diagnosis.\n", understandingWording, doctor.Name[0].Family)
	fmt.Printf("    Your feeling about this diagnosis: %s\n", feeling)

}
