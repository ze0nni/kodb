Vue.component("kodb-library-expanded", {
    props: [
        "depth",
        "schema",
        "libraryName",
        
        "parentLibraryName",
        "parentColumnId",
        "parentRowId",
    ],
    data() {
        return {
                expandedLibraryName: null,
                expandedColumnId: null,
        }
    },
    computed: {
        columns() {
                return this.getLibraryColumns(this.libraryName)
        }
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
            filterItems(rows) {
                const columns = this.columns
                const parentLibraryId = this.parentLibraryName
                const parentRowId = this.parentRowId
                const parentColumnId = this.parentColumnId
                return (rows||[])
                        .filter(r => this.isParentOf(
                                parentLibraryId,
                                parentRowId,
                                parentColumnId,
                                r
                        ))
                        .map(r => [
                                {extendedRow: false, row: r, columns: columns},
                                { extendedRow: true, row: r, columns:[], numColumns: columns.length }
                        ])
                        .reduce((a, b) => a.concat(b), [])
            },
            expandRow(library, columnId) {
                    return () => {
                        if (this.expandedLibraryName == library) {
                                this.expandedLibraryName = null
                        } else {
                                this.expandedLibraryName = library
                                this.expandedColumnId = columnId
                        }
                    }
            },
            colorFromDepth(depth) {
                    return 'grey lighten-2'
            }
    },
    template:
`
<v-card dense tile :color=colorFromDepth(depth)>
        <v-simple-table dense>
                <thead>
                        <tr color="red">
                                <th :colspan="columns.length + 1"
                                        style="margin:0; padding:0"
                                >
                                        <v-card tile dark>
                                                <v-col>
                                                {{ getColumnName(parentLibraryName, parentColumnId) }}
                                                </v-col>
                                        </v-card>
                                </th>
                        </tr>
                        <tr>
                                <th style="width:1em; min-wdth:1em">
                                        <v-icon>
                                                mdi-chevron-up
                                        </v-icon>
                                </th>
                                <th v-for="col in columns"
                                >
                                        {{ col.name }}
                                </th>
                        </tr>
                </thead>
                <tbody>
                        <tr v-for="r in filterItems(schema.rowsMap[libraryName])"
                        >
                                <td v-if="!r.extendedRow"></td>
                                <td v-for="col in r.columns"
                                >
                                        <kodb-library-cell

                                                :schema="schema"
                                                :libraryName="libraryName"
                                                :rowId="r.row.rowId"
                                                :columnId="col.id"

                                                :rowData="r.row"
                                                :cellData="r.row.data[col.value]"

                                                :expandRow="expandRow(col.reference, col.id)"
                                                :isExpanded="false"
                                        >
                                        </kodb-library-cell>
                                </td>

                                <td v-if="r.extendedRow && r.extendedRow && expandedLibraryName"></td>
                                <td v-if="r.extendedRow && expandedLibraryName"
                                        :colspan="r.numColumns">
                                        <kodb-library-expanded
                                                v-if="expandedLibraryName"

                                                :depth="Number(depth)+1"
                                                :schema="schema"
                                                :libraryName="expandedLibraryName"
                                                
                                                :parentLibraryName="libraryName"
                                                :parentColumnId="expandedColumnId"
                                                :parentRowId="r.row.rowId"
                                        >
                                        </kodb-library-expanded>
                                </td>
                        </tr>
                </tbody>
                <tfoot>
                        <tr>
                                <td
                                        :colspan="columns.length + 1"
                                >
                                        <v-btn small text
                                                v-on:click="newRow"
                                        >
                                                New row
                                        </v-btn>
                                </td>
                        </tr>
                </tfoot>
        </v-simple-table>
</v-card>
`
});