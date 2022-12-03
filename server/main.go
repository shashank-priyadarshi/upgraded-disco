package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type dummyData struct {
	Header struct {
		Name string   `json:"name"`
		Role []string `json:"role"`
	} `json:"header"`
	Footer struct {
		Links []struct {
			Linkedin string `json:"linkedin"`
			Github   string `json:"github"`
		} `json:"links"`
	} `json:"footer"`
	Body struct {
		Summary         []string `json:"summary"`
		CareerObjective string   `json:"career_objective"`
		Skillsets       struct {
			TechnicalSkillGrid []string `json:"technical_skill_grid"`
		} `json:"skillsets"`
		Education struct {
			College   []string `json:"college"`
			Secondary []string `json:"secondary"`
		} `json:"education"`
		EmploymentProfile struct {
			CurrentEmployer  []string      `json:"current_employer"`
			PreviousEmployer []interface{} `json:"previous_employer"`
		} `json:"employment_profile"`
		Projects       [][]string    `json:"projects"`
		Certifications []interface{} `json:"Certifications"`
	} `json:"body"`
	Keys struct {
		Header []string `json:"header"`
		Footer []string `json:"footer"`
		Links  []string `json:"links"`
		Body   struct {
			Skillsets []struct {
				TechnicalSkillGrid []string `json:"technical_skill_grid"`
			} `json:"skillsets"`
			Education []struct {
				College   []string `json:"college,omitempty"`
				Secondary []string `json:"secondary,omitempty"`
			} `json:"education"`
			EmploymentProfile struct {
				Employer []string `json:"employer"`
			} `json:"employment_profile"`
			Projects       []string `json:"projects"`
			Certifications struct {
			} `json:"Certifications"`
		} `json:"body"`
	} `json:"keys"`
	Assets            []interface{} `json:"assets"`
	EmploymentProfile struct {
		CurrentEmployer struct {
			Employer       string `json:"employer"`
			Designation    string `json:"designation"`
			DateOfJoining  string `json:"date_of_joining"`
			OfficeLocation string `json:"office_location"`
			Description    string `json:"description"`
		} `json:"current_employer"`
		PreviousEmployer []interface{} `json:"previous_employer"`
	} `json:"employment_profile"`
}

func handleRequests() {
	//myRouter := mux.NewRouter().StrictSlash(true)
	//mux.NewRouter().StrictSlash(true)
	//http.HandleFunc("/", homePage)
	http.HandleFunc("/biodata", returnBiodata)
	//myRouter.HandleFunc("/newdata", editData).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func readData() dummyData {
	rawData, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println(err)
	}
	var unmarshalledData dummyData
	err = json.Unmarshal(rawData, &unmarshalledData)
	if err != nil {
		fmt.Println(err)
	}
	return unmarshalledData
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllData")
	json.NewEncoder(w).Encode(readData())
}

//func editData(w http.ResponseWriter, r *http.Request) {
//	// get the body of our POST request
//	// return the string response containing the request body
//	reqBody, _ := io.ReadAll(r.Body)
//	fmt.Fprintf(w, "%+v", string(reqBody))
//}

func main() {
	handleRequests()
}
