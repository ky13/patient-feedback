package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {

	inputPath := flag.String("i", "input.txt", "input file")
	flag.Parse()

	bundle := NewBundle(*inputPath)
	patient := GetPatient(bundle)
	doctor := GetDoctor(bundle)
	appointment := GetAppointment(bundle)
	diagnosis := GetDiagnosis(bundle)

	reader := bufio.NewReader(os.Stdin)

	t, _ := time.Parse(time.RFC3339, appointment.Period.Start)
	fmt.Printf("\nHello %s, we are reaching out to you after your appointment with Dr %s on %s at %s. We would greatly appreciate if you could take some time to provide some feedback. We assure you that all of your answers are confidential.\n\n", patient.Name[0].Given[0], doctor.Name[0].Family, t.Format("Monday Jan 06"), t.Format("3:04PM"))

	fmt.Printf("\nHi %s, on a scale of 1-10, would you recommend Dr %s to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend\n", patient.Name[0].Given[0], doctor.Name[0].Family)

	rating := ""
	for rating == "" {
		fmt.Print("> ")
		ratingInput, _ := reader.ReadString('\n')
		ratingInput = strings.TrimSuffix(ratingInput, "\n")
		ratingInt, _ := strconv.Atoi(ratingInput)
		if 0 < ratingInt && ratingInt < 11 {
			rating = ratingInput
		} else {
			fmt.Println("Please, only respond with a number between 1 and 10.")
		}
	}

	fmt.Printf("\nThank you. You were diagnosed with %s. Did Dr %s explain how to manage this diagnosis in a way you could understand? (yes/no)\n", diagnosis.Code.Coding[0].Name, doctor.Name[0].Family)

	var understandingBool bool
	var understandingWording string
	understanding := ""
	for understanding == "" {
		fmt.Print("> ")
		understanding, _ = reader.ReadString('\n')
		yesRegex, _ := regexp.Compile("(y|yes)")
		noRegex, _ := regexp.Compile("(n|no)")

		if yesRegex.MatchString(strings.ToLower((understanding))) == true {
			understandingBool = true
			understandingWording = "do"
		} else if noRegex.MatchString(strings.ToLower((understanding))) == true {
			understandingBool = false
			understandingWording = "do not"
		} else {
			understanding = ""
			fmt.Println("Please, only respond with yes or no.")
		}
	}

	fmt.Printf("\nWe appreciate the feedback, one last question: how do you feel about being diagnosed with diagnosis %s?\n", diagnosis.Code.Coding[0].Name)

	fmt.Print("> ")
	feeling, _ := reader.ReadString('\n')
	feeling = strings.TrimSuffix(feeling, "\n")

	fmt.Println("\n\nThanks again! Here’s what we heard:\n")
	fmt.Printf("    You said you were %s out of 10 likely to recommend Dr %s\n", rating, doctor.Name[0].Family)
	fmt.Printf("    You felt that you %s understand Dr %s's explanation for how to manage this diagnosis.\n", understandingWording, doctor.Name[0].Family)
	fmt.Printf("    Your feeling about this diagnosis: %s\n", feeling)

	feedback := make(map[string]string)
	feedback["rating"] = rating
	feedback["understanding"] = strconv.FormatBool(understandingBool)
	feedback["feeling"] = string(feeling)
	SaveFeedback(feedback)
}
