module Model.Schema exposing (..)

type alias  FieldId = String

type alias Field =
    { id: FieldId
    }

type alias Type =
    { fields: List(Field)
    }

type alias Table = 
    { name: String
    , tableType: Type
    }

type alias Schema =
    { name: String
    , tables: List(Table)
    }