Vue.component("kodb-library-expanded", {
    props: [
            "schema",
            "librarisData",
            "parentRow",
            "columns",
            "librarySchema",
            "rows"
    ],
    data() {
        return {
                expandedLibraryName: null
        }
    },
    methods: {
            filterItems(rows) {
                    const parentId = this.parentRow.rowId
                    return (rows||[])
                            .filter(r => r.data.parent.value == parentId)
                            .map(r => [
                                    {extendedRow: false, row: r, columns: this.librarySchema.columns },
                                    {extendedRow: true, row: r, rcolumns:[] }
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
        <table>
                <thead>
                        <tr>
                                <th v-for="col in librarySchema.columns"
                                >
                                        {{ col.name }}
                                </th>
                        </tr>
                </thead>
                <tbody>
                        <tr v-for="r in filterItems(rows)"
                        >
                                <td v-for="col in r.columns"
                                >
                                        <kodb-library-cell
                                                :libraryName="librarySchema.name"
                                                :rowId="r.row.rowId"
                                                :column="col"
                                                :data="r.row.data"
                                                :librarisData="librarisData"

                                                :expandRow="expandRow(col.reference)"
                                                :isExpanded="false"
                                        >
                                        </kodb-library-cell>
                                </td>

                                <td v-if="r.extendedRow" :colspan="librarySchema.columns.length">
                                        <kodb-library-expanded
                                                v-if="expandedLibraryName"

                                                :schema="schema"
                                                :librarisData="librarisData"
                                                :parentRow="r.row"
                                                :librarySchema="schema[expandedLibraryName]"
                                                :columns="librarySchema"
                                                :rows="librarisData[expandedLibraryName]"
                                        >
                                        </kodb-library-expanded>
                        </kodb-library-expanded>
                                </td>
                        </tr>
                </tbody>
        </table>
</div>
`
});