Vue.component("kodb-library", {
        props: [
                "schema",
                "libraryName",
        ],
        data() {
                return {
                        
                }
        },
        computed: {
                type() {
                        const table = this.schema.map[this.libraryName]
                        if (null == table) {
                                return null
                        }
                        return this.schema.types[table.type]
                },
                cases() {
                        return this.type.cases
                },
                multiCase() {
                        return  1 < this.cases.length
                },
                fields() {
                        return Object.values(this.type.fields).filter(f => f['case'] == "")
                },
                originRows() {
                        return this.schema.rowsMap[this.libraryName]
                },
                rows() {
                        return this.originRows
                                .map(r => [
                                        {
                                                expanded: false,
                                                fields: this.fields,
                                                row: r,
                                        },
                                        {
                                                expanded: true,
                                                fields: [],
                                                row: r,
                                        }
                                ])
                                .reduce((a,b) => a.concat(b), [])
                },
                selectedRows() {
                        return []
                }
        },
        methods: {
                isRowSelected(row) {
                        return false
                },
                rowFields(row) {
                        const rowCase = row['case'] || ""
                        const fields = this.type.fields
                        return Object.values(fields).filter(f => f['case'] == rowCase)
                },
                canDisplayRow(row) {
                        if (false == row.expanded) {
                                return true
                        }
                        return false
                }
        },
        template:
`
<v-data-iterator
        :items="rows"
>
        <template v-slot:header>
                <v-toolbar>
                        <kodb-library-rows-menu
                                :libraryName="libraryName"
                                :libraryRows="originRows"

                                :selectedRows="selectedRows"
                        >
                        </kodb-library-rows-menu>
                </v-toolbar>
        </template>

        <template v-slot:default="{ items }">
                <v-card tile>
                        <v-simple-table
                                dense
                        >
                                <thead>
                                        <tr>
                                                <th width="1em">
                                                </th>

                                                <th v-for="f in fields">
                                                        {{ f.name }}
                                                </th>

                                                <th v-if="multiCase"
                                                        width="1em"
                                                >
                                                        Case
                                                </th>

                                                <th v-if="multiCase">
                                                        Type
                                                </th>
                                        </tr>
                                </thead>
                                <tbody>
                                        <tr v-for="i in items" v-if="canDisplayRow(i)">
                                                <td v-if="i.expanded"
                                                        :colspan="multiCase ? 3 : fields.length + 1"
                                                >
                                                </td>

                                                <td v-if="false == i.expanded">
                                                        <v-icon
                                                        >
                                                                {{ isRowSelected(i.row) ? "mdi-check-box-outline" : "mdi-checkbox-blank-outline" }}
                                                        </v-icon>
                                                </td>

                                                <td v-for="f in i.fields"
                                                >
                                                        <kodb-library-cell
                                                                :schema="schema"
                                                                :libraryName="libraryName"
                                                                
                                                                :row="i.row"
                                                                :field="f"
                                                        >
                                                        </kodb-library-cell>
                                                </td>

                                                <td v-if="false == i.expanded && multiCase" 
                                                >
                                                        <kodb-library-card-case
                                                                :libraryName="libraryName"

                                                                :row="i.row"
                                                                :type="type"
                                                        >
                                                        </kodb-library-card-case>
                                                </td>

                                                <td v-if="false == i.expanded && multiCase" 
                                                >
                                                        <kodb-library-card
                                                                :schema="schema"
                                                                :libraryName="libraryName"

                                                                :row="i.row"
                                                                :type="type"
                                                        >
                                                        </kodb-library-card>
                                                </td>
                                        </tr>
                                </tbody>
                        </v-simple-table>
                </v-card>
        </template>
</v-data-iterator>
`
});

