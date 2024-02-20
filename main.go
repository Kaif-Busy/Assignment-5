package main

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
)

type EmpManager struct {
	EmpID     uint64
	ManagerID uint64
}

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "kaif",
		Password: "kaif",
		Database: "mydb",
		Addr:     "localhost:5432",
	}

	db := pg.Connect(opts)
	if db == nil {
		fmt.Print("failed to connect to the database")
		os.Exit(100)
	} else {

		fmt.Println("Hello, you are connected to the database")
	}
	return db
}

func main() {
	db := Connect()
	defer db.Close()
	var emps []EmpManager
	err := db.Model(&emps).Select()
	if err != nil {
		panic(err)
	}
	m := make(map[uint64]uint64)
	for _, emp := range emps {
		m[emp.EmpID] = emp.ManagerID
	}
	fmt.Println("Enter the Employee and Manager Codes:")
	var a, b uint64
	fmt.Scan(&a, &b)
	if x, exists := m[a]; exists && x == b {
		temp := make(map[uint64]bool)
		cycle := dfsDetectCycle(a, b, m, temp)
		if cycle {
			fmt.Println("Cycle Exists")
		} else {
			fmt.Println("Cycle does not Exists")
		}

	} else {
		fmt.Println("No such employee-manager combination exists")
	}

}

func dfsDetectCycle(EmpId, ManagerId uint64, adj map[uint64]uint64, vis map[uint64]bool) bool {
	if _, exists := vis[EmpId]; exists {
		return true
	} else {
		vis[EmpId] = true
	}
	if value, exists := adj[ManagerId]; exists {
		return dfsDetectCycle(ManagerId, value, adj, vis)
	}
	return false
}
