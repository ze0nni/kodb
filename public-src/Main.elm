port module Main exposing (..)

import Browser
import Html exposing (Html, button, div, text)
import Html.Events exposing (onClick)
import Html.Attributes exposing (id, class)
import Model.Schema exposing (..)
import Platform.Sub exposing (Sub)
import List

main : Program () Model Msg
main =
  Browser.element
  { init = init
  , update = update
  , view = view
  , subscriptions = subscriptions 
  }

port sendMessage : String -> Cmd msg
port messageReceiver : (String -> msg) -> Sub msg

type Model
  = Loading
  | ErrorPage String
  | Connected ContentModel


type alias ContentModel =
  { selectedTable: String
  , schema: Model.Schema.Schema
  }

init : () -> (Model, Cmd Msg)
init _ =
  ( Loading
  , sendMessage ""
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
  | Request String
  | GotContentMsg ContentMsg

type ContentMsg
  = SelectTable String


update : Msg -> Model -> (Model, Cmd Msg)
update msg model = case msg of
  Init -> (model, Cmd.none)
  Request _-> (model, Cmd.none)
  GotContentMsg innerMsg -> case model of 
    Connected innerModel -> case (updateContent innerMsg innerModel) of
      (innerModel1, cmd) -> (Connected innerModel1, Cmd.map GotContentMsg cmd)
    _ -> (model, Cmd.none)

updateContent : ContentMsg -> ContentModel -> (ContentModel, Cmd ContentMsg)
updateContent msg model =
  case msg of
    SelectTable name ->
      ( { model | selectedTable = name },  Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions model = messageReceiver Request

view : Model -> Html Msg
view model = case model of 
    Loading -> Html.h1 [] [text "Loading..."]
    ErrorPage msg -> Html.h1 [] [text msg]
    Connected m -> contentView m |> Html.map GotContentMsg 


contentView : ContentModel -> (Html ContentMsg)
contentView model =
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