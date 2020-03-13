package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/validate"
	"github.com/ze0nni/kodb/run/web"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	eng := engine.New(driver.InMemory())
	userLib, _ := eng.AddLibrary(engine.LibraryName("user"))
	firstname, _ := userLib.NewColumn("firstName")
	secondName, _ := userLib.NewColumn("secondName")
	age, _ := userLib.NewColumn("age")
	for i := 0; i < 10; i++ {
		row, _ := userLib.NewRow()
		userLib.UpdateValue(row, firstname, randomdata.FirstName(0))
		userLib.UpdateValue(row, secondName, randomdata.LastName())
		userLib.UpdateValue(row, age, strconv.Itoa(randomdata.Number(16, 40)))
	}

	invLib, _ := eng.AddLibrary(engine.LibraryName("location"))
	name, _ := invLib.NewColumn("name")
	invLib.NewColumn("type")
	invLib.NewColumn("title")
	invLib.NewColumn("picture")
	owner, _ := invLib.NewRefColumn("owner", userLib.Name())
	for i := 0; i < 20; i++ {
		row, _ := invLib.NewRow()
		invLib.UpdateValue(row, name, randomdata.Adjective())
		invLib.UpdateValue(row, owner, randomdata.Alphanumeric(32))
	}

	questLib, _ := eng.AddLibrary("quest")
	questName, _ := questLib.NewColumn("name")
	questLib.NewListColumn(eng, "tasks", engine.LibraryName("tasks"))

	tasksLib, _ := eng.Library(engine.LibraryName("tasks"))
	tasksLib.NewListColumn(eng, "rewards", engine.LibraryName("rewards"))

	rewardsLib, _ := eng.Library(engine.LibraryName("rewards"))

	for i := 1; i <= 3; i++ {
		questRow, _ := questLib.NewRow()
		questLib.UpdateValue(questRow, questName, "quest_00"+strconv.Itoa(i))

		for j := 1; j <= 1+i; j++ {
			taskRow, _ := tasksLib.NewRow()
			tasksLib.UpdateValue(
				taskRow,
				engine.ColumnID("parent"),
				questRow.ToString(),
			)

			for k := 1; k < i*j; k++ {
				rewardRow, _ := rewardsLib.NewRow()
				rewardsLib.UpdateValue(
					rewardRow,
					engine.ColumnID("parent"),
					taskRow.ToString(),
				)
			}
		}
	}

	validate.Validate(eng, func(
		l engine.LibraryName, r engine.RowID, c engine.ColumnID, err error,
	) {
		fmt.Printf("Error in %s:%s%s: %s\n", l, r, c, err)
	})

	err := web.Run(eng)
	if nil != err {
		log.Fatal(err)
	}
}
