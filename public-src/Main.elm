import Browser
import Html exposing (Html, button, div, text)
import Html.Events exposing (onClick)
import Html.Attributes exposing (id, class)
import Model.Schema exposing (..)
import Platform.Sub exposing (Sub)
import List

main =
  Browser.element
  { init = init
  , update = update
  , view = view
  , subscriptions = subscriptions 
  }

type alias Model =
  { selectedTable: String
  , schema: Model.Schema.Schema
  }
init : () -> (Model, Cmd Msg)
init _ =
  ( { selectedTable = "Users"
    , schema = newSchema
    }
  , Cmd.none
  )

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

type Msg
  = Init
  | Recieve
  | SelectTable String

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of 
    Init -> (model, Cmd.none)
    Recieve -> (model, Cmd.none)
    SelectTable name ->
      ( { model | selectedTable = name },  Cmd.none )

subscriptions : Model -> Sub Msg
subscriptions model = Sub.none

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