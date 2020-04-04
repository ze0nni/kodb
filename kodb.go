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

	userType, _ := eng.Types().New(types.TypeName("user"))
	user_picture, _ := userType.New(types.NewValueFieldData("picture"))
	user_firstName, _ := userType.New(types.NewValueFieldData("firstName"))
	user_secondName, _ := userType.New(types.NewValueFieldData("secondName"))
	user_ht, _ := userType.New(types.NewValueFieldData("ht"))
	user_dx, _ := userType.New(types.NewValueFieldData("dx"))
	user_iq, _ := userType.New(types.NewValueFieldData("iq"))
	user_age, _ := userType.New(types.NewValueFieldData("age"))

	userLib, _ := eng.AddLibrary(engine.LibraryName("users"), userType.Name())

	for i := 0; i < 10; i++ {
		row, _ := userLib.NewRow()
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_picture.ID().String()),
			randomdata.PhoneNumber(),
		)
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_firstName.ID().String()),
			randomdata.FirstName(0),
		)
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_secondName.ID().String()),
			randomdata.LastName(),
		)
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_ht.ID().String()),
			randomdata.StringNumber(2, "-"),
		)
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_dx.ID().String()),
			randomdata.StringNumber(2, "-"),
		)
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_iq.ID().String()),
			randomdata.StringNumber(2, "-"),
		)
		userLib.UpdateValue(
			row,
			engine.ColumnID(user_age.ID().String()),
			randomdata.StringNumber(2, "-"),
		)
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
		l engine.LibraryName, r engine.RowID, c engine.ColumnID, err error,
	) {
		fmt.Printf("Error in %s:%s%s: %s\n", l, r, c, err)
	})

	err := web.Run(eng)
	if nil != err {
		log.Fatal(err)
	}
}
