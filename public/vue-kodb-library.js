Vue.component("kodb-library", {
        props: [
                "librarySchema"
        ],
        data: function() {
                return {
                        rows:[],
                }
        },
        methods: {
                mapColumns(columns) {
                        return columns.map(col => {
                                return {
                                        text: col.name,
                                        value: col.id
                                }
                        })
                },
                isRowExists(item, colName) {
                        return item
                                && item.data
                                && item.data[colName]
                                && item.data[colName].exists
                },
                }
        },
        webSockets: {
                connected() {
                        this.$wsocket.send({
                                "command": "getLibraryRows",
                                "library": this.librarySchema.name
                        })
                },
                command: {
                        setLibraryRows(msg) {
                                if (msg.library == this.librarySchema.name) {
                                        this.rows = msg.rows
                                }
                        }
                }
        },
        template:
`
<v-data-table
        :headers="mapColumns(librarySchema.columns)"
        :items="rows"
        :items-per-page="10"
        item-key="rowId"
>
        <template v-slot:item="{ item,headers }">
                <tr>
                        <td v-for="col in headers"
                        >
                                <div
                                        v-if="isRowExists(item, col.value)"
                                >
                                        {{item.data[col.value].value}}
                                </div>
                                <v-chip 
                                        v-else
                                        color="red">
                                        NIL
                                </v-chip>
                        </td>
                </tr>
        </template>
</v-data-table>
`
});