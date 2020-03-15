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
	firstname, _ := userLib.NewColumn(engine.NewLiteralColumn("firstName"))
	secondName, _ := userLib.NewColumn(engine.NewLiteralColumn("secondName"))
	age, _ := userLib.NewColumn(engine.NewLiteralColumn("age"))
	for i := 0; i < 10; i++ {
		row, _ := userLib.NewRow()
		userLib.UpdateValue(row, firstname.ID(), randomdata.FirstName(0))
		userLib.UpdateValue(row, secondName.ID(), randomdata.LastName())
		userLib.UpdateValue(row, age.ID(), strconv.Itoa(randomdata.Number(16, 40)))
	}

	invLib, _ := eng.AddLibrary(engine.LibraryName("location"))
	name, _ := invLib.NewColumn(engine.NewLiteralColumn("name"))
	invLib.NewColumn(engine.NewLiteralColumn("type"))
	invLib.NewColumn(engine.NewLiteralColumn("title"))
	invLib.NewColumn(engine.NewLiteralColumn("picture"))
	owner, _ := invLib.NewColumn(engine.NewRefColumn("owner", userLib.Name()))
	for i := 0; i < 20; i++ {
		row, _ := invLib.NewRow()
		invLib.UpdateValue(row, name.ID(), randomdata.Adjective())
		invLib.UpdateValue(row, owner.ID(), randomdata.Alphanumeric(32))
	}

	rewardsLib, _ := eng.AddLibrary(engine.LibraryName("rewards"))
	rewardsLib.AddColumn(engine.ColumnID("parent"), engine.NewLiteralColumn("parent").SetHidden(true))
	rwTitle, _ := rewardsLib.NewColumn(engine.NewLiteralColumn("title"))
	rwType, _ := rewardsLib.NewColumn(engine.NewLiteralColumn("type"))
	rwPrice, _ := rewardsLib.NewColumn(engine.NewLiteralColumn("price"))

	tasksLib, _ := eng.AddLibrary(engine.LibraryName("tasks"))
	tasksLib.AddColumn(engine.ColumnID("parent"), engine.NewLiteralColumn("parent").SetHidden(true))
	tasksLib.NewColumn(engine.NewListColumn("rewards", rewardsLib.Name()))

	questLib, _ := eng.AddLibrary("quest")
	questName, _ := questLib.NewColumn(engine.NewLiteralColumn("name"))
	questLib.NewColumn(engine.NewListColumn("tasks", engine.LibraryName("tasks")))

	for i := 1; i <= 3; i++ {
		questRow, _ := questLib.NewRow()
		questLib.UpdateValue(questRow, questName.ID(), "quest_00"+strconv.Itoa(i))

		for j := 1; j <= 1+i; j++ {
			taskRow, _ := tasksLib.NewRow()
			tasksLib.UpdateValue(
				taskRow,
				engine.ColumnID("parent"),
				questRow.ToString(),
			)

			for k := 1; k < 4; k++ {
				rewardRow, _ := rewardsLib.NewRow()
				rewardsLib.UpdateValue(
					rewardRow,
					engine.ColumnID("parent"),
					taskRow.ToString(),
				)

				rewardsLib.UpdateValue(
					rewardRow,
					rwTitle.ID(),
					randomdata.Email(),
				)

				rewardsLib.UpdateValue(
					rewardRow,
					rwType.ID(),
					randomdata.Adjective(),
				)

				rewardsLib.UpdateValue(
					rewardRow,
					rwPrice.ID(),
					randomdata.StringNumber(2, "."),
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
