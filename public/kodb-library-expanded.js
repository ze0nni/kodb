Vue.component("kodb-library-expanded", {
    props: [
            "schema",
            "librarisData",
            "parentRow",
            "columns",
            "librarySchema",
            "rows"
    ],
    methods: {
            filterItems(rows) {
                    const parentId = this.parentRow.rowId
                    return (rows||[])
                            .filter(r => r.data.parent.value == parentId)
                            .map(r => [
                                    {extendedRow: false, row: r, columns: this.librarySchema.columns },
                                    {extendedRow: true, columns:[] }
                                ])
                            .reduce((a, b) => a.concat(b), [])
            },
            expandRow() {
                    console.log(123)
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

                                                :expandRow="function() {}"
                                                :isExpanded="false"
                                        >
                                        </kodb-library-cell>
                                </td>

                                <td v-if="r.extendedRow" :colspan="librarySchema.columns.length">
                                        expanded!
                                </td>
                        </tr>
                </tbody>
        </table>
</div>
`
});