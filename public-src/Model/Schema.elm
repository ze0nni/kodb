module Model.Schema exposing (..)

type alias  FieldId = String

type alias Field =
    { id: FieldId
    }

type alias Type =
    { fields: List(Field)
    }

type alias Row = 
    {
    }

type alias Table = 
    { name: String
    , tableType: Type
    , rows: List(Row)
    }

type alias Schema =
    { name: String
    , tables: List(Table)
    }