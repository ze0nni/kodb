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
                                <td v-for="col in librarySchema.columns"
                                >
                                        <kodb-library-cell
                                                :libraryName="librarySchema.name"
                                                :rowId="r.rowId"
                                                :column="col"
                                                :data="r.data"
                                                :librarisData="librarisData"

                                                :expandRow="function() {}"
                                                :isExpanded="false"
                                        >
                                        </kodb-library-cell>
                                </td>
                        </tr>
                </tbody>
        </table>
</div>
`
});