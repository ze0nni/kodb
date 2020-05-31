import Browser
import Html exposing (Html, button, div, text)
import Html.Events exposing (onClick)
import Model.Schema exposing (..)
import List

main =
  Browser.sandbox { init = newModel, update = update, view = view }

newModel =
  { selectedTable = "Users"
  , schema = newSchema
  }

newSchema : Schema
newSchema =
  { name = "Test"
  , tables =
    [ { name = "Users"
      , tableType = usersType
      }
    , { name = "Inventory"
      , tableType = inventoryType
      }
    ]
  }

usersType : Model.Schema.Type
usersType =
  { fields = 
    [ { id = "Id" }
    , { id = "First name" }
    , { id = "Second name" }
    , { id = "Age" }
    ]
  }

inventoryType : Model.Schema.Type
inventoryType =
  { fields = 
    [ { id = "Id" }
    , { id = "Title" }
    , { id = "Price" }
    , { id = "Kind" }
    ]
  }

type Msg = SelectTable String

type alias Model =
  { name: String
  }

update msg model =
  case msg of 
    SelectTable name ->
      { model | selectedTable = name }

view model =
  Html.header []
  [ tableTabs model]

tableTabs model=
  Html.section []
  [ Html.header []
    [ Html.nav []
      [ Html.ul []
        (List.map (tableHeader model.selectedTable) model.schema.tables)
      ]
    ]
  , Html.article []
    (model.schema.tables |> List.filter (\t -> t.name == model.selectedTable) |> List.map tableView)
  , Html.footer []
    []
  ]

tableHeader selected table =
  if table.name == selected then
    Html.li [] [text table.name]
  else
    Html.li [Html.Events.onClick (SelectTable table.name) ] [text table.name]


tableView: Model.Schema.Table -> Html msg
tableView table =
  Html.table []
  [ Html.caption []
    [text table.name]
  , Html.thead []
    [ Html.tr []
      (List.map headerView table.tableType.fields)
    ]
  , Html.tbody []
    []
  , Html.tfoot []
    []
  ]

headerView: Model.Schema.Field -> Html msg
headerView field =
  Html.th [] [text field.id]