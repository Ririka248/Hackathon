package models

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

type FoodPFC struct {
	Category     int
	ID           int
	Name         string
	Protein      string
	Fat          string
	Carbohydrate string
}

type Sports struct {
	ID   int
	Mets float64
	Name string
}

func ReadFoodfiles() {
	files1, _ := os.ReadDir("./data/food")
	var filePathList1 []string
	for _, file := range files1 {
		if filepath.Ext(file.Name()) == ".csv" {
			filePathList1 = append(filePathList1, "./data/food/"+file.Name())
		}
	}

	for _, v := range filePathList1 {
		file, err := os.Open(v)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		r := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

		for {
			records, err := r.Read()
			//fmt.Println(records)

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			var catcast int
			catcast, err = strconv.Atoi(records[0])
			if err != nil {
				log.Println(err)
			}
			var idcast int
			idcast, err = strconv.Atoi(records[2])
			if err != nil {
				log.Println(err)
			}
			foodPFC := FoodPFC{
				Category:     catcast,
				ID:           idcast,
				Name:         records[3],
				Protein:      records[8],
				Fat:          records[10],
				Carbohydrate: records[16]}

			cmd := `insert into foods (
				category, 
				id, 
				name, 
				protein, 
				fat, 
				carbohydrate) values (?, ?, ?, ?, ?, ?)`

			_, err = Db.Exec(cmd,
				foodPFC.Category,
				foodPFC.ID,
				foodPFC.Name,
				foodPFC.Protein,
				foodPFC.Fat,
				foodPFC.Carbohydrate)
			if err != nil {
				log.Println(err)
			}
		}
	}

	files2, _ := os.ReadDir("./data/sports")
	var filePathList2 []string
	for _, file := range files2 {
		if filepath.Ext(file.Name()) == ".csv" {
			filePathList2 = append(filePathList2, "./data/sports/"+file.Name())
		}
	}
	for _, v := range filePathList2 {
		file, err := os.Open(v)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		r := csv.NewReader(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

		for {
			records, err := r.Read()
			//fmt.Println(records)

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println(err)
			}

			var metscast float64
			metscast, err = strconv.ParseFloat(records[0], 64)
			if err != nil {
				log.Println(err)
			}

			sports := Sports{
				Mets: metscast,
				Name: records[1]}

			cmd := `insert into sports (
					mets, 
					sports_name) values (?, ?)`

			_, err = Db.Exec(cmd,
				sports.Mets,
				sports.Name)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func GetFood() (foodPFC FoodPFC, err error) {
	foodPFC = FoodPFC{}
	cmd := `select category, id, name, protein,	fat, carbohydrate
	from foods`
	err = Db.QueryRow(cmd).Scan(
		&foodPFC.Category,
		&foodPFC.ID,
		&foodPFC.Name,
		&foodPFC.Protein,
		&foodPFC.Fat,
		&foodPFC.Carbohydrate)
	return foodPFC, err
}

func GetSports() (sports []Sports, err error) {
	cmd := `select id, mets, sports_name from sports`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var sport Sports
		err = rows.Scan(
			&sport.ID,
			&sport.Mets,
			&sport.Name)
		if err != nil {
			log.Fatalln(err)
		}
		sports = append(sports, sport)
	}
	rows.Close()

	return sports, err
}
