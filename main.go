package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	var students []student

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Incorrect file path")
		return []student{}
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return []student{}
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			break
		}

		stud := student{
			firstName:  record[0],
			lastName:   record[1],
			university: record[2],
			test1Score: studentScores(record[3]),
			test2Score: studentScores(record[4]),
			test3Score: studentScores(record[5]),
			test4Score: studentScores(record[6]),
		}
		students = append(students, stud)
	}
	return students
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}

func studentScores(record string) int {
	score, err := strconv.Atoi(record)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return score
}

func main() {
	filePath := "grades.csv"
	fmt.Print(parseCSV(filePath))

}
