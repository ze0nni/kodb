Vue.component("kodb-library-expanded", {
    props: [
            "parentRow",
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
            <tr v-for="r in filterItems(rows)"
            >
                    <td>{{ r.rowId }}</td>
            </tr>
    </table>
</div>
`
});