Vue.component("kodb-library-card-case", {
        props: [
                "libraryName",

                "row",
                "type"
        ],
        methods: {
                updateCase(c) {
                        this.$wsocket.send({
                                "command": "updateRowCase",
                                "library": this.libraryName,
                                "rowId": this.row.rowId,
                                "case": c
                        })
                }
        },
        template:
`
<v-menu offset-y>
        <template v-slot:activator="{ on }">
                <v-btn outlined block text v-on="on">
                        {{ row['case'] }}
                </v-btn>
        </template>

        <v-list>
                <v-list-item v-for="c in type.cases"
                        :key="c"
                        v-on:click="updateCase(c)"
                >
                        {{ c }}
                </v-list-item>
        </v-list>
</v-menu>
`
})

Vue.component("kodb-library-card", {
        props: [
                "schema",
                "libraryName",

                "row",
                "type"
        ],
        computed: {
                fields() {
                        const rowCase = this.row['case'] || "Mult"
                        const fields = this.type.fields
                        
                        return Object.values(fields).filter(f => f['case'] == rowCase)
                },
                columns() {
                        return [
                                this.fields
                        ]
                }
        },
        template:
`
<v-row justify="space-between">
        <v-col v-for="col in columns">
                <v-row v-for="f in col"
                        :key="f.id"
                >
                        <kodb-library-cell
                                :schema="schema"
                                :libraryName="libraryName"
                                
                                :row="row"
                                :field="f"
                        >
                        </kodb-library-cell>
                </v-row>
        </v-col>
</v-row>
`
})

Vue.component("kodb-library-rows-menu", {
        props: [
                "libraryName",
                "libraryRows",

                "parentLibraryName",
                "parentRowId",
                "parentColumnId",

                "selectedRows"
        ],
        computed: {
                canMoveUp() {
                        if (0 == this.selectedRows.length) {
                                return false
                        }
                        const rows = this.libraryRows
                        for (let r of this.selectedRows) {
                                if (0 == rows.indexOf(r)) {
                                        return false
                                }
                        }
                        return true

                },
                canMoveDown() {
                        if (0 == this.selectedRows.length) {
                                return false
                        }
                        const rows = this.libraryRows
                        const lastRowId = rows.length - 1
                        for (let r of this.selectedRows) {
                                if (lastRowId == rows.indexOf(r)) {
                                        return false
                                }
                        }
                        return true
                },
        },
        methods: {
                newRow() {
                        this.$wsocket.send({
                                "command": "newRow",
                                "library": this.libraryName,

                                "parentLibrary": this.parentLibraryName,
                                "parentRow": this.parentRowId,
                                "parentColumn": this.parentColumnId
                        })
                },
                deleteSelectedRows() {
                        for (let row of this.selectedRows) {
                                this.$wsocket.send({
                                        "command": "deleteRow",
                                        "library": this.libraryName,
                                        "rowId": row.rowId
                                })
                        }
                        //this.selectedRows = []
                },
                swapWith(direction) {
                        if (1 != this.selectedRows.length) {
                                return
                        }
                        const row = this.selectedRows[0]
                        const rowIndex = this.libraryRows.indexOf(row)
                        const row0 = this.libraryRows[rowIndex + direction]

                        this.$wsocket.send({
                                "command": "swapRows",
                                "library": this.libraryName,
                                "row": row.rowId,
                                "row0": row0.rowId
                        })
                },
                moveUp() {
                        this.swapWith(-1)
                },
                moveDown(direction) {
                        this.swapWith(+1)
                }
        },
        template:
`
<v-item-group>
        <v-btn v-on:click="newRow"
                icon color="primary"
        >
                <v-icon>mdi-plus</v-icon>
        </v-btn>
        <v-btn v-on:click="deleteSelectedRows"
                :disabled="selectedRows.length == 0"
                icon color="error"
        >
                <v-icon>mdi-delete</v-icon>
        </v-btn>

        <v-btn :disabled="false == canMoveUp"
                v-on:click="moveUp"
                icon color="primary"

        >
                <v-icon>mdi-arrow-up</v-icon>
        </v-btn>
        <v-btn :disabled="false == canMoveDown"
                v-on:click="moveDown"
                icon color="primary"
        >
                <v-icon>mdi-arrow-down</v-icon>
        </v-btn>
</v-item-group>
`
});