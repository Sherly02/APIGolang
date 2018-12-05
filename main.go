package main

import "fmt"
import "net/http"
import "encoding/json"	
import "strconv"

type lagu struct {
	Id int
	Judul string
	Genre string
	Pencipta string
}

var datalagu = []lagu {
	lagu{1,"RocketMan1","RockNRoll1","Tanyouka1"},
	lagu{2,"RocketMan2","RockNRoll2","Tanyouka2"},
	lagu{3,"RocketMan3","RockNRoll3","Tanyouka3"},
	lagu{4,"RocketMan4","RockNRoll4","Tanyouka4"},
	lagu{5,"RocketMan5","RockNRoll5","Tanyouka5"},
	lagu{6,"RocketMan6","RockNRoll6","Tanyouka6"},
	lagu{7,"RocketMan7","RockNRoll7","Tanyouka7"},
	lagu{8,"RocketMan8","RockNRoll8","Tanyouka8"},
	lagu{9,"RocketMan9","RockNRoll9","Tanyouka9"},
	lagu{10,"RocketMan10","RockNRoll10","Tanyouka10"},
	lagu{11,"RocketMan11","RockNRoll11","Tanyouka11"},
	lagu{12,"RocketMan12","RockNRoll12","Tanyouka12"},
	lagu{13,"RocketMan13","RockNRoll13","Tanyouka13"},
	lagu{14,"RocketMan14","RockNRoll14","Tanyouka14"},
	lagu{15,"RocketMan15","RockNRoll15","Tanyouka15"},
	lagu{16,"RocketMan16","RockNRoll16","Tanyouka16"},
	lagu{17,"RocketMan17","RockNRoll17","Tanyouka17"},
	lagu{18,"RocketMan18","RockNRoll18","Tanyouka18"},
	lagu{19,"RocketMan19","RockNRoll19","Tanyouka19"},
	lagu{20,"RocketMan20","RockNRoll20","Tanyouka20"},
	lagu{21,"RocketMan21","RockNRoll21","Tanyouka21"},
	lagu{22,"RocketMan22","RockNRoll22","Tanyouka22"},
	lagu{23,"RocketMan23","RockNRoll23","Tanyouka23"},
	lagu{24,"RocketMan24","RockNRoll24","Tanyouka24"},
	lagu{25,"RocketMan25","RockNRoll25","Tanyouka25"},
	lagu{26,"RocketMan26","RockNRoll26","Tanyouka26"},
	lagu{27,"RocketMan27","RockNRoll27","Tanyouka27"},
	lagu{28,"RocketMan28","RockNRoll28","Tanyouka28"},
	lagu{29,"RocketMan29","RockNRoll29","Tanyouka29"},
	lagu{30,"RocketMan30","RockNRoll30","Tanyouka30"},
	lagu{31,"RocketMan31","RockNRoll31","Tanyouka31"},
	lagu{32,"RocketMan32","RockNRoll32","Tanyouka32"},
	lagu{33,"RocketMan33","RockNRoll33","Tanyouka33"},
	lagu{34,"RocketMan34","RockNRoll34","Tanyouka34"},
	lagu{35,"RocketMan35","RockNRoll35","Tanyouka35"},
	lagu{36,"RocketMan36","RockNRoll36","Tanyouka36"},
	lagu{37,"RocketMan37","RockNRoll37","Tanyouka37"},
	lagu{38,"RocketMan38","RockNRoll38","Tanyouka38"},
	lagu{39,"RocketMan39","RockNRoll39","Tanyouka39"},
	lagu{40,"RocketMan40","RockNRoll40","Tanyouka40"},
	lagu{41,"RocketMan41","RockNRoll41","Tanyouka41"},
	lagu{42,"RocketMan42","RockNRoll42","Tanyouka42"},
	lagu{43,"RocketMan43","RockNRoll43","Tanyouka43"},
	lagu{44,"RocketMan44","RockNRoll44","Tanyouka44"},
	lagu{45,"RocketMan45","RockNRoll45","Tanyouka45"},
	lagu{46,"RocketMan46","RockNRoll46","Tanyouka46"},
	lagu{47,"RocketMan47","RockNRoll47","Tanyouka47"},
	lagu{48,"RocketMan48","RockNRoll48","Tanyouka48"},
	lagu{49,"RocketMan49","RockNRoll49","Tanyouka49"},
	lagu{50,"RocketMan50","RockNRoll50","Tanyouka50"},
}

func getlagus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {

		var page,err1 = strconv.Atoi(r.URL.Query()["page"][0])
		var page_size,err2 = strconv.Atoi(r.URL.Query()["page_size"][0])

		if(err1 != nil || err2 != nil){
			http.Error(w, "Internal Serive Error", http.StatusInternalServerError)
			return
		}
		var datalagus = datalagu[(page - 1) * page_size : page * page_size]
		var result, err = json.Marshal(datalagus)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	fmt.Println("starting web server at http://localhost:8080/")
	http.HandleFunc("/lagus", getlagus)
	http.ListenAndServe(":8080", nil)
}