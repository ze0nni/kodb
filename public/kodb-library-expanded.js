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
            filterItems(rows) {
                const columns = this.columns
                const parentId = this.parentRowId
                return (rows||[])
                        .filter(r => {
                                const parent = r.data.parent
                                if (null == parent) {
                                        console.warn(`parentNotExists`)
                                        return false
                                }
                                return parentId == parent.value
                        })
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
<v-card dense outlined :color=colorFromDepth(depth)>
        <v-simple-table dense>
                <thead>
                        <tr color="red">
                                <th :colspan="columns.length + 1">
                                        <v-card>
                                        {{ parentColumnId }}
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
                                        <v-btn small text>
                                                New row
                                        </v-btn>
                                </td>
                        </tr>
                </tfoot>
        </v-simple-table>
</v-card>
`
});