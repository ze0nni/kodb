package main

import (
	"fmt"
	"log"

	"github.com/ze0nni/kodb/internal/driver"
	"github.com/ze0nni/kodb/internal/engine"
	"github.com/ze0nni/kodb/internal/types"
	"github.com/ze0nni/kodb/internal/validate"
	"github.com/ze0nni/kodb/run/web"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	eng := engine.New(driver.InMemory())

	userType, err := eng.Types().New(types.TypeName("user"))
	logError(err)
	user_picture, err := userType.New(types.NewValueFieldData("picture"))
	logError(err)
	user_firstName, err := userType.New(types.NewValueFieldData("firstName"))
	logError(err)
	user_secondName, err := userType.New(types.NewValueFieldData("secondName"))
	logError(err)
	user_ht, err := userType.New(types.NewValueFieldData("ht"))
	logError(err)
	user_dx, err := userType.New(types.NewValueFieldData("dx"))
	logError(err)
	user_iq, err := userType.New(types.NewValueFieldData("iq"))
	logError(err)
	user_age, err := userType.New(types.NewValueFieldData("age"))
	logError(err)

	userLib, _ := eng.AddLibrary(engine.LibraryName("users"), userType.Name())

	for i := 0; i < 10; i++ {
		row, _ := userLib.NewRow()
		logError(userLib.UpdateValue(
			row,
			user_picture.ID(),
			randomdata.PhoneNumber(),
		))
		logError(userLib.UpdateValue(
			row,
			user_firstName.ID(),
			randomdata.FirstName(0),
		))
		logError(userLib.UpdateValue(
			row,
			user_secondName.ID(),
			randomdata.LastName(),
		))
		logError(userLib.UpdateValue(
			row,
			user_ht.ID(),
			randomdata.StringNumber(2, "-"),
		))
		logError(userLib.UpdateValue(
			row,
			user_dx.ID(),
			randomdata.StringNumber(2, "-"),
		))
		logError(userLib.UpdateValue(
			row,
			user_iq.ID(),
			randomdata.StringNumber(2, "-"),
		))
		logError(userLib.UpdateValue(
			row,
			user_age.ID(),
			randomdata.StringNumber(2, "-"),
		))
	}

	mathOp, _ := eng.Types().New(types.TypeName("mathOp"))
	mathOp.UpdateCases([]types.FieldCase{
		types.FieldCase("Const"),
		types.FieldCase("Sum"),
		types.FieldCase("Mult"),
	})

	constValue, _ := mathOp.New(types.NewValueFieldData("value"))
	constValue.SetCase(types.FieldCase("Const"))

	opSumLeft, _ := mathOp.New(types.NewValueFieldData("left"))
	opSumRight, _ := mathOp.New(types.NewValueFieldData("right"))
	opSumLeft.SetCase(types.FieldCase("Sum"))
	opSumRight.SetCase(types.FieldCase("Sum"))

	opMultLeft, _ := mathOp.New(types.NewValueFieldData("left"))
	opMultRight, _ := mathOp.New(types.NewValueFieldData("right"))
	opMultLeft.SetCase(types.FieldCase("Mult"))
	opMultRight.SetCase(types.FieldCase("Mult"))

	operationsLib, _ := eng.AddLibrary(engine.LibraryName("operations"), mathOp.Name())
	operationsLib.NewRow()
	operationsLib.NewRow()
	operationsLib.NewRow()

	validate.Validate(eng, func(
		l engine.LibraryName, r engine.RowID, f engine.FieldID, err error,
	) {
		fmt.Printf("Error in %s:%s%s: %s\n", l, r, f, err)
	})

	err = web.Run(eng)
	if nil != err {
		log.Fatal(err)
	}
}

func logError(err error) {
	if nil != err {
		log.Panic(err)
	}
}
