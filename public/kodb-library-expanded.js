Vue.component("kodb-library-expanded", {
    props: [
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
            }
    },
    template:
`
<div>
        <v-simple-table>
                <thead>
                        <tr>
                                <th v-for="col in schema.map[libraryName].columns"
                                >
                                        {{ col.name }}
                                </th>
                        </tr>
                </thead>
                <tbody>
                        <tr v-for="r in filterItems(schema.rowsMap[libraryName])"
                        >
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

                                <td v-if="r.extendedRow && expandedLibraryName" :colspan="schema.map[libraryName].columns.length">
                                        <kodb-library-expanded
                                                v-if="expandedLibraryName"

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
</div>
`
});