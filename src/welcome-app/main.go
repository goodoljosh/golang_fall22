package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}
//Note:Work on class videos commented out for focus on assignment
// type JsonResponse struct{
// 	Value1 string `json:"key1"`
// 	Value2 string `json:"key2"`
// 	JsonNested JsonNested `json:"JsonNested`
// }


// type JsonNested struct {
// 	NestedValue1 string `json:"nestedkey1"`
// 	NestedValue2 string `json:"nestedkey2"`
// }

type JsonContact struct{
    Value3 string `json:"key3"`
	Value4 string `json:"key4"`
	Value5 string `json:"key5"`
	JsonNestedAdd JsonNestedAdd `json:"JsonNestedAdd`
	Value6 string `json:"key6"`
	JsonNestedConInfo JsonNestedConInfo `json: "JsonNestedConInfo`
}

type JsonNestedAdd struct {
	NestedValue3 string `json:"nestedkey3"`
	NestedValue4 string `json:"nestedkey4"`
}


type JsonNestedConInfo struct{
	NestedValue5 string `json:"nestedkey5"`
	NestedValue6 string `json:"nestedkey6"`
}

// Go application entrypoint
func main() {
	//Instantiate a Welcome struct object and pass in some random information.
	//We shall get the name of the user as a query parameter from the URL
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	//We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	// the relative path). We wrap it in a call to template.Must() which handles any errors and halts if there are fatal errors

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
    //Note:Work on class videos commented out for focus on assignment
	// nested := JsonNested{
	// 	NestedValue1: "first nested value",
	// 	NestedValue2: "second nested value",
	// }

	// jsonResp := JsonResponse{
	// 	Value1: "some Data",
	// 	Value2: "other Data",
	// 	JsonNested: nested,
	// }

	nestedAdd := JsonNestedAdd{
		NestedValue3: "245 Faux Avenue",
		NestedValue4: "Blueberrysburg, Michigan",
	}

	NestedConInfo := JsonNestedConInfo{
		NestedValue5: "Email: eatatjoes@barnsandnoble.com",
		NestedValue6: "Phone: 309-111-2321",
	}

	jsonCon := JsonContact{
		Value3: "FirstName: John ",
		Value4: "LastName: Doe ",
		Value5: "Address:",
		JsonNestedAdd: nestedAdd,
		Value6: "Contact Information:",
		JsonNestedConInfo: NestedConInfo,

	}

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative "static" directory first using http.FileServer(), then matches it to a

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		//If errors show an internal server error message
		//I also pass the welcome struct to the welcome-template.html file.
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
    
	//Start the web server, set the port to listen to 8080. Without a path it assumes localhost
	//Print any errors from starting the webserver using fmt
	http.HandleFunc("/jsonContact", func(w http.ResponseWriter, r *http.Request){
		//json.NewEncoder(w).Encode(jsonResp)
		json.NewEncoder(w).Encode(jsonCon)
	})

    // third path, get/fetch, return an json object like an API with 2 nested objects {firstname:"", lastname:"", address:(street"", city...), contactInfo:{email:"", phone:""}}

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
