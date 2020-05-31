module Content exposing (Model, Msg, view, update, newSchema)

import Html exposing (Html, button, div, text)
import Html.Events exposing (onClick)
import Html.Attributes exposing (id, class)
import Model.Schema exposing (..)
import Platform.Sub exposing (Sub)

-- TEST


newSchema : Schema
newSchema =
  { name = "Test"
  , tables =
    [ { name = "Users"
      , tableType = usersType
      , rows = []
      }
    , { name = "Inventory"
      , tableType = inventoryType
      , rows = []
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

-- TEST

type alias Model =
  { selectedTable: String
  , schema: Model.Schema.Schema
  }

type Msg = SelectTable String


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    SelectTable name ->
      ( { model | selectedTable = name },  Cmd.none )

view : Model -> (Html Msg)
view model =
  Html.header []
  [ tableTabs model]

tableTabs model=
  Html.section []
  [ Html.header []
    [ Html.nav []
      [ Html.div [class "nav-wrapper"]
        [ Html.ul [id "nav-mobile", class "right"]
          (List.map (tableHeader model.selectedTable) model.schema.tables)
        ]
      ]
    ]
  , Html.article []
    (model.schema.tables |> List.filter (\t -> t.name == model.selectedTable) |> List.map tableView)
  , Html.footer []
    []
  ]

tableHeader selected table =
  if table.name == selected then
    Html.li [class "active"] [Html.a [] [text table.name]]
  else
    Html.li [] [Html.a [Html.Events.onClick (SelectTable table.name) ] [text table.name]]


tableView: Model.Schema.Table -> Html msg
tableView table =
  Html.table []
  [ Html.thead []
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