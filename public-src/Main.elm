port module Main exposing (..)

import Browser
import Html exposing (Html, button, div, text)
import Html.Events exposing (onClick)
import Html.Attributes exposing (id, class)
import Model.Schema exposing (..)
import Platform.Sub exposing (Sub)
import Content
import List

main : Program () Model Msg
main =
  Browser.element
  { init = init
  , update = update
  , view = view
  , subscriptions = subscriptions 
  }

port wsConnect : () -> Cmd msg
port wsDisconnect : () -> Cmd msg
port wsSendMessage : String -> Cmd msg

port wsConnected : (() -> msg) -> Sub msg
port wsDisconnected : (() -> msg) -> Sub msg
port wsMessageReceiver : (String -> msg) -> Sub msg

type Model
  = Disconnected
  | Loading
  | ErrorPage String
  | Content Content.Model

init : () -> (Model, Cmd Msg)
init _ =
  ( Disconnected
  , wsConnect ()
  )

type Msg
  = WSConnected
  | WSDisconnected
  | Request String
  | GotContentMsg Content.Msg

update : Msg -> Model -> (Model, Cmd Msg)
update msg model = case msg of
  WSConnected -> (Loading, wsSendMessage "getSchema")
  WSDisconnected -> (Disconnected, Cmd.none)
  Request _-> (Content { selectedTable ="", schema = Content.newSchema }, Cmd.none)
  GotContentMsg innerMsg -> case model of 
    Content innerModel -> case (Content.update innerMsg innerModel) of
      (innerModel1, cmd) -> (Content innerModel1, Cmd.map GotContentMsg cmd)
    _ -> (model, Cmd.none)



subscriptions : Model -> Sub Msg
subscriptions model = Platform.Sub.batch
  [ wsConnected (\_ -> WSConnected)
  , wsDisconnected (\_ -> WSDisconnected)
  , wsMessageReceiver Request
  ]

view : Model -> Html Msg
view model = case model of 
    Disconnected -> Html.h1 [] [text "Connecting..."]
    Loading -> Html.h1 [] [text "Loading..."]
    ErrorPage msg -> Html.h1 [] [text msg]
    Content m -> Content.view m |> Html.map GotContentMsg 
