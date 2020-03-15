Vue.component("kodb-library-expanded", {
    props: [
        "depth",
        "schema",
        "libraryName",
        
        "parentLibraryName",
        "parentRowId",
    ],
    data() {
        return {
                expandedLibraryName: null
        }
    },
    methods: {
            filterItems(rows) {
                const columns = this.schema.map[this.libraryName].columns
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
                                {extendedRow: false, row: r, columns: columns },
                                {extendedRow: true, row: r, columns:[] }
                        ])
                        .reduce((a, b) => a.concat(b), [])
            },
            expandRow(library) {
                    return () => {
                        if (this.expandedLibraryName == library) {
                                this.expandedLibraryName = null
                        } else {
                                this.expandedLibraryName = library
                        }
                    }
            },
            colorFromDepth(depth) {
                    return 'grey lighten-2'
            }
    },
    template:
`
<v-card outlined :color=colorFromDepth(depth)>
        <v-simple-table>
                <thead>
                        <tr>
                                <th style="width:1em; min-wdth:1em">
                                        <v-icon>
                                                mdi-chevron-up
                                        </v-icon>
                                </th>
                                <th v-for="col in schema.map[libraryName].columns"
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

                                                :expandRow="expandRow(col.reference)"
                                                :isExpanded="false"
                                        >
                                        </kodb-library-cell>
                                </td>

                                <td v-if="r.extendedRow && r.extendedRow && expandedLibraryName"></td>
                                <td v-if="r.extendedRow && expandedLibraryName"
                                        :colspan="schema.map[libraryName].columns.length">
                                        <kodb-library-expanded
                                                v-if="expandedLibraryName"

                                                :depth="Number(depth)+1"
                                                :schema="schema"
                                                :libraryName="expandedLibraryName"
                                                
                                                :parentLibraryName="libraryName"
                                                :parentRowId="r.row.rowId"
                                        >
                                        </kodb-library-expanded>
                                </td>
                        </tr>
                </tbody>
        </v-simple-table>
</v-card>
`
});