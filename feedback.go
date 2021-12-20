package main

import (
	"bufio"
	"flag"
	"fmt"
	//"github.com/k0kubun/pp"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func VerifyRating(rating string) bool {
	rating = strings.TrimSuffix(rating, "\n")
	ratingInt, _ := strconv.Atoi(rating)
	return 0 < ratingInt && ratingInt < 11
}

func IsYes(input string) bool {
	yesRegex, _ := regexp.Compile("^(y|yes)$")
	return yesRegex.MatchString(strings.ToLower((input)))
}

func IsNo(input string) bool {
	noRegex, _ := regexp.Compile("^(n|no)$")
	return noRegex.MatchString(strings.ToLower((input)))
}

func Prompt() {
	fmt.Print("> ")
}

func Ask(reader *bufio.Reader) string {
	Prompt()
	response, _ := reader.ReadString('\n')
	return strings.TrimSuffix(response, "\n")
}

func GetUnderstandingWording(understanding string) string {
	var understandingWording string
	if understanding == "true" {
		understandingWording = "do"
	} else {
		understandingWording = "do not"
	}
	return understandingWording
}

func AskRating(reader *bufio.Reader, bundle Bundle) string {

	patient := GetPatient(bundle)
	doctor := GetDoctor(bundle)

	fmt.Printf("\nHi %s, on a scale of 1-10, would you recommend Dr %s to a friend or family member? 1 = Would not recommend, 10 = Would strongly recommend\n", patient.Name[0].Given[0], doctor.Name[0].Family)

	var rating string
	done := false
	for done == false {
		rating = Ask(reader)
		if VerifyRating(rating) {
			done = true
		} else {
			fmt.Println("Please, only respond with a number between 1 and 10.")
		}
	}
	return rating
}

func AskUnderstanding(reader *bufio.Reader, bundle Bundle) bool {

	diagnosis := GetDiagnosis(bundle)
	doctor := GetDoctor(bundle)

	fmt.Printf("\nThank you. You were diagnosed with %s. Did Dr %s explain how to manage this diagnosis in a way you could understand? (yes/no)\n", diagnosis.Code.Coding[0].Name, doctor.Name[0].Family)

	var understandingBool bool
	done := false
	for done == false {
		understanding := Ask(reader)

		if IsYes(understanding) {
			done = true
			understandingBool = true
		} else if IsNo(understanding) {
			done = true
			understandingBool = false
		} else {
			fmt.Println("Please, only respond with yes or no.")
		}
	}
	return understandingBool
}

func AskFeeling(reader *bufio.Reader, bundle Bundle) string {
	diagnosis := GetDiagnosis(bundle)
	fmt.Printf("\nWe appreciate the feedback, one last question: how do you feel about being diagnosed with diagnosis %s?\n", diagnosis.Code.Coding[0].Name)
	return Ask(reader)
}

func Greeting(bundle Bundle) {

	patient := GetPatient(bundle)
	doctor := GetDoctor(bundle)
	appointment := GetAppointment(bundle)

	t, _ := time.Parse(time.RFC3339, appointment.Period.Start)
	fmt.Printf("\nHello %s, we are reaching out to you after your appointment with Dr %s on %s at %s. We would greatly appreciate if you could take some time to provide some feedback. We assure you that all of your answers are confidential.\n\n", patient.Name[0].Given[0], doctor.Name[0].Family, t.Format("Monday Jan 06"), t.Format("3:04PM"))
}

func Summary(bundle Bundle, feedback map[string]string) {

	doctor := GetDoctor(bundle)

	fmt.Print("\n\nThanks again! Here’s what we heard:\n\n")
	fmt.Printf("    You said you were %s out of 10 likely to recommend Dr %s\n", feedback["rating"], doctor.Name[0].Family)
	fmt.Printf("    You felt that you %s understand Dr %s's explanation for how to manage this diagnosis.\n", GetUnderstandingWording(feedback["understanding"]), doctor.Name[0].Family)
	fmt.Printf("    Your feeling about this diagnosis: %s\n", feedback["feeling"])
}

func main() {

	inputPath := flag.String("i", "input.txt", "input file")
	flag.Parse()

	bundle := NewBundle(*inputPath)

	reader := bufio.NewReader(os.Stdin)

	Greeting(bundle)
	rating := AskRating(reader, bundle)
	understandingBool := AskUnderstanding(reader, bundle)
	feeling := AskFeeling(reader, bundle)

	feedback := make(map[string]string)
	feedback["rating"] = rating
	feedback["understanding"] = strconv.FormatBool(understandingBool)
	feedback["feeling"] = feeling
	SaveFeedback(feedback)

	Summary(bundle, feedback)

}
