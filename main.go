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

func ConnectToDB() *pg.DB {
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
	db := ConnectToDB()
	defer db.Close()

	var employeeManagers []EmpManager
	err := db.Model(&employeeManagers).Select()
	if err != nil {
		panic(err)
	}

	adj := make(map[uint64][]uint64)
	for _, emp := range employeeManagers {
		adj[emp.EmpID] = append(adj[emp.EmpID],emp.ManagerID)
	}

	fmt.Println("Enter the Employee and Manager Codes:")
	var empID, mgrID uint64
	fmt.Scan(&empID, &mgrID)

	adj[empID] = append(adj[empID],mgrID)

	visited := make(map[uint64]bool)
	isCycleDetected := dfsDetectCycle(empID, mgrID, adj, visited)

	if isCycleDetected {
		fmt.Println("Cycle exists")
	} else {
		fmt.Println("Cycle Does Not exists")
	}

}

func dfsDetectCycle(empId, managerId uint64, adj map[uint64][]uint64, visited map[uint64]bool) bool {
	if _, exists := visited[empId]; exists {
		return true
	} else {
		visited[empId] = true
	}
	for _,manager:=range adj[managerId]{
		if dfsDetectCycle(managerId,manager,adj,visited){
			return true
		}
	}
	return false
}